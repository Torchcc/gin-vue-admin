package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreatePkgCategory
// @description   create a PkgCategory
// @param     pkgCtg               model.PkgCategory
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreatePkgCategory(pkgCtg model.PkgCategory) (err error) {
	err = global.BIZ_DB.Create(&pkgCtg).Error
	return err
}

// @title    DeletePkgCategory
// @description   delete a PkgCategory
// @auth                     （2020/04/05  20:22）
// @param     pkgCtg               model.PkgCategory
// @return                    error

func DeletePkgCategory(pkgCtg model.PkgCategory) (err error) {
	err = global.BIZ_DB.Delete(pkgCtg).Error
	return err
}

// @title    UpdatePkgCategory
// @description   update a PkgCategory
// @param     pkgCtg          *model.PkgCategory
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdatePkgCategory(pkgCtg *model.PkgCategory) (err error) {
	err = global.BIZ_DB.Save(pkgCtg).Error
	return err
}

// @title    GetPkgCategory
// @description   get the info of a PkgCategory
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    PkgCategory        PkgCategory

func GetPkgCategory(id uint) (err error, pkgCtg model.PkgCategory) {
	err = global.BIZ_DB.Where("id = ?", id).First(&pkgCtg).Error
	return
}

// @title    GetPkgCategoryInfoList
// @description   get PkgCategory list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetPkgCategoryInfoList(info request.PkgCategorySearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.PkgCategory{}).Where("is_deleted = ?", 0)
	var pkgCtgs []model.PkgCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if limit == 0 && offset == 0 {
		err = db.Find(&pkgCtgs).Error
	} else {
		err = db.Limit(limit).Offset(offset).Find(&pkgCtgs).Error
	}
	return err, pkgCtgs, total
}
