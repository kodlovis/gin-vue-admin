package pas

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model/response"
    sp "gin-vue-admin/service/pas"
    rp "gin-vue-admin/model/request/pas"
    "github.com/gin-gonic/gin"
    //mp "gin-vue-admin/model/pas"
	"go.uber.org/zap"
)

func CreateKpiTag(c *gin.Context) {
	var res rp.KpiTagList
	_ = c.ShouldBindJSON(&res)
	if err := sp.CreateKpiTag(res.KpiTagList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// func UpdatePerformanceReviewItemByInfo(c *gin.Context) {
// 	var PerformanceReviewItem mp.PerformanceReviewItem
// 	_ = c.ShouldBindJSON(&PerformanceReviewItem)
// 	if err := sp.UpdatePerformanceReviewItemByInfo(PerformanceReviewItem.ID,PerformanceReviewItem.Score,PerformanceReviewItem.UserId); err != nil {
//         global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
// 		response.FailWithMessage("更新失败", c)
// 	} else {
// 		response.OkWithMessage("更新成功", c)
// 	}
// }