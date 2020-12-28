package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateEvaluationForm
//@description: 创建EvaluationForm记录
//@param: EvaluationForm model.EvaluationForm
//@return: err error

func CreateEvaluationForm(EvaluationForm mp.EvaluationForm) (err error) {
	err = global.GVA_DB.Create(&EvaluationForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEvaluationForm
//@description: 删除EvaluationForm记录
//@param: EvaluationForm model.EvaluationForm
//@return: err error

func DeleteEvaluationForm(EvaluationForm mp.EvaluationForm) (err error) {
	err = global.GVA_DB.Delete(EvaluationForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEvaluationFormByIds
//@description: 批量删除EvaluationForm记录
//@param: ids request.IdsReq
//@return: err error

func DeleteEvaluationFormByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.EvaluationForm{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateEvaluationForm
//@description: 更新EvaluationForm记录
//@param: EvaluationForm *model.EvaluationForm
//@return: err error

func UpdateEvaluationForm(EvaluationForm *mp.EvaluationForm) (err error) {
	err = global.GVA_DB.Save(EvaluationForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEvaluationForm
//@description: 根据id获取EvaluationForm记录
//@param: id uint
//@return: err error, EvaluationForm model.EvaluationForm

func GetEvaluationForm(id uint) (err error, EvaluationForm mp.EvaluationForm) {
	err = global.GVA_DB.Where("id = ?", id).First(&EvaluationForm).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEvaluationFormInfoList
//@description: 分页获取EvaluationForm记录
//@param: info request.EvaluationFormSearch
//@return: err error, list interface{}, total int64

func GetEvaluationFormInfoList(info rp.EvaluationFormSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.EvaluationForm{})
    var EvaluationForms []mp.EvaluationForm
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&EvaluationForms).Error
	return err, EvaluationForms, total
}