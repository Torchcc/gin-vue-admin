import service from '@/utils/request'

// @Tags Disease
// @Summary 创建Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "创建Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /disease/createDisease [post]
export const createDisease = (data) => {
     return service({
         url: "/disease/createDisease",
         method: 'post',
         data
     })
 }


// @Tags Disease
// @Summary 删除Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "删除Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disease/deleteDisease [delete]
 export const deleteDisease = (data) => {
     return service({
         url: "/disease/deleteDisease",
         method: 'delete',
         data
     })
 }

// @Tags Disease
// @Summary 删除Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disease/deleteDisease [delete]
 export const deleteDiseaseByIds = (data) => {
     return service({
         url: "/disease/deleteDiseaseByIds",
         method: 'delete',
         data
     })
 }

// @Tags Disease
// @Summary 更新Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "更新Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /disease/updateDisease [put]
 export const updateDisease = (data) => {
     return service({
         url: "/disease/updateDisease",
         method: 'put',
         data
     })
 }


// @Tags Disease
// @Summary 用id查询Disease
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Disease true "用id查询Disease"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /disease/findDisease [get]
 export const findDisease = (params) => {
     return service({
         url: "/disease/findDisease",
         method: 'get',
         params
     })
 }


// @Tags Disease
// @Summary 分页获取Disease列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Disease列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /disease/getDiseaseList [get]
 export const getDiseaseList = (params) => {
     return service({
         url: "/disease/getDiseaseList",
         method: 'get',
         params
     })
 }