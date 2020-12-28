package pas

import (
    "gin-vue-admin/model/pas"
)
type KpiSearch struct{
    pas.Kpi
    Kpis []pas.Kpi
    ID uint
    PageInfo
}


