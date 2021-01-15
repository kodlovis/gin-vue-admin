// 自动生成模板KpiTag
package pas
import (
  "gin-vue-admin/model"
)

// 如果含有time.Time 请自行import time包
type EvaluationKpi struct {
      Kpis  []Kpi  `json:"kpis" gorm:"ForeignKey:id;AssociationForeignKey:KpiId;References:kpi_id"`
      ID uint  `gorm:"primarykey"`
      KpiId uint `json:"kpiId" gorm:"column:kpi_id;primarykey;comment:指标ID"`
      EvaluationId uint  `json:"evaluationId"  gorm:"column:evaluation_id;primarykey;comment:方案ID"`
      KpiScore  float64 `json:"kpiScore" form:"kpiScore" gorm:"column:kpi_score;comment:指标分数;type:float;"`
      UserId uint `json:"userId" gorm:"column:user_id;primarykey;comment:用户ID"`
      User  model.SysUser  `json:"user" gorm:"ForeignKey:id;AssociationForeignKey:UserId;References:user_id"`
      Evaluation  Evaluation  `json:"evaluation" gorm:"ForeignKey:id;AssociationForeignKey:EvaluationId;References:evaluation_id"`
    }


func (EvaluationKpi) TableName() string {
  return "evaluation_kpi"
}
