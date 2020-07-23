// 自动生成模板Disease
package model

// 如果含有time.Time 请自行import time包
type Disease struct {
	ID         uint   `json:"id" form:"id" db:"id" gorm:"column:id;comment:'自增id';primary_key"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:'高发疾病名称';type:varchar(128)"`
	CreateTime int64  `json:"create_time" form:"create_time" gorm:"column:create_time;comment:'创建时间';type:int(11)"`
	UpdateTime int64  `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间';type:int(11)"`
	IsDeleted  int8   `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:'是否删除';type:tinyint"`
}

func (Disease) TableName() string {
	return "mkp_disease"
}
