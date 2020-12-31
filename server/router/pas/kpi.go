package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitKpiRouter(Router *gin.RouterGroup) {
	KpiRouter := Router.Group("Kpi").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		KpiRouter.POST("createKpi", pas.CreateKpi)   // 新建Kpi
		KpiRouter.DELETE("deleteKpi", pas.DeleteKpi) // 删除Kpi
		KpiRouter.DELETE("deleteKpiByIds", pas.DeleteKpiByIds) // 批量删除Kpi
		KpiRouter.PUT("updateKpi", pas.UpdateKpi)    // 更新Kpi
		KpiRouter.GET("findKpi", pas.FindKpi)        // 根据ID获取Kpi
		KpiRouter.GET("getKpiList", pas.GetKpiList)  // 获取Kpi列表
		KpiRouter.DELETE("removeKpiTags",pas.RemoveKpiTags)
		//KpiRouter.POST("assignedKpiEvaluation", pas.AssignedKpiEvaluation)
		KpiRouter.POST("getKpiEvaluation", pas.GetKpiEvaluation)
		KpiRouter.POST("getKpiByIds", pas.GetKpiByIds)
		KpiRouter.POST("getKpiScoreByIds", pas.GetKpiScoreByIds)
	}
}
