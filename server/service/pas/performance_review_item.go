package pas

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/pas"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

func CreatePerformanceReviewItem(list []pas.PerformanceReviewItem) (err error) {
	for i := 0; i < len(list); i++ {
		err = global.GVA_DB.Create(&list[i]).Error
	}
	return err
}

func DeletePerformanceReviewItem(PerformanceReviewItem mp.PerformanceReviewItem) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{},"pr_id = ?",PerformanceReviewItem.PRId).Error
	return err
}

func DeletePerformanceReviewItemByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{},"pr_id in ?",ids.Ids).Error
	return err
}