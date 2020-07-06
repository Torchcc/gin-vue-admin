package initialize

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"

	"os"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.GVA_CONFIG.Mysql
	biz := global.GVA_CONFIG.BizMysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GVA_DB.LogMode(admin.LogMode)
	}

	// 业务数据库
	if bizDb, err := gorm.Open("mysql", biz.Username+":"+biz.Password+"@("+biz.Path+")/"+biz.Dbname+"?"+biz.Config); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", err)
		os.Exit(0)

	} else {
		bizDb.Callback().Delete().Replace("gorm:delete", LogicDelete)
		bizDb.Callback().Create().Before("gorm:create").Register("update_create_n_update_time", BeforeCreate)
		bizDb.Callback().Update().Before("gorm:update").Register("update_update_time", BeforeUpdate)
		bizDb.Callback().Query().Before("gorm:query").Register("exclude_is_deleted_item", BeforeQuery)

		global.BIZ_DB = bizDb
		global.BIZ_DB.DB().SetMaxIdleConns(biz.MaxIdleConns)
		global.BIZ_DB.DB().SetMaxOpenConns(biz.MaxOpenConns)
		global.BIZ_DB.LogMode(biz.LogMode)
	}

	// mysqlx
	dsn := biz.Username + ":" + biz.Password + "@tcp(" + biz.Path + ")/" + biz.Dbname + "?charset=utf8mb4&autocommit=true"
	if db, err := sqlx.Connect("mysql", dsn); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.BIZ_DBX = db
		db.SetMaxOpenConns(biz.MaxOpenConns)
		db.SetMaxIdleConns(biz.MaxIdleConns)
	}
}
