import service from '@/utils/request'

// @Tags PasKpi
// @Summary 创建PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "创建PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasKpi/createPasKpi [post]
export const createPasKpi = (data) => {
     return service({
         url: "/pasKpi/createPasKpi",
         method: 'post',
         data
     })
 }


// @Tags PasKpi
// @Summary 删除PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "删除PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasKpi/deletePasKpi [delete]
 export const deletePasKpi = (data) => {
     return service({
         url: "/pasKpi/deletePasKpi",
         method: 'delete',
         data
     })
 }

// @Tags PasKpi
// @Summary 删除PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasKpi/deletePasKpi [delete]
 export const deletePasKpiByIds = (data) => {
     return service({
         url: "/pasKpi/deletePasKpiByIds",
         method: 'delete',
         data
     })
 }

// 樊新增批量归档
//  export const updatePasKpiByIds = (data) => {
//     return service({
//         url: "/pasKpi/updatePasKpiByIds",
//         method: 'update',
//         data
//     })
// }


// @Tags PasKpi
// @Summary 更新PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "更新PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasKpi/updatePasKpi [put]
 export const updatePasKpi = (data) => {
     return service({
         url: "/pasKpi/updatePasKpi",
         method: 'put',
         data
     })
 }


// @Tags PasKpi
// @Summary 用id查询PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "用id查询PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasKpi/findPasKpi [get]
 export const findPasKpi = (params) => {
     return service({
         url: "/pasKpi/findPasKpi",
         method: 'get',
         params
     })
 }


// @Tags PasKpi
// @Summary 分页获取PasKpi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PasKpi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasKpi/getPasKpiList [get]
 export const getPasKpiList = (params) => {
     return service({
         url: "/pasKpi/getPasKpiList",
         method: 'get',
         params
     })
 }