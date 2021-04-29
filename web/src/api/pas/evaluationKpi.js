import service from '@/utils/request'

 export const addUserEvaluationKpi = (data) => {
    return service({
        url: "/EvaluationKpi/addUserEvaluationKpi",
        method: 'post',
        data
    })
}

export const getEvaluationKpiList = (params) => {
    return service({
        url: "/EvaluationKpi/getEvaluationKpiList",
        method: 'get',
        params
    })
}

export const setUserEvaluation= (data) => {
    return service({
        url: "/EvaluationKpi/setUserEvaluation",
        method: 'post',
        data
    })
}

export const removeEvaluationKpi = (data) => {
    return service({
        url: "/EvaluationKpi/removeEvaluationKpi",
        method: 'delete',
        data
    })
}
export const createEvaluationKpi = (data) => {
     return service({
         url: "/EvaluationKpi/createEvaluationKpi",
         method: 'post',
         data
     })
 }
 
 export const removeEvaluationKpiByIds = (data) => {
    return service({
        url: "/EvaluationKpi/removeEvaluationKpiByIds",
        method: 'delete',
        data
    })
}

export const getEvaluationKpiById = (data) => {
    return service({
        url: "/EvaluationKpi/getEvaluationKpiById",
        method: 'post',
        data
    })
}
export const updateEvaluationKpi = (data) => {
    return service({
        url: "/EvaluationKpi/updateEvaluationKpi",
        method: 'PUT',
        data
    })
}

export const getLastEvaluationKpi = (params) => {
    return service({
        url: "/EvaluationKpi/getLastEvaluationKpi",
        method: 'get',
        params
    })
}

export const getEvaluationKpi = (params) => {
    return service({
        url: "/EvaluationKpi/getEvaluationKpi",
        method: 'get',
        params
    })
}