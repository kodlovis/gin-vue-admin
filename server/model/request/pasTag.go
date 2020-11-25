package request

import "gin-vue-admin/model"

type PasTagSearch struct{
    model.PasTag
    PageInfo
}