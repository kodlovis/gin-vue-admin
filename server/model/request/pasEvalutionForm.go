package request

import "gin-vue-admin/model"

type PasEvalutionFormSearch struct{
    model.PasEvalutionForm
    PageInfo
}