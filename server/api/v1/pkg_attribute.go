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

// @Tags PkgAttr
// @Summary 创建PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "创建PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgAttr/createPkgAttr [post]
func CreatePkgAttr(c *gin.Context) {
	var pkgAttr model.PkgAttr
	_ = c.ShouldBindJSON(&pkgAttr)
	err := service.CreatePkgAttr(pkgAttr)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PkgAttr
// @Summary 删除PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "删除PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkgAttr/deletePkgAttr [delete]
func DeletePkgAttr(c *gin.Context) {
	var pkgAttr model.PkgAttr
	_ = c.ShouldBindJSON(&pkgAttr)
	err := service.DeletePkgAttr(pkgAttr)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PkgAttr
// @Summary 更新PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "更新PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkgAttr/updatePkgAttr [put]
func UpdatePkgAttr(c *gin.Context) {
	var pkgAttr model.PkgAttr
	_ = c.ShouldBindJSON(&pkgAttr)
	err := service.UpdatePkgAttr(&pkgAttr)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PkgAttr
// @Summary 用id查询PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "用id查询PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkgAttr/findPkgAttr [get]
func FindPkgAttr(c *gin.Context) {
	var pkgAttr model.PkgAttr
	_ = c.ShouldBindQuery(&pkgAttr)
	err, repkgAttr := service.GetPkgAttr(pkgAttr.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"repkgAttr": repkgAttr}, c)
	}
}

// @Tags PkgAttr
// @Summary 分页获取PkgAttr列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取PkgAttr列表page"
// @Param pageSize query string true "分页获取PkgAttr列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgAttr/getPkgAttrList [get]
func GetPkgAttrList(c *gin.Context) {
	var pageInfo request.PkgAttrSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetPkgAttrInfoList(pageInfo)
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

//
// func GetPkgAttrList(c *gin.Context) {
// 	var pkgAttr model.PkgAttr
// 	_ = c.ShouldBindQuery(&pkgAttr)
// 	pkgAttrs, err := service.GetPkgAttrList(pkgAttr.ID)
// 	if err != nil {
// 		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
// 	} else {
// 		response.OkWithData(gin.H{"list": pkgAttrs, "id": pkgAttr.ID}, c)
// 	}
// }
