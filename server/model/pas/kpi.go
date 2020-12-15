// 自动生成模板Kpi
package pas

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Kpi struct {
      global.GVA_MODEL
      Name  string `json:"Name" form:"Name" gorm:"column:Name;comment:指标名称;type:varchar(255);size:255;"`
      Description  string `json:"Description" form:"Description" gorm:"column:Description;comment:指标说明;type:varchar(255);size:255;"`
      Status  bool `json:"Status" form:"Status" gorm:"column:Status;comment:指标状态;type:varchar(255);size:255;"`
      Category  string `json:"Category" form:"Category" gorm:"column:Category;comment:指标类型;type:varchar(255);size:255;"`
      Tags  []Tag `gorm:"many2many:kpi_tag;"`
      Evaluations  []Evaluation `gorm:"many2many:evaluation_kpi;"`
}

func (Kpi) TableName() string {
  return "kpi"
}

