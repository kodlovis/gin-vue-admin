import service from '@/utils/request'

// @Tags KpiEvaluation
// @Summary 创建KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "创建KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kpiEvaluation/createKpiEvaluation [post]
export const createKpiEvaluation = (data) => {
     return service({
         url: "/kpiEvaluation/createKpiEvaluation",
         method: 'post',
         data
     })
 }


// @Tags KpiEvaluation
// @Summary 删除KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "删除KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kpiEvaluation/deleteKpiEvaluation [delete]
 export const deleteKpiEvaluation = (data) => {
     return service({
         url: "/kpiEvaluation/deleteKpiEvaluation",
         method: 'delete',
         data
     })
 }

// @Tags KpiEvaluation
// @Summary 删除KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kpiEvaluation/deleteKpiEvaluation [delete]
 export const deleteKpiEvaluationByIds = (data) => {
     return service({
         url: "/kpiEvaluation/deleteKpiEvaluationByIds",
         method: 'delete',
         data
     })
 }

// @Tags KpiEvaluation
// @Summary 更新KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "更新KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /kpiEvaluation/updateKpiEvaluation [put]
 export const updateKpiEvaluation = (data) => {
     return service({
         url: "/kpiEvaluation/updateKpiEvaluation",
         method: 'put',
         data
     })
 }


// @Tags KpiEvaluation
// @Summary 用id查询KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "用id查询KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /kpiEvaluation/findKpiEvaluation [get]
 export const findKpiEvaluation = (params) => {
     return service({
         url: "/kpiEvaluation/findKpiEvaluation",
         method: 'get',
         params
     })
 }


// @Tags KpiEvaluation
// @Summary 分页获取KpiEvaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取KpiEvaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kpiEvaluation/getKpiEvaluationList [get]
 export const getKpiEvaluationList = (params) => {
     return service({
         url: "/kpiEvaluation/getKpiEvaluationList",
         method: 'get',
         params
     })
 }