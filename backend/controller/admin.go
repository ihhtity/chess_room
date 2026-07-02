package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"

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
