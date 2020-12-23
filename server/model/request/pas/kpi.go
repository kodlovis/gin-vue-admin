package pas

import (
    "gin-vue-admin/model/pas"
    "gin-vue-admin/model"
)
type KpiSearch struct{
    pas.Kpi
    Kpis []pas.Kpi
    ID uint
    PageInfo
}

type AddKpiEvaluationInfo struct{
    Kpis []pas.Kpi
    ID uint
    KpiScore []float64
    Users []model.SysUser
    NickName []string
}

