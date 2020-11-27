package request

import "gin-vue-admin/model"

type PasEvaluationSearch struct{
    model.PasEvaluation
    PageInfo
}