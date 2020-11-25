package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateKpiEvaluation
//@description: 创建KpiEvaluation记录
//@param: kpiEvaluation model.KpiEvaluation
//@return: err error

func CreateKpiEvaluation(kpiEvaluation model.KpiEvaluation) (err error) {
	err = global.GVA_DB.Create(&kpiEvaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKpiEvaluation
//@description: 删除KpiEvaluation记录
//@param: kpiEvaluation model.KpiEvaluation
//@return: err error

func DeleteKpiEvaluation(kpiEvaluation model.KpiEvaluation) (err error) {
	err = global.GVA_DB.Delete(kpiEvaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKpiEvaluationByIds
//@description: 批量删除KpiEvaluation记录
//@param: ids request.IdsReq
//@return: err error

func DeleteKpiEvaluationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.KpiEvaluation{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateKpiEvaluation
//@description: 更新KpiEvaluation记录
//@param: kpiEvaluation *model.KpiEvaluation
//@return: err error

func UpdateKpiEvaluation(kpiEvaluation *model.KpiEvaluation) (err error) {
	err = global.GVA_DB.Save(kpiEvaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKpiEvaluation
//@description: 根据id获取KpiEvaluation记录
//@param: id uint
//@return: err error, kpiEvaluation model.KpiEvaluation

func GetKpiEvaluation(id uint) (err error, kpiEvaluation model.KpiEvaluation) {
	err = global.GVA_DB.Where("id = ?", id).First(&kpiEvaluation).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKpiEvaluationInfoList
//@description: 分页获取KpiEvaluation记录
//@param: info request.KpiEvaluationSearch
//@return: err error, list interface{}, total int64

func GetKpiEvaluationInfoList(info request.KpiEvaluationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.KpiEvaluation{})
    var kpiEvaluations []model.KpiEvaluation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&kpiEvaluations).Error
	return err, kpiEvaluations, total
}