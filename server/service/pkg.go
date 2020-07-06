package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreatePackage
// @description   create a Package
// @param     pkg               model.Package
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreatePackage(pkg model.Package) (err error) {
	err = global.BIZ_DB.Create(&pkg).Error
	return err
}

// @title    DeletePackage
// @description   delete a Package
// @auth                     （2020/04/05  20:22）
// @param     pkg               model.Package
// @return                    error

func DeletePackage(pkg *model.Package) (err error) {
	err = global.BIZ_DB.Delete(pkg).Error
	return err
}

// @title    UpdatePackage
// @description   update a Package
// @param     pkg          *model.Package
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdatePackage(pkg *model.Package) (err error) {
	err = global.BIZ_DB.Save(pkg).Error
	return err
}

// @title    GetPackage
// @description   get the info of a Package
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Package        Package

func GetPackage(id uint) (err error, pkg model.Package) {
	err = global.BIZ_DB.Where("id = ?", id).First(&pkg).Error
	return
}

// @title    GetPackageInfoList
// @description   get Package list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetPackageInfoList(info request.PackageSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.Package{})
	var pkgs []model.Package
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.HospitalId != 0 {
		db = db.Where("hospital_id = ?", info.HospitalId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Target != 0 {
		db = db.Where("target = ?", info.Target)
	}
	if info.Disease != 0 {
		db = db.Where("disease = ?", info.Disease)
	}
	if info.PriceOriginal != 0 {
		db = db.Where("price_original <> ?", info.PriceOriginal)
	}
	if info.PriceReal != 0 {
		db = db.Where("price_real <> ?", info.PriceReal)
	}
	if info.AvatarUrl != "" {
		db = db.Where("avatar_url = ?", info.AvatarUrl)
	}
	if info.CreateTime != 0 {
		db = db.Where("create_time <> ?", info.CreateTime)
	}
	if info.UpdateTime != 0 {
		db = db.Where("update_time <> ?", info.UpdateTime)
	}
	if info.IsDeleted != 0 {
		db = db.Where("is_deleted = ?", info.IsDeleted)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pkgs).Error
	return err, pkgs, total
}
