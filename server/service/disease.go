package service

import (
	"fmt"

	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateDisease
// @description   create a Disease
// @param     disease               model.Disease
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDisease(disease model.Disease) (err error) {
	err = global.BIZ_DB.Create(&disease).Error
	return err
}

// @title    DeleteDisease
// @description   delete a Disease
// @auth                     （2020/04/05  20:22）
// @param     disease               model.Disease
// @return                    error

func DeleteDisease(disease model.Disease) (err error) {
	err = global.BIZ_DB.Delete(disease).Error
	return err
}

// @title    DeleteDiseaseByIds
// @description   delete Diseases
// @auth                     （2020/04/05  20:22）
// @param     disease               model.Disease
// @return                    error

func DeleteDiseaseByIds(ids request.IdsReq) (err error) {
	db := global.BIZ_DB.Model(&model.Disease{})
	cmd := fmt.Sprintf("UPDATE %s SET is_deleted = 1 WHERE id IN (?)", model.Disease{}.TableName())
	err = db.Exec(cmd, ids.Ids).Error
	return err
}

// @title    UpdateDisease
// @description   update a Disease
// @param     disease          *model.Disease
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDisease(disease *model.Disease) (err error) {
	err = global.BIZ_DB.Save(disease).Error
	return err
}

// @title    GetDisease
// @description   get the info of a Disease
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Disease        Disease

func GetDisease(id uint) (err error, disease model.Disease) {
	err = global.BIZ_DB.Where("id = ?", id).First(&disease).Error
	return
}

// @title    GetDiseaseInfoList
// @description   get Disease list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDiseaseInfoList(info request.DiseaseSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.Disease{}).Where("is_deleted = ?", 0)
	var diseases []model.Disease
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&diseases).Error
	return err, diseases, total
}
