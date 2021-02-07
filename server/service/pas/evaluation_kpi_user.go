package pas

import (
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
	return err
}

func GetEKUByEKID(EKU rp.EKU, info rp.EvaluationKpiUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.EvaluationKpiUser{}).Where("ek_id = ?", EKU.EKID)
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