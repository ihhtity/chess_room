package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RechargePackageCreateRequest struct {
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"`
	GiftAmount  float64 `json:"gift_amount"`
	GiftPoints  int     `json:"gift_points"`
	Description string  `json:"description"`
}

type RechargePackageUpdateRequest struct {
	Name        string   `json:"name"`
	Amount      *float64 `json:"amount"`
	GiftAmount  *float64 `json:"gift_amount"`
	GiftPoints  *int     `json:"gift_points"`
	Description string   `json:"description"`
	Status      *int     `json:"status"`
}

// @Summary 获取储值套餐列表
// @Description 获取储值套餐列表，支持按名称、状态筛选
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param name query string false "套餐名称"
// @Param status query string false "套餐状态"
// @Success 200 {object} response.Response{data=[]model.RechargePackage}
// @Failure 400 {object} response.Response
// @Router /recharge-packages [get]
func GetRechargePackageList(c *gin.Context) {
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

	packages, total, err := logic.GetRechargePackageListFiltered(name, status, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  packages,
		"total": total,
	})
}

// @Summary 获取储值套餐详情
// @Description 根据套餐ID获取储值套餐详情
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param id path string true "套餐ID"
// @Success 200 {object} response.Response{data=model.RechargePackage}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /recharge-packages/{id} [get]
func GetRechargePackageDetail(c *gin.Context) {
	id := c.Param("id")
	pkg, err := logic.GetRechargePackageByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, pkg)
}

// @Summary 创建储值套餐
// @Description 管理员创建储值套餐
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param body body RechargePackageCreateRequest true "储值套餐信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.RechargePackage}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /recharge-packages [post]
func CreateRechargePackage(c *gin.Context) {
	var req RechargePackageCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	pkg := &model.RechargePackage{
		Name:        req.Name,
		Amount:      req.Amount,
		GiftAmount:  req.GiftAmount,
		GiftPoints:  req.GiftPoints,
		Description: req.Description,
		Status:      1,
	}

	if err := logic.CreateRechargePackage(pkg); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", pkg)
}

// @Summary 更新储值套餐
// @Description 管理员更新储值套餐信息
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param id path string true "套餐ID"
// @Param body body RechargePackageUpdateRequest true "储值套餐更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.RechargePackage}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /recharge-packages/{id} [put]
func UpdateRechargePackage(c *gin.Context) {
	id := c.Param("id")
	var req RechargePackageUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	pkg, err := logic.GetRechargePackageByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Name != "" {
		pkg.Name = req.Name
	}
	if req.Amount != nil {
		pkg.Amount = *req.Amount
	}
	if req.GiftAmount != nil {
		pkg.GiftAmount = *req.GiftAmount
	}
	if req.GiftPoints != nil {
		pkg.GiftPoints = *req.GiftPoints
	}
	if req.Description != "" {
		pkg.Description = req.Description
	}
	if req.Status != nil {
		pkg.Status = *req.Status
	}

	if err := logic.UpdateRechargePackage(id, pkg); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", pkg)
}

// @Summary 删除储值套餐
// @Description 管理员删除储值套餐
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param id path string true "套餐ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /recharge-packages/{id} [delete]
func DeleteRechargePackage(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRechargePackage(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 批量删除储值套餐
// @Description 管理员批量删除储值套餐
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "套餐ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /recharge-packages/batch [delete]
func BatchDeleteRechargePackage(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的储值套餐")
		return
	}
	if err := logic.BatchDeleteRechargePackage(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新储值套餐
// @Description 管理员批量更新储值套餐状态
// @Tags 储值套餐
// @Accept json
// @Produce json
// @Param body body []object{id=int64,status=int} true "储值套餐更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /recharge-packages/batch [put]
func BatchUpdateRechargePackage(c *gin.Context) {
	var reqs []struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateRechargePackage(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
