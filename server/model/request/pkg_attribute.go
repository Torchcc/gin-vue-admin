package request

import "gin-vue-admin/model"

type PkgAttrSearch struct {
	model.PkgAttr
	PageInfo
}
