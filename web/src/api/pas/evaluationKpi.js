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