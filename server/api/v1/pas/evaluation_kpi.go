package pas

import (
	"gin-vue-admin/global"
    rp "gin-vue-admin/model/request/pas"
    "gin-vue-admin/model/response"
    mp "gin-vue-admin/model/pas"
    sp "gin-vue-admin/service/pas"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

func GetEvaluationKpiList(c *gin.Context) {
	var pageInfo rp.EvaluationKpiSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetEvaluationKpiInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

func SetUserEvaluation(c *gin.Context){
    var evaluationKpi rp.AssignedKpiEvaluationInfo
	_ = c.ShouldBindJSON(&evaluationKpi)
	if err := sp.SetUserEvaluation(evaluationKpi.ID, evaluationKpi.Users); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func CreateEvaluationKpi(c *gin.Context) {
	var EvaluationKpi mp.EvaluationKpi
	_ = c.ShouldBindJSON(&EvaluationKpi)
	if err := sp.CreateEvaluationKpi(EvaluationKpi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func RemoveEvaluationKpi(c *gin.Context) {
	var evaluationKpi rp.AssignedKpiEvaluationInfo
	_ = c.ShouldBindJSON(&evaluationKpi)
	if err := sp.RemoveEvaluationKpi(evaluationKpi.ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func RemoveEvaluationKpiByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.RemoveEvaluationKpiByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func GetEvaluationKpiById(c *gin.Context) {
	var EvaluationKpi mp.EvaluationKpi
	var pageInfo rp.EvaluationKpiSearch
    _ = c.ShouldBindJSON(&EvaluationKpi)
	if err, list, total := sp.GetEvaluationKpiById(EvaluationKpi.ID,pageInfo); err != nil {
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