package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTag
//@description: 创建Tag记录
//@param: Tag model.Tag
//@return: err error

func CreateTag(Tag mp.Tag) (err error) {
	err = global.GVA_DB.Create(&Tag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTag
//@description: 删除Tag记录
//@param: Tag model.Tag
//@return: err error

func DeleteTag(Tag mp.Tag) (err error) {
	err = global.GVA_DB.Delete(Tag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTagByIds
//@description: 批量删除Tag记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTagByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Tag{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTag
//@description: 更新Tag记录
//@param: Tag *model.Tag
//@return: err error

func UpdateTag(Tag *mp.Tag) (err error) {
	err = global.GVA_DB.Save(Tag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTag
//@description: 根据id获取Tag记录
//@param: id uint
//@return: err error, Tag model.Tag

func GetTag(id uint) (err error, Tag mp.Tag) {
	err = global.GVA_DB.Where("id = ?", id).First(&Tag).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTagInfoList
//@description: 分页获取Tag记录
//@param: info request.TagSearch
//@return: err error, list interface{}, total int64

func GetTagInfoList(info rp.TagSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Tag{})
    var Tags []mp.Tag
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Tags).Error
	return err, Tags, total
}