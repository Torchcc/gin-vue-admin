package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/sms"
)

func SmsNotifyAppointmentOk(input *model.SmsNotifyAppointmentOkInput) (err error) {
	var hosInf model.HospitalInfo
	const cmd = `SELECT
					mh.name,
					mh.address
				FROM 
					mkh_hospital mh INNER JOIN mkp_package mp 
					ON mp.hospital_id = mh.id AND mh.is_deleted = 0
				WHERE
					mp.id = ?
`
	if err = global.BIZ_DBX.Get(&hosInf, cmd, input.PkgId); err != nil {
		global.GVA_LOG.Errorf("failed to notify appointment ok, input is %v, err: [%s]", input, err.Error())
		return
	}
	if err = sms.SendAppointmentOkMsg(input.ExamineeMobile, input.ExamineeName, input.OutTradeNo, hosInf.Name, hosInf.Address, input.ExamineDate); err != nil {
		global.GVA_LOG.Errorf("failed to notify appointment ok, input is %v, err: [%s]", input, err.Error())
		return
	}
	return
}

func SmsNotifyRefundOk(input *model.SmsNotifyRefundOkInput) (err error) {
	var pkgName string
	const cmd = `SELECT name FROM mkp_package WHERE id = ?`
	if err = global.BIZ_DBX.Get(&pkgName, cmd, input.PkgId); err != nil {
		global.GVA_LOG.Errorf("failed to notify refund ok, input is %v, err: [%s]", input, err.Error())
		return
	}
	if err = sms.SendRefundOkMsg(input.Mobile, input.OutTradeNo, pkgName, input.Amount, input.PkgCount); err != nil {
		global.GVA_LOG.Errorf("failed to notify refund ok, input is %v, err: [%s]", input, err.Error())
		return
	}
	return
}
