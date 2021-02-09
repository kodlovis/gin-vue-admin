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
func DeletePRItemById(c *gin.Context) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PerformanceReviewItem)
	if err := sp.DeletePRItemById(PerformanceReviewItem); err != nil {
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

func GetPerformanceReviewItemListById(c *gin.Context) {
	var PerformanceReviewItem mp.PerformanceReviewItem
	var pageInfo rp.PerformanceReviewItemSearch
	_ = c.ShouldBindJSON(&PerformanceReviewItem)
	if err, list, total := sp.GetPerformanceReviewItemListById(PerformanceReviewItem.PRId, pageInfo); err != nil {
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

func UpdatePerformanceReviewItemByInfo(c *gin.Context) {
	var PRI mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PRI)
	if err := sp.UpdatePerformanceReviewItemByInfo(PRI.ID, PRI.Score, PRI.Status); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func GetPRItemListByUser(c *gin.Context) {
	var info rp.PerformanceReviewItemSearch
	_ = c.ShouldBindJSON(&info)
	if err, list, total := sp.GetPRItemListByUser(info.ID, info.Status, info); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}

func UpdatePRItemStatusById(c *gin.Context) {
	var PRItem mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PRItem)
	if err := sp.UpdatePRItemStatusById(PRItem.ID, PRItem.Status, PRItem.Result); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
func UpdatePRItemStatusByPrId(c *gin.Context) {
	var PRItem mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PRItem)
	if err := sp.UpdatePRItemStatusByPrId(PRItem.PRId, PRItem.Status); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
func GetPRItemCount(c *gin.Context) {
	var PRItem mp.PerformanceReviewItem
	_ = c.ShouldBindJSON(&PRItem)
	if err, count := sp.GetPRItemCount(PRItem.PRId, PRItem.Status); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			Count: count,
		}, "获取成功", c)
	}
}
func UpdatePRItemStatysByIds(c *gin.Context) {
	var PRItemInfo rp.PerformanceReviewInfo
	_ = c.ShouldBindJSON(&PRItemInfo)
	if err := sp.UpdatePRItemStatysByIds(PRItemInfo.Ids, PRItemInfo.Status); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
func GetPRItemListByStatusPrid(c *gin.Context) {
	var info rp.PerformanceReviewItemSearch
	_ = c.ShouldBindJSON(&info)
	if err, list, total := sp.GetPRItemListByStatusPrid(info.Status, info.Ids, info); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}
func GetPRItemListByStatus(c *gin.Context) {
	var info rp.PerformanceReviewItemSearch
	_ = c.ShouldBindJSON(&info)
	if err, list, total := sp.GetPRItemListByStatus(info.Status, info); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}

func GetPRItemListByPrids(c *gin.Context) {
	var pageInfo rp.PerformanceReviewItemSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := sp.GetPRItemListByPrids(pageInfo.Ids, pageInfo); err != nil {
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
