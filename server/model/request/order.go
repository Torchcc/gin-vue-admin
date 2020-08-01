package request

import "gin-vue-admin/model"

type OrderSearch struct {
	model.Order
	PageInfo
	IsRefundApplied    bool   `json:"is_refund_applied" form:"is_refund_applied"`
	HandleStatusFilter *int64 `json:"handle_status_filter" form:"handle_status_filter"`
}
