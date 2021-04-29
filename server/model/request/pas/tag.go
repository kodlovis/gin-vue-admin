package pas

import "gin-vue-admin/model/pas"

type TagSearch struct{
    pas.Tag
    PageInfo
}