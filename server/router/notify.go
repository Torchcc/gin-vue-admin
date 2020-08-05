package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitNotifyRouter(Router *gin.RouterGroup) {
	PackageRouter := Router.Group("notify").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		PackageRouter.POST("/sms/appointment_ok/", v1.SmsNotifyAppointmentOk) // 短信通知体检人预约成功
		PackageRouter.POST("/sms/refund_ok/", v1.SmsNotifyRefundOk)           // 短信通知体检人预约成功
	}
}
