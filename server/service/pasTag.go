package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePasTag
//@description: 创建PasTag记录
//@param: pasTag model.PasTag
//@return: err error

func CreatePasTag(pasTag model.PasTag) (err error) {
	err = global.GVA_DB.Create(&pasTag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasTag
//@description: 删除PasTag记录
//@param: pasTag model.PasTag
//@return: err error

func DeletePasTag(pasTag model.PasTag) (err error) {
	err = global.GVA_DB.Delete(pasTag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePasTagByIds
//@description: 批量删除PasTag记录
//@param: ids request.IdsReq
//@return: err error

func DeletePasTagByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PasTag{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePasTag
//@description: 更新PasTag记录
//@param: pasTag *model.PasTag
//@return: err error

func UpdatePasTag(pasTag *model.PasTag) (err error) {
	err = global.GVA_DB.Save(pasTag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasTag
//@description: 根据id获取PasTag记录
//@param: id uint
//@return: err error, pasTag model.PasTag

func GetPasTag(id uint) (err error, pasTag model.PasTag) {
	err = global.GVA_DB.Where("id = ?", id).First(&pasTag).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPasTagInfoList
//@description: 分页获取PasTag记录
//@param: info request.PasTagSearch
//@return: err error, list interface{}, total int64

func GetPasTagInfoList(info request.PasTagSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PasTag{})
    var pasTags []model.PasTag
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pasTags).Error
	return err, pasTags, total
}