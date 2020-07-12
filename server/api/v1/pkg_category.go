package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags PkgCategory
// @Summary 创建PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "创建PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgCtg/createPkgCategory [post]
func CreatePkgCategory(c *gin.Context) {
	var pkgCtg model.PkgCategory
	_ = c.ShouldBindJSON(&pkgCtg)
	err := service.CreatePkgCategory(pkgCtg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PkgCategory
// @Summary 删除PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "删除PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkgCtg/deletePkgCategory [delete]
func DeletePkgCategory(c *gin.Context) {
	var pkgCtg model.PkgCategory
	_ = c.ShouldBindJSON(&pkgCtg)
	err := service.DeletePkgCategory(pkgCtg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PkgCategory
// @Summary 更新PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "更新PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkgCtg/updatePkgCategory [put]
func UpdatePkgCategory(c *gin.Context) {
	var pkgCtg model.PkgCategory
	_ = c.ShouldBindJSON(&pkgCtg)
	err := service.UpdatePkgCategory(&pkgCtg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PkgCategory
// @Summary 用id查询PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "用id查询PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkgCtg/findPkgCategory [get]
func FindPkgCategory(c *gin.Context) {
	var pkgCtg model.PkgCategory
	_ = c.ShouldBindQuery(&pkgCtg)
	err, repkgCtg := service.GetPkgCategory(pkgCtg.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"repkgCtg": repkgCtg}, c)
	}
}

// @Tags PkgCategory
// @Summary 分页获取PkgCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取PkgCategory列表page"
// @Param pageSize query string true "分页获取PkgCategory列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgCtg/getPkgCategoryList [get]
func GetPkgCategoryList(c *gin.Context) {
	var pageInfo request.PkgCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetPkgCategoryInfoList(pageInfo)
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
