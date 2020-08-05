package core

import (
	"fmt"

	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/superconf"
)

func init() {
	global.GVA_CONFIG = config.Server{}
	var allConfigs = make(map[string]interface{})
	allConfigs["/superconf/admin/mysql/admin_db"] = &global.GVA_CONFIG.Mysql
	allConfigs["/superconf/admin/mysql/biz_db"] = &global.GVA_CONFIG.BizMysql
	allConfigs["/superconf/third_party/qiniu"] = &global.GVA_CONFIG.Qiniu
	allConfigs["/superconf/admin/casbin"] = &global.GVA_CONFIG.Casbin
	allConfigs["/superconf/admin/redis/admin_redis"] = &global.GVA_CONFIG.Redis
	allConfigs["/superconf/admin/system"] = &global.GVA_CONFIG.System
	allConfigs["/superconf/admin/jwt"] = &global.GVA_CONFIG.JWT
	allConfigs["/superconf/admin/captcha"] = &global.GVA_CONFIG.Captcha
	allConfigs["/superconf/admin/log"] = &global.GVA_CONFIG.Log
	allConfigs["/superconf/third_party/cos"] = &global.GVA_CONFIG.Cos
	allConfigs["/superconf/third_party/sms/apmt_ok_msg_tmpl"] = &global.GVA_CONFIG.ApmtOkMsgTmpl
	allConfigs["/superconf/third_party/sms/refund_ok_msg_tmpl"] = &global.GVA_CONFIG.RefundOkMsgTmpl
	superconf.NewSuperConfig(&allConfigs)
	fmt.Printf("the conf for admin sys is: [%#v]", global.GVA_CONFIG)
}
