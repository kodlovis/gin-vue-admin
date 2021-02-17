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
func GetLastPRICreatePRIU(c *gin.Context) {
	var res rp.PerformanceReviewItemUserSearch
	_ = c.ShouldBindJSON(&res)
	if err := sp.GetLastPRICreatePRIU(res.EKUID,res.PRID); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
func GetPRIUByPRIID(c *gin.Context) {
	var pageInfo rp.PerformanceReviewItemUserSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := sp.GetPRIUByPRIID(pageInfo.PRIID,pageInfo); err != nil {
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

func UpdatePRIU(c *gin.Context) {
	var PRIU mp.PerformanceReviewItemUser
	_ = c.ShouldBindJSON(&PRIU)
	if err := sp.UpdatePRIU(PRIU.ID,PRIU.Status,PRIU.Result,PRIU.UserId,PRIU.Score); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func RemovePRIU(c *gin.Context) {
	var PerformanceReviewItemUser mp.PerformanceReviewItemUser
	_ = c.ShouldBindJSON(&PerformanceReviewItemUser)
	if err := sp.RemovePRIU(PerformanceReviewItemUser.ID,PerformanceReviewItemUser.PRIID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func CreatePRIU(c *gin.Context) {
	var PRIU rp.PRIU
	_ = c.ShouldBindJSON(&PRIU)
	if err := sp.CreatePRIU(PRIU.PerformanceReviewItemUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func GetPRIUListByUser(c *gin.Context) {
	var info rp.PerformanceReviewItemUserSearch
	_ = c.ShouldBindJSON(&info)
	if err, list, total := sp.GetPRIUListByUser(info.ID, info.Status, info); err != nil {
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

func UpdatePRIUStatusByPRIID(c *gin.Context) {
	var PRIU mp.PerformanceReviewItemUser
	_ = c.ShouldBindJSON(&PRIU)
	if err := sp.UpdatePRIUStatusByPRIID(PRIU.ID, PRIU.Status, PRIU.Result); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}