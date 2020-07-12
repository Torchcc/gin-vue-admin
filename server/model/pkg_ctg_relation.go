package model

type PkgCtgRelation struct {
	ID        uint  `json:"id" form:"id" db:"id" gorm:"column:id;comment:'id';primary_key"`
	PkgId     uint  `json:"pkg_id" db:"pkg_id" gorm:"column:pkg_id"`
	CtgId     uint  `json:"category_id" db:"category_id" gorm:"column:category_id"`
	IsDeleted uint8 `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:'是否删除'"`
}

func (PkgCtgRelation) TableName() string {
	return "mkp_package_category"
}
