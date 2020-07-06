package request

import "gin-vue-admin/model"

type PackageSearch struct {
	model.Package
	PageInfo
}
