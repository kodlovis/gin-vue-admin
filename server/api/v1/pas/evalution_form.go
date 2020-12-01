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

// @Tags EvalutionForm
// @Summary 创建EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "创建EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvalutionForm/createEvalutionForm [post]
func CreateEvalutionForm(c *gin.Context) {
	var EvalutionForm mp.EvalutionForm
	_ = c.ShouldBindJSON(&EvalutionForm)
	if err := sp.CreateEvalutionForm(EvalutionForm); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags EvalutionForm
// @Summary 删除EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "删除EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EvalutionForm/deleteEvalutionForm [delete]
func DeleteEvalutionForm(c *gin.Context) {
	var EvalutionForm mp.EvalutionForm
	_ = c.ShouldBindJSON(&EvalutionForm)
	if err := sp.DeleteEvalutionForm(EvalutionForm); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags EvalutionForm
// @Summary 批量删除EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /EvalutionForm/deleteEvalutionFormByIds [delete]
func DeleteEvalutionFormByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteEvalutionFormByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags EvalutionForm
// @Summary 更新EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "更新EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EvalutionForm/updateEvalutionForm [put]
func UpdateEvalutionForm(c *gin.Context) {
	var EvalutionForm mp.EvalutionForm
	_ = c.ShouldBindJSON(&EvalutionForm)
	if err := sp.UpdateEvalutionForm(&EvalutionForm); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags EvalutionForm
// @Summary 用id查询EvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvalutionForm true "用id查询EvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EvalutionForm/findEvalutionForm [get]
func FindEvalutionForm(c *gin.Context) {
	var EvalutionForm mp.EvalutionForm
	_ = c.ShouldBindQuery(&EvalutionForm)
	if err, reEvalutionForm := sp.GetEvalutionForm(EvalutionForm.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reEvalutionForm": reEvalutionForm}, c)
	}
}

// @Tags EvalutionForm
// @Summary 分页获取EvalutionForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.EvalutionFormSearch true "分页获取EvalutionForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EvalutionForm/getEvalutionFormList [get]
func GetEvalutionFormList(c *gin.Context) {
	var pageInfo rp.EvalutionFormSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetEvalutionFormInfoList(pageInfo); err != nil {
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
