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

// @Tags Disease
// @Summary 创建Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "创建Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /disease/createDisease [post]
func CreateDisease(c *gin.Context) {
	var disease model.Disease
	_ = c.ShouldBindJSON(&disease)
	err := service.CreateDisease(disease)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Disease
// @Summary 删除Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "删除Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disease/deleteDisease [delete]
func DeleteDisease(c *gin.Context) {
	var disease model.Disease
	_ = c.ShouldBindJSON(&disease)
	err := service.DeleteDisease(disease)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Disease
// @Summary 批量删除Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disease/deleteDiseaseByIds [delete]
func DeleteDiseaseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDiseaseByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Disease
// @Summary 更新Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "更新Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /disease/updateDisease [put]
func UpdateDisease(c *gin.Context) {
	var disease model.Disease
	_ = c.ShouldBindJSON(&disease)
	err := service.UpdateDisease(&disease)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Disease
// @Summary 用id查询Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "用id查询Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /disease/findDisease [get]
func FindDisease(c *gin.Context) {
	var disease model.Disease
	_ = c.ShouldBindQuery(&disease)
	err, redisease := service.GetDisease(disease.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redisease": redisease}, c)
	}
}

// @Tags Disease
// @Summary 分页获取Disease列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取Disease列表page"
// @Param pageSize query string true "分页获取Disease列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /disease/getDiseaseList [get]
func GetDiseaseList(c *gin.Context) {
	var pageInfo request.DiseaseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDiseaseInfoList(pageInfo)
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
