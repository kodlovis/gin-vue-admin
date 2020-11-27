import service from '@/utils/request'

// @Tags PasEvaluation
// @Summary 创建PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "创建PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvaluation/createPasEvaluation [post]
export const createPasEvaluation = (data) => {
     return service({
         url: "/pasEvaluation/createPasEvaluation",
         method: 'post',
         data
     })
 }


// @Tags PasEvaluation
// @Summary 删除PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "删除PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasEvaluation/deletePasEvaluation [delete]
 export const deletePasEvaluation = (data) => {
     return service({
         url: "/pasEvaluation/deletePasEvaluation",
         method: 'delete',
         data
     })
 }

// @Tags PasEvaluation
// @Summary 删除PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasEvaluation/deletePasEvaluation [delete]
 export const deletePasEvaluationByIds = (data) => {
     return service({
         url: "/pasEvaluation/deletePasEvaluationByIds",
         method: 'delete',
         data
     })
 }

// @Tags PasEvaluation
// @Summary 更新PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "更新PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasEvaluation/updatePasEvaluation [put]
 export const updatePasEvaluation = (data) => {
     return service({
         url: "/pasEvaluation/updatePasEvaluation",
         method: 'put',
         data
     })
 }


// @Tags PasEvaluation
// @Summary 用id查询PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "用id查询PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasEvaluation/findPasEvaluation [get]
 export const findPasEvaluation = (params) => {
     return service({
         url: "/pasEvaluation/findPasEvaluation",
         method: 'get',
         params
     })
 }


// @Tags PasEvaluation
// @Summary 分页获取PasEvaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PasEvaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvaluation/getPasEvaluationList [get]
 export const getPasEvaluationList = (params) => {
     return service({
         url: "/pasEvaluation/getPasEvaluationList",
         method: 'get',
         params
     })
 }