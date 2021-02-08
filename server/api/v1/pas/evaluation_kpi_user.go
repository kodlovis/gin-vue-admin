package pas

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/pas"
	rp "gin-vue-admin/model/request/pas"
	"gin-vue-admin/model/response"
	sp "gin-vue-admin/service/pas"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
func CreateEKU(c *gin.Context) {
	var EKU rp.EKU
	_ = c.ShouldBindJSON(&EKU)
	if err := sp.CreateEKU(EKU.EvaluationKpiUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


func GetEKUByEKID(c *gin.Context) {
	var EKU rp.EKU
	var pageInfo rp.EvaluationKpiUserSearch
	_ = c.ShouldBindJSON(&EKU)
	if err, list, total := sp.GetEKUByEKID(EKU,pageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func UpdateEKU(c *gin.Context) {
	var EvaluationKpiUser mp.EvaluationKpiUser
	_ = c.ShouldBindJSON(&EvaluationKpiUser)
	if err := sp.UpdateEKU(&EvaluationKpiUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func RemoveEKU(c *gin.Context) {
	var EvaluationKpiUser mp.EvaluationKpiUser
	_ = c.ShouldBindJSON(&EvaluationKpiUser)
	if err := sp.RemoveEKU(EvaluationKpiUser.ID,EvaluationKpiUser.EKID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}