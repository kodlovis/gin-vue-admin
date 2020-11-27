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

// @Tags PasEvaluation
// @Summary 创建PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "创建PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvaluation/createPasEvaluation [post]
func CreatePasEvaluation(c *gin.Context) {
	var pasEvaluation model.PasEvaluation
	_ = c.ShouldBindJSON(&pasEvaluation)
	if err := service.CreatePasEvaluation(pasEvaluation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PasEvaluation
// @Summary 删除PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "删除PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasEvaluation/deletePasEvaluation [delete]
func DeletePasEvaluation(c *gin.Context) {
	var pasEvaluation model.PasEvaluation
	_ = c.ShouldBindJSON(&pasEvaluation)
	if err := service.DeletePasEvaluation(pasEvaluation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PasEvaluation
// @Summary 批量删除PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pasEvaluation/deletePasEvaluationByIds [delete]
func DeletePasEvaluationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePasEvaluationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PasEvaluation
// @Summary 更新PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "更新PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasEvaluation/updatePasEvaluation [put]
func UpdatePasEvaluation(c *gin.Context) {
	var pasEvaluation model.PasEvaluation
	_ = c.ShouldBindJSON(&pasEvaluation)
	if err := service.UpdatePasEvaluation(&pasEvaluation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PasEvaluation
// @Summary 用id查询PasEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasEvaluation true "用id查询PasEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasEvaluation/findPasEvaluation [get]
func FindPasEvaluation(c *gin.Context) {
	var pasEvaluation model.PasEvaluation
	_ = c.ShouldBindQuery(&pasEvaluation)
	if err, repasEvaluation := service.GetPasEvaluation(pasEvaluation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repasEvaluation": repasEvaluation}, c)
	}
}

// @Tags PasEvaluation
// @Summary 分页获取PasEvaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PasEvaluationSearch true "分页获取PasEvaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasEvaluation/getPasEvaluationList [get]
func GetPasEvaluationList(c *gin.Context) {
	var pageInfo request.PasEvaluationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPasEvaluationInfoList(pageInfo); err != nil {
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
