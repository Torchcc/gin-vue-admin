package request

import "gin-vue-admin/model"

type PkgCategorySearch struct {
	model.PkgCategory
	PageInfo
}
