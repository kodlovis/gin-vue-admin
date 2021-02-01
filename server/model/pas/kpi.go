// 自动生成模板Kpi
package pas

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Kpi struct {
	global.GVA_MODEL
	Name           string        `json:"name" form:"name" gorm:"column:name;comment:指标名称;type:varchar(255);size:255;"`
	Description    string        `json:"description" form:"description" gorm:"column:description;comment:指标说明;type:varchar(255);"`
	Status         uint          `json:"status" form:"status" gorm:"column:status;comment:指标状态;type:varchar(255);size:255;"`
	Category       string        `json:"category" form:"category" gorm:"column:category;comment:指标类型;type:varchar(255);"`
	Tags           []Tag         `gorm:"many2many:kpi_tag;"`
	Evaluations    []Evaluation  `json:"evaluations" gorm:"many2many:evaluation_kpi;"`
	EvaluationKpis EvaluationKpi `json:"evaluationKpis" gorm:"ForeignKey:KpiId;AssociationForeignKey:ID"`
}

func (Kpi) TableName() string {
	return "kpi"
}
