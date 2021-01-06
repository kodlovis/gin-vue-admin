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
      KpiScore  float64 `json:"kpiScore" form:"kpiScore" gorm:"column:KpiScore;comment:指标分数;type:float;"`
      Users []model.SysUser  `gorm:"many2many:evaluationKpi_user;foreignKey:ID;References:ID"`
    }


func (EvaluationKpi) TableName() string {
  return "evaluation_kpi"
}
