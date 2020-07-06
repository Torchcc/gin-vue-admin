package v1

import (
	"fmt"
	"strconv"

	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Package
// @Summary 创建Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "创建Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkg/createPackage [post]
func CreatePackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindJSON(&pkg)
	err := service.CreatePackage(pkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Package
// @Summary 删除Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "删除Package, data 只需要传主键id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkg/deletePackage [delete]
func DeletePackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindJSON(&pkg)
	err := service.DeletePackage(&pkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Package
// @Summary 更新Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "更新Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkg/updatePackage [put]
func UpdatePackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindJSON(&pkg)
	err := service.UpdatePackage(&pkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Package
// @Summary 用id查询Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "用id查询Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkg/findPackage/{id} [get]
func FindPackage(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err, repkg := service.GetPackage(uint(id))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"repkg": repkg}, c)
	}
}

// @Tags Package
// @Summary 分页获取Package列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取Package列表page"
// @Param pageSize query string true "分页获取Package列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkg/getPackageList [get]
func GetPackageList(c *gin.Context) {
	var pageInfo request.PackageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetPackageInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
