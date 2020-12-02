// 自动生成模板Allocation
package pas

import (
	"gin-vue-admin/global"
)
import (
  "time"
)

// 如果含有time.Time 请自行import time包
type Allocation struct {
      global.GVA_MODEL
      EvaluationId  int `json:"EvaluationId" form:"EvaluationId" gorm:"column:EvaluationId;comment:方案ID;type:varchar(255);size:255;"`
      UID  int `json:"UID" form:"UID" gorm:"column:UID;comment:用户ID;type:varchar(255);size:255;"`
      Scorer  string `json:"Scorer" form:"Scorer" gorm:"column:Scorer;comment:评分人;type:varchar(255);size:255;"`
      OriginalScore  float64 `json:"OriginalScore" form:"OriginalScore" gorm:"column:OriginalScore;comment:原始分数;type:varchar(255);size:255;"`
      AdjustedScore  float64 `json:"AdjustedScore" form:"AdjustedScore" gorm:"column:AdjustedScore;comment:修改后的分数;type:varchar(255);size:255;"`
      StartDate  time.Time `json:"StartDate" form:"StartDate" gorm:"column:StartDate;comment:开始日期;type:datetime;"`
      EndingDate  time.Time `json:"EndingDate" form:"EndingDate" gorm:"column:EndingDate;comment:结束日期;type:datetime;"`
}


func (Allocation) TableName() string {
  return "allocation"
}
