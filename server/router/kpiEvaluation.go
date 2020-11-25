package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitKpiEvaluationRouter(Router *gin.RouterGroup) {
	KpiEvaluationRouter := Router.Group("kpiEvaluation").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		KpiEvaluationRouter.POST("createKpiEvaluation", v1.CreateKpiEvaluation)   // 新建KpiEvaluation
		KpiEvaluationRouter.DELETE("deleteKpiEvaluation", v1.DeleteKpiEvaluation) // 删除KpiEvaluation
		KpiEvaluationRouter.DELETE("deleteKpiEvaluationByIds", v1.DeleteKpiEvaluationByIds) // 批量删除KpiEvaluation
		KpiEvaluationRouter.PUT("updateKpiEvaluation", v1.UpdateKpiEvaluation)    // 更新KpiEvaluation
		KpiEvaluationRouter.GET("findKpiEvaluation", v1.FindKpiEvaluation)        // 根据ID获取KpiEvaluation
		KpiEvaluationRouter.GET("getKpiEvaluationList", v1.GetKpiEvaluationList)  // 获取KpiEvaluation列表
	}
}
