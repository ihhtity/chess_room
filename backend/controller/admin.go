package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

func GetAdminProfile(c *gin.Context) {
	adminID := c.GetInt64("admin_id")
	admin, err := logic.GetAdminByID(adminID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, admin)
}