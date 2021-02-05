package pas

import (
	//"gin-vue-admin/global"
	"gin-vue-admin/global"
	//"gin-vue-admin/model"
	mp "gin-vue-admin/model/pas"
	//rp "gin-vue-admin/model/request/pas"
)
func CreateEKU(EvaluationKpiUser mp.EvaluationKpiUser) (err error) {
	err = global.GVA_DB.Create(&EvaluationKpiUser).Error
	return err
}
