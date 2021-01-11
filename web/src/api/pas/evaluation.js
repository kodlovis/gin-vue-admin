import service from '@/utils/request'

// @Tags Evaluation
// @Summary 创建Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "创建Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Evaluation/createEvaluation [post]
export const createEvaluation = (data) => {
     return service({
         url: "/Evaluation/createEvaluation",
         method: 'post',
         data
     })
 }

// @Tags Evaluation
// @Summary 删除Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "删除Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Evaluation/deleteEvaluation [delete]
 export const deleteEvaluation = (data) => {
     return service({
         url: "/Evaluation/deleteEvaluation",
         method: 'delete',
         data
     })
 }

// @Tags Evaluation
// @Summary 删除Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Evaluation/deleteEvaluation [delete]
 export const deleteEvaluationByIds = (data) => {
     return service({
         url: "/Evaluation/deleteEvaluationByIds",
         method: 'delete',
         data
     })
 }

// @Tags Evaluation
// @Summary 更新Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "更新Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Evaluation/updateEvaluation [put]
 export const updateEvaluation = (data) => {
     return service({
         url: "/Evaluation/updateEvaluation",
         method: 'put',
         data
     })
 }


// @Tags Evaluation
// @Summary 用id查询Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "用id查询Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Evaluation/findEvaluation [get]
 export const findEvaluation = (data) => {
     return service({
         url: "/Evaluation/findEvaluation",
         method: 'post',
         data
     })
 }


// @Tags Evaluation
// @Summary 分页获取Evaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Evaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Evaluation/getEvaluationList [get]
 export const getEvaluationList = (params) => {
     return service({
         url: "/Evaluation/getEvaluationList",
         method: 'get',
         params
     })
 }

 export const updateEvaluationByInfo = (data) => {
    return service({
        url: "/Evaluation/updateEvaluationByInfo",
        method: 'post',
        data
    })
}