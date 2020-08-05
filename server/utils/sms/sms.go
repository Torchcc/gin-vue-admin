package sms

import (
	"encoding/json"
	"time"

	"gin-vue-admin/global"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

const (
	ENDPOINT = "sms.tencentcloudapi.com"
	SMS_SIGN = "迈康体检网"
)

type smsParam struct {
	// 格式 ["+8618520456660", ..., ]
	PhoneNumberSet   []string
	TemplateID       string   // zk 读取
	Sign             string   // 固定为"迈康体检网"
	TemplateParamSet []string // 模板参数
	SmsSdkAppid      string   // zk读取

}

// 发送预约成功的通知短信
func SendAppointmentOkMsg(mobile, examineeName, outTradeNo, hospitalName, Addr string, examDate int64) (err error) {
	examDateStr := time.Unix(examDate, 0).Format("2006年01月02日 ") + "07:30至12：00"

	credential := common.NewCredential(
		global.GVA_CONFIG.Cos.SecretID,
		global.GVA_CONFIG.Cos.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = ENDPOINT
	client, _ := sms.NewClient(credential, global.GVA_CONFIG.Cos.Region, cpf)

	request := sms.NewSendSmsRequest()

	payload := smsParam{
		PhoneNumberSet:   []string{"+86" + mobile},
		TemplateID:       global.GVA_CONFIG.ApmtOkMsgTmpl.TemplateID,
		Sign:             SMS_SIGN,
		TemplateParamSet: []string{examineeName, outTradeNo, examDateStr, hospitalName, Addr},
		SmsSdkAppid:      global.GVA_CONFIG.ApmtOkMsgTmpl.SmsSdkAppid,
	}

	params, _ := json.Marshal(payload)
	err = request.FromJsonString(string(params))
	if err != nil {
		return
	}
	_, err = client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return
	}
	if err != nil {
		return
	}
	return
}
