package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionCreateRequest struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Group       string `json:"group"`
	Module      string `json:"module"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

type PermissionUpdateRequest struct {
	Name        string `json:"name"`
	Group       string `json:"group"`
	Module      string `json:"module"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

type AssignPermissionRequest struct {
	PermissionIDs []int64 `json:"permission_ids" binding:"required"`
}

// @Summary 获取权限列表
// @Description 获取所有权限列表
// @Tags 权限管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.Permission}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/permissions [get]
func GetPermissionList(c *gin.Context) {
	permissions, err := logic.GetPermissionList()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, permissions)
}

// @Summary 获取权限列表（按分组）
// @Description 获取所有权限列表，按分组返回
// @Tags 权限管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.Permission}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/permissions/grouped [get]
func GetPermissionListByGroup(c *gin.Context) {
	permissions, err := logic.GetPermissionListByGroup()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, permissions)
}

// @Summary 获取权限详情
// @Description 根据权限ID获取权限详情
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param id path string true "权限ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Permission}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/permissions/{id} [get]
func GetPermissionDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	permission, err := logic.GetPermissionByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, permission)
}

// @Summary 创建权限
// @Description 创建新权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param body body PermissionCreateRequest true "权限信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Permission}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/permissions [post]
func CreatePermission(c *gin.Context) {
	var req PermissionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	permission := &model.Permission{
		Code:        req.Code,
		Name:        req.Name,
		Group:       req.Group,
		Module:      req.Module,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}

	if err := logic.CreatePermission(permission); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", permission)
}

// @Summary 更新权限
// @Description 更新权限信息
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param id path string true "权限ID"
// @Param body body PermissionUpdateRequest true "权限信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Permission}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/permissions/{id} [put]
func UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req PermissionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	permission := &model.Permission{}
	if req.Name != "" {
		permission.Name = req.Name
	}
	if req.Group != "" {
		permission.Group = req.Group
	}
	if req.Module != "" {
		permission.Module = req.Module
	}
	if req.Description != "" {
		permission.Description = req.Description
	}
	if req.SortOrder != 0 {
		permission.SortOrder = req.SortOrder
	}

	if err := logic.UpdatePermission(idInt, permission); err != nil {
		response.HandleError(c, err)
		return
	}

	updatedPermission, _ := logic.GetPermissionByID(idInt)
	response.SuccessWithMsg(c, "更新成功", updatedPermission)
}

// @Summary 删除权限
// @Description 删除权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param id path string true "权限ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/permissions/{id} [delete]
func DeletePermission(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.DeletePermission(idInt); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}

// @Summary 获取角色权限
// @Description 根据角色ID获取该角色的所有权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param role_id query string true "角色ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.Permission}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/permissions/role [get]
func GetRolePermissions(c *gin.Context) {
	roleIDStr := c.Query("role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	permissions, err := logic.GetPermissionsByRoleID(roleID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, permissions)
}

// @Summary 分配权限给角色
// @Description 为角色分配权限（只能分配自己拥有的权限给层级低于自己的角色）
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param role_id path string true "角色ID"
// @Param body body AssignPermissionRequest true "权限ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /admin/permissions/role/{role_id} [post]
func AssignPermissions(c *gin.Context) {
	roleIDStr := c.Param("role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req AssignPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	currentRoleID := c.GetInt64("role_id")

	if err := logic.AssignPermissions(currentRoleID, roleID, req.PermissionIDs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "权限分配成功", nil)
}

// @Summary 获取当前管理员的权限列表
// @Description 获取当前登录管理员的所有权限编码
// @Tags 权限管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]string}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /admin/permissions/mine [get]
func GetMyPermissions(c *gin.Context) {
	roleID := c.GetInt64("role_id")

	codes, err := logic.GetPermissionCodesByRoleID(roleID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, codes)
}