// 自动生成模板Package
package model

import "gin-vue-admin/global"

// 如果含有time.Time 请自行import time包
type Package struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	HospitalId    int     `json:"hospital_id" form:"hospital_id" gorm:"column:hospital_id;comment:'医院id'"`
	Name          string  `json:"name" form:"name" gorm:"column:name;comment:'套餐名字'"`
	Target        int     `json:"target" form:"target" gorm:"column:target;comment:'套餐目标人群'"`
	Disease       int     `json:"disease" form:"disease" gorm:"column:disease;comment:''"`
	PriceOriginal float64 `json:"price_original" form:"price_original" gorm:"column:price_original;comment:'门市价'"`
	PriceReal     float64 `json:"price_real" form:"price_real" gorm:"column:price_real;comment:''"`
	AvatarUrl     string  `json:"avatar_url" form:"avatar_url" gorm:"column:avatar_url;comment:''"`
	Brief         string  `json:"brief" form:"brief" gorm:"column:brief;comment:''"`
	Comment       string  `json:"comment" form:"comment" gorm:"column:comment;comment:'套餐备注'"`
	Tips          string  `json:"tips" form:"tips" gorm:"column:tips;comment:'温馨提示'"`
	CreateTime    int64   `json:"create_time" form:"create_time" gorm:"column:create_time;type:int;comment:''"`
	UpdateTime    int64   `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'';type:int"`
	IsDeleted     int     `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:''"`
}

func (Package) TableName() string {
	return "mkp_package"
}

func (p *Package) Delete() error {
	const cmd = "UPDATE mkp_package SET is_deleted = 1 WHERE id = ?"
	_, err := global.BIZ_DBX.Exec(cmd, p.ID)
	return err
}
