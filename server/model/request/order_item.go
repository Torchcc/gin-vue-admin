package request

import "gin-vue-admin/model"

type OrderItemSearch struct {
	model.OrderItem
	PageInfo
}
