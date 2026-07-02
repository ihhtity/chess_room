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

// @Summary 用户登录
// @Description 用户登录，支持微信登录、手机号登录和密码登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body UserLoginRequest true "登录信息"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /user/login [post]
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

// @Summary 获取用户信息
// @Description 获取当前登录用户的个人信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.User}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /user/info [get]
func GetUserProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	user, err := logic.GetUserByID(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, user)
}

// @Summary 更新用户信息
// @Description 用户更新个人信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body UserUpdateRequest true "用户更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.User}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /user/info [put]
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

// @Summary 修改密码
// @Description 用户修改登录密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body ChangePasswordRequest true "密码修改信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /user/change-password [post]
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

// @Summary 获取用户列表
// @Description 管理员获取用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.User}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /users [get]
func GetUserList(c *gin.Context) {
	var users []model.User
	err := logic.GetUserList(&users)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, users)
}

type UpdateUserStatusRequest struct {
	Status int `json:"status"`
}

// @Summary 更新用户状态
// @Description 管理员更新用户状态
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Param body body UpdateUserStatusRequest true "用户状态更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.User}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /users/{id}/status [put]
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
