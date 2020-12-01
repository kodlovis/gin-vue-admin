// 自动生成模板EvalutionForm
package pas

import (
  "gin-vue-admin/global"
  "time"
)

// 如果含有time.Time 请自行import time包
type EvalutionForm struct {
      global.GVA_MODEL
      AllocationId  int `json:"AllocationId" form:"AllocationId" gorm:"column:AllocationId;comment:;type:int;size:10;"`
      Score  int `json:"Score" form:"Score" gorm:"column:Score;comment:;type:int;size:10;"`
      ScoreDate  time.Time `json:"ScoreDate" form:"ScoreDate" gorm:"column:ScoreDate;comment:;type:datetime;"`
      //Allocation Allocation
}


func (EvalutionForm) TableName() string {
  return "EvalutionForm"
}
