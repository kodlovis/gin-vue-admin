import service from '@/utils/request'

// @Tags EvaluationForm
// @Summary 创建EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "创建EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvaluationForm/createEvaluationForm [post]
export const createEvaluationForm = (data) => {
     return service({
         url: "/EvaluationForm/createEvaluationForm",
         method: 'post',
         data
     })
 }


// @Tags EvaluationForm
// @Summary 删除EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "删除EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EvaluationForm/deleteEvaluationForm [delete]
 export const deleteEvaluationForm = (data) => {
     return service({
         url: "/EvaluationForm/deleteEvaluationForm",
         method: 'delete',
         data
     })
 }

// @Tags EvaluationForm
// @Summary 删除EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EvaluationForm/deleteEvaluationForm [delete]
 export const deleteEvaluationFormByIds = (data) => {
     return service({
         url: "/EvaluationForm/deleteEvaluationFormByIds",
         method: 'delete',
         data
     })
 }

// @Tags EvaluationForm
// @Summary 更新EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "更新EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EvaluationForm/updateEvaluationForm [put]
 export const updateEvaluationForm = (data) => {
     return service({
         url: "/EvaluationForm/updateEvaluationForm",
         method: 'put',
         data
     })
 }


// @Tags EvaluationForm
// @Summary 用id查询EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "用id查询EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EvaluationForm/findEvaluationForm [get]
 export const findEvaluationForm = (params) => {
     return service({
         url: "/EvaluationForm/findEvaluationForm",
         method: 'get',
         params
     })
 }


// @Tags EvaluationForm
// @Summary 分页获取EvaluationForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取EvaluationForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvaluationForm/getEvaluationFormList [get]
 export const getEvaluationFormList = (params) => {
     return service({
         url: "/EvaluationForm/getEvaluationFormList",
         method: 'get',
         params
     })
 }