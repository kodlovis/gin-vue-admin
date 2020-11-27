package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePasAllocation
//@description: 创建PasAllocation记录
//@param: pasAllocation model.PasAllocation
//@return: err error

func CreatePasAllocation(pasAllocation model.PasAllocation) (err error) {
	err = global.GVA_DB.Create(&pasAllocation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasAllocation
//@description: 删除PasAllocation记录
//@param: pasAllocation model.PasAllocation
//@return: err error

func DeletePasAllocation(pasAllocation model.PasAllocation) (err error) {
	err = global.GVA_DB.Delete(pasAllocation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasAllocationByIds
//@description: 批量删除PasAllocation记录
//@param: ids request.IdsReq
//@return: err error

func DeletePasAllocationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PasAllocation{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePasAllocation
//@description: 更新PasAllocation记录
//@param: pasAllocation *model.PasAllocation
//@return: err error

func UpdatePasAllocation(pasAllocation *model.PasAllocation) (err error) {
	err = global.GVA_DB.Save(pasAllocation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasAllocation
//@description: 根据id获取PasAllocation记录
//@param: id uint
//@return: err error, pasAllocation model.PasAllocation

func GetPasAllocation(id uint) (err error, pasAllocation model.PasAllocation) {
	err = global.GVA_DB.Where("id = ?", id).First(&pasAllocation).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasAllocationInfoList
//@description: 分页获取PasAllocation记录
//@param: info request.PasAllocationSearch
//@return: err error, list interface{}, total int64

func GetPasAllocationInfoList(info request.PasAllocationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PasAllocation{})
    var pasAllocations []model.PasAllocation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pasAllocations).Error
	return err, pasAllocations, total
}