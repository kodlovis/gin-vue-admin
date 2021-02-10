package pas

import (
	"gin-vue-admin/global"
	//"gin-vue-admin/model/pas"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
	"fmt"
)

func GetLastPRICreatePRIU(ekuid uint,prid uint) (err error) {
	var PRI []mp.PerformanceReviewItem
	var EK []mp.EvaluationKpi
	err = global.GVA_DB.Where("pr_id = ?", prid).Find(&PRI).Error
	err = global.GVA_DB.Where("evaluation_id = ?", ekuid).Find(&EK).Error
	var EKU []mp.EvaluationKpiUser
	for i := 0; i < len(EK); i++ {
		err = global.GVA_DB.Where("ek_id = ?", EK[i].ID).Find(&EKU).Error
		for j := 0; j < len(EKU); j++ {
			err = global.GVA_DB.Create(&mp.PerformanceReviewItemUser{
				PRIID:PRI[i].ID,
				Score:EKU[j].Score,
				UserId:EKU[j].UserId,
				Status:100,
			}).Error
		}
	}
	return err
}
func GetPRIUByPRIID(PRIID uint, info rp.PerformanceReviewItemUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItemUser{}).Where("pri_id = ?", PRIID)
	var PerformanceReviewItemUser []mp.PerformanceReviewItemUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItemUser).Error
	err = db.Preload("PRI").Preload("User").Find(&PerformanceReviewItemUser).Error
	return err, PerformanceReviewItemUser, total
}

func UpdatePRIU(ID uint,status uint,result float64,uid uint,score float64) (err error) {
	//err = global.GVA_DB.Save(PerformanceReviewItemUser).Error
	var PRIU mp.PerformanceReviewItemUser
	err = global.GVA_DB.Model(&PRIU).Where("id = ?",ID).Select("result", "status","user_id","score").Updates(map[string]interface{}{"result": result, "status": status,"user_id":uid,"score":score}).Error
	return err
}

func RemovePRIU(ID uint,priid uint) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItemUser{}, "id = ?", ID).Error
	var PerformanceReviewItemUser mp.PerformanceReviewItemUser
	var total int64
	err = global.GVA_DB.Model(&PerformanceReviewItemUser).Where("pri_id = ?", priid).Count(&total).Error
	var num = float64(1/float64(total))
	var score = fmt.Sprintf("%.5f", num)
	err = global.GVA_DB.Model(&PerformanceReviewItemUser).Where("pri_id = ?", priid).Update("score",score).Error
	return err
}
func CreatePRIU(info []mp.PerformanceReviewItemUser) (err error) {
	for i := 0; i < len(info); i++ {
		err = global.GVA_DB.Create(&info[i]).Error
	}
	var PerformanceReviewItemUser mp.PerformanceReviewItemUser
	var total int64
	err = global.GVA_DB.Model(&PerformanceReviewItemUser).Where("pri_id = ?", info[0].PRIID).Count(&total).Error
	var num = float64(1/float64(total))
	var score = fmt.Sprintf("%.5f", num)
	err = global.GVA_DB.Model(&PerformanceReviewItemUser).Where("pri_id = ?", info[0].PRIID).Update("score",score).Error
	return err
}