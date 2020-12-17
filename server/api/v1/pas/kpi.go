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

// @Tags Kpi
// @Summary 创建Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "创建Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Kpi/createKpi [post]
func CreateKpi(c *gin.Context) {
	var Kpi mp.Kpi
	_ = c.ShouldBindJSON(&Kpi)
	if err := sp.CreateKpi(Kpi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Kpi
// @Summary 删除Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "删除Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Kpi/deleteKpi [delete]
func DeleteKpi(c *gin.Context) {
	var Kpi mp.Kpi
	_ = c.ShouldBindJSON(&Kpi)
	if err := sp.DeleteKpi(Kpi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func RemoveKpiTags(c *gin.Context) {
	var Kpi mp.Kpi
	_ = c.ShouldBindJSON(&Kpi)
	if err := sp.RemoveKpiTags(Kpi); err != nil {
        global.GVA_LOG.Error("清除失败!", zap.Any("err", err))
		response.FailWithMessage("清除失败", c)
	} else {
		response.OkWithMessage("清除成功", c)
	}
}
// @Tags Kpi
// @Summary 批量删除Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Kpi/deleteKpiByIds [delete]
func DeleteKpiByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteKpiByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}


// @Tags Kpi
// @Summary 更新Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "更新Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Kpi/updateKpi [put]
func UpdateKpi(c *gin.Context) {
	var Kpi mp.Kpi
	_ = c.ShouldBindJSON(&Kpi)
	if err := sp.UpdateKpi(&Kpi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Kpi
// @Summary 用id查询Kpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Kpi true "用id查询Kpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Kpi/findKpi [get]
func FindKpi(c *gin.Context) {
	var Kpi mp.Kpi
	_ = c.ShouldBindQuery(&Kpi)
	if err, reKpi := sp.GetKpi(Kpi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reKpi": reKpi}, c)
	}
}

// @Tags Kpi
// @Summary 分页获取Kpi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.KpiSearch true "分页获取Kpi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Kpi/getKpiList [get]
func GetKpiList(c *gin.Context) {
	var pageInfo rp.KpiSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetKpiInfoList(pageInfo); err != nil {
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
func GetKpiByIds(c *gin.Context) {
	var IDS rp.IdsReq
	var pageInfo rp.KpiSearch
    _ = c.ShouldBindJSON(&IDS)
	if err, list, total := sp.GetKpiByIds(IDS,pageInfo); err != nil {
        global.GVA_LOG.Error("批量查询失败!", zap.Any("err", err))
		response.FailWithMessage("批量查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
	}
}
func AddKpiEvaluation(c *gin.Context) {
	var evaluationKpi rp.AddKpiEvaluationInfo
	_ = c.ShouldBindJSON(&evaluationKpi)
	if err := sp.AddKpiEvaluation(evaluationKpi.Kpis, evaluationKpi.ID); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func GetKpiEvaluation(c *gin.Context) {
	var param rp.GetEvaluationId
	var pageInfo rp.KpiSearch
	_ = c.ShouldBindJSON(&param)
	if err,list := sp.GetKpiEvaluation(&param,pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": list}, "获取成功", c)
	}
}
