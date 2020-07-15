// 自动生成模板OrderItem
package model

// 如果含有time.Time 请自行import time包
type OrderItem struct {
	ID             uint    `json:"id" form:"id" gorm:"column:id;comment:'id'"`
	UserId         uint    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:'下单用户的id'"`
	OrderId        uint    `json:"order_id" form:"order_id" gorm:"column:order_id;comment:'订单id'"`
	PkgId          uint    `json:"pkg_id" form:"pkg_id" gorm:"column:pkg_id;comment:'套餐id'"`
	PkgPrice       float64 `json:"pkg_price" form:"pkg_price" gorm:"column:pkg_price;comment:'下单时套餐的单价'"`
	ExamineeName   string  `json:"examinee_name" form:"examinee_name" gorm:"column:examinee_name;comment:'体检人姓名'"`
	ExamineeMobile string  `json:"examinee_mobile" form:"examinee_mobile" gorm:"column:examinee_mobile;comment:'体检人电话'"`
	IdCardNo       string  `json:"id_card_no" form:"id_card_no" gorm:"column:id_card_no;comment:'体检人身份证号码'"`
	IsMarried      int     `json:"is_married" form:"is_married" gorm:"column:is_married;comment:'体检人是否已婚'"`
	Gender         int     `json:"gender" form:"gender" gorm:"column:gender;comment:'性别 1-男, 2-女'"`
	ExamineDate    int     `json:"examine_date" form:"examine_date" gorm:"column:examine_date;comment:'体检日期,精确到日'"`
	CreateTime     int     `json:"create_time" form:"create_time" gorm:"column:create_time;comment:'创建时间'"`
	UpdateTime     int     `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间'"`
}

func (OrderItem) TableName() string {
	return "mko_order_item"
}
