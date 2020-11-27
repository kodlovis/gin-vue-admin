// 自动生成模板PasEvaluation
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type PasEvaluation struct {
      global.GVA_MODEL
      Name  string `json:"Name" form:"Name" gorm:"column:Name;comment:;type:varchar(255);size:255;"`
      Category  string `json:"Category" form:"Category" gorm:"column:Category;comment:;type:varchar(255);size:255;"`
      Status  string `json:"Status" form:"Status" gorm:"column:Status;comment:;type:varchar(255);size:255;"`
      Description  string `json:"Description" form:"Description" gorm:"column:Description;comment:;type:varchar(255);size:255;"`
      Score  string `json:"Score" form:"Score" gorm:"column:Score;comment:;type:varchar(255);size:255;"`
}


func (PasEvaluation) TableName() string {
  return "pasEvaluation"
}
