package pas

import (
	"gin-vue-admin/model"
)

// 如果含有time.Time 请自行import time包
type PerformanceReviewItemUser struct {
	ID     uint          `json:"id" gorm:"primarykey"`
	PRIID   uint         `json:"priid" gorm:"column:pri_id;primarykey;comment:考核表指标ID"`
	Score  float64       `json:"score" gorm:"column:score;comment:考核指标的分数;type:float;"`
	Result  float64      `json:"result" gorm:"column:result;comment:得到的分数;type:float;"`
	UserId uint          `json:"userId" gorm:"column:user_id;primarykey;comment:用户ID"`
	Comment string       `json:"comment" form:"comment" gorm:"column:comment;comment:评分人上传备注;"`
	User   model.SysUser `json:"user" gorm:"ForeignKey:id;AssociationForeignKey:UserId;References:user_id"`
	Status  uint         `json:"status" form:"status" gorm:"column:status;comment:考核指标中每个用户的状态;type:varchar(255);size:255;"`
	PRI PerformanceReviewItem `json:"performanceReviewItem" gorm:"ForeignKey:ID;AssociationForeignKey:PRIID;References:PRIID"`
	//EvaluationKpi EvaluationKpi `json:"evaluationKpi" gorm:"ForeignKey:id;AssociationForeignKey:EKID;References:ek_id"`
}

func (PerformanceReviewItemUser) TableName() string {
	return "performance_review_item_user"
}
