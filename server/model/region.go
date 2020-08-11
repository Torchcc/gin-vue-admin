package model

import (
	"strconv"
	"sync"

	"gin-vue-admin/global"
	"github.com/patrickmn/go-cache"
)

type Region struct {
	// 省，市， 区/县， 镇/街道 的行政区域id
	Id int64 `json:"id" db:"id"`
	// 行政区域名称
	Name string `json:"name" db:"name"`
	// 该行政区域的父级id， 1级行政区域的父id 是 0
	ParentId int64 `json:"parent_id" db:"parent_id"`
	// 行政区域等级, 1-省，2-市， 3-区/县， 4-镇/街道
	Level int8 `json:"level" db:"level"`
}

type RegionIdName struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func FindRegionsByParentId(parentId int64) (output []*Region, err error) {
	// 增加内存缓存
	key := "region" + strconv.FormatInt(parentId, 10)
	if x, found := global.GoCache.Get(key); found {
		output = x.([]*Region)
		return
	}

	const cmd = `SELECT id, name, parent_id, level FROM mkm_region WHERE parent_id = ? AND is_deleted = 0`
	err = global.BIZ_DBX.Select(&output, cmd, parentId)

	if err != nil {
		global.GVA_LOG.Errorf("failed to save data to memory, err: %s", err.Error())
	} else {
		global.GoCache.Set(key, output, cache.DefaultExpiration)
	}
	return
}

var regionId2NameMap map[int64]string
var once sync.Once

// check lock check的once 实现单例模式
func GetRegionIdNameMap() (id2name map[int64]string, err error) {
	once.Do(func() {
		var idNames []RegionIdName
		cmd := `SELECT id, name FROM mkm_region WHERE is_deleted = 0`
		err = global.BIZ_DBX.Select(&idNames, cmd)
		if err != nil {
			global.GVA_LOG.Errorf("查询region出错，err: [%s]", err)
			return
		}
		regionId2NameMap = make(map[int64]string)
		for _, idName := range idNames {
			regionId2NameMap[idName.Id] = idName.Name
		}
	})
	return regionId2NameMap, err
}
