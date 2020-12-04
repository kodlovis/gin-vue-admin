// 自动生成模板Tag
package pas

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Tag struct {
      global.GVA_MODEL
      Name  string `json:"Name" form:"Name" gorm:"column:Name;comment:标签名称;type:varchar(255);size:255;"`
      Category  string `json:"Category" form:"Category" gorm:"column:Category;comment:标签分类;type:varchar(255);size:255;"`
      Parentid  int `json:"Parentid" form:"Parentid" gorm:"column:Parentid;comment:树形结构;type:int;size:10;"`
      Kpis  []Kpi `gorm:"many2many:kpi_tag;"`
}


func (Tag) TableName() string {
  return "tag"
}
