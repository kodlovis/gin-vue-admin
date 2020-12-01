package pas

import "gin-vue-admin/model/pas"

type AllocationSearch struct{
    pas.Allocation
    PageInfo
}