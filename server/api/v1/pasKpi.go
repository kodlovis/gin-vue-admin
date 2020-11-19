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

// @Tags PasKpi
// @Summary 创建PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "创建PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasKpi/createPasKpi [post]
func CreatePasKpi(c *gin.Context) {
	var pasKpi model.PasKpi
	_ = c.ShouldBindJSON(&pasKpi)
	if err := service.CreatePasKpi(pasKpi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PasKpi
// @Summary 删除PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "删除PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasKpi/deletePasKpi [delete]
func DeletePasKpi(c *gin.Context) {
	var pasKpi model.PasKpi
	_ = c.ShouldBindJSON(&pasKpi)
	if err := service.DeletePasKpi(pasKpi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PasKpi
// @Summary 批量删除PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pasKpi/deletePasKpiByIds [delete]
func DeletePasKpiByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePasKpiByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

//樊新增批量归档
// func UpdatePasKpiByIds(c *gin.Context) {
// 	var IDS request.IdsReq
//     _ = c.ShouldBindJSON(&IDS)
// 	if err := service.UpdatePasKpiByIds(IDS); err != nil {
//         global.GVA_LOG.Error("批量归档失败!", zap.Any("err", err))
// 		response.FailWithMessage("批量归档失败", c)
// 	} else {
// 		response.OkWithMessage("批量归档成功", c)
// 	}
// }

// @Tags PasKpi
// @Summary 更新PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "更新PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasKpi/updatePasKpi [put]
func UpdatePasKpi(c *gin.Context) {
	var pasKpi model.PasKpi
	_ = c.ShouldBindJSON(&pasKpi)
	if err := service.UpdatePasKpi(&pasKpi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PasKpi
// @Summary 用id查询PasKpi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasKpi true "用id查询PasKpi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasKpi/findPasKpi [get]
func FindPasKpi(c *gin.Context) {
	var pasKpi model.PasKpi
	_ = c.ShouldBindQuery(&pasKpi)
	if err, repasKpi := service.GetPasKpi(pasKpi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repasKpi": repasKpi}, c)
	}
}

// @Tags PasKpi
// @Summary 分页获取PasKpi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PasKpiSearch true "分页获取PasKpi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasKpi/getPasKpiList [get]
func GetPasKpiList(c *gin.Context) {
	var pageInfo request.PasKpiSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPasKpiInfoList(pageInfo); err != nil {
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
