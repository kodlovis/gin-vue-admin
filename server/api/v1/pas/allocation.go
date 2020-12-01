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

// @Tags Allocation
// @Summary 创建Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "创建Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Allocation/createAllocation [post]
func CreateAllocation(c *gin.Context) {
	var Allocation mp.Allocation
	_ = c.ShouldBindJSON(&Allocation)
	if err := sp.CreateAllocation(Allocation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Allocation
// @Summary 删除Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "删除Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Allocation/deleteAllocation [delete]
func DeleteAllocation(c *gin.Context) {
	var Allocation mp.Allocation
	_ = c.ShouldBindJSON(&Allocation)
	if err := sp.DeleteAllocation(Allocation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Allocation
// @Summary 批量删除Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Allocation/deleteAllocationByIds [delete]
func DeleteAllocationByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteAllocationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Allocation
// @Summary 更新Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "更新Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Allocation/updateAllocation [put]
func UpdateAllocation(c *gin.Context) {
	var Allocation mp.Allocation
	_ = c.ShouldBindJSON(&Allocation)
	if err := sp.UpdateAllocation(&Allocation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Allocation
// @Summary 用id查询Allocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Allocation true "用id查询Allocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Allocation/findAllocation [get]
func FindAllocation(c *gin.Context) {
	var Allocation mp.Allocation
	_ = c.ShouldBindQuery(&Allocation)
	if err, reAllocation := sp.GetAllocation(Allocation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reAllocation": reAllocation}, c)
	}
}

// @Tags Allocation
// @Summary 分页获取Allocation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AllocationSearch true "分页获取Allocation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Allocation/getAllocationList [get]
func GetAllocationList(c *gin.Context) {
	var pageInfo rp.AllocationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetAllocationInfoList(pageInfo); err != nil {
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
