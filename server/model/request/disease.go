package request

import "gin-vue-admin/model"

type DiseaseSearch struct {
	model.Disease
	PageInfo
}
