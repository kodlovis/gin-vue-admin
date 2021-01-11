package pas

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model/response"
    sp "gin-vue-admin/service/pas"
    rp "gin-vue-admin/model/request/pas"
    "github.com/gin-gonic/gin"
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