package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"

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
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"`
	GiftAmount  float64 `json:"gift_amount"`
	GiftPoints  int     `json:"gift_points"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}

func GetRechargePackageList(c *gin.Context) {
	packages, err := logic.GetRechargePackageList()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, packages)
}

func GetRechargePackageDetail(c *gin.Context) {
	id := c.Param("id")
	pkg, err := logic.GetRechargePackageByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, pkg)
}

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

func UpdateRechargePackage(c *gin.Context) {
	id := c.Param("id")
	var req RechargePackageUpdateRequest
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
		Status:      req.Status,
	}

	if err := logic.UpdateRechargePackage(id, pkg); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", nil)
}

func DeleteRechargePackage(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRechargePackage(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
