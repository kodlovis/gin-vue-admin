// 自动生成模板PerformanceReview
package pas

import (
	"gin-vue-admin/model"
	"time"
)

// 如果含有time.Time 请自行import time包
type PerformanceReview struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `json:"name" form:"name" gorm:"column:name;comment:绩效考核表名称;"`
	Status uint   `json:"status" form:"status" gorm:"column:status;comment:考核表状态;type:varchar(255);size:255;"`
	//
	User       model.SysUser           `json:"user" gorm:"foreignKey:ID;References:EmployeeId;AssociationForeignKey:EmployeeId;"`
	EmployeeId uint                    `json:"employeeId" gorm:"column:employee_id;comment:员工ID关联User表"`
	Result     float64                 `json:"result" form:"result" gorm:"column:result;comment:得分"`
	PRItems    []PerformanceReviewItem `json:"pRItems" gorm:"foreignKey:PRId;References:ID;AssociationForeignKey:ID;"`
	StartDate  time.Time               `json:"startDate" form:"startDate" gorm:"column:startDate;comment:开始日期;type:datetime;"`
	EndingDate time.Time               `json:"endingDate" form:"endingDate" gorm:"column:endingDate;comment:结束日期;type:datetime;"`
	Comment    string                  `json:"comment" form:"comment" gorm:"column:comment;comment:员工上传备注;"`
	//
	Evaluation   Evaluation `json:"evaluation" gorm:"foreignKey:id;References:evaluation_id;AssociationForeignKey:EvaluationId;"`
	EvaluationId uint       `json:"evaluationId" form:"evaluationId" gorm:"column:evaluation_id;comment:关联EvaluationKpi表"`
	Score        float64    `json:"score" form:"score" gorm:"column:score;comment:更改后的分数"`
}

func (PerformanceReview) TableName() string {
	return "performanceReview"
}
