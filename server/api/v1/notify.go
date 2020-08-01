package v1

import (
	"fmt"

	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

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
