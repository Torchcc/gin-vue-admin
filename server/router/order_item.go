package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOrderItemRouter(Router *gin.RouterGroup) {
	OrderItemRouter := Router.Group("orderItem").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		OrderItemRouter.POST("createOrderItem", v1.CreateOrderItem)   // 新建OrderItem
		OrderItemRouter.DELETE("deleteOrderItem", v1.DeleteOrderItem) // 删除OrderItem
		OrderItemRouter.PUT("updateOrderItem", v1.UpdateOrderItem)    // 更新OrderItem
		OrderItemRouter.GET("findOrderItem", v1.FindOrderItem)        // 根据ID获取OrderItem
		OrderItemRouter.GET("getOrderItemList", v1.GetOrderItemList)  // 获取OrderItem列表
	}
}
