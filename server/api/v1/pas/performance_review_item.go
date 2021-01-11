package pas

import (
	"gin-vue-admin/global"
    mp "gin-vue-admin/model/pas"
    "gin-vue-admin/model/response"
    sp "gin-vue-admin/service/pas"
    "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePerformanceReviewItem(c *gin.Context) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PerformanceReviewItem)
	if err := sp.CreatePerformanceReviewItem(PerformanceReviewItem); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}