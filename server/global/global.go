package global

import (
	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	BIZ_DB     *gorm.DB
	BIZ_DBX    *sqlx.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *oplogging.Logger
)
