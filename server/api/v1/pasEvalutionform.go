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

// @Tags PasEvalutionForm
// @Summary 创建PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "创建PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvalutionform/createPasEvalutionForm [post]
func CreatePasEvalutionForm(c *gin.Context) {
	var pasEvalutionform model.PasEvalutionForm
	_ = c.ShouldBindJSON(&pasEvalutionform)
	if err := service.CreatePasEvalutionForm(pasEvalutionform); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PasEvalutionForm
// @Summary 删除PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "删除PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasEvalutionform/deletePasEvalutionForm [delete]
func DeletePasEvalutionForm(c *gin.Context) {
	var pasEvalutionform model.PasEvalutionForm
	_ = c.ShouldBindJSON(&pasEvalutionform)
	if err := service.DeletePasEvalutionForm(pasEvalutionform); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PasEvalutionForm
// @Summary 批量删除PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pasEvalutionform/deletePasEvalutionFormByIds [delete]
func DeletePasEvalutionFormByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePasEvalutionFormByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PasEvalutionForm
// @Summary 更新PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "更新PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasEvalutionform/updatePasEvalutionForm [put]
func UpdatePasEvalutionForm(c *gin.Context) {
	var pasEvalutionform model.PasEvalutionForm
	_ = c.ShouldBindJSON(&pasEvalutionform)
	if err := service.UpdatePasEvalutionForm(&pasEvalutionform); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PasEvalutionForm
// @Summary 用id查询PasEvalutionForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvalutionForm true "用id查询PasEvalutionForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasEvalutionform/findPasEvalutionForm [get]
func FindPasEvalutionForm(c *gin.Context) {
	var pasEvalutionform model.PasEvalutionForm
	_ = c.ShouldBindQuery(&pasEvalutionform)
	if err, repasEvalutionform := service.GetPasEvalutionForm(pasEvalutionform.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repasEvalutionform": repasEvalutionform}, c)
	}
}

// @Tags PasEvalutionForm
// @Summary 分页获取PasEvalutionForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PasEvalutionFormSearch true "分页获取PasEvalutionForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvalutionform/getPasEvalutionFormList [get]
func GetPasEvalutionFormList(c *gin.Context) {
	var pageInfo request.PasEvalutionFormSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPasEvalutionFormInfoList(pageInfo); err != nil {
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
