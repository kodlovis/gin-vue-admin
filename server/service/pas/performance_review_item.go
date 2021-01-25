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

func GetPerformanceReviewItemListById(id uint, info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64){
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

func UpdatePerformanceReviewItemByInfo(id uint , score float64 ,uid uint,status uint)(err error){
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("id = ?", id).Select("score","user_id","status").Updates(map[string]interface{}{"score": score,"user_id":uid,"status":status}).Error
	return err
}

func GetPRItemListByUser(id uint,status uint,info rp.PerformanceReviewItemSearch) (err error, list interface{}, total int64){
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("user_id = ? AND status = ?", id,status)
	var PerformanceReviewItems []mp.PerformanceReviewItem
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviewItems).Error
	err = db.Preload("User").Preload("Kpi").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}

func UpdatePRItemStatusById(id uint , status uint)(err error){
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("id = ?", id).Update("status", status).Error
	return err
}

func UpdatePRItemStatusByPrId(id uint , status uint)(err error){
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("pr_id = ?", id).Update("status", status).Error
	return err
}

func GetPRItemCount(prid uint , status uint)(err error,count int64){
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("pr_id = ? AND status = ?",prid, status).Count(&count).Error
	return err,count
}
func UpdatePRItemStatysByIds(Ids []int,status uint)(err error){
	for i := 0; i < len(Ids); i++ {
		err = global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id = ?",Ids[i]).Update("status",status).Error
	}
	return err
}