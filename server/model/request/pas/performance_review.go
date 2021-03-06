package pas

import "gin-vue-admin/model/pas"
import "time"

type PerformanceReviewSearch struct{
    pas.PerformanceReview
    PageInfo
    Status uint `json:"status"`
    ID uint `json:"ID"`
    Ids []int `json:"ids"`
    NickName string `json:"nickName" form:"nickName"` 
}
type PerformanceReviewInfo struct {
    ID uint `json:"ID"`
    EmployeeId uint `json:"employeeId"`
    EvaluationId uint `json:"evaluationId"` 
    Name string `json:"name"`
    Status uint `json:"status"`
    StartDate time.Time `json:"startDate"`
    EndingDate time.Time `json:"endingDate"`
    Score float64 `json:"score"`
    Ids []int `json:"ids"`
}
