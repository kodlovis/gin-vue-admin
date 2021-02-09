package pas

import (
	"gin-vue-admin/global"
	//"gin-vue-admin/model/pas"
	mp "gin-vue-admin/model/pas"
	//rp "gin-vue-admin/model/request/pas"
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