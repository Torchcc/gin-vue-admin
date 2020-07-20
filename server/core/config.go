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
	allConfigs["/superconf/union/mysql/admin"] = &global.GVA_CONFIG.Mysql
	allConfigs["/superconf/union/mysql/biz"] = &global.GVA_CONFIG.BizMysql
	allConfigs["/superconf/third_party/qiniu"] = &global.GVA_CONFIG.Qiniu
	allConfigs["/superconf/admin/casbin"] = &global.GVA_CONFIG.Casbin
	allConfigs["/superconf/admin/redis"] = &global.GVA_CONFIG.Redis
	allConfigs["/superconf/admin/system"] = &global.GVA_CONFIG.System
	allConfigs["/superconf/admin/jwt"] = &global.GVA_CONFIG.JWT
	allConfigs["/superconf/admin/captcha"] = &global.GVA_CONFIG.Captcha
	allConfigs["/superconf/admin/log"] = &global.GVA_CONFIG.Log
	superconf.NewSuperConfig(&allConfigs)
	fmt.Printf("the conf for admin sys is: [%#v]", global.GVA_CONFIG)
}
