import service from '@/utils/request'

// @Tags OrderItem
// @Summary 创建OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "创建OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderItem/createOrderItem [post]
export const createOrderItem = (data) => {
     return service({
         url: "/orderItem/createOrderItem",
         method: 'post',
         data
     })
 }


// @Tags OrderItem
// @Summary 删除OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "删除OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderItem/deleteOrderItem [delete]
 export const deleteOrderItem = (data) => {
     return service({
         url: "/orderItem/deleteOrderItem",
         method: 'delete',
         data
     })
 }

// @Tags OrderItem
// @Summary 更新OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "更新OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderItem/updateOrderItem [put]
 export const updateOrderItem = (data) => {
     return service({
         url: "/orderItem/updateOrderItem",
         method: 'put',
         data
     })
 }


// @Tags OrderItem
// @Summary 用id查询OrderItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "用id查询OrderItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderItem/findOrderItem [get]
 export const findOrderItem = (params) => {
     return service({
         url: "/orderItem/findOrderItem",
         method: 'get',
         params
     })
 }


// @Tags OrderItem
// @Summary 分页获取OrderItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取OrderItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderItem/getOrderItemList [get]
 export const getOrderItemList = (params) => {
     return service({
         url: "/orderItem/getOrderItemList",
         method: 'get',
         params
     })
 }