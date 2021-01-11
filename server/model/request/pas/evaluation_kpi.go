package pas

import "gin-vue-admin/model/pas"
import "gin-vue-admin/model"

type EvaluationKpiSearch struct{
    pas.EvaluationKpi
    PageInfo
}

type AssignedKpiEvaluationInfo struct{
    ID uint `json:"id"`
    User model.SysUser   `json:"user"`
    Kpis []pas.Kpi  `json:"kpis"`
    KpiScore float64  `json:"score"`
}