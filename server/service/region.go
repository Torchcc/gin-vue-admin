package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func RetrieveRegionsByParentId(parentId int64) (output []*model.Region, err error) {
	output, err = model.FindRegionsByParentId(parentId)
	if err != nil {
		global.GVA_LOG.Errorf("list region failed, err: [%s]", err.Error())
	}
	return
}
