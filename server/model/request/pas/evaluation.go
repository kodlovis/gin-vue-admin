package pas

import "gin-vue-admin/model/pas"

type EvaluationSearch struct{
    pas.Evaluation
    PageInfo
}