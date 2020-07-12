// 自动生成模板PkgCategory
package model

// 如果含有time.Time 请自行import time包
type PkgCategory struct {
	ID         uint   `json:"id" form:"id" db:"id" gorm:"column:id;comment:'自增id';primary_key"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:'套餐类别名称'"`
	CreateTime int64  `json:"create_time" form:"create_time" gorm:"column:create_time;comment:'创建时间'"`
	UpdateTime int64  `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间'"`
	IsDeleted  int8   `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:'是否删除'"`
}

func (PkgCategory) TableName() string {
	return "mkp_category"
}
