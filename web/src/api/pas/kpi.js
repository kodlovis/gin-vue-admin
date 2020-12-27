import service from '@/utils/request'

// @Tags Kpi
// @Summary 创建Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "创建Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Kpi/createKpi [post]
export const createKpi = (data) => {
     return service({
         url: "/Kpi/createKpi",
         method: 'post',
         data
     })
 }


// @Tags Kpi
// @Summary 删除Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "删除Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Kpi/deleteKpi [delete]
 export const deleteKpi = (data) => {
     return service({
         url: "/Kpi/deleteKpi",
         method: 'delete',
         data
     })
 }

 export const removeKpiTags = (data) => {
    return service({
        url: "/Kpi/removeKpiTags",
        method: 'delete',
        data
    })
}
// @Tags Kpi
// @Summary 删除Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Kpi/deleteKpi [delete]
 export const deleteKpiByIds = (data) => {
     return service({
         url: "/Kpi/deleteKpiByIds",
         method: 'delete',
         data
     })
 }


// @Tags Kpi
// @Summary 更新Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "更新Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Kpi/updateKpi [put]
 export const updateKpi = (data) => {
     return service({
         url: "/Kpi/updateKpi",
         method: 'put',
         data
     })
 }


// @Tags Kpi
// @Summary 用id查询Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "用id查询Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Kpi/findKpi [get]
 export const findKpi = (params) => {
     return service({
         url: "/Kpi/findKpi",
         method: 'get',
         params
     })
 }

 export const getKpiByIds = (data) => {
    return service({
        url: "/Kpi/getKpiByIds",
        method: 'post',
        data
    })
}

export const getKpiScoreByIds = (data) => {
    return service({
        url: "/Kpi/getKpiScoreByIds",
        method: 'post',
        data
    })
}
// @Tags Kpi
// @Summary 分页获取Kpi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Kpi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Kpi/getKpiList [get]
 export const getKpiList = (params) => {
     return service({
         url: "/Kpi/getKpiList",
         method: 'get',
         params
     })
 }

 export const assignedKpiEvaluation = (data) => {
    return service({
        url: "/Kpi/assignedKpiEvaluation",
        method: 'post',
        data
    })
}

export const getKpiEvaluation = (data) => {
    return service({
        url: "/Kpi/getKpiEvaluation",
        method: 'post',
        data
    })
}