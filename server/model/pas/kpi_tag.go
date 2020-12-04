// 自动生成模板KpiTag
package pas

// 如果含有time.Time 请自行import time包
type KpiTag struct {
      KpiId  int `json:"kpiId" form:"kpiId" gorm:"column:kpiId;comment:指标ID;type:int;size:10;"`
      TagId  int `json:"tagId" form:"tagId" gorm:"column:tagId;comment:标签ID;type:int;size:10;"`
}


func (KpiTag) TableName() string {
  return "kpi_tag"
}
