import service from '@/utils/request'

// @Tags EvalutionForm
// @Summary 创建EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "创建EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvalutionForm/createEvalutionForm [post]
export const createEvalutionForm = (data) => {
     return service({
         url: "/EvalutionForm/createEvalutionForm",
         method: 'post',
         data
     })
 }


// @Tags EvalutionForm
// @Summary 删除EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "删除EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EvalutionForm/deleteEvalutionForm [delete]
 export const deleteEvalutionForm = (data) => {
     return service({
         url: "/EvalutionForm/deleteEvalutionForm",
         method: 'delete',
         data
     })
 }

// @Tags EvalutionForm
// @Summary 删除EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EvalutionForm/deleteEvalutionForm [delete]
 export const deleteEvalutionFormByIds = (data) => {
     return service({
         url: "/EvalutionForm/deleteEvalutionFormByIds",
         method: 'delete',
         data
     })
 }

// @Tags EvalutionForm
// @Summary 更新EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "更新EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EvalutionForm/updateEvalutionForm [put]
 export const updateEvalutionForm = (data) => {
     return service({
         url: "/EvalutionForm/updateEvalutionForm",
         method: 'put',
         data
     })
 }


// @Tags EvalutionForm
// @Summary 用id查询EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "用id查询EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EvalutionForm/findEvalutionForm [get]
 export const findEvalutionForm = (params) => {
     return service({
         url: "/EvalutionForm/findEvalutionForm",
         method: 'get',
         params
     })
 }


// @Tags EvalutionForm
// @Summary 分页获取EvalutionForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取EvalutionForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvalutionForm/getEvalutionFormList [get]
 export const getEvalutionFormList = (params) => {
     return service({
         url: "/EvalutionForm/getEvalutionFormList",
         method: 'get',
         params
     })
 }