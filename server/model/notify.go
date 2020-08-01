package model

type SmsNotifyAppointmentOkInput struct {
	OutTradeNo     string `json:"out_trade_no" form:"out_trade_no" db:"out_trade_no" binding:"required"`
	PkgId          int64  `json:"pkg_id" form:"pkg_id" db:"pkg_id" binding:"required"`
	ExamineeName   string `json:"examinee_name" form:"examinee_name" db:"examinee_name" binding:"required"`
	ExamineeMobile string `json:"examinee_mobile" form:"examinee_mobile" db:"examinee_mobile" binding:"required"`
	ExamineDate    int64  `json:"examine_date" form:"examine_date" db:"examine_date" binding:"required"`
}

type HospitalInfo struct {
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}
