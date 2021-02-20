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
	var PRI []mp.PerformanceReviewItem
	err = global.GVA_DB.Where("pr_id = ?", PerformanceReviewItem.PRId).Find(&PRI).Error
	for i := 0; i < len(PRI); i++ {
		err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItemUser{}, "pri_id = ?", PRI[i].ID).Error
	}
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{}, "pr_id = ?", PerformanceReviewItem.PRId).Error
	return err
}
func DeletePRItemById(PerformanceReviewItem mp.PerformanceReviewItem) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItem{}, "id = ?", PerformanceReviewItem.ID).Error
	err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItemUser{}, "pri_id = ?", PerformanceReviewItem.ID).Error
	return err
}

func DeletePerformanceReviewItemByIds(ids rp.IdsReq) (err error) {
	var PRI []mp.PerformanceReviewItem
	for i := 0; i < len(ids.Ids); i++ {
		err = global.GVA_DB.Where("pr_id = ?", ids.Ids[i]).Find(&PRI).Error
		for j := 0; j < len(PRI); j++ {
			err = global.GVA_DB.Delete(&[]mp.PerformanceReviewItemUser{}, "pri_id = ?", PRI[j].ID).Error
		}
		
	}
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
	err = db.Preload("PRIUs.User").Preload("Kpi").Preload("Kpi.Tags").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}

func UpdatePerformanceReviewItemByInfo(id uint, score float64, status uint) (err error) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("id = ?", id).Select("score", "status").Updates(map[string]interface{}{"score": score,  "status": status}).Error
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
	err = db.Preload("Kpi").Preload("PRs.User").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}

func UpdatePRItemStatusById(id uint, status uint, result float64, comment string) (err error) {
	var PRIU mp.PerformanceReviewItemUser
	err = global.GVA_DB.Model(&PRIU).Where("id = ?", id).Select("result", "status","comment").Updates(map[string]interface{}{"result": result, "status": status,"comment":comment}).Error
	
	var PRI mp.PerformanceReviewItem
	var total = int64(0)
	var num = int64(0)
	err = global.GVA_DB.Where("id = ?", id).Find(&PRIU).Error
	err = global.GVA_DB.Model(&PRIU).Where("pri_id = ? AND status = ?", PRIU.PRIID, status).Count(&num).Error
	err = global.GVA_DB.Model(&PRIU).Where("pri_id = ?",PRIU.PRIID).Count(&total).Error
	var count = float64(num) / float64(total)
	err = global.GVA_DB.Where("id = ?",PRIU.PRIID).First(&PRI).Error
	result=float64(result)*float64(PRIU.Score)+PRI.Result
	if count ==1  {
		err = global.GVA_DB.Model(&PRI).Where("id = ?", PRIU.PRIID).Select("result", "status").Updates(map[string]interface{}{"result": result, "status": status}).Error
	}else{
		err = global.GVA_DB.Model(&PRI).Where("id = ?", PRIU.PRIID).Select("result").Updates(map[string]interface{}{"result": result}).Error
	}
	return err
}

func UpdatePRItemStatusByPrId(id uint, status uint) (err error) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	var PRIs []mp.PerformanceReviewItem
	err = global.GVA_DB.Model(&PerformanceReviewItem).Where("pr_id = ?", id).Update("status", status).Error
	err = global.GVA_DB.Where("pr_id = ?",id).Find(&PRIs).Error
	var PRIU mp.PerformanceReviewItemUser
	for i := 0; i < len(PRIs) ;i++ {
		err = global.GVA_DB.Model(&PRIU).Where("pri_id = ?", PRIs[i].ID).Update("status", status).Error
	}
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
	var PRIs []mp.PerformanceReviewItem
	for i := 0; i < len(Ids); i++ {
		err = global.GVA_DB.Model(&mp.PerformanceReviewItem{}).Where("pr_id = ?", Ids[i]).Update("status", status).Error
		err = global.GVA_DB.Where("pr_id = ?", Ids[i]).Find(&PRIs).Error
		for j := 0; j < len(PRIs); j++ {
			err = global.GVA_DB.Model(&mp.PerformanceReviewItemUser{}).Where("pri_id = ?", PRIs[j].ID).Update("status", status).Error
		}
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
	err = db.Preload("Kpi").Preload("PRs.User").Preload("PRIUs.User").Find(&PerformanceReviewItems).Error
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
	err = db.Preload("Kpi").Preload("PRs.User").Preload("PRIUs.User").Find(&PerformanceReviewItems).Error
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
	err = db.Preload("Kpi").Preload("PRs.User").Preload("PRIUs.User").Find(&PerformanceReviewItems).Error
	return err, PerformanceReviewItems, total
}
func GetLastPRI() (err error, PerformanceReviewItem mp.PerformanceReviewItem) {
	err = global.GVA_DB.Last(&PerformanceReviewItem).Error
	return
}

func UpdatePRI(PRI *mp.PerformanceReviewItem) (err error) {
	err = global.GVA_DB.Save(PRI).Error
	return err
}