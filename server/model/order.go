// 自动生成模板Order
package model

// 如果含有time.Time 请自行import time包
type Order struct {
	ID                 uint    `json:"id" form:"id" gorm:"column:id;comment:'主键id'"`
	OutTradeNo         string  `json:"out_trade_no" form:"out_trade_no" gorm:"column:out_trade_no;comment:'订单号'"`
	UserId             uint    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:'下单的用户id'"`
	Mobile             string  `json:"mobile" form:"mobile" gorm:"column:mobile;comment:'下单用户的电话'"`
	OpenId             string  `json:"open_id" form:"open_id" gorm:"column:open_id;comment:'微信open_id'"`
	Amount             float64 `json:"amount" form:"amount" gorm:"column:amount;comment:'交易总金额'"`
	RefundTime         int     `json:"refund_time" form:"refund_time" gorm:"column:refund_time;comment:'退款时间'"`
	Status             int     `json:"status" form:"status" gorm:"column:status;comment:'支付状态: 0待支付 1处理中 2付款成功，待预约 3已退款 4已过期/关闭 5 体检完毕，待评价 6 评价完毕'"`
	Remark             string  `json:"remark" form:"remark" gorm:"column:remark;comment:'订单备注'"`
	HandleStatus       int8    `json:"handle_status" form:"handle_status" db:"handle_status" gorm:"column:handle_status;comment:'处理状态:0-待处理,1-跟进中 2-跟进完毕'"`
	HandleRemark       string  `json:"handle_remark" form:"handle_remark" db:"handle_remark" gorm:"column:handle_remark"`
	RefundReasonId     uint    `json:"refund_reason_id" form:"refund_reason_id" db:"refund_reason_id" gorm:"refund_reason_id"`
	RefundReasonRemark string  `json:"refund_reason_remark" form:"refund_reason_remark" db:"refund_reason_remark" gorm:"refund_reason_remark"`
	CreateTime         int     `json:"create_time" form:"create_time" gorm:"column:create_time;comment:'创建时间'"`
	UpdateTime         int     `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间'"`
}

func (Order) TableName() string {
	return "mko_order"
}
