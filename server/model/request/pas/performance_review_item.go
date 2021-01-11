package pas

import "gin-vue-admin/model/pas"

type PerformanceReviewItemSearch struct{
    pas.PerformanceReviewItem
    PageInfo
}

type PerformanceReviewItemList struct{
	PerformanceReviewItemList []pas.PerformanceReviewItem `json:"item"`
}