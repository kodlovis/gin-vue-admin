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

// @Tags Evaluation
// @Summary 创建Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "创建Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Evaluation/createEvaluation [post]
func CreateEvaluation(c *gin.Context) {
	var Evaluation mp.Evaluation
	_ = c.ShouldBindJSON(&Evaluation)
	if err := sp.CreateEvaluation(Evaluation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Evaluation
// @Summary 删除Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "删除Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Evaluation/deleteEvaluation [delete]
func DeleteEvaluation(c *gin.Context) {
	var Evaluation mp.Evaluation
	_ = c.ShouldBindJSON(&Evaluation)
	if err := sp.DeleteEvaluation(Evaluation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Evaluation
// @Summary 批量删除Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Evaluation/deleteEvaluationByIds [delete]
func DeleteEvaluationByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteEvaluationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Evaluation
// @Summary 更新Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "更新Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Evaluation/updateEvaluation [put]
func UpdateEvaluation(c *gin.Context) {
	var Evaluation mp.Evaluation
	_ = c.ShouldBindJSON(&Evaluation)
	if err := sp.UpdateEvaluation(&Evaluation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
func UpdateEvaluationByInfo(c *gin.Context) {
	var Evaluation rp.EvaluationInfo
	_ = c.ShouldBindJSON(&Evaluation)
	if err := sp.UpdateEvaluationByInfo(Evaluation.ID,Evaluation.TotalScore); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
// @Tags Evaluation
// @Summary 用id查询Evaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluation true "用id查询Evaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Evaluation/findEvaluation [get]
func FindEvaluation(c *gin.Context) {
	var Evaluation mp.Evaluation
	_ = c.ShouldBindJSON(&Evaluation)
	if err, reEvaluation := sp.GetEvaluation(Evaluation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reEvaluation": reEvaluation}, c)
	}
}

// @Tags Evaluation
// @Summary 分页获取Evaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.EvaluationSearch true "分页获取Evaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Evaluation/getEvaluationList [get]
func GetEvaluationList(c *gin.Context) {
	var pageInfo rp.EvaluationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetEvaluationInfoList(pageInfo); err != nil {
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
