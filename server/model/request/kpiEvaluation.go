package request

import "gin-vue-admin/model"

type KpiEvaluationSearch struct{
    model.KpiEvaluation
    PageInfo
}