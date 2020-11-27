package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePasEvaluation
//@description: 创建PasEvaluation记录
//@param: pasEvaluation model.PasEvaluation
//@return: err error

func CreatePasEvaluation(pasEvaluation model.PasEvaluation) (err error) {
	err = global.GVA_DB.Create(&pasEvaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasEvaluation
//@description: 删除PasEvaluation记录
//@param: pasEvaluation model.PasEvaluation
//@return: err error

func DeletePasEvaluation(pasEvaluation model.PasEvaluation) (err error) {
	err = global.GVA_DB.Delete(pasEvaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasEvaluationByIds
//@description: 批量删除PasEvaluation记录
//@param: ids request.IdsReq
//@return: err error

func DeletePasEvaluationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PasEvaluation{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePasEvaluation
//@description: 更新PasEvaluation记录
//@param: pasEvaluation *model.PasEvaluation
//@return: err error

func UpdatePasEvaluation(pasEvaluation *model.PasEvaluation) (err error) {
	err = global.GVA_DB.Save(pasEvaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasEvaluation
//@description: 根据id获取PasEvaluation记录
//@param: id uint
//@return: err error, pasEvaluation model.PasEvaluation

func GetPasEvaluation(id uint) (err error, pasEvaluation model.PasEvaluation) {
	err = global.GVA_DB.Where("id = ?", id).First(&pasEvaluation).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasEvaluationInfoList
//@description: 分页获取PasEvaluation记录
//@param: info request.PasEvaluationSearch
//@return: err error, list interface{}, total int64

func GetPasEvaluationInfoList(info request.PasEvaluationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PasEvaluation{})
    var pasEvaluations []model.PasEvaluation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pasEvaluations).Error
	return err, pasEvaluations, total
}