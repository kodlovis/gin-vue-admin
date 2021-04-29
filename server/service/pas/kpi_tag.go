package pas

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/pas"
	// mp "gin-vue-admin/model/pas"
	// rp "gin-vue-admin/model/request/pas"
)

func CreateKpiTag(list []pas.KpiTag) (err error) {
	for i := 0; i < len(list); i++ {
		err = global.GVA_DB.Create(&list[i]).Error
	}
	return err
}

// func UpdatePerformanceReviewItemByInfo(id uint , score float64 ,uid uint)(err error){
// 	var PerformanceReviewItem mp.PerformanceReviewItem
// 	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("id = ?", id).Select("score","user_id").Updates(map[string]interface{}{"score": score,"user_id":uid}).Error
// 	return err
// }