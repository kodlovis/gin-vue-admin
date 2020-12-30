package pas

import (
	//"gin-vue-admin/global"
	"gin-vue-admin/model"
	mp "gin-vue-admin/model/pas"
    rp "gin-vue-admin/model/request/pas"
	"gin-vue-admin/global"
)

func AddUserEvaluationKpi(Users []model.SysUser, ID uint) (err error){
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
	err = db.Preload("Kpis").Preload("Users").Find(&EvaluationKpis).Error
	return err, EvaluationKpis, total
}
// func SetUserEvaluation(evaluation *mp.EvaluationKpi) error {
// 	var s mp.EvaluationKpi
// 	global.GVA_DB.Preload("Users").First(&s, "evaluation_id = ?", evaluation.EvaluationId)
// 	//err := global.GVA_DB.Model(&s).Association("Users").Replace(&evaluation.Users)
// 	return err
// }


func SetUserEvaluation(ID uint,Users []model.SysUser)error{
	var evaluation mp.EvaluationKpi
	evaluation.EvaluationId = ID
	evaluation.Users = Users
	var s mp.EvaluationKpi
	global.GVA_DB.Preload("Users").First(&s, "id = ?", evaluation.EvaluationId)
	err := global.GVA_DB.Model(&s).Association("Users").Replace(&evaluation.Users)
	return err
}

func CreateEvaluationKpi(EvaluationKpi mp.EvaluationKpi) (err error) {
	err = global.GVA_DB.Create(&EvaluationKpi).Error
	return err
}

func DeleteEvaluationKpi(EvaluationKpi mp.EvaluationKpi) (err error) {
	err = global.GVA_DB.Delete(EvaluationKpi).Error
	return err
}