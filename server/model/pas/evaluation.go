// 自动生成模板Evaluation
package pas

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Evaluation struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:方案名称;type:varchar(255);size:255;"`
      Category  string `json:"category" form:"category" gorm:"column:category;comment:方案类型;type:varchar(255);size:255;"`
      Status  bool `json:"status" form:"status" gorm:"column:status;comment:方案状态;type:varchar(255);size:255;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:方案描述;type:varchar(255);size:255;"`
      Score  float64 `json:"score" form:"score" gorm:"column:score;comment:方案总分;type:float;"`
      Kpis []Kpi `gorm:"many2many:evaluation_kpi;"`
      //EvaluationKpis  []EvaluationKpi  `json:"evaluationKpis" gorm:"ForeignKey:EvaluationId;AssociationForeignKey:ID"`
}


func (Evaluation) TableName() string {
  return "evaluation"
}
