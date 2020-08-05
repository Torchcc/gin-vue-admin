package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateOrderItem
// @description   create a OrderItem
// @param     orderItem               model.OrderItem
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateOrderItem(orderItem model.OrderItem) (err error) {
	err = global.BIZ_DB.Create(&orderItem).Error
	return err
}

// @title    DeleteOrderItem
// @description   delete a OrderItem
// @auth                     （2020/04/05  20:22）
// @param     orderItem               model.OrderItem
// @return                    error

func DeleteOrderItem(orderItem model.OrderItem) (err error) {
	err = global.BIZ_DB.Delete(orderItem).Error
	return err
}

// @title    UpdateOrderItem
// @description   update a OrderItem
// @param     orderItem          *model.OrderItem
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateOrderItem(orderItem *model.OrderItem) (err error) {
	err = global.BIZ_DB.Save(orderItem).Error
	return err
}

// @title    GetOrderItem
// @description   get the info of a OrderItem
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    OrderItem        OrderItem

func GetOrderItem(id uint) (err error, orderItem model.OrderItem) {
	err = global.BIZ_DB.Where("id = ?", id).First(&orderItem).Error
	return
}

// @title    GetOrderItemInfoList
// @description   get OrderItem list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetOrderItemInfoList(info request.OrderItemSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.OrderItem{}).Where("is_deleted = ?", 0)
	var orderItems []model.OrderItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.OrderId != 0 {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.PkgId != 0 {
		db = db.Where("pkg_id = ?", info.PkgId)
	}
	if info.ExamineeName != "" {
		db = db.Where("examinee_name LIKE ?", "%"+info.ExamineeName+"%")
	}
	if info.ExamineeMobile != "" {
		db = db.Where("examinee_mobile = ?", info.ExamineeMobile)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&orderItems).Error
	return err, orderItems, total
}
