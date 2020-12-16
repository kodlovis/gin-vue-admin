package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateEvaluation
//@description: 创建Evaluation记录
//@param: Evaluation model.Evaluation
//@return: err error

func CreateEvaluation(Evaluation mp.Evaluation) (err error) {
	err = global.GVA_DB.Create(&Evaluation).Error
	return err
}

func SetKpiEvaluation(evaluation *mp.Evaluation) error {
	var s mp.Evaluation
	global.GVA_DB.Preload("Kpis").First(&s, "id = ?", evaluation.ID)
	err := global.GVA_DB.Model(&s).Association("Kpis").Replace(&evaluation.Kpis)
	return err
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEvaluation
//@description: 删除Evaluation记录
//@param: Evaluation model.Evaluation
//@return: err error

func DeleteEvaluation(Evaluation mp.Evaluation) (err error) {
	err = global.GVA_DB.Delete(Evaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEvaluationByIds
//@description: 批量删除Evaluation记录
//@param: ids request.IdsReq
//@return: err error

func DeleteEvaluationByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Evaluation{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateEvaluation
//@description: 更新Evaluation记录
//@param: Evaluation *model.Evaluation
//@return: err error

func UpdateEvaluation(Evaluation *mp.Evaluation) (err error) {
	err = global.GVA_DB.Save(Evaluation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEvaluation
//@description: 根据id获取Evaluation记录
//@param: id uint
//@return: err error, Evaluation model.Evaluation

func GetEvaluation(id uint) (err error, Evaluation mp.Evaluation) {
	err = global.GVA_DB.Where("id = ?", id).First(&Evaluation).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEvaluationInfoList
//@description: 分页获取Evaluation记录
//@param: info request.EvaluationSearch
//@return: err error, list interface{}, total int64

func GetEvaluationInfoList(info rp.EvaluationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Evaluation{})
    var Evaluations []mp.Evaluation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Evaluations).Error
	err = db.Preload("Kpis").Find(&Evaluations).Error
	return err, Evaluations, total
}