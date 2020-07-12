package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPkgCategoryRouter(Router *gin.RouterGroup) {
	PkgCategoryRouter := Router.Group("pkgCtg").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		PkgCategoryRouter.POST("createPkgCategory", v1.CreatePkgCategory)   // 新建PkgCategory
		PkgCategoryRouter.DELETE("deletePkgCategory", v1.DeletePkgCategory) // 删除PkgCategory
		PkgCategoryRouter.PUT("updatePkgCategory", v1.UpdatePkgCategory)    // 更新PkgCategory
		PkgCategoryRouter.GET("findPkgCategory", v1.FindPkgCategory)        // 根据ID获取PkgCategory
		PkgCategoryRouter.GET("getPkgCategoryList", v1.GetPkgCategoryList)  // 获取PkgCategory列表
	}
}
