package middleware

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, exists := c.Get("admin_id")
		if !exists {
			c.Next()
			return
		}

		adminIDInt, ok := adminID.(int64)
		if !ok {
			c.Next()
			return
		}

		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()

		var bodyData []byte
		if method == "POST" || method == "PUT" {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err == nil {
				bodyData = body
				c.Request.Body.Close()
				c.Request.Body = ioutil.NopCloser(strings.NewReader(string(body)))
			}
		}

		c.Next()

		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			return
		}

		module := extractModule(path)
		action := extractAction(method, path)
		targetID := extractTargetID(path)

		var content string
		if len(bodyData) > 0 {
			content = string(bodyData)
			if len(content) > 500 {
				content = content[:500] + "..."
			}
		}

		if module != "" && action != "" {
			log := &model.OperationLog{
				AdminID:  adminIDInt,
				Action:   action,
				Module:   module,
				TargetID: targetID,
				Content:  content,
				IP:       ip,
			}
			go logic.CreateOperationLog(log)
		}
	}
}

func extractModule(path string) string {
	pathParts := strings.Split(path, "/")
	if len(pathParts) >= 3 {
		module := pathParts[2]
		if module == "admin" && len(pathParts) >= 4 {
			switch pathParts[3] {
			case "login":
				return "admin"
			case "profile":
				return "admin"
			case "reset-password":
				return "admin"
			case "change-password":
				return "admin"
			default:
				return module
			}
		}
		return module
	}
	return ""
}

func extractAction(method, path string) string {
	switch method {
	case "POST":
		if strings.HasSuffix(path, "/login") {
			return "login"
		}
		if strings.Contains(path, "/recharge") {
			return "recharge"
		}
		if strings.Contains(path, "/assign") {
			return "assign"
		}
		return "create"
	case "PUT":
		if strings.Contains(path, "/reset-password") {
			return "reset_password"
		}
		if strings.Contains(path, "/change-password") {
			return "change_password"
		}
		if strings.Contains(path, "/status") {
			return "update_status"
		}
		if strings.Contains(path, "/confirm") {
			return "confirm"
		}
		if strings.Contains(path, "/complete") {
			return "complete"
		}
		if strings.Contains(path, "/cancel") {
			return "cancel"
		}
		if strings.Contains(path, "/assign") {
			return "assign"
		}
		return "update"
	case "DELETE":
		return "delete"
	case "GET":
		if strings.Contains(path, "/profile") {
			return "view_profile"
		}
		if strings.Contains(path, "/mine") {
			return "view_permissions"
		}
		return ""
	default:
		return ""
	}
}

func extractTargetID(path string) int64 {
	pathParts := strings.Split(path, "/")
	for i, part := range pathParts {
		if i > 0 && pathParts[i-1] != "admin" && pathParts[i-1] != "admins" {
			if id, err := json.Number(part).Int64(); err == nil {
				return id
			}
		}
	}
	return 0
}
