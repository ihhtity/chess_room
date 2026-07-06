package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	OpenID   string `json:"open_id"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
}

type UserUpdateRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Realname string `json:"realname"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func UserLogin(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	user, token, err := logic.UserLogin(req.OpenID, req.Phone, req.Password, req.Nickname, req.Avatar, req.Gender)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"user":  user,
		"token": token,
	})
}

func GetUserProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	user, err := logic.GetUserByID(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, user)
}

func UpdateUserProfile(c *gin.Context) {
	var req UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	user, err := logic.GetUserByID(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Realname != "" {
		user.Realname = req.Realname
	}

	if err := logic.UpdateUser(user); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, user)
}

func ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		response.Fail(c, 400, "密码不能为空")
		return
	}

	if len(req.NewPassword) < 6 {
		response.Fail(c, 400, "新密码长度不能少于6位")
		return
	}

	userID := c.GetInt64("user_id")
	if err := logic.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "密码修改成功", nil)
}

func GetUserList(c *gin.Context) {
	nickname := c.Query("nickname")
	phone := c.Query("phone")
	statusStr := c.Query("status")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	status := -1
	var err error
	if statusStr != "" {
		status, err = strconv.Atoi(statusStr)
		if err != nil {
			response.Fail(c, 400, "状态格式错误")
			return
		}
	}

	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			response.Fail(c, 400, "页码格式错误")
			return
		}
	}

	pageSize := 10
	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			response.Fail(c, 400, "每页数量格式错误")
			return
		}
	}

	users, total, err := logic.GetUserListFiltered(nickname, phone, status, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"data":  users,
		"total": total,
	})
}

func GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	user, err := logic.GetUserByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, user)
}

type UpdateUserStatusRequest struct {
	Status int `json:"status"`
}

func UpdateUserStatus(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req UpdateUserStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	user, err := logic.UpdateUserStatus(idInt, req.Status)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", user)
}

type AdminCreateUserRequest struct {
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Realname string `json:"realname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
}

func CreateUser(c *gin.Context) {
	var req AdminCreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	user := &model.User{
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Realname: req.Realname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		Status:   req.Status,
	}

	if err := logic.CreateUser(user); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", user)
}

type AdminUpdateUserRequest struct {
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Realname string `json:"realname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req AdminUpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	user, err := logic.GetUserByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Realname != "" {
		user.Realname = req.Realname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Status >= 0 {
		user.Status = req.Status
	}

	if err := logic.UpdateUser(user); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.DeleteUser(idInt); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}

type BatchDeleteRequest struct {
	IDs []int64 `json:"ids"`
}

func BatchDeleteUser(c *gin.Context) {
	var req BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的用户")
		return
	}

	if err := logic.BatchDeleteUser(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量删除成功", nil)
}

func BatchUpdateUser(c *gin.Context) {
	var reqs []struct {
		ID       int    `json:"id"`
		Nickname string `json:"nickname"`
		Phone    string `json:"phone"`
		Realname string `json:"realname"`
		Gender   int    `json:"gender"`
		Status   int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if len(reqs) == 0 {
		response.Fail(c, 400, "请选择要更新的用户")
		return
	}

	if err := logic.BatchUpdateUser(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
