package pas

import "gin-vue-admin/model/pas"

type PerformanceReviewSearch struct{
    pas.PerformanceReview
    PageInfo
}