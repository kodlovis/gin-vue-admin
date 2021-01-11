package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
)

func CreatePerformanceReviewItem(PerformanceReviewItem mp.PerformanceReviewItem) (err error) {
	err = global.GVA_DB.Create(&PerformanceReviewItem).Error
	return err
}