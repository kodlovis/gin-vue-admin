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

func GetPerformanceReviewListById(id uint, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64){
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id = ?", id)
	var PerformanceReviewItems []mp.PerformanceReviewItem
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}