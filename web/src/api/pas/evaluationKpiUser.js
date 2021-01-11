import service from '@/utils/request'

 export const findEvaluationKpiUser = (data) => {
     return service({
         url: "/EvaluationKpiUser/findEvaluationKpiUser",
         method: 'post',
         data
     })
 }

 
 export const removeEvaluationUsersByIds = (data) => {
    return service({
        url: "/EvaluationKpiUser/removeEvaluationUsersByIds",
        method: 'delete',
        data
    })
}
