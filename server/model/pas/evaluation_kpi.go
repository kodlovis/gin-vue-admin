// 自动生成模板KpiTag
package pas

// 如果含有time.Time 请自行import time包
type EvaluationKpi struct {
      KpiId  int `json:"kpiId" form:"kpiId" gorm:"column:kpiId;comment:指标ID;type:int;size:10;"`
      EvaluationId  int `json:"evaluationId" form:"evaluationId" gorm:"column:evaluationId;comment:方案ID;type:int;size:10;"`
}


func (EvaluationKpi) TableName() string {
  return "evaluation_kpi"
}
