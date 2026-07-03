package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminRoleCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Level       int    `json:"level" binding:"required"`
	Description string `json:"description"`
}

type AdminRoleUpdateRequest struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

// @Summary 获取角色列表
// @Description 获取所有角色列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.AdminRole}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/roles [get]
func GetAdminRoleList(c *gin.Context) {
	roles, err := logic.GetAdminRoleList()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, roles)
}

// @Summary 获取角色详情
// @Description 根据角色ID获取角色详情
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path string true "角色ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.AdminRole}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/roles/{id} [get]
func GetAdminRoleDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	role, err := logic.GetAdminRoleByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, role)
}

// @Summary 创建角色
// @Description 创建新角色（只能创建层级低于自己的角色）
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param body body AdminRoleCreateRequest true "角色信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.AdminRole}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /admin/roles [post]
func CreateAdminRole(c *gin.Context) {
	var req AdminRoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	currentRoleID := c.GetInt64("role_id")

	role := &model.AdminRole{
		Name:        req.Name,
		Level:       req.Level,
		Description: req.Description,
		Status:      1,
	}

	if err := logic.CreateAdminRole(currentRoleID, role); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", role)
}

// @Summary 更新角色
// @Description 更新角色信息（只能修改层级低于自己的角色）
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path string true "角色ID"
// @Param body body AdminRoleUpdateRequest true "角色信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.AdminRole}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/roles/{id} [put]
func UpdateAdminRole(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req AdminRoleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	currentRoleID := c.GetInt64("role_id")

	role := &model.AdminRole{}
	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Level != 0 {
		role.Level = req.Level
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	if req.Status != 0 {
		role.Status = req.Status
	}

	if err := logic.UpdateAdminRole(currentRoleID, idInt, role); err != nil {
		response.HandleError(c, err)
		return
	}

	updatedRole, _ := logic.GetAdminRoleByID(idInt)
	response.SuccessWithMsg(c, "更新成功", updatedRole)
}

// @Summary 删除角色
// @Description 删除角色（只能删除层级低于自己的角色）
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path string true "角色ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/roles/{id} [delete]
func DeleteAdminRole(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	currentRoleID := c.GetInt64("role_id")

	if err := logic.DeleteAdminRole(currentRoleID, idInt); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}

// @Summary 获取当前管理员可管理的角色列表
// @Description 获取当前管理员可以分配的角色列表（层级低于自己的角色）
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.AdminRole}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/roles/available [get]
func GetAvailableRoles(c *gin.Context) {
	currentRoleID := c.GetInt64("role_id")
	currentRole, err := logic.GetAdminRoleByID(currentRoleID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	roles, err := logic.GetRolesByLevel(currentRole.Level + 1)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, roles)
}
