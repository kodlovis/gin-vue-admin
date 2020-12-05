package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateKpi
//@description: 创建Kpi记录
//@param: Kpi model.Kpi
//@return: err error

func CreateKpi(Kpi mp.Kpi) (err error) {
	err = global.GVA_DB.Create(&Kpi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKpi
//@description: 删除Kpi记录
//@param: Kpi model.Kpi
//@return: err error

func DeleteKpi(Kpi mp.Kpi) (err error) {
	err = global.GVA_DB.Delete(Kpi).Error
	err = global.GVA_DB.Table("kpi_tag").Where("kpi_id = ?",Kpi.ID).Delete(&[]mp.KpiTag{}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKpiByIds
//@description: 批量删除Kpi记录
//@param: ids request.IdsReq
//@return: err error

func DeleteKpiByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Kpi{},"id in ?",ids.Ids).Error
	err = global.GVA_DB.Table("kpi_tag").Where("kpi_id = ?",ids.Ids).Delete(&[]mp.KpiTag{}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateKpi
//@description: 更新Kpi记录
//@param: Kpi *model.Kpi
//@return: err error

func UpdateKpi(Kpi *mp.Kpi) (err error) {
	err = global.GVA_DB.Save(Kpi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKpi
//@description: 根据id获取Kpi记录
//@param: id uint
//@return: err error, Kpi model.Kpi

func GetKpi(id uint) (err error, Kpi mp.Kpi) {
	err = global.GVA_DB.Where("id = ?", id).First(&Kpi).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKpiInfoList
//@description: 分页获取Kpi记录
//@param: info request.KpiSearch
//@return: err error, list interface{}, total int64

func GetKpiInfoList(info rp.KpiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Kpi{})
    var Kpis []mp.Kpi
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Kpis).Error
	err = db.Preload("Tags").Find(&Kpis).Error
	return err, Kpis, total
}
