package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHospitalRouter(Router *gin.RouterGroup) {
	HospitalRouter := Router.Group("hospital").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HospitalRouter.POST("createHospital", v1.CreateHospital)             // 新建Hospital
		HospitalRouter.POST("uploadAvatar", v1.UploadHospitalAvatar)         // 上传医院头像
		HospitalRouter.DELETE("deleteHospital", v1.DeleteHospital)           // 删除Hospital
		HospitalRouter.DELETE("deleteHospitalByIds", v1.DeleteHospitalByIds) // 批量删除Hospital
		HospitalRouter.PUT("updateHospital", v1.UpdateHospital)              // 更新Hospital
		HospitalRouter.GET("findHospital", v1.FindHospital)                  // 根据ID获取Hospital
		HospitalRouter.GET("getHospitalList", v1.GetHospitalList)            // 获取Hospital列表
	}
}
