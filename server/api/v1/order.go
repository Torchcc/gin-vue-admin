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

// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
func CreateOrder(c *gin.Context) {
	var order model.Order
	_ = c.ShouldBindJSON(&order)
	err := service.CreateOrder(order)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func DeleteOrder(c *gin.Context) {
	var order model.Order
	_ = c.ShouldBindJSON(&order)
	err := service.DeleteOrder(order)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func UpdateOrder(c *gin.Context) {
	var order model.Order
	_ = c.ShouldBindJSON(&order)
	err := service.UpdateOrder(&order)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func FindOrder(c *gin.Context) {
	var order model.Order
	_ = c.ShouldBindQuery(&order)
	err, reorder := service.GetOrder(order.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取Order列表page"
// @Param pageSize query string true "分页获取Order列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func GetOrderList(c *gin.Context) {
	var pageInfo request.OrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetOrderInfoList(pageInfo)
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
