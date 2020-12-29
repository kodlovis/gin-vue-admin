package pas

import "gin-vue-admin/model/pas"
import "gin-vue-admin/model"

type EvaluationKpiSearch struct{
    pas.EvaluationKpi
    PageInfo
}

type AssignedKpiEvaluationInfo struct{
    ID uint `json:"id"`
    Users []model.SysUser   `json:"users"`
    Kpis []pas.Kpi
    KpiScore []float64
}