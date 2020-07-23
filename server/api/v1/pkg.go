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

func fen2yuan(p *model.Package) {
	p.PriceOriginal *= 0.01
	p.PriceReal *= 0.01
}

func yuan2fen(p *model.Package) {
	p.PriceOriginal *= 100
	p.PriceReal *= 100
}

// @Tags Package
// @Summary 创建Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "创建Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkg/createPackage [post]
func CreatePackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindJSON(&pkg)
	yuan2fen(&pkg)
	err := service.CreatePackage(pkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Package
// @Summary 删除Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "删除Package, data 只需要传主键id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkg/deletePackage [delete]
func DeletePackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindJSON(&pkg)
	err := service.DeletePackage(&pkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Package
// @Summary 更新Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "更新Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkg/updatePackage [put]
func UpdatePackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindJSON(&pkg)
	yuan2fen(&pkg)
	err := service.UpdatePackage(&pkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Package
// @Summary 用id查询Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "用id查询Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkg/findPackage/ [get]
func FindPackage(c *gin.Context) {
	var pkg model.Package
	_ = c.ShouldBindQuery(&pkg)
	err, repkg := service.GetPackage(pkg.ID)
	fen2yuan(&repkg)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"repkg": repkg}, c)
	}
}

// @Tags Package
// @Summary 分页获取Package列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query string true "分页获取Package列表page"
// @Param pageSize query string true "分页获取Package列表pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkg/getPackageList [get]
func GetPackageList(c *gin.Context) {
	var pageInfo request.PackageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetPackageInfoList(pageInfo)
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

func UploadPackageAvatar(c *gin.Context) {
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
			err = service.UploadPkgAvatar(uint(id), filePath)
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

func UpdatePkgCtgRelation(c *gin.Context) {
	var pkgCtg model.PkgWithCtgNDisease
	_ = c.ShouldBindJSON(&pkgCtg)
	err := service.UpdatePkgCtgRelation(&pkgCtg)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}

func UpdatePkgDiseaseRelation(c *gin.Context) {
	var pkgDisease model.PkgWithCtgNDisease
	_ = c.ShouldBindJSON(&pkgDisease)
	err := service.UpdatePkgDiseaseRelation(&pkgDisease)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}
