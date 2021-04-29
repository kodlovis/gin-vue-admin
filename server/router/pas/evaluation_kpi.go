package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEvaluationKpiRouter(Router *gin.RouterGroup) {
	EvaluationKpiRouter := Router.Group("EvaluationKpi").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		EvaluationKpiRouter.GET("getEvaluationKpiList", pas.GetEvaluationKpiList)
		EvaluationKpiRouter.POST("setUserEvaluation", pas.SetUserEvaluation)
		EvaluationKpiRouter.POST("createEvaluationKpi", pas.CreateEvaluationKpi)   // 新建Kpi
		EvaluationKpiRouter.DELETE("removeEvaluationKpi", pas.RemoveEvaluationKpi) // 删除Kpi
		EvaluationKpiRouter.DELETE("removeEvaluationKpiByIds", pas.RemoveEvaluationKpiByIds) 
		EvaluationKpiRouter.POST("getEvaluationKpiById", pas.GetEvaluationKpiById) 
		EvaluationKpiRouter.PUT("updateEvaluationKpi", pas.UpdateEvaluationKpi) 
		EvaluationKpiRouter.GET("getLastEvaluationKpi", pas.GetLastEvaluationKpi) 
		EvaluationKpiRouter.GET("getEvaluationKpi", pas.GetEvaluationKpi) 
		
	}
}
