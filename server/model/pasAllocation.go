// 自动生成模板PasAllocation
package model

import (
	"gin-vue-admin/global"
)
import (
  "time"
)

// 如果含有time.Time 请自行import time包
type PasAllocation struct {
      global.GVA_MODEL
      EvaluationId  string `json:"EvaluationId" form:"EvaluationId" gorm:"column:EvaluationId;comment:;type:varchar(255);size:255;"`
      UID  string `json:"UID" form:"UID" gorm:"column:UID;comment:;type:varchar(255);size:255;"`
      Scorer  string `json:"Scorer" form:"Scorer" gorm:"column:Scorer;comment:;type:varchar(255);size:255;"`
      OriginalScore  string `json:"OriginalScore" form:"OriginalScore" gorm:"column:OriginalScore;comment:;type:varchar(255);size:255;"`
      AdjustedScore  string `json:"AdjustedScore" form:"AdjustedScore" gorm:"column:AdjustedScore;comment:;type:varchar(255);size:255;"`
      StartDate  time.Time `json:"StartDate" form:"StartDate" gorm:"column:StartDate;comment:;type:datetime;"`
      EndingDate  time.Time `json:"EndingDate" form:"EndingDate" gorm:"column:EndingDate;comment:;type:datetime;"`
}


func (PasAllocation) TableName() string {
  return "pasAllocation"
}
