package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags KpiEvaluation
// @Summary 创建KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "创建KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kpiEvaluation/createKpiEvaluation [post]
func CreateKpiEvaluation(c *gin.Context) {
	var kpiEvaluation model.KpiEvaluation
	_ = c.ShouldBindJSON(&kpiEvaluation)
	if err := service.CreateKpiEvaluation(kpiEvaluation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags KpiEvaluation
// @Summary 删除KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "删除KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kpiEvaluation/deleteKpiEvaluation [delete]
func DeleteKpiEvaluation(c *gin.Context) {
	var kpiEvaluation model.KpiEvaluation
	_ = c.ShouldBindJSON(&kpiEvaluation)
	if err := service.DeleteKpiEvaluation(kpiEvaluation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags KpiEvaluation
// @Summary 批量删除KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /kpiEvaluation/deleteKpiEvaluationByIds [delete]
func DeleteKpiEvaluationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteKpiEvaluationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags KpiEvaluation
// @Summary 更新KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "更新KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /kpiEvaluation/updateKpiEvaluation [put]
func UpdateKpiEvaluation(c *gin.Context) {
	var kpiEvaluation model.KpiEvaluation
	_ = c.ShouldBindJSON(&kpiEvaluation)
	if err := service.UpdateKpiEvaluation(&kpiEvaluation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags KpiEvaluation
// @Summary 用id查询KpiEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KpiEvaluation true "用id查询KpiEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /kpiEvaluation/findKpiEvaluation [get]
func FindKpiEvaluation(c *gin.Context) {
	var kpiEvaluation model.KpiEvaluation
	_ = c.ShouldBindQuery(&kpiEvaluation)
	if err, rekpiEvaluation := service.GetKpiEvaluation(kpiEvaluation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rekpiEvaluation": rekpiEvaluation}, c)
	}
}

// @Tags KpiEvaluation
// @Summary 分页获取KpiEvaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.KpiEvaluationSearch true "分页获取KpiEvaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kpiEvaluation/getKpiEvaluationList [get]
func GetKpiEvaluationList(c *gin.Context) {
	var pageInfo request.KpiEvaluationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetKpiEvaluationInfoList(pageInfo); err != nil {
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
