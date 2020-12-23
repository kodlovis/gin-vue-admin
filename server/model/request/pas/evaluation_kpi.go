package pas

import "gin-vue-admin/model/pas"
import "gin-vue-admin/model"

type EvaluationKpiSearch struct{
    pas.EvaluationKpi
    PageInfo
}

type AddUserEvaluationKpiInfo struct{
    ID uint
    Users []model.SysUser
    NickName []string
}