// 自动生成模板KpiTag
package pas
import (
  "gin-vue-admin/model"
)

// 如果含有time.Time 请自行import time包
type EvaluationKpi struct {
      Kpis  []Kpi  `json:"Kpis" gorm:"ForeignKey:id;References:kpi_id"`
      ID uint  `gorm:"primarykey"`
      KpiId uint `json:"KpiId" gorm:"column:kpi_id;primarykey"`
      EvaluationId uint  `json:"EvaluationId"  gorm:"column:evaluation_id;primarykey"`
      KpiScore  float64 `json:"KpiScore" form:"KpiScore" gorm:"column:KpiScore;comment:指标分数;type:float;"`
      Users []model.SysUser  `gorm:"many2many:evaluationKpi_user;foreignKey:ID;References:ID"`
    }


func (EvaluationKpi) TableName() string {
  return "evaluation_kpi"
}
