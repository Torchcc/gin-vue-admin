package service

import (
	"fmt"

	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHospital
// @description   create a Hospital
// @param     hospital               model.Hospital
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHospital(hospital model.Hospital) (err error) {
	err = global.BIZ_DB.Create(&hospital).Error
	return err
}

// @title    DeleteHospital
// @description   delete a Hospital
// @auth                     （2020/04/05  20:22）
// @param     hospital               model.Hospital
// @return                    error

func DeleteHospital(hospital model.Hospital) (err error) {
	err = global.BIZ_DB.Delete(hospital).Error
	return err
}

// @title    DeleteHospitalByIds
// @description   delete Hospitals
// @auth                     （2020/04/05  20:22）
// @param     hospital               model.Hospital
// @return                    error

func DeleteHospitalByIds(ids request.IdsReq) (err error) {
	db := global.BIZ_DB.Model(&model.Hospital{})
	cmd := fmt.Sprintf("UPDATE %s SET is_deleted = 1 WHERE id IN (?)", model.Hospital{}.TableName())
	err = db.Exec(cmd, ids.Ids).Error
	return err
}

// @title    UpdateHospital
// @description   update a Hospital
// @param     hospital          *model.Hospital
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHospital(hospital *model.Hospital) (err error) {
	err = global.BIZ_DB.Save(hospital).Error
	return err
}

// @title    GetHospital
// @description   get the info of a Hospital
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Hospital        Hospital

func GetHospital(id uint) (err error, hospital model.Hospital) {
	err = global.BIZ_DB.Where("id = ?", id).First(&hospital).Error
	return
}

// @title    GetHospitalInfoList
// @description   get Hospital list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHospitalInfoList(info request.HospitalSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.Hospital{}).Where("is_deleted = ?", 0)
	var hospitals []model.ListHospitalOutputEle
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Level != 0 {
		db = db.Where("level = ?", info.Level)
	}
	if info.ProvinceId != 0 {
		db = db.Where("province_id = ?", info.ProvinceId)
	}
	if info.CityId != 0 {
		db = db.Where("city_id = ?", info.CityId)
	}
	if info.CountyId != 0 {
		db = db.Where("county_id = ?", info.CountyId)
	}
	if info.TownId != 0 {
		db = db.Where("town_id = ?", info.TownId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hospitals).Error
	serializeRegionName(hospitals)
	serializeLevelName(hospitals)
	return err, hospitals, total
}

func serializeLevelName(hospitals []model.ListHospitalOutputEle) {
	levelNameMap := map[int]string{
		0: "无级别",
		1: "公立三甲",
		2: "公立医院",
		3: "民营医院",
		4: "专业机构",
	}
	for idx, h := range hospitals {
		hospitals[idx].LevelName = levelNameMap[h.Level]
	}
}

func serializeRegionName(hospitals []model.ListHospitalOutputEle) {
	idNameMap, _ := model.GetRegionIdNameMap()

	for idx, h := range hospitals {
		hospitals[idx].ProvinceName = idNameMap[int64(h.ProvinceId)]
		hospitals[idx].CityName = idNameMap[int64(h.CityId)]
		hospitals[idx].CountyName = idNameMap[int64(h.CountyId)]
		hospitals[idx].TownName = idNameMap[int64(h.TownId)]
	}
}

func UploadHospitalAvatar(id uint, filePath string) (err error) {
	var hospital model.Hospital
	err = global.BIZ_DB.Where("id = ?", id).First(&hospital).Update("avatar_url", filePath).Error
	return err
}
