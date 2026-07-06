package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OperationLogCreateRequest struct {
	AdminID  int64  `json:"admin_id" binding:"required" comment:"管理员ID"`
	Action   string `json:"action" binding:"required" comment:"操作类型"`
	Module   string `json:"module" comment:"模块"`
	TargetID int64  `json:"target_id" comment:"目标ID"`
	Content  string `json:"content" comment:"操作内容"`
	IP       string `json:"ip" comment:"IP"`
}

// @Summary 获取操作日志列表
// @Description 获取所有操作日志列表
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param admin_id query string false "管理员ID"
// @Param module query string false "模块"
// @Success 200 {object} response.Response
// @Router /operation-logs [get]
func GetOperationLogList(c *gin.Context) {
	adminID := c.Query("admin_id")
	module := c.Query("module")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	var page int = 1
	var pageSize int = 10
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10
		}
	}

	logs, total, err := logic.GetOperationLogList(adminID, module, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  logs,
		"total": total,
	})
}

// @Summary 获取操作日志详情
// @Description 根据ID获取操作日志详情
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param id path string true "操作日志ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /operation-logs/{id} [get]
func GetOperationLogDetail(c *gin.Context) {
	id := c.Param("id")
	log, err := logic.GetOperationLogByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, log)
}

// @Summary 创建操作日志
// @Description 创建新的操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param body body OperationLogCreateRequest true "操作日志信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /operation-logs [post]
func CreateOperationLog(c *gin.Context) {
	var req OperationLogCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	log := &model.OperationLog{
		AdminID:  req.AdminID,
		Action:   req.Action,
		Module:   req.Module,
		TargetID: req.TargetID,
		Content:  req.Content,
		IP:       req.IP,
	}

	if err := logic.CreateOperationLog(log); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, log)
}

// @Summary 删除操作日志
// @Description 删除操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param id path string true "操作日志ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /operation-logs/{id} [delete]
func DeleteOperationLog(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteOperationLog(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 批量删除操作日志
// @Description 批量删除操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param body body struct{ IDs []string `json:"ids"` } true "操作日志ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /operation-logs/batch [delete]
func BatchDeleteOperationLog(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if err := logic.BatchDeleteOperationLog(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新操作日志
// @Description 批量更新操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param body body []model.OperationLog true "操作日志列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /operation-logs/batch [put]
func BatchUpdateOperationLog(c *gin.Context) {
	var logs []model.OperationLog
	if err := c.ShouldBindJSON(&logs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if err := logic.BatchUpdateOperationLog(logs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量更新成功", nil)
}
