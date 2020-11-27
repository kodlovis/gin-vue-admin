import service from '@/utils/request'

// @Tags PasEvalutionForm
// @Summary 创建PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "创建PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvalutionform/createPasEvalutionForm [post]
export const createPasEvalutionForm = (data) => {
     return service({
         url: "/pasEvalutionform/createPasEvalutionForm",
         method: 'post',
         data
     })
 }


// @Tags PasEvalutionForm
// @Summary 删除PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "删除PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasEvalutionform/deletePasEvalutionForm [delete]
 export const deletePasEvalutionForm = (data) => {
     return service({
         url: "/pasEvalutionform/deletePasEvalutionForm",
         method: 'delete',
         data
     })
 }

// @Tags PasEvalutionForm
// @Summary 删除PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasEvalutionform/deletePasEvalutionForm [delete]
 export const deletePasEvalutionFormByIds = (data) => {
     return service({
         url: "/pasEvalutionform/deletePasEvalutionFormByIds",
         method: 'delete',
         data
     })
 }

// @Tags PasEvalutionForm
// @Summary 更新PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "更新PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasEvalutionform/updatePasEvalutionForm [put]
 export const updatePasEvalutionForm = (data) => {
     return service({
         url: "/pasEvalutionform/updatePasEvalutionForm",
         method: 'put',
         data
     })
 }


// @Tags PasEvalutionForm
// @Summary 用id查询PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "用id查询PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasEvalutionform/findPasEvalutionForm [get]
 export const findPasEvalutionForm = (params) => {
     return service({
         url: "/pasEvalutionform/findPasEvalutionForm",
         method: 'get',
         params
     })
 }


// @Tags PasEvalutionForm
// @Summary 分页获取PasEvalutionForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PasEvalutionForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvalutionform/getPasEvalutionFormList [get]
 export const getPasEvalutionFormList = (params) => {
     return service({
         url: "/pasEvalutionform/getPasEvalutionFormList",
         method: 'get',
         params
     })
 }