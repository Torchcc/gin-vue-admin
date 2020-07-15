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

// @Tags OrderItem
// @Summary 创建OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "创建OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderItem/createOrderItem [post]
func CreateOrderItem(c *gin.Context) {
	var orderItem model.OrderItem
	_ = c.ShouldBindJSON(&orderItem)
	err := service.CreateOrderItem(orderItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags OrderItem
// @Summary 删除OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "删除OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderItem/deleteOrderItem [delete]
func DeleteOrderItem(c *gin.Context) {
	var orderItem model.OrderItem
	_ = c.ShouldBindJSON(&orderItem)
	err := service.DeleteOrderItem(orderItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags OrderItem
// @Summary 更新OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "更新OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderItem/updateOrderItem [put]
func UpdateOrderItem(c *gin.Context) {
	var orderItem model.OrderItem
	_ = c.ShouldBindJSON(&orderItem)
	err := service.UpdateOrderItem(&orderItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags OrderItem
// @Summary 用id查询OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "用id查询OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderItem/findOrderItem [get]
func FindOrderItem(c *gin.Context) {
	var orderItem model.OrderItem
	_ = c.ShouldBindQuery(&orderItem)
	err, reorderItem := service.GetOrderItem(orderItem.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reorderItem": reorderItem}, c)
	}
}

// @Tags OrderItem
// @Summary 分页获取OrderItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取OrderItem列表page"
// @Param pageSize query string true "分页获取OrderItem列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderItem/getOrderItemList [get]
func GetOrderItemList(c *gin.Context) {
	var pageInfo request.OrderItemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetOrderItemInfoList(pageInfo)
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
