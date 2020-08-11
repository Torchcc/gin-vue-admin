package initialize

import (
	"time"

	"gin-vue-admin/global"
	"github.com/patrickmn/go-cache"
)

func NewGoCache() *cache.Cache {
	return cache.New(15*time.Minute, 20*time.Minute)
}

func init() {
	global.GoCache = NewGoCache()
}
