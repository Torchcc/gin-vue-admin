import service from '@/utils/request'

// @Tags PkgAttr
// @Summary 创建PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "创建PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgAttr/createPkgAttr [post]
export const createPkgAttr = (data) => {
     return service({
         url: "/pkgAttr/createPkgAttr",
         method: 'post',
         data
     })
 }


// @Tags PkgAttr
// @Summary 删除PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "删除PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkgAttr/deletePkgAttr [delete]
 export const deletePkgAttr = (data) => {
     return service({
         url: "/pkgAttr/deletePkgAttr",
         method: 'delete',
         data
     })
 }

// @Tags PkgAttr
// @Summary 更新PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "更新PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkgAttr/updatePkgAttr [put]
 export const updatePkgAttr = (data) => {
     return service({
         url: "/pkgAttr/updatePkgAttr",
         method: 'put',
         data
     })
 }


// @Tags PkgAttr
// @Summary 用id查询PkgAttr
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgAttr true "用id查询PkgAttr"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkgAttr/findPkgAttr [get]
 export const findPkgAttr = (params) => {
     return service({
         url: "/pkgAttr/findPkgAttr",
         method: 'get',
         params
     })
 }


// @Tags PkgAttr
// @Summary 分页获取PkgAttr列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PkgAttr列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkgAttr/getPkgAttrList [get]
 export const getPkgAttrList = (params) => {
     return service({
         url: "/pkgAttr/getPkgAttrList",
         method: 'get',
         params
     })
 }