package request

import "gin-vue-admin/model"

type PasKpiSearch struct{
    model.PasKpi
    PageInfo
}