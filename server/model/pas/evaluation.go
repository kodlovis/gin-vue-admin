// 自动生成模板Evaluation
package pas

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Evaluation struct {
      global.GVA_MODEL
      Name  string `json:"Name" form:"Name" gorm:"column:Name;comment:方案名称;type:varchar(255);size:255;"`
      Category  string `json:"Category" form:"Category" gorm:"column:Category;comment:方案类型;type:varchar(255);size:255;"`
      Status  bool `json:"Status" form:"Status" gorm:"column:Status;comment:方案状态;type:varchar(255);size:255;"`
      Description  string `json:"Description" form:"Description" gorm:"column:Description;comment:方案描述;type:varchar(255);size:255;"`
      Score  float64 `json:"Score" form:"Score" gorm:"column:Score;comment:方案总分;type:float;"`
}


func (Evaluation) TableName() string {
  return "evaluation"
}
