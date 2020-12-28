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

// @Tags EvaluationForm
// @Summary 创建EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "创建EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvaluationForm/createEvaluationForm [post]
func CreateEvaluationForm(c *gin.Context) {
	var EvaluationForm mp.EvaluationForm
	_ = c.ShouldBindJSON(&EvaluationForm)
	if err := sp.CreateEvaluationForm(EvaluationForm); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags EvaluationForm
// @Summary 删除EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "删除EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EvaluationForm/deleteEvaluationForm [delete]
func DeleteEvaluationForm(c *gin.Context) {
	var EvaluationForm mp.EvaluationForm
	_ = c.ShouldBindJSON(&EvaluationForm)
	if err := sp.DeleteEvaluationForm(EvaluationForm); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags EvaluationForm
// @Summary 批量删除EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /EvaluationForm/deleteEvaluationFormByIds [delete]
func DeleteEvaluationFormByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteEvaluationFormByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags EvaluationForm
// @Summary 更新EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "更新EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EvaluationForm/updateEvaluationForm [put]
func UpdateEvaluationForm(c *gin.Context) {
	var EvaluationForm mp.EvaluationForm
	_ = c.ShouldBindJSON(&EvaluationForm)
	if err := sp.UpdateEvaluationForm(&EvaluationForm); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags EvaluationForm
// @Summary 用id查询EvaluationForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluationForm true "用id查询EvaluationForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EvaluationForm/findEvaluationForm [get]
func FindEvaluationForm(c *gin.Context) {
	var EvaluationForm mp.EvaluationForm
	_ = c.ShouldBindQuery(&EvaluationForm)
	if err, reEvaluationForm := sp.GetEvaluationForm(EvaluationForm.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reEvaluationForm": reEvaluationForm}, c)
	}
}

// @Tags EvaluationForm
// @Summary 分页获取EvaluationForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.EvaluationFormSearch true "分页获取EvaluationForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvaluationForm/getEvaluationFormList [get]
func GetEvaluationFormList(c *gin.Context) {
	var pageInfo rp.EvaluationFormSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetEvaluationFormInfoList(pageInfo); err != nil {
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
