package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPackageRouter(Router *gin.RouterGroup) {
	PackageRouter := Router.Group("pkg") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())  TODO 调试完umcomment
	{
		PackageRouter.POST("createPackage", v1.CreatePackage)   // 新建Package
		PackageRouter.DELETE("deletePackage", v1.DeletePackage) // 删除Package
		PackageRouter.PUT("updatePackage", v1.UpdatePackage)    // 更新Package
		PackageRouter.GET("findPackage/:id", v1.FindPackage)    // 根据ID获取Package
		PackageRouter.GET("getPackageList", v1.GetPackageList)  // 获取Package列表
	}
}
