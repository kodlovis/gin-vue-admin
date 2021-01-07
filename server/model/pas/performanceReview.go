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
      User model.SysUser  `gorm:"foreignKey:ID;References:EmployeeId"`
      EmployeeId  uint `json:"employee" form:"employee" gorm:"column:employee_id;comment:员工ID关联User表"`

      PRItemId  uint  `json:"pRItemId" form:"pRItemId" gorm:"column:pRItem_id;comment:关联prItem表"`
      StartDate  time.Time `json:"startDate" form:"startDate" gorm:"column:startDate;comment:开始日期;type:datetime;"`
      EndingDate  time.Time `json:"endingDate" form:"endingDate" gorm:"column:endingDate;comment:结束日期;type:datetime;"`
      EvaluationKpi EvaluationKpi `json:"evaluationKpi" gorm:"foreignKey:EvaluationId;references:EvaluationKpiId;comment:方案"`
      EvaluationKpiId  uint `json:"evaluationKpiId" form:"evaluationKpiId" gorm:"column:evaluationKpi_id;comment:关联EvaluationKpi表"`

}


func (PerformanceReview) TableName() string {
  return "performanceReview"
}
