import service from '@/utils/request'

// @Tags PasAllocation
// @Summary 创建PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "创建PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasAllocation/createPasAllocation [post]
export const createPasAllocation = (data) => {
     return service({
         url: "/pasAllocation/createPasAllocation",
         method: 'post',
         data
     })
 }


// @Tags PasAllocation
// @Summary 删除PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "删除PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasAllocation/deletePasAllocation [delete]
 export const deletePasAllocation = (data) => {
     return service({
         url: "/pasAllocation/deletePasAllocation",
         method: 'delete',
         data
     })
 }

// @Tags PasAllocation
// @Summary 删除PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasAllocation/deletePasAllocation [delete]
 export const deletePasAllocationByIds = (data) => {
     return service({
         url: "/pasAllocation/deletePasAllocationByIds",
         method: 'delete',
         data
     })
 }

// @Tags PasAllocation
// @Summary 更新PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "更新PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasAllocation/updatePasAllocation [put]
 export const updatePasAllocation = (data) => {
     return service({
         url: "/pasAllocation/updatePasAllocation",
         method: 'put',
         data
     })
 }


// @Tags PasAllocation
// @Summary 用id查询PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "用id查询PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasAllocation/findPasAllocation [get]
 export const findPasAllocation = (params) => {
     return service({
         url: "/pasAllocation/findPasAllocation",
         method: 'get',
         params
     })
 }


// @Tags PasAllocation
// @Summary 分页获取PasAllocation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PasAllocation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasAllocation/getPasAllocationList [get]
 export const getPasAllocationList = (params) => {
     return service({
         url: "/pasAllocation/getPasAllocationList",
         method: 'get',
         params
     })
 }