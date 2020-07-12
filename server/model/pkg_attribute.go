package model

type PkgAttr struct {
	ID         uint   `json:"id" form:"id" db:"id" gorm:"column:id;comment:'属性id';primary_key"`
	PkgId      uint   `json:"pkg_id" form:"pkg_id" gorm:"column:pkg_id;comment:'套餐id'"`
	AttrType   uint8  `json:"attr_type" form:"attr_type" gorm:"column:attr_type;comment:'属性类型'"`
	OrderNo    uint8  `json:"order_no" form:"order_no" gorm:"column:order_no;comment:'属性类型'"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:'属性名称'"`
	Brief      string `json:"brief" form:"brief" gorm:"column:brief;comment:'简介'"`
	Comment    string `json:"comment" form:"comment" gorm:"column:comment;comment:'详细描述'"`
	CreateTime int64  `json:"create_time" form:"create_time" gorm:"column:create_time;type:int;comment:''"`
	UpdateTime int64  `json:"update_time" form:"update_time" gorm:"column:update_time;comment:''"`
	IsDeleted  int    `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:''"`
}

func (PkgAttr) TableName() string {
	return "mkp_package_attribute"
}
