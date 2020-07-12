import service from '@/utils/request'

// @Tags PkgCategory
// @Summary 创建PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "创建PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgCtg/createPkgCategory [post]
export const createPkgCategory = (data) => {
     return service({
         url: "/pkgCtg/createPkgCategory",
         method: 'post',
         data
     })
 }


// @Tags PkgCategory
// @Summary 删除PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "删除PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkgCtg/deletePkgCategory [delete]
 export const deletePkgCategory = (data) => {
     return service({
         url: "/pkgCtg/deletePkgCategory",
         method: 'delete',
         data
     })
 }

// @Tags PkgCategory
// @Summary 更新PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "更新PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkgCtg/updatePkgCategory [put]
 export const updatePkgCategory = (data) => {
     return service({
         url: "/pkgCtg/updatePkgCategory",
         method: 'put',
         data
     })
 }


// @Tags PkgCategory
// @Summary 用id查询PkgCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgCategory true "用id查询PkgCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkgCtg/findPkgCategory [get]
 export const findPkgCategory = (params) => {
     return service({
         url: "/pkgCtg/findPkgCategory",
         method: 'get',
         params
     })
 }


// @Tags PkgCategory
// @Summary 分页获取PkgCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PkgCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgCtg/getPkgCategoryList [get]
 export const getPkgCategoryList = (params) => {
     return service({
         url: "/pkgCtg/getPkgCategoryList",
         method: 'get',
         params
     })
 }