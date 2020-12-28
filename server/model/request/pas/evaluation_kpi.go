package pas

import "gin-vue-admin/model/pas"
import "gin-vue-admin/model"

type EvaluationKpiSearch struct{
    pas.EvaluationKpi
    PageInfo
}

type AssignedKpiEvaluationInfo struct{
    ID uint
    Users []model.SysUser
    Ids []uint
    Kpis []pas.Kpi
    KpiScore []float64
}