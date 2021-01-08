package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePerformanceReview
//@description: 创建PerformanceReview记录
//@param: PerformanceReview model.PerformanceReview
//@return: err error

func CreatePerformanceReview(PerformanceReview mp.PerformanceReview) (err error) {
	err = global.GVA_DB.Create(&PerformanceReview).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePerformanceReview
//@description: 删除PerformanceReview记录
//@param: PerformanceReview model.PerformanceReview
//@return: err error

func DeletePerformanceReview(PerformanceReview mp.PerformanceReview) (err error) {
	//err = global.GVA_DB.Delete(PerformanceReview).Error
	err = global.GVA_DB.Delete(&[]mp.PerformanceReview{},"id = ?",PerformanceReview.ID).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePerformanceReviewByIds
//@description: 批量删除PerformanceReview记录
//@param: ids request.IdsReq
//@return: err error

func DeletePerformanceReviewByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReview{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePerformanceReview
//@description: 更新PerformanceReview记录
//@param: PerformanceReview *model.PerformanceReview
//@return: err error

func UpdatePerformanceReview(PerformanceReview *mp.PerformanceReview) (err error) {
	err = global.GVA_DB.Save(PerformanceReview).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPerformanceReview
//@description: 根据id获取PerformanceReview记录
//@param: id uint
//@return: err error, PerformanceReview model.PerformanceReview

func GetPerformanceReview(id uint) (err error, PerformanceReview mp.PerformanceReview) {
	err = global.GVA_DB.Where("id = ?", id).First(&PerformanceReview).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPerformanceReviewInfoList
//@description: 分页获取PerformanceReview记录
//@param: info request.PerformanceReviewSearch
//@return: err error, list interface{}, total int64

func GetPerformanceReviewInfoList(info rp.PerformanceReviewSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReview{})
    var PerformanceReviews []mp.PerformanceReview
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
	err = db.Preload("Evaluation").Preload("User").Preload("PRItem.Kpi").Preload("PRItem.User").Find(&PerformanceReviews).Error
	return err, PerformanceReviews, total
}

func UpdatePerformanceReviewByInfo(PRInfo rp.PerformanceReviewInfo)(err error){
	err = global.GVA_DB.Model(&mp.PerformanceReview{}).Where("id = ?", PRInfo.ID).Updates(mp.PerformanceReview{
		EmployeeId:PRInfo.EmployeeId,
		EvaluationId:PRInfo.EvaluationId,
		Name:PRInfo.Name,
		Status:PRInfo.Status,
		StartDate:PRInfo.StartDate,
		EndingDate:PRInfo.EndingDate,
		}).Error
	return err
}