package service

import (
	"fmt"
	"time"

	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/deckarep/golang-set"
	"github.com/jmoiron/sqlx"
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
	var pkgs []*model.PkgWithCategories
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
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pkgs).Error

	pkgDic := make(map[uint]*model.PkgWithCategories)

	pkgIds := make([]uint, len(pkgs))
	for idx, p := range pkgs {
		p.CtgIds = make([]uint, 0, 4)

		pkgIds[idx] = p.ID
		pkgDic[p.ID] = p
	}

	var allCtgs []*model.PkgCtgRelation
	db = global.BIZ_DB.Model(&model.PkgCtgRelation{})
	err = db.Where("pkg_id IN (?)", pkgIds).Where("is_deleted = ?", 0).Find(&allCtgs).Error

	for _, c := range allCtgs {
		pkgDic[c.PkgId].CtgIds = append(pkgDic[c.PkgId].CtgIds, c.CtgId)
	}

	return err, pkgs, total
}

func UploadPkgAvatar(id uint, filePath string) (err error) {
	var pkg model.Package
	err = global.BIZ_DB.Where("id = ?", id).First(&pkg).Update("avatar_url", filePath).Error
	return err
}

func UpdatePkgCtgRelation(pkgCtg *model.PkgWithCategories) error {
	oldCtgIds := make([]uint, 0, 4)
	db := global.BIZ_DB.Model(&model.PkgCtgRelation{})
	db.Where("pkg_id = ?", pkgCtg.ID).Where("is_deleted = ?", 0).Pluck("category_id", &oldCtgIds)
	oldCtgIdsSet := mapset.NewSet()
	for _, id := range oldCtgIds {
		oldCtgIdsSet.Add(id)
	}

	newCtgIdsSet := mapset.NewSet()
	for _, id := range pkgCtg.CtgIds {
		newCtgIdsSet.Add(id)
	}

	toCreateSet := newCtgIdsSet.Difference(oldCtgIdsSet)
	toDeleteSet := oldCtgIdsSet.Difference(newCtgIdsSet)

	toCreateIds := make([]uint, 0, 4)
	for item := range toCreateSet.Iter() {
		toCreateIds = append(toCreateIds, item.(uint))
	}

	toDeleteIds := make([]uint, 0, 4)
	for item := range toDeleteSet.Iter() {
		toDeleteIds = append(toDeleteIds, item.(uint))
	}

	if len(toDeleteIds) != 0 {
		cmd, args, err := sqlx.In(
			fmt.Sprintf(`UPDATE mkp_package_category SET is_deleted = 1, update_time = UNIX_TIMESTAMP(now()) WHERE pkg_id = %d AND category_id IN (?)`, pkgCtg.ID),
			toDeleteIds)
		if err != nil {
			global.GVA_LOG.Error(err)
			return err
		}
		cmd = global.BIZ_DBX.Rebind(cmd)
		_, err = global.BIZ_DBX.Exec(cmd, args...)
		if err != nil {
			global.GVA_LOG.Error(err)
			return err
		}
	}

	if len(toCreateIds) != 0 {
		cmd := "INSERT INTO mkp_package_category (pkg_id, category_id, create_time, update_time) VALUES "
		for _, id := range toCreateIds {
			cmd += fmt.Sprintf("(%d, %d, %d, %d),", pkgCtg.ID, id, time.Now().Unix(), time.Now().Unix())
		}
		cmd = cmd[:len(cmd)-1]
		_, err := global.BIZ_DBX.Exec(cmd)
		if err != nil {
			global.GVA_LOG.Error(err)
			return err
		}
	}
	return nil
}

func GetPkgAttrList(pkgId uint) ([]*model.PkgAttr, error) {
	db := global.BIZ_DB.Model(&model.PkgAttr{})
	var pkgAttrs []*model.PkgAttr
	err := db.Where("pkg_id = ?", pkgId).Order("attr_type, order_no").Find(&pkgAttrs).Error
	return pkgAttrs, err
}
