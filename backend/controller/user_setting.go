package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func GetUserSetting(c *gin.Context) {
	userID := c.GetInt64("user_id")
	setting, err := logic.GetUserSetting(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, setting)
}

type UpdateUserSettingRequest struct {
	Notifications  int    `json:"notifications"`
	Sound          int    `json:"sound"`
	Vibrate        int    `json:"vibrate"`
	Language       string `json:"language"`
	Theme          string `json:"theme"`
	AllowPush      int    `json:"allow_push"`
	AllowMarketing int    `json:"allow_marketing"`
}

func UpdateUserSetting(c *gin.Context) {
	userID := c.GetInt64("user_id")
	
	var req UpdateUserSettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	
	updates := make(map[string]interface{})
	
	if req.Notifications >= 0 {
		updates["notifications"] = req.Notifications
	}
	if req.Sound >= 0 {
		updates["sound"] = req.Sound
	}
	if req.Vibrate >= 0 {
		updates["vibrate"] = req.Vibrate
	}
	if req.Language != "" {
		updates["language"] = req.Language
	}
	if req.Theme != "" {
		updates["theme"] = req.Theme
	}
	if req.AllowPush >= 0 {
		updates["allow_push"] = req.AllowPush
	}
	if req.AllowMarketing >= 0 {
		updates["allow_marketing"] = req.AllowMarketing
	}
	
	setting, err := logic.UpdateUserSetting(userID, updates)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	
	response.Success(c, setting)
}

type ToggleSettingRequest struct {
	Key string `json:"key" binding:"required"`
}

func ToggleUserSetting(c *gin.Context) {
	userID := c.GetInt64("user_id")
	
	var req ToggleSettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	
	setting, err := logic.ToggleSetting(userID, req.Key)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	
	response.Success(c, setting)
}
