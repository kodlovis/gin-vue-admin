// 自动生成模板KpiTag
package pas

import (
	"gin-vue-admin/model"
)

// 如果含有time.Time 请自行import time包
type EvaluationKpiUser struct {
	ID     uint          `json:"id" gorm:"primarykey"`
	EKID   uint          `json:"ekid" gorm:"column:ek_id;primarykey;comment:方案指标库的ID"`
	Score  float64       `json:"score" gorm:"column:score;comment:指标分数;type:float;"`
	UserId uint          `json:"userId" gorm:"column:user_id;primarykey;comment:用户ID"`
	User   model.SysUser `json:"user" gorm:"ForeignKey:id;AssociationForeignKey:UserId;References:user_id"`
	//EvaluationKpi EvaluationKpi `json:"evaluationKpi" gorm:"ForeignKey:id;AssociationForeignKey:EKID;References:ek_id"`
}

func (EvaluationKpiUser) TableName() string {
	return "evaluation_kpi_user"
}
