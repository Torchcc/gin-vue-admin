import service from '@/utils/request'

// @Tags Package
// @Summary 创建Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "创建Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkg/createPackage [post]
export const createPackage = (data) => {
     return service({
         url: "/pkg/createPackage",
         method: 'post',
         data
     })
 }


// @Tags Package
// @Summary 删除Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "删除Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pkg/deletePackage [delete]
 export const deletePackage = (data) => {
     return service({
         url: "/pkg/deletePackage",
         method: 'delete',
         data
     })
 }

// @Tags Package
// @Summary 更新Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "更新Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pkg/updatePackage [put]
 export const updatePackage = (data) => {
     return service({
         url: "/pkg/updatePackage",
         method: 'put',
         data
     })
 }


// @Tags Package
// @Summary 用id查询Package
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Package true "用id查询Package"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pkg/findPackage [get]
 export const findPackage = (params) => {
     return service({
         url: "/pkg/findPackage",
         method: 'get',
         params
     })
 }


// @Tags Package
// @Summary 分页获取Package列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Package列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pkg/getPackageList [get]
 export const getPackageList = (params) => {
     return service({
         url: "/pkg/getPackageList",
         method: 'get',
         params
     })
 }

export const upatePkgCtgRelation = (data) => {
  return service({
    url: "/pkg/updatePkgCtgRelation",
    method: 'put',
    data
  })
}

export const upatePkgDiseaseRelation = (data) => {
  return service({
    url: "/pkg/updatePkgDiseaseRelation",
    method: 'put',
    data
  })
}

export const getPkgAttrByPkgId = (params) => {
  return service({
    url: "/pkg/getPkgAttr",
    method: "get",
    params
  })
}