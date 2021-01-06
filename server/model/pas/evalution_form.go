// 自动生成模板EvaluationForm
package pas

import (
  "gin-vue-admin/global"
  "time"
)

// 如果含有time.Time 请自行import time包
type EvaluationForm struct {
      global.GVA_MODEL
      AllocationId  int `json:"allocationId" form:"allocationId" gorm:"column:AllocationId;comment:考核分配ID;type:int;size:10;"`
      Score  float64 `json:"score" form:"score" gorm:"column:Score;comment:得分;type:int;size:10;"`
      ScoreDate  time.Time `json:"scoreDate" form:"scoreDate" gorm:"column:ScoreDate;comment:考核日期;type:datetime;"`
      //Allocation Allocation
}


func (EvaluationForm) TableName() string {
  return "evaluationForm"
}
