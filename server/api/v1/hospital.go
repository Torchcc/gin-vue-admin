package v1

import (
	"fmt"
	"strconv"

	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags Hospital
// @Summary 创建Hospital
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hospital true "创建Hospital"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hospital/createHospital [post]
func CreateHospital(c *gin.Context) {
	var hospital model.Hospital
	_ = c.ShouldBindJSON(&hospital)
	err := service.CreateHospital(hospital)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Hospital
// @Summary 删除Hospital
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hospital true "删除Hospital"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hospital/deleteHospital [delete]
func DeleteHospital(c *gin.Context) {
	var hospital model.Hospital
	_ = c.ShouldBindJSON(&hospital)
	err := service.DeleteHospital(hospital)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Hospital
// @Summary 批量删除Hospital
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hospital"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hospital/deleteHospitalByIds [delete]
func DeleteHospitalByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHospitalByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Hospital
// @Summary 更新Hospital
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hospital true "更新Hospital"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hospital/updateHospital [put]
func UpdateHospital(c *gin.Context) {
	var hospital model.Hospital
	_ = c.ShouldBindJSON(&hospital)
	err := service.UpdateHospital(&hospital)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Hospital
// @Summary 用id查询Hospital
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "用id查询Hospital"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hospital/findHospital [get]
func FindHospital(c *gin.Context) {
	var hospital model.Hospital
	_ = c.ShouldBindQuery(&hospital)
	err, rehospital := service.GetHospital(hospital.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehospital": rehospital}, c)
	}
}

// @Tags Hospital
// @Summary 分页获取Hospital列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取Hospital列表page"
// @Param pageSize query string true "分页获取Hospital列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hospital/getHospitalList [get]
func GetHospitalList(c *gin.Context) {
	var pageInfo request.HospitalSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHospitalInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

func UploadHospitalAvatar(c *gin.Context) {
	_, avatar, err := c.Request.FormFile("avatar")
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	// 便于找到用户 以后从jwt中取
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		// 文件上传后拿到文件路径
		err, filePath, _ := utils.Upload(avatar)
		fmt.Print("文件路径: ", filePath)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			err = service.UploadHospitalAvatar(uint(id), filePath)
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				response.OkWithData(struct {
					AvatarUrl string `json:"avatar_url"`
					Id        int    `json:"id"`
				}{AvatarUrl: filePath, Id: id}, c)
			}
		}
	}
}
