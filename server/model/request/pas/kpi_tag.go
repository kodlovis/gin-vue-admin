package pas

import "gin-vue-admin/model/pas"

type KpiTagSearch struct{
    pas.KpiTag
    PageInfo
}

type KpiTagList struct{
	KpiTagList []pas.KpiTag `json:"item"`
}