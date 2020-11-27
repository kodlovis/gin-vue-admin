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

// @Tags PasAllocation
// @Summary 创建PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "创建PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasAllocation/createPasAllocation [post]
func CreatePasAllocation(c *gin.Context) {
	var pasAllocation model.PasAllocation
	_ = c.ShouldBindJSON(&pasAllocation)
	if err := service.CreatePasAllocation(pasAllocation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PasAllocation
// @Summary 删除PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "删除PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasAllocation/deletePasAllocation [delete]
func DeletePasAllocation(c *gin.Context) {
	var pasAllocation model.PasAllocation
	_ = c.ShouldBindJSON(&pasAllocation)
	if err := service.DeletePasAllocation(pasAllocation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PasAllocation
// @Summary 批量删除PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pasAllocation/deletePasAllocationByIds [delete]
func DeletePasAllocationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePasAllocationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PasAllocation
// @Summary 更新PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "更新PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasAllocation/updatePasAllocation [put]
func UpdatePasAllocation(c *gin.Context) {
	var pasAllocation model.PasAllocation
	_ = c.ShouldBindJSON(&pasAllocation)
	if err := service.UpdatePasAllocation(&pasAllocation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PasAllocation
// @Summary 用id查询PasAllocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasAllocation true "用id查询PasAllocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasAllocation/findPasAllocation [get]
func FindPasAllocation(c *gin.Context) {
	var pasAllocation model.PasAllocation
	_ = c.ShouldBindQuery(&pasAllocation)
	if err, repasAllocation := service.GetPasAllocation(pasAllocation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repasAllocation": repasAllocation}, c)
	}
}

// @Tags PasAllocation
// @Summary 分页获取PasAllocation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PasAllocationSearch true "分页获取PasAllocation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasAllocation/getPasAllocationList [get]
func GetPasAllocationList(c *gin.Context) {
	var pageInfo request.PasAllocationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPasAllocationInfoList(pageInfo); err != nil {
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
