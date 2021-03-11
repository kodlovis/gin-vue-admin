package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
	"fmt"
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
	err = global.GVA_DB.Delete(&[]mp.PerformanceReview{}, "id = ?", PerformanceReview.ID).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePerformanceReviewByIds
//@description: 批量删除PerformanceReview记录
//@param: ids request.IdsReq
//@return: err error

func DeletePerformanceReviewByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PerformanceReview{}, "id in ?", ids.Ids).Error
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
	err = global.GVA_DB.Where("id = ?", id).Preload("User").First(&PerformanceReview).Error
	return
}
func GetPRBystatus(status uint, info rp.PerformanceReviewSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&mp.PerformanceReview{}).Where("status = ?", status)
	var PerformanceReviews []mp.PerformanceReview
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
	err = db.Preload("Evaluation").Preload("User").Preload("PRItems").Find(&PerformanceReviews).Error
	return err, PerformanceReviews, total
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
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.PerformanceReview.Status != 0 {
        db = db.Where("`status` = ?",info.PerformanceReview.Status)
    }
    if info.NickName != "" {
		err = db.Count(&total).Error
		err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
		err = db.Preload("Evaluation").Preload("User").Preload("PRItems").Joins("INNER JOIN sys_users AS `user` ON performance_review.employee_id = `user`.id").Find(&PerformanceReviews,"`user`.nick_name LIKE ?","%"+ info.NickName+"%").Error
		return err, PerformanceReviews, total
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
	err = db.Preload("Evaluation").Preload("User").Preload("PRItems").Find(&PerformanceReviews).Error
	return err, PerformanceReviews, total
}
func GetPRListWithoutFinishedStatus(info rp.PerformanceReviewSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PerformanceReview{})
	var PerformanceReviews []mp.PerformanceReview
	// 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.PerformanceReview.Status != 0 {
        db = db.Where("`status` = ?",info.PerformanceReview.Status)
    }
    if info.PerformanceReview.Status == 0 {
        db = db.Where("`status` != ?",9)
    }
    if info.NickName != "" {
		err = db.Count(&total).Error
		err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
		err = db.Preload("Evaluation").Preload("User").Preload("PRItems").Joins("INNER JOIN sys_users AS `user` ON performance_review.employee_id = `user`.id").Find(&PerformanceReviews,"`user`.nick_name LIKE ?","%"+ info.NickName+"%").Error
		return err, PerformanceReviews, total
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
	err = db.Preload("Evaluation").Preload("User").Preload("PRItems").Find(&PerformanceReviews).Error
	return err, PerformanceReviews, total
}

func UpdatePerformanceReviewByInfo(PRInfo rp.PerformanceReviewInfo) (err error) {
	var PRIs []mp.PerformanceReviewItem
	err = global.GVA_DB.Where("pr_id = ?", PRInfo.ID).Find(&PRIs).Error
	var sum float64
	for i := 0; i < len(PRIs); i++ {
		sum=sum+PRIs[i].Score
	}
	err = global.GVA_DB.Model(&mp.PerformanceReview{}).Where("id = ?", PRInfo.ID).Updates(mp.PerformanceReview{
		EmployeeId:   PRInfo.EmployeeId,
		EvaluationId: PRInfo.EvaluationId,
		Name:         PRInfo.Name,
		Status:       PRInfo.Status,
		StartDate:    PRInfo.StartDate,
		EndingDate:   PRInfo.EndingDate,
		Score:        sum,
	}).Error
	return err
}

func GetLastPerformanceReview() (err error, PerformanceReview mp.PerformanceReview) {
	err = global.GVA_DB.Last(&PerformanceReview).Error
	return
}

func UpdatePRStatusById(id uint, status uint, result float64) (err error) {
	err = global.GVA_DB.Model(&mp.PerformanceReview{}).Where("id = ?", id).Update("status", status).Error
	var pr mp.PerformanceReview
	if result==0 {
		err = global.GVA_DB.Model(&pr).Where("id = ?", id).Select("status").Updates(map[string]interface{}{ "status": status}).Error
	}else {
		err = global.GVA_DB.Model(&pr).Where("id = ?", id).Select("result", "status").Updates(map[string]interface{}{"result": result, "status": status}).Error
	}
	return err
}

func UpdatePRStatysByIds(Ids []int, status uint) (err error) {
	for i := 0; i < len(Ids); i++ {
		err = global.GVA_DB.Model(&mp.PerformanceReview{}).Where("id = ?", Ids[i]).Update("status", status).Error
	}
	return err
}

func GetPRByUser(id uint, status uint) (err error, PerformanceReview mp.PerformanceReview) {
	db := global.GVA_DB.Model(&PerformanceReview).Where("employee_id = ? AND status = ?", id, status)
	err = db.Preload("Evaluation").Preload("User").Preload("PRItems").First(&PerformanceReview).Error
	return
}
func GetPRListByUser(id uint, Ids []int, info rp.PerformanceReviewSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	if Ids==nil {
		db := global.GVA_DB.Model(&mp.PerformanceReview{}).Where("employee_id = ?", id)
		var PerformanceReviews []mp.PerformanceReview
		// 如果有条件搜索 下方会自动创建搜索语句
		if info.Name != "" {
			db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
		}
		if info.Status != 0 {
			db = db.Where("`status` = ?",info.Status)
		}
		err = db.Count(&total).Error
		err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
		err = db.Preload("User").Preload("PRItems").Find(&PerformanceReviews).Error
		return err, PerformanceReviews, total
	}else{
		db := global.GVA_DB.Model(&mp.PerformanceReview{}).Where("employee_id = ? AND status in ?", id, Ids)
		var PerformanceReviews []mp.PerformanceReview
		// 如果有条件搜索 下方会自动创建搜索语句
		if info.Name != "" {
			db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
		}
		if info.Status != 0 {
			db = db.Where("`status` = ?",info.Status)
		}
		err = db.Count(&total).Error
		err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
		err = db.Preload("User").Preload("PRItems").Find(&PerformanceReviews).Error
		return err, PerformanceReviews, total
	}
}

func UpdatePRResult(ID uint) (err error) {
	var PRIs []mp.PerformanceReviewItem
	err = global.GVA_DB.Where("pr_id = ?", ID).Find(&PRIs).Error
	var total float64 
	var PR mp.PerformanceReview
	for i := 0; i < len(PRIs); i++ {
		total=total+PRIs[i].Result
	}
	var re=fmt.Sprintf("%.1f", total)
	err = global.GVA_DB.Model(&PR).Where("id = ?", ID).Select("result").Updates(map[string]interface{}{"result": re}).Error
	return err
}
func GetPRByID(id uint, info rp.PerformanceReviewSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&mp.PerformanceReview{}).Where("id = ?", id)
	var PerformanceReviews []mp.PerformanceReview
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PerformanceReviews).Error
	err = db.Preload("User").Find(&PerformanceReviews).Error
	return err, PerformanceReviews, total
}