import service from '@/utils/request'

// @Tags PasTag
// @Summary 创建PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "创建PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasTag/createPasTag [post]
export const createPasTag = (data) => {
     return service({
         url: "/pasTag/createPasTag",
         method: 'post',
         data
     })
 }


// @Tags PasTag
// @Summary 删除PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "删除PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasTag/deletePasTag [delete]
 export const deletePasTag = (data) => {
     return service({
         url: "/pasTag/deletePasTag",
         method: 'delete',
         data
     })
 }

// @Tags PasTag
// @Summary 删除PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasTag/deletePasTag [delete]
 export const deletePasTagByIds = (data) => {
     return service({
         url: "/pasTag/deletePasTagByIds",
         method: 'delete',
         data
     })
 }

// @Tags PasTag
// @Summary 更新PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "更新PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasTag/updatePasTag [put]
 export const updatePasTag = (data) => {
     return service({
         url: "/pasTag/updatePasTag",
         method: 'put',
         data
     })
 }


// @Tags PasTag
// @Summary 用id查询PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "用id查询PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasTag/findPasTag [get]
 export const findPasTag = (params) => {
     return service({
         url: "/pasTag/findPasTag",
         method: 'get',
         params
     })
 }


// @Tags PasTag
// @Summary 分页获取PasTag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PasTag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasTag/getPasTagList [get]
 export const getPasTagList = (params) => {
     return service({
         url: "/pasTag/getPasTagList",
         method: 'get',
         params
     })
 }