package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAllocation
//@description: 创建Allocation记录
//@param: Allocation model.Allocation
//@return: err error

func CreateAllocation(Allocation mp.Allocation) (err error) {
	err = global.GVA_DB.Create(&Allocation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAllocation
//@description: 删除Allocation记录
//@param: Allocation model.Allocation
//@return: err error

func DeleteAllocation(Allocation mp.Allocation) (err error) {
	err = global.GVA_DB.Delete(Allocation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAllocationByIds
//@description: 批量删除Allocation记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAllocationByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Allocation{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAllocation
//@description: 更新Allocation记录
//@param: Allocation *model.Allocation
//@return: err error

func UpdateAllocation(Allocation *mp.Allocation) (err error) {
	err = global.GVA_DB.Save(Allocation).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllocation
//@description: 根据id获取Allocation记录
//@param: id uint
//@return: err error, Allocation model.Allocation

func GetAllocation(id uint) (err error, Allocation mp.Allocation) {
	err = global.GVA_DB.Where("id = ?", id).First(&Allocation).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllocationInfoList
//@description: 分页获取Allocation记录
//@param: info request.AllocationSearch
//@return: err error, list interface{}, total int64

func GetAllocationInfoList(info rp.AllocationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Allocation{})
    var Allocations []mp.Allocation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Allocations).Error
	err = db.Preload("Evaluation").Find(&Allocations).Error
	return err, Allocations, total
}