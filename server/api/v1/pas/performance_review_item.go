package pas

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model/response"
    sp "gin-vue-admin/service/pas"
    rp "gin-vue-admin/model/request/pas"
    "github.com/gin-gonic/gin"
    mp "gin-vue-admin/model/pas"
	"go.uber.org/zap"
)

func CreatePerformanceReviewItem(c *gin.Context) {
	var res rp.PerformanceReviewItemList
	_ = c.ShouldBindJSON(&res)
	if err := sp.CreatePerformanceReviewItem(res.PerformanceReviewItemList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeletePerformanceReviewItem(c *gin.Context) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PerformanceReviewItem)
	if err := sp.DeletePerformanceReviewItem(PerformanceReviewItem); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func DeletePerformanceReviewItemByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeletePerformanceReviewItemByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func GetPerformanceReviewListById(c *gin.Context) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	var pageInfo rp.PerformanceReviewItemSearch
    _ = c.ShouldBindJSON(&PerformanceReviewItem)
	if err, list, total := sp.GetPerformanceReviewListById(PerformanceReviewItem.PRId,pageInfo); err != nil {
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