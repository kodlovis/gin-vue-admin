import service from '@/utils/request'

 export const findEvaluationKpiUser = (data) => {
     return service({
         url: "/EvaluationKpiUser/findEvaluationKpiUser",
         method: 'post',
         data
     })
 }

 export const createEKU = (data) => {
    return service({
        url: "/EvaluationKpiUser/createEKU",
        method: 'post',
        data
    })
}
