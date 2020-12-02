// 自动生成模板EvalutionForm
package pas

import (
  "gin-vue-admin/global"
  "time"
)

// 如果含有time.Time 请自行import time包
type EvalutionForm struct {
      global.GVA_MODEL
      AllocationId  int `json:"AllocationId" form:"AllocationId" gorm:"column:AllocationId;comment:考核分配ID;type:int;size:10;"`
      Score  float64 `json:"Score" form:"Score" gorm:"column:Score;comment:得分;type:int;size:10;"`
      ScoreDate  time.Time `json:"ScoreDate" form:"ScoreDate" gorm:"column:ScoreDate;comment:考核日期;type:datetime;"`
      //Allocation Allocation
}


func (EvalutionForm) TableName() string {
  return "evalutionForm"
}
