package pas

import (
	"gorm.io/gorm"
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
func RemoveKpiTags(Kpi mp.Kpi) (err error) {
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
	err = db.Preload("Tags").Preload("EvaluationKpis","evaluation_id = ?",info.ID).Find(&Kpis).Error
	return err, Kpis, total
}

func GetKpiScoreInfoList(info rp.KpiSearch) (err error, EvaluationKpi []mp.EvaluationKpi) {
	err = global.GVA_DB.Where("kpi_id = ? And evaluation_id = ?", "23", "1").First(&EvaluationKpi).Error
	return 
}

func GetKpiByIds(ids rp.IdsReq,info rp.KpiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Kpi{}).Where("id in ?",ids.Ids)
	var Kpis []mp.Kpi
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Kpis).Error
	return err, Kpis, total
}
func GetKpiScoreByIds(id rp.GetEvaluationId,info rp.KpiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Kpi{})
	var Kpis []mp.Kpi
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Kpis).Error
	err = db.Preload("Tags").Preload("EvaluationKpis.Users").Preload("EvaluationKpis","evaluation_id = ?",id.ID).Find(&Kpis).Error
	return err, Kpis, total
}
func AssignedKpiEvaluation(Kpis []mp.Kpi, ID uint) (err error) {
	var evaluation mp.EvaluationKpi
	evaluation.EvaluationId = ID
	var s mp.EvaluationKpi
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Kpis").Find(&s).Where("evaluation_id = ?", evaluation.EvaluationId).Error; err != nil {
			return err
		  }
		if err := tx.Model(&s).Association("Kpis").Replace(&evaluation.Kpis); err != nil {
		 	return err
		  }
		return nil
		})
		return err
}


func GetKpiEvaluation(id *rp.GetEvaluationId,info rp.EvaluationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.EvaluationKpi{}).Where("evaluation_id = ?",id.ID)
	var Evaluations []mp.EvaluationKpi
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Evaluations).Error
	err = db.Preload("Kpis").Preload("Users").Find(&Evaluations).Error
	return err, Evaluations, total
}
