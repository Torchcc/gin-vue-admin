package v1

import (
	"fmt"

	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

const Refund int8 = 3

func SmsNotifyAppointmentOk(c *gin.Context) {
	var input model.SmsNotifyAppointmentOkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithMessage(fmt.Sprintf("发送失败，%v", err), c)
		return
	}
	err := service.SmsNotifyAppointmentOk(&input)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("发送失败，%v", err), c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}

func SmsNotifyRefundOk(c *gin.Context) {
	var input model.SmsNotifyRefundOkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithMessage(fmt.Sprintf("发送退款短信绑定参数失败失败，%v", err), c)
		return
	}
	err := service.SmsNotifyRefundOk(&input)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("发送退款短信失败，%v", err), c)
		return
	}

	if err := service.UpdateOrderStatus(input.OrderId, Refund); err != nil {
		response.FailWithMessage(fmt.Sprintf("更改订单状态失败，请手动将该订单设置为已退款，%v", err), c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}
