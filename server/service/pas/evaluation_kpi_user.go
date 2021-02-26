package pas

import (
	//"fmt"
	//"gin-vue-admin/global"
	"gin-vue-admin/global"
	//"gin-vue-admin/model"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)
func CreateEKU(info []mp.EvaluationKpiUser) (err error) {
	for i := 0; i < len(info); i++ {
		err = global.GVA_DB.Create(&info[i]).Error
	}
	var EvaluationKpiUsers mp.EvaluationKpiUser
	var EvaluationKpi mp.EvaluationKpi
	var total int64
	err = global.GVA_DB.Model(&EvaluationKpiUsers).Where("ek_id = ?", info[0].EKID).Count(&total).Error
	err = global.GVA_DB.Where("ek_id = ?", info[0].EKID).Find(&EvaluationKpiUsers).Error
	err = global.GVA_DB.Where("id = ?", EvaluationKpiUsers.EKID).Find(&EvaluationKpi).Error
	// var num = float64(total)/float64(EvaluationKpi.KpiScore)
	// var score = fmt.Sprintf("%.5f", num)
	err = global.GVA_DB.Model(&mp.EvaluationKpiUser{}).Where("ek_id = ?", info[0].EKID).Update("score",EvaluationKpi.KpiScore).Error
	return err
}

func GetEKUByEKID(info rp.EvaluationKpiUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.EvaluationKpiUser{}).Where("ek_id = ?", info.EKID)
	var EvaluationKpiUsers []mp.EvaluationKpiUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&EvaluationKpiUsers).Error
	err = db.Preload("User").Find(&EvaluationKpiUsers).Error
	return err, EvaluationKpiUsers, total
}

func UpdateEKU(EvaluationKpiUser *mp.EvaluationKpiUser) (err error) {
	err = global.GVA_DB.Save(EvaluationKpiUser).Error
	return err
}

func RemoveEKU(ID uint,ekid uint) (err error) {
	err = global.GVA_DB.Delete(&[]mp.EvaluationKpiUser{}, "id = ?", ID).Error
	var EvaluationKpiUsers mp.EvaluationKpiUser
	var EvaluationKpi mp.EvaluationKpi
	var total int64
	err = global.GVA_DB.Model(&EvaluationKpiUsers).Where("ek_id = ?", ekid).Count(&total).Error
	err = global.GVA_DB.Where("ek_id = ?", ekid).Find(&EvaluationKpiUsers).Error
	err = global.GVA_DB.Where("id = ?", EvaluationKpiUsers.EKID).Find(&EvaluationKpi).Error
	// var num = float64(total)/float64(EvaluationKpi.KpiScore)
	// var score = fmt.Sprintf("%.5f", num)
	err = global.GVA_DB.Model(&mp.EvaluationKpiUser{}).Where("ek_id = ?", ekid).Update("score",EvaluationKpi.KpiScore).Error
	return err
}