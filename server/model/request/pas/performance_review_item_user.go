package pas

import "gin-vue-admin/model/pas"

type PerformanceReviewItemUserSearch struct {
	pas.PerformanceReviewItemUser
	PageInfo
	EKUID uint `json:"ekuid"`
	PRID uint `json:"prid"`
	
}