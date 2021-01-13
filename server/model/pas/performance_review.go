// 自动生成模板PerformanceReview
package pas

import (
  "time"
  "gin-vue-admin/model"
)

// 如果含有time.Time 请自行import time包
type PerformanceReview struct {
      ID  uint `gorm:"primarykey"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:绩效考核表名称;"`
      Status  string `json:"status" form:"status" gorm:"column:status;comment:考核表状态;type:varchar(255);size:255;"`
      //
      User model.SysUser `json:"user" gorm:"foreignKey:id;References:employee_id;AssociationForeignKey:EmployeeId;"`
      EmployeeId  uint `json:"employeeId" gorm:"column:employee_id;comment:员工ID关联User表"`

      PRItems  []PerformanceReviewItem  `json:"pRItems" gorm:"foreignKey:PRId;References:ID;AssociationForeignKey:ID;"`
      StartDate  time.Time `json:"startDate" form:"startDate" gorm:"column:startDate;comment:开始日期;type:datetime;"`
      EndingDate  time.Time `json:"endingDate" form:"endingDate" gorm:"column:endingDate;comment:结束日期;type:datetime;"`
      //
      Evaluation Evaluation `json:"evaluation" gorm:"foreignKey:id;References:evaluation_id;AssociationForeignKey:EvaluationId;"`
      EvaluationId  uint `json:"evaluationId" form:"evaluationId" gorm:"column:evaluation_id;comment:关联EvaluationKpi表"`

}


func (PerformanceReview) TableName() string {
  return "performanceReview"
}
