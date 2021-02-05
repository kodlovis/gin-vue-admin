package pas

import "gin-vue-admin/model/pas"

type EvaluationKpiSearch struct {
	pas.EvaluationKpi
	PageInfo
}

type AssignedKpiEvaluationInfo struct {
	ID                 uint                    `json:"id"`
	EvaluationKpiUsers []pas.EvaluationKpiUser `json:"evaluationKpiUsers"`
	Kpis               []pas.Kpi               `json:"kpis"`
	KpiScore           float64                 `json:"score"`
}
