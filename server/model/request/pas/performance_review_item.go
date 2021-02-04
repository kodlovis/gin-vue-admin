package pas

import "gin-vue-admin/model/pas"

type PerformanceReviewItemSearch struct {
	pas.PerformanceReviewItem
	PageInfo
	ID     uint  `json:"ID"`
	Status uint  `json:"status"`
	PRID   uint  `json:"prid"`
	Ids    []int `json:"ids"`
}

type PerformanceReviewItemList struct {
	PerformanceReviewItemList []pas.PerformanceReviewItem `json:"item"`
}
type PRItemInfo struct {
	ID     uint `json:"ID"`
	Status uint `json:"status"`
	Ids    int  `json:"ids"`
}
