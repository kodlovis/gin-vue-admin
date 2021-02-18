package pas

import "gin-vue-admin/model/pas"

type PerformanceReviewItemUserSearch struct {
	pas.PerformanceReviewItemUser
	PageInfo
	EKUID uint `json:"ekuid"`
	PRID uint `json:"prid"`
	PRIID uint `json:"priid"`
	ID     uint  `json:"ID"`
	Status uint  `json:"status"`
	Result float64  `json:"result"`
	PRStatus uint  `json:"prstatus"`
}
type PRIU struct{
	PerformanceReviewItemUser []pas.PerformanceReviewItemUser `json:"items"`
	AuthorityId	string `json:"authorityId"`
	PRIID uint `json:"priid"`
}