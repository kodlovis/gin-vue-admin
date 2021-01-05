package pas

import "gin-vue-admin/model/pas"

type EvaluationSearch struct{
    pas.Evaluation
    PageInfo
}
type EvaluationInfo struct {
    ID uint `json:"Id"`
    TotalScore float64  `json:"Score"`
}