package request

import "gin-vue-admin/model"

type PasAllocationSearch struct{
    model.PasAllocation
    PageInfo
}