// 自动生成模板KpiTag
package pas
import "gin-vue-admin/model"

// 如果含有time.Time 请自行import time包
type EvaluationKpi struct {
      Kpis []Kpi `json:"Kpis" gorm:"ForeignKey:ID;AssociationForeignKey:KpiId"`
      KpiId uint `json:"KpiId" gorm:"primarykey"`
      EvaluationId uint  `json:"EvaluationId"`
      KpiScore  float64 `json:"KpiScore" form:"KpiScore" gorm:"column:KpiScore;comment:指标分数;type:float;"`
      UserId   uint  `json:"UserId" gorm:"column:UserId;comment:用户ID;"`
      Users []model.SysUser  `json:"Users" gorm:"ForeignKey:ID;references:UserId"`
}


func (EvaluationKpi) TableName() string {
  return "evaluation_kpi"
}
