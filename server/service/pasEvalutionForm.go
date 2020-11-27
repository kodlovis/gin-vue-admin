package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePasEvalutionForm
//@description: 创建PasEvalutionForm记录
//@param: pasEvalutionForm model.PasEvalutionForm
//@return: err error

func CreatePasEvalutionForm(pasEvalutionForm model.PasEvalutionForm) (err error) {
	err = global.GVA_DB.Create(&pasEvalutionForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasEvalutionForm
//@description: 删除PasEvalutionForm记录
//@param: pasEvalutionForm model.PasEvalutionForm
//@return: err error

func DeletePasEvalutionForm(pasEvalutionForm model.PasEvalutionForm) (err error) {
	err = global.GVA_DB.Delete(pasEvalutionForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasEvalutionFormByIds
//@description: 批量删除PasEvalutionForm记录
//@param: ids request.IdsReq
//@return: err error

func DeletePasEvalutionFormByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PasEvalutionForm{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePasEvalutionForm
//@description: 更新PasEvalutionForm记录
//@param: pasEvalutionForm *model.PasEvalutionForm
//@return: err error

func UpdatePasEvalutionForm(pasEvalutionForm *model.PasEvalutionForm) (err error) {
	err = global.GVA_DB.Save(pasEvalutionForm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasEvalutionForm
//@description: 根据id获取PasEvalutionForm记录
//@param: id uint
//@return: err error, pasEvalutionForm model.PasEvalutionForm

func GetPasEvalutionForm(id uint) (err error, pasEvalutionForm model.PasEvalutionForm) {
	err = global.GVA_DB.Where("id = ?", id).First(&pasEvalutionForm).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasEvalutionFormInfoList
//@description: 分页获取PasEvalutionForm记录
//@param: info request.PasEvalutionFormSearch
//@return: err error, list interface{}, total int64

func GetPasEvalutionFormInfoList(info request.PasEvalutionFormSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PasEvalutionForm{})
    var pasEvalutionForms []model.PasEvalutionForm
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pasEvalutionForms).Error
	return err, pasEvalutionForms, total
}