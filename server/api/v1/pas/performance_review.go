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

// @Tags PerformanceReview
// @Summary 创建PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "创建PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PerformanceReview/createPerformanceReview [post]
func CreatePerformanceReview(c *gin.Context) {
	var PerformanceReview mp.PerformanceReview
	_ = c.ShouldBindJSON(&PerformanceReview)
	if err := sp.CreatePerformanceReview(PerformanceReview); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PerformanceReview
// @Summary 删除PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "删除PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PerformanceReview/deletePerformanceReview [delete]
func DeletePerformanceReview(c *gin.Context) {
	var PerformanceReview mp.PerformanceReview
	_ = c.ShouldBindJSON(&PerformanceReview)
	if err := sp.DeletePerformanceReview(PerformanceReview); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PerformanceReview
// @Summary 批量删除PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PerformanceReview/deletePerformanceReviewByIds [delete]
func DeletePerformanceReviewByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeletePerformanceReviewByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PerformanceReview
// @Summary 更新PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "更新PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PerformanceReview/updatePerformanceReview [put]
func UpdatePerformanceReview(c *gin.Context) {
	var PerformanceReview mp.PerformanceReview
	_ = c.ShouldBindJSON(&PerformanceReview)
	if err := sp.UpdatePerformanceReview(&PerformanceReview); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PerformanceReview
// @Summary 用id查询PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "用id查询PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PerformanceReview/findPerformanceReview [get]
func FindPerformanceReview(c *gin.Context) {
	var PerformanceReview mp.PerformanceReview
	_ = c.ShouldBindQuery(&PerformanceReview)
	if err, rePerformanceReview := sp.GetPerformanceReview(PerformanceReview.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePerformanceReview": rePerformanceReview}, c)
	}
}

func GetPRBystatus(c *gin.Context) {
	var PerformanceReview mp.PerformanceReview
	var pageInfo rp.PerformanceReviewSearch
	_ = c.ShouldBindQuery(&PerformanceReview)
	if err, list, total := sp.GetPRBystatus(PerformanceReview.Status,pageInfo); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
	}
}
// @Tags PerformanceReview
// @Summary 分页获取PerformanceReview列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PerformanceReviewSearch true "分页获取PerformanceReview列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PerformanceReview/getPerformanceReviewList [get]
func GetPerformanceReviewList(c *gin.Context) {
	var pageInfo rp.PerformanceReviewSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetPerformanceReviewInfoList(pageInfo); err != nil {
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

func UpdatePerformanceReviewByInfo(c *gin.Context) {
	var PRInfo rp.PerformanceReviewInfo
	_ = c.ShouldBindJSON(&PRInfo)
	if err := sp.UpdatePerformanceReviewByInfo(PRInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


func GetLastPerformanceReview(c *gin.Context) {
	var PerformanceReview mp.PerformanceReview
	_ = c.ShouldBindQuery(&PerformanceReview)
	if err, rePR := sp.GetLastPerformanceReview(); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePR": rePR}, c)
	}
}

func UpdatePRStatusById(c *gin.Context) {
	var PRInfo  mp.PerformanceReview
	_ = c.ShouldBindJSON(&PRInfo)
	if err := sp.UpdatePRStatusById(PRInfo.ID,PRInfo.Status); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}