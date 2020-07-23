package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDiseaseRouter(Router *gin.RouterGroup) {
	DiseaseRouter := Router.Group("disease").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DiseaseRouter.POST("createDisease", v1.CreateDisease)             // 新建Disease
		DiseaseRouter.DELETE("deleteDisease", v1.DeleteDisease)           // 删除Disease
		DiseaseRouter.DELETE("deleteDiseaseByIds", v1.DeleteDiseaseByIds) // 批量删除Disease
		DiseaseRouter.PUT("updateDisease", v1.UpdateDisease)              // 更新Disease
		DiseaseRouter.GET("findDisease", v1.FindDisease)                  // 根据ID获取Disease
		DiseaseRouter.GET("getDiseaseList", v1.GetDiseaseList)            // 获取Disease列表
	}
}
