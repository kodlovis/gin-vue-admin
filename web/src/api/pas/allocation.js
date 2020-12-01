import service from '@/utils/request'

// @Tags Allocation
// @Summary 创建Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "创建Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Allocation/createAllocation [post]
export const createAllocation = (data) => {
     return service({
         url: "/Allocation/createAllocation",
         method: 'post',
         data
     })
 }


// @Tags Allocation
// @Summary 删除Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "删除Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Allocation/deleteAllocation [delete]
 export const deleteAllocation = (data) => {
     return service({
         url: "/Allocation/deleteAllocation",
         method: 'delete',
         data
     })
 }

// @Tags Allocation
// @Summary 删除Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Allocation/deleteAllocation [delete]
 export const deleteAllocationByIds = (data) => {
     return service({
         url: "/Allocation/deleteAllocationByIds",
         method: 'delete',
         data
     })
 }

// @Tags Allocation
// @Summary 更新Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "更新Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Allocation/updateAllocation [put]
 export const updateAllocation = (data) => {
     return service({
         url: "/Allocation/updateAllocation",
         method: 'put',
         data
     })
 }


// @Tags Allocation
// @Summary 用id查询Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "用id查询Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Allocation/findAllocation [get]
 export const findAllocation = (params) => {
     return service({
         url: "/Allocation/findAllocation",
         method: 'get',
         params
     })
 }


// @Tags Allocation
// @Summary 分页获取Allocation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Allocation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Allocation/getAllocationList [get]
 export const getAllocationList = (params) => {
     return service({
         url: "/Allocation/getAllocationList",
         method: 'get',
         params
     })
 }