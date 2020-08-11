package request

import "gin-vue-admin/model"

type HospitalSearch struct {
	model.Hospital
	PageInfo
}
