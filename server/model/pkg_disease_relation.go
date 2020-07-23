package model

type PkgDiseaseRelation struct {
	ID        uint  `json:"id" form:"id" db:"id" gorm:"column:id;comment:'id';primary_key"`
	PkgId     uint  `json:"pkg_id" db:"pkg_id" gorm:"column:pkg_id"`
	DiseaseId uint  `json:"disease_id" db:"disease_id" gorm:"column:disease_id"`
	IsDeleted uint8 `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:'是否删除'"`
}

func (PkgDiseaseRelation) TableName() string {
	return "mkp_package_disease"
}
