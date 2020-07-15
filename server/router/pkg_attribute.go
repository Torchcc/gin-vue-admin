package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPkgAttrRouter(Router *gin.RouterGroup) {
	PkgAttrRouter := Router.Group("pkgAttr").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		PkgAttrRouter.POST("createPkgAttr", v1.CreatePkgAttr)   // 新建PkgAttr
		PkgAttrRouter.DELETE("deletePkgAttr", v1.DeletePkgAttr) // 删除PkgAttr
		PkgAttrRouter.PUT("updatePkgAttr", v1.UpdatePkgAttr)    // 更新PkgAttr
		PkgAttrRouter.GET("findPkgAttr", v1.FindPkgAttr)        // 根据ID获取PkgAttr
		PkgAttrRouter.GET("getPkgAttrList", v1.GetPkgAttrList)  // 获取PkgAttr列表
	}
}
