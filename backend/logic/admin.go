package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"chess-room-backend/pkg/jwt"
	"chess-room-backend/pkg/utils"
	"encoding/json"

	"github.com/jinzhu/gorm"
)

func AdminLogin(username, password string) (*model.Admin, string, error) {
	admin, err := mysql.GetAdminByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "", errno.New(errno.AdminNotFound)
		}
		return nil, "", errno.New(errno.InternalError)
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return nil, "", errno.New(errno.PasswordError)
	}

	if admin.Status != 1 {
		return nil, "", errno.New(errno.AdminDisabled)
	}

	token, err := jwt.GenerateToken(admin.ID, admin.Username, 1, admin.RoleID)
	if err != nil {
		return nil, "", errno.New(errno.InternalError)
	}

	return admin, token, nil
}

func GetAdminByID(id int64) (*model.Admin, error) {
	admin, err := mysql.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.AdminNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return admin, nil
}

func UpdateAdminProfile(id int64, username, realname string) (*model.Admin, error) {
	admin, err := mysql.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.AdminNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	oldAdmin := *admin
	if username != "" {
		admin.Username = username
	}
	if realname != "" {
		admin.Realname = realname
	}

	if err := mysql.UpdateAdmin(admin); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	go RecordOperationLog(id, "admin", "update_profile", admin.ID, "更新管理员个人信息: "+jsonDiff(oldAdmin, *admin))

	return admin, nil
}

func AdminChangePassword(id int64, oldPassword, newPassword string) error {
	admin, err := mysql.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.AdminNotFound)
		}
		return errno.New(errno.InternalError)
	}

	if !utils.CheckPasswordHash(oldPassword, admin.Password) {
		return errno.New(errno.PasswordError)
	}

	newPasswordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	admin.Password = newPasswordHash

	if err := mysql.UpdateAdmin(admin); err != nil {
		return errno.New(errno.InternalError)
	}

	go RecordOperationLog(id, "admin", "change_password", admin.ID, "修改管理员密码")

	return nil
}

func getAdminRoleLevel(adminID int64) (int, error) {
	admin, err := mysql.GetAdminByID(adminID)
	if err != nil {
		return 0, err
	}

	if admin.RoleID == 0 {
		return 999, nil
	}

	role, err := mysql.GetAdminRoleByID(admin.RoleID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 999, nil
		}
		return 0, err
	}

	return role.Level, nil
}

func canManageAdmin(currentAdminID, targetAdminID int64) (bool, error) {
	if currentAdminID == targetAdminID {
		return true, nil
	}

	currentLevel, err := getAdminRoleLevel(currentAdminID)
	if err != nil {
		return false, err
	}

	targetLevel, err := getAdminRoleLevel(targetAdminID)
	if err != nil {
		return false, err
	}

	return currentLevel < targetLevel, nil
}

func GetAdminList(currentAdminID int64, username, realname string, roleID int64, status int) ([]model.Admin, error) {
	currentLevel, err := getAdminRoleLevel(currentAdminID)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}

	admins, err := mysql.GetAdminList(username, realname, roleID, status, currentLevel)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return admins, nil
}

func CreateAdmin(currentAdminID int64, username, password, realname string, roleID int64, status int) (*model.Admin, error) {
	if roleID == 0 {
		roleID = 1
	}

	targetRole, err := mysql.GetAdminRoleByID(roleID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.RoleNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	currentLevel, err := getAdminRoleLevel(currentAdminID)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}

	if currentLevel >= targetRole.Level {
		return nil, errno.New(errno.PermissionDenied)
	}

	existingAdmin, _ := mysql.GetAdminByUsername(username)
	if existingAdmin.ID != 0 {
		return nil, errno.New(errno.UsernameExists)
	}

	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}

	if status == 0 {
		status = 1
	}

	admin := &model.Admin{
		Username: username,
		Password: passwordHash,
		Realname: realname,
		RoleID:   roleID,
		Status:   status,
	}

	if err := mysql.CreateAdmin(admin); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	go RecordOperationLog(currentAdminID, "admin", "create", admin.ID, "创建管理员: "+username)

	return admin, nil
}

func UpdateAdmin(currentAdminID, targetAdminID int64, username, realname string, roleID int64, status int) (*model.Admin, error) {
	canManage, err := canManageAdmin(currentAdminID, targetAdminID)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	if !canManage {
		return nil, errno.New(errno.PermissionDenied)
	}

	admin, err := mysql.GetAdminByID(targetAdminID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.AdminNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	oldAdmin := *admin

	if username != "" {
		if username != admin.Username {
			existingAdmin, _ := mysql.GetAdminByUsername(username)
			if existingAdmin.ID != 0 && existingAdmin.ID != admin.ID {
				return nil, errno.New(errno.UsernameExists)
			}
			admin.Username = username
		}
	}
	if realname != "" {
		admin.Realname = realname
	}
	if roleID != 0 {
		targetRole, err := mysql.GetAdminRoleByID(roleID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errno.New(errno.RoleNotFound)
			}
			return nil, errno.New(errno.InternalError)
		}

		currentLevel, err := getAdminRoleLevel(currentAdminID)
		if err != nil {
			return nil, errno.New(errno.InternalError)
		}

		if currentLevel >= targetRole.Level {
			return nil, errno.New(errno.PermissionDenied)
		}

		admin.RoleID = roleID
	}
	if status >= 0 {
		admin.Status = status
	}

	if err := mysql.UpdateAdmin(admin); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	go RecordOperationLog(currentAdminID, "admin", "update", admin.ID, "更新管理员信息: "+jsonDiff(oldAdmin, *admin))

	return admin, nil
}

func DeleteAdmin(currentAdminID, targetAdminID int64) error {
	if currentAdminID == targetAdminID {
		return errno.New(errno.CannotDeleteSelf)
	}

	canManage, err := canManageAdmin(currentAdminID, targetAdminID)
	if err != nil {
		return errno.New(errno.InternalError)
	}
	if !canManage {
		return errno.New(errno.PermissionDenied)
	}

	admin, err := mysql.GetAdminByID(targetAdminID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.AdminNotFound)
		}
		return errno.New(errno.InternalError)
	}

	if err := mysql.DeleteAdmin(targetAdminID); err != nil {
		return errno.New(errno.InternalError)
	}

	go RecordOperationLog(currentAdminID, "admin", "delete", targetAdminID, "删除管理员: "+admin.Username)

	return nil
}

func ResetAdminPassword(currentAdminID, targetAdminID int64, newPassword string) error {
	canManage, err := canManageAdmin(currentAdminID, targetAdminID)
	if err != nil {
		return errno.New(errno.InternalError)
	}
	if !canManage {
		return errno.New(errno.PermissionDenied)
	}

	admin, err := mysql.GetAdminByID(targetAdminID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.AdminNotFound)
		}
		return errno.New(errno.InternalError)
	}

	newPasswordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	admin.Password = newPasswordHash

	if err := mysql.UpdateAdmin(admin); err != nil {
		return errno.New(errno.InternalError)
	}

	go RecordOperationLog(currentAdminID, "admin", "reset_password", targetAdminID, "重置管理员密码: "+admin.Username)

	return nil
}

func jsonDiff(old, new model.Admin) string {
	oldBytes, _ := json.Marshal(map[string]interface{}{
		"username": old.Username,
		"realname": old.Realname,
		"role_id":  old.RoleID,
		"status":   old.Status,
	})
	newBytes, _ := json.Marshal(map[string]interface{}{
		"username": new.Username,
		"realname": new.Realname,
		"role_id":  new.RoleID,
		"status":   new.Status,
	})
	return string(oldBytes) + " -> " + string(newBytes)
}

func RecordOperationLog(adminID int64, module, action string, targetID int64, content string) {
	log := &model.OperationLog{
		AdminID:  adminID,
		Action:   action,
		Module:   module,
		TargetID: targetID,
		Content:  content,
	}
	_ = CreateOperationLog(log)
}
