package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CronJobCreateRequest struct {
	Name           string `json:"name" binding:"required"`
	CronExpression string `json:"cron_expression" binding:"required"`
	Handler        string `json:"handler" binding:"required"`
	Params         string `json:"params"`
	Status         int    `json:"status"`
	Description    string `json:"description"`
}

type CronJobUpdateRequest struct {
	Name           *string `json:"name"`
	CronExpression *string `json:"cron_expression"`
	Handler        *string `json:"handler"`
	Params         *string `json:"params"`
	Status         *int    `json:"status"`
	Description    *string `json:"description"`
}

// @Summary 获取定时任务列表
// @Description 获取定时任务列表，支持按名称、状态筛选
// @Tags 定时任务
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "任务名称"
// @Param status query string false "任务状态"
// @Success 200 {object} response.Response{data=[]model.CronJob}
// @Failure 400 {object} response.Response
// @Router /admin/cron-jobs [get]
func GetCronJobList(c *gin.Context) {
	name := c.Query("name")
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
	pageSize := 10
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

	cronJobs, total, err := logic.GetCronJobListAdminFiltered(name, status, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  cronJobs,
		"total": total,
	})
}

// @Summary 获取定时任务详情
// @Description 根据定时任务ID获取任务详情
// @Tags 定时任务
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "任务ID"
// @Success 200 {object} response.Response{data=model.CronJob}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/cron-jobs/{id} [get]
func GetCronJobDetail(c *gin.Context) {
	id := c.Param("id")
	cronJob, err := logic.GetCronJobByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, cronJob)
}

// @Summary 创建定时任务
// @Description 管理员创建新定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param body body CronJobCreateRequest true "定时任务信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.CronJob}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/cron-jobs [post]
func CreateCronJob(c *gin.Context) {
	var req CronJobCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	cronJob := &model.CronJob{
		Name:           req.Name,
		CronExpression: req.CronExpression,
		Handler:        req.Handler,
		Params:         req.Params,
		Status:         req.Status,
		Description:    req.Description,
	}

	if err := logic.CreateCronJob(cronJob); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", cronJob)
}

// @Summary 更新定时任务
// @Description 管理员更新定时任务信息
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param id path string true "任务ID"
// @Param body body CronJobUpdateRequest true "定时任务更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.CronJob}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/cron-jobs/{id} [put]
func UpdateCronJob(c *gin.Context) {
	id := c.Param("id")
	var req CronJobUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	cronJob, err := logic.GetCronJobByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Name != nil {
		cronJob.Name = *req.Name
	}
	if req.CronExpression != nil {
		cronJob.CronExpression = *req.CronExpression
	}
	if req.Handler != nil {
		cronJob.Handler = *req.Handler
	}
	if req.Params != nil {
		cronJob.Params = *req.Params
	}
	if req.Status != nil {
		cronJob.Status = *req.Status
	}
	if req.Description != nil {
		cronJob.Description = *req.Description
	}

	if err := logic.UpdateCronJob(id, cronJob); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", cronJob)
}

// @Summary 删除定时任务
// @Description 管理员删除定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param id path string true "任务ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/cron-jobs/{id} [delete]
func DeleteCronJob(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteCronJob(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "删除成功", nil)
}

// @Summary 批量删除定时任务
// @Description 管理员批量删除定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "任务ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/cron-jobs/batch [delete]
func BatchDeleteCronJob(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的定时任务")
		return
	}
	if err := logic.BatchDeleteCronJob(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新定时任务
// @Description 管理员批量更新定时任务状态
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param body body []object{id=int64,status=int} true "定时任务更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/cron-jobs/batch [put]
func BatchUpdateCronJob(c *gin.Context) {
	var reqs []struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateCronJob(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
