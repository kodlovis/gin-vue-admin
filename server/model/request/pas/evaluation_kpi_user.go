package pas

import "gin-vue-admin/model/pas"

type EvaluationKpiUserSearch struct {
	pas.EvaluationKpiUser
	PageInfo
	AuthorityId	string `json:"authorityId"`
	EKID uint `json:"ekid"`
}
type EKU struct{
	EvaluationKpiUser []pas.EvaluationKpiUser `json:"items"`
	AuthorityId	string `json:"authorityId"`
	EKID uint `json:"ekid"`
}