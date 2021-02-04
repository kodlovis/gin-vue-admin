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
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{}, "pr_id = ?", PerformanceReviewItem.PRId).Error
	return err
}
func DeletePRItemById(PerformanceReviewItem mp.PerformanceReviewItem) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{}, "id = ?", PerformanceReviewItem.ID).Error
	return err
}

func DeletePerformanceReviewItemByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{}, "pr_id in ?", ids.Ids).Error
	return err
}

func GetPerformanceReviewItemListById(id uint, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id = ?", id)
	var PerformanceReviewItems []mp.PerformanceReviewItem
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Preload("Kpi.Tags").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}

func UpdatePerformanceReviewItemByInfo(id uint, score float64, uid uint, status uint) (err error) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("id = ?", id).Select("score", "user_id", "status").Updates(map[string]interface{}{"score": score, "user_id": uid, "status": status}).Error
	return err
}

func GetPRItemListByUser(id uint, status uint, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("user_id = ? AND status = ?", id, status)
	var PerformanceReviewItems []mp.PerformanceReviewItem
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Preload("PRs.User").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}

func UpdatePRItemStatusById(id uint, status uint, result float64, comment string) (err error) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("id = ?", id).Select("result", "status", "comment").Updates(map[string]interface{}{"result": result, "status": status, "comment": comment}).Error
	return err
}

func UpdatePRItemStatusByPrId(id uint, status uint) (err error) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("pr_id = ?", id).Update("status", status).Error
	return err
}

func GetPRItemCount(prid uint, status uint) (err error, count float64) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	var total = int64(0)
	var num = int64(0)
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("pr_id = ? AND status = ?", prid, status).Count(&num).Error
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("pr_id = ?", prid).Count(&total).Error
	count = float64(num) / float64(total)
	return err, count
}
func UpdatePRItemStatysByIds(Ids []int, status uint) (err error) {
	for i := 0; i < len(Ids); i++ {
		err = global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id = ?", Ids[i]).Update("status", status).Error
	}
	return err
}
func GetPRItemListByStatusPrid(status uint, Ids []int, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id in ? AND status = ?", Ids, status)
	var PerformanceReviewItems []mp.PerformanceReviewItem
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Preload("PRs.User").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}
func GetPRItemListByStatus(status uint, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("Status = ?", status)
	var PerformanceReviewItems []mp.PerformanceReviewItem
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Preload("PRs.User").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}
func GetPRItemListByPrids(ids []int, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id in ?", ids)
	var PerformanceReviewItems []mp.PerformanceReviewItem
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Preload("PRs.User").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}
