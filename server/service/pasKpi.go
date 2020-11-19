package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePasKpi
//@description: 创建PasKpi记录
//@param: pasKpi model.PasKpi
//@return: err error

func CreatePasKpi(pasKpi model.PasKpi) (err error) {
	err = global.GVA_DB.Create(&pasKpi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasKpi
//@description: 删除PasKpi记录
//@param: pasKpi model.PasKpi
//@return: err error

func DeletePasKpi(pasKpi model.PasKpi) (err error) {
	err = global.GVA_DB.Delete(pasKpi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasKpiByIds
//@description: 批量删除PasKpi记录
//@param: ids request.IdsReq
//@return: err error

func DeletePasKpiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PasKpi{},"id in ?",ids.Ids).Error
	return err
}

//樊新增，记录：批量删除括号中的是根据id去删除，批量更新的语句应该不同。待研习
// func UpdatePasKpiByIds(ids request.IdsReq) (err error) {
// 	err = global.GVA_DB.Save(&[]model.PasKpi{},"id in ?",ids.Ids).Error
// 	return err
// }

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePasKpi
//@description: 更新PasKpi记录
//@param: pasKpi *model.PasKpi
//@return: err error

func UpdatePasKpi(pasKpi *model.PasKpi) (err error) {
	err = global.GVA_DB.Save(pasKpi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasKpi
//@description: 根据id获取PasKpi记录
//@param: id uint
//@return: err error, pasKpi model.PasKpi

func GetPasKpi(id uint) (err error, pasKpi model.PasKpi) {
	err = global.GVA_DB.Where("id = ?", id).First(&pasKpi).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasKpiInfoList
//@description: 分页获取PasKpi记录
//@param: info request.PasKpiSearch
//@return: err error, list interface{}, total int64

func GetPasKpiInfoList(info request.PasKpiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PasKpi{})
    var pasKpis []model.PasKpi
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pasKpis).Error
	return err, pasKpis, total
}