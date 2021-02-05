package pas

import (
	//"gin-vue-admin/global"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

func AddUserEvaluationKpi(User model.SysUser, ID uint) (err error) {
	var evaluation mp.EvaluationKpi
	evaluation.EvaluationId = ID
	//evaluation.Users = Users
	//err = SetUserEvaluation(&evaluation)
	return err
}

func GetEvaluationKpiInfoList(info rp.EvaluationKpiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.EvaluationKpi{})
	var EvaluationKpis []mp.EvaluationKpi
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&EvaluationKpis).Error
	err = db.Preload("Kpis").Preload("EvaluationKpiUsers").Find(&EvaluationKpis).Error
	return err, EvaluationKpis, total
}

// func SetUserEvaluation(evaluation *mp.EvaluationKpi) error {
// 	var s mp.EvaluationKpi
// 	global.GVA_DB.Preload("Users").First(&s, "evaluation_id = ?", evaluation.EvaluationId)
// 	//err := global.GVA_DB.Model(&s).Association("Users").Replace(&evaluation.Users)
// 	return err
// }

func SetUserEvaluation(ID uint, EvaluationKpiUsers []mp.EvaluationKpiUser) error {
	var evaluation mp.EvaluationKpi
	evaluation.EvaluationId = ID
	evaluation.EvaluationKpiUsers = EvaluationKpiUsers
	var s mp.EvaluationKpi
	global.GVA_DB.Preload("EvaluationKpiUsers").First(&s, "id = ?", evaluation.EvaluationId)
	err := global.GVA_DB.Model(&s).Association("EvaluationKpiUsers").Replace(&evaluation.EvaluationKpiUsers)
	return err
}

func CreateEvaluationKpi(EvaluationKpi mp.EvaluationKpi) (err error) {
	err = global.GVA_DB.Create(&EvaluationKpi).Error
	return err
}

func RemoveEvaluationKpi(ID uint) (err error) {
	err = global.GVA_DB.Delete(&[]mp.EvaluationKpi{}, "id = ?", ID).Error
	return err
}

func RemoveEvaluationKpiByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.EvaluationKpi{}, "id in ?", ids.Ids).Error
	//global.GVA_DB.Delete(&[]mp.EvaluationKpiUser{},"evaluation_kpi_id in ?",ids.Ids)
	return err
}

func GetEvaluationKpiById(id uint, info rp.EvaluationKpiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.EvaluationKpi{}).Where("evaluation_id = ?", id)
	var EvaluationKpis []mp.EvaluationKpi
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&EvaluationKpis).Error
	err = db.Preload("EvaluationKpiUsers").Preload("Kpis").Find(&EvaluationKpis).Error
	return err, EvaluationKpis, total
}

func UpdateEvaluationKpi(EvaluationKpi *mp.EvaluationKpi) (err error) {
	err = global.GVA_DB.Save(EvaluationKpi).Error
	return err
}