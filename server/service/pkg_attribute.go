package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreatePkgAttr
// @description   create a PkgAttr
// @param     pkgAttr               model.PkgAttr
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreatePkgAttr(pkgAttr model.PkgAttr) (err error) {
	err = global.BIZ_DB.Create(&pkgAttr).Error
	return err
}

// @title    DeletePkgAttr
// @description   delete a PkgAttr
// @auth                     （2020/04/05  20:22）
// @param     pkgAttr               model.PkgAttr
// @return                    error

func DeletePkgAttr(pkgAttr model.PkgAttr) (err error) {
	err = global.BIZ_DB.Delete(pkgAttr).Error
	return err
}

// @title    UpdatePkgAttr
// @description   update a PkgAttr
// @param     pkgAttr          *model.PkgAttr
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdatePkgAttr(pkgAttr *model.PkgAttr) (err error) {
	err = global.BIZ_DB.Save(pkgAttr).Error
	return err
}

// @title    GetPkgAttr
// @description   get the info of a PkgAttr
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    PkgAttr        PkgAttr

func GetPkgAttr(id uint) (err error, pkgAttr model.PkgAttr) {
	err = global.BIZ_DB.Where("id = ?", id).First(&pkgAttr).Error
	return
}

// @title    GetPkgAttrInfoList
// @description   get PkgAttr list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetPkgAttrInfoList(info request.PkgAttrSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.PkgAttr{}).Where("is_deleted = ?", 0)
	var pkgAttrs []model.PkgAttr
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.PkgId != 0 {
		db = db.Where("pkg_id = ?", info.PkgId)
	}
	if info.AttrType != 0 {
		db = db.Where("attr_type = ?", info.AttrType)
	}
	if info.OrderNo != 0 {
		db = db.Where("order_no < ?", info.OrderNo)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	err = db.Order("attr_type, order_no").Limit(limit).Offset(offset).Find(&pkgAttrs).Error
	return err, pkgAttrs, total
}

// func GetPkgAttrList(pkgId uint) ([]*model.PkgAttr, error) {
// 	db := global.BIZ_DB.Model(&model.PkgAttr{})
// 	var pkgAttrs []*model.PkgAttr
// 	err := db.Where("pkg_id = ?", pkgId).Order("attr_type, order_no").Find(&pkgAttrs).Error
// 	return pkgAttrs, err
// }
