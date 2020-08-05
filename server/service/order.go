package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateOrder
// @description   create a Order
// @param     order               model.Order
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateOrder(order model.Order) (err error) {
	err = global.BIZ_DB.Create(&order).Error
	return err
}

// @title    DeleteOrder
// @description   delete a Order
// @auth                     （2020/04/05  20:22）
// @param     order               model.Order
// @return                    error

func DeleteOrder(order model.Order) (err error) {
	err = global.BIZ_DB.Delete(order).Error
	return err
}

// @title    UpdateOrder
// @description   update a Order
// @param     order          *model.Order
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateOrder(order *model.Order) (err error) {
	err = global.BIZ_DB.Save(order).Error
	return err
}

func UpdateOrderStatus(orderId int64, status int8) (err error) {
	const cmd = `UPDATE mko_order SET status = ? WHERE id = ? AND is_deleted = 0`
	_, err = global.BIZ_DBX.Exec(cmd, status, orderId)
	return
}

// @title    GetOrder
// @description   get the info of a Order
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Order        Order

func GetOrder(id uint) (err error, order model.Order) {
	err = global.BIZ_DB.Where("id = ?", id).First(&order).Error
	return
}

// @title    GetOrderInfoList
// @description   get Order list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetOrderInfoList(info request.OrderSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.BIZ_DB.Model(&model.Order{}).Where("is_deleted = ?", 0)
	var orders []model.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.OutTradeNo != "" {
		db = db.Where("out_trade_no = ?", info.OutTradeNo)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+info.Mobile+"%")
	}
	if info.OpenId != "" {
		db = db.Where("open_id = ?", info.OpenId)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.IsRefundApplied {
		db = db.Where("refund_reason_id != ?", 0)
	}
	if info.HandleStatusFilter != nil {
		db = db.Where("handle_status = ?", *info.HandleStatusFilter)
	}
	err = db.Count(&total).Error
	err = db.Order("create_time DESC").Limit(limit).Offset(offset).Find(&orders).Error
	// 分->元
	for idx := range orders {
		orders[idx].Amount *= 0.01
	}
	return err, orders, total
}
