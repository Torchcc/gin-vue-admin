package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRegionRouter(Router *gin.RouterGroup) {
	RegionRouter := Router.Group("region").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{

		RegionRouter.GET("getRegionList", v1.GetRegionsByParentId) // 根据父级ID获取区域
	}
}
