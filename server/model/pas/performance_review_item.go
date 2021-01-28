package pas

import (
	"gin-vue-admin/model"
)

// 如果含有time.Time 请自行import time包
type PerformanceReviewItem struct {
	ID      uint              `gorm:"primarykey"`
	PRId    uint              `json:"PRId" form:"PRId" gorm:"column:pr_id;comment:关联performanceReview表"`
	Kpi     Kpi               `json:"kpi" gorm:"foreignKey:ID;References:KpiId;AssociationForeignKey:KpiId;"`
	KpiId   uint              `json:"kpiId" form:"kpiId" gorm:"column:kpi_id;comment:关联Kpi表"`
	UserId  uint              `json:"userId" form:"userId" gorm:"column:user_id;comment:关联User表"`
	Status  uint              `json:"status" form:"status" gorm:"column:status;comment:考核指标流程状态;type:varchar(255);size:255;"`
	Result  float64           `json:"result" form:"result" gorm:"column:result;comment:得分"`
	Score   float64           `json:"score" form:"score" gorm:"column:score;comment:更改后的分数"`
	Comment string            `json:"comment" form:"comment" gorm:"column:comment;comment:评分人上传备注;"`
	User    model.SysUser     `json:"user" gorm:"foreignKey:ID;References:UserId;AssociationForeignKey:UserId;"`
	PRs     PerformanceReview `json:"prs" gorm:"foreignKey:ID;References:PRId;AssociationForeignKey:PRId;"`
}

func (PerformanceReviewItem) TableName() string {
	return "performanceReviewItem"
}
