package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateEvalutionForm
//@description: 创建EvalutionForm记录
//@param: EvalutionForm model.EvalutionForm
//@return: err error

func CreateEvalutionForm(EvalutionForm mp.EvalutionForm) (err error) {
	err = global.GVA_DB.Create(&EvalutionForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEvalutionForm
//@description: 删除EvalutionForm记录
//@param: EvalutionForm model.EvalutionForm
//@return: err error

func DeleteEvalutionForm(EvalutionForm mp.EvalutionForm) (err error) {
	err = global.GVA_DB.Delete(EvalutionForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEvalutionFormByIds
//@description: 批量删除EvalutionForm记录
//@param: ids request.IdsReq
//@return: err error

func DeleteEvalutionFormByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.EvalutionForm{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateEvalutionForm
//@description: 更新EvalutionForm记录
//@param: EvalutionForm *model.EvalutionForm
//@return: err error

func UpdateEvalutionForm(EvalutionForm *mp.EvalutionForm) (err error) {
	err = global.GVA_DB.Save(EvalutionForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEvalutionForm
//@description: 根据id获取EvalutionForm记录
//@param: id uint
//@return: err error, EvalutionForm model.EvalutionForm

func GetEvalutionForm(id uint) (err error, EvalutionForm mp.EvalutionForm) {
	err = global.GVA_DB.Where("id = ?", id).First(&EvalutionForm).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEvalutionFormInfoList
//@description: 分页获取EvalutionForm记录
//@param: info request.EvalutionFormSearch
//@return: err error, list interface{}, total int64

func GetEvalutionFormInfoList(info rp.EvalutionFormSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.EvalutionForm{})
    var EvalutionForms []mp.EvalutionForm
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&EvalutionForms).Error
	return err, EvalutionForms, total
}