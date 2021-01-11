package pas

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/pas"
)

func CreatePerformanceReviewItem(list []pas.PerformanceReviewItem) (err error) {
	for i := 0; i < len(list); i++ {
		err = global.GVA_DB.Create(&list[i]).Error
	}
	return err
}