package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary 管理员登录
// @Description 管理员账号密码登录
// @Tags 管理员
// @Accept json
// @Produce json
// @Param body body AdminLoginRequest true "登录信息"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/login [post]
func AdminLogin(c *gin.Context) {
	var req AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	admin, token, err := logic.AdminLogin(req.Username, req.Password)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"admin": admin,
		"token": token,
	})
}

// @Summary 获取管理员信息
// @Description 获取当前登录管理员的个人信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Admin}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/profile [get]
func GetAdminProfile(c *gin.Context) {
	adminID := c.GetInt64("admin_id")
	admin, err := logic.GetAdminByID(adminID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, admin)
}

type AdminUpdateProfileRequest struct {
	Username string `json:"username"`
	Realname string `json:"realname"`
}

// @Summary 更新管理员信息
// @Description 更新当前登录管理员的个人信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Param body body AdminUpdateProfileRequest true "管理员更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Admin}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/profile [put]
func UpdateAdminProfile(c *gin.Context) {
	adminID := c.GetInt64("admin_id")

	var req AdminUpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	admin, err := logic.UpdateAdminProfile(adminID, req.Username, req.Realname)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", admin)
}

type AdminChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// @Summary 管理员修改密码
// @Description 管理员修改登录密码
// @Tags 管理员
// @Accept json
// @Produce json
// @Param body body AdminChangePasswordRequest true "密码修改信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/change-password [post]
func AdminChangePassword(c *gin.Context) {
	adminID := c.GetInt64("admin_id")

	var req AdminChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.AdminChangePassword(adminID, req.OldPassword, req.NewPassword); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "密码修改成功", nil)
}

// @Summary 获取管理员列表
// @Description 获取管理员列表（只能查看自己和比自己等级低的管理员）
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param username query string false "用户名"
// @Param realname query string false "真实姓名"
// @Param role_id query string false "角色ID"
// @Param status query string false "状态"
// @Success 200 {object} response.Response{data=[]model.Admin}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/admins [get]
func GetAdminList(c *gin.Context) {
	currentAdminID := c.GetInt64("admin_id")

	username := c.Query("username")
	realname := c.Query("realname")
	roleIDStr := c.Query("role_id")
	statusStr := c.Query("status")

	roleID := int64(0)
	if roleIDStr != "" {
		var err error
		roleID, err = strconv.ParseInt(roleIDStr, 10, 64)
		if err != nil {
			response.Fail(c, 400, "角色ID格式错误")
			return
		}
	}

	status := -1
	if statusStr != "" {
		var err error
		status, err = strconv.Atoi(statusStr)
		if err != nil {
			response.Fail(c, 400, "状态格式错误")
			return
		}
	}

	admins, err := logic.GetAdminList(currentAdminID, username, realname, roleID, status)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, admins)
}

type AdminCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Realname string `json:"realname"`
	RoleID   int64  `json:"role_id" binding:"required"`
	Status   int    `json:"status"`
}

// @Summary 创建管理员
// @Description 创建新的管理员（只能创建比自己等级低的管理员）
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body AdminCreateRequest true "管理员信息"
// @Success 200 {object} response.Response{data=model.Admin}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /admin/admins [post]
func CreateAdmin(c *gin.Context) {
	currentAdminID := c.GetInt64("admin_id")

	var req AdminCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	admin, err := logic.CreateAdmin(currentAdminID, req.Username, req.Password, req.Realname, req.RoleID, req.Status)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", admin)
}

type AdminUpdateRequest struct {
	Username string `json:"username"`
	Realname string `json:"realname"`
	RoleID   int64  `json:"role_id"`
	Status   int    `json:"status"`
}

// @Summary 更新管理员信息
// @Description 更新管理员信息（只能更新自己和比自己等级低的管理员）
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "管理员ID"
// @Param body body AdminUpdateRequest true "管理员更新信息"
// @Success 200 {object} response.Response{data=model.Admin}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /admin/admins/{id} [put]
func UpdateAdmin(c *gin.Context) {
	currentAdminID := c.GetInt64("admin_id")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "ID格式错误")
		return
	}

	var req AdminUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	admin, err := logic.UpdateAdmin(currentAdminID, id, req.Username, req.Realname, req.RoleID, req.Status)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", admin)
}

// @Summary 删除管理员
// @Description 删除管理员（只能删除比自己等级低的管理员）
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "管理员ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /admin/admins/{id} [delete]
func DeleteAdmin(c *gin.Context) {
	currentAdminID := c.GetInt64("admin_id")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "ID格式错误")
		return
	}

	if err := logic.DeleteAdmin(currentAdminID, id); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}

type AdminResetPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

// @Summary 重置管理员密码
// @Description 重置管理员密码（只能重置比自己等级低的管理员）
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "管理员ID"
// @Param body body AdminResetPasswordRequest true "新密码"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /admin/admins/{id}/reset-password [post]
func ResetAdminPassword(c *gin.Context) {
	currentAdminID := c.GetInt64("admin_id")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "ID格式错误")
		return
	}

	var req AdminResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.ResetAdminPassword(currentAdminID, id, req.NewPassword); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "密码重置成功", nil)
}
