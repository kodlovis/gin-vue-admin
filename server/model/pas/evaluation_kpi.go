// 自动生成模板KpiTag
package pas

// 如果含有time.Time 请自行import time包
type EvaluationKpi struct {
      KpiId uint `json:"KpiId" gorm:"primarykey"`
      EvaluationId uint  `json:"EvaluationId"`
      KpiScore  float64 `json:"KpiScore" form:"KpiScore" gorm:"column:KpiScore;comment:指标分数;type:float;"`
}


func (EvaluationKpi) TableName() string {
  return "evaluation_kpi"
}
