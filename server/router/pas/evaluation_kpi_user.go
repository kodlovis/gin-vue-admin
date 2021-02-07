package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEvaluationKpiUserRouter(Router *gin.RouterGroup) {
	EvaluationKpiUserRouter := Router.Group("EvaluationKpiUser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		EvaluationKpiUserRouter.POST("createEKU", pas.CreateEKU)   // 新建Kpi
		EvaluationKpiUserRouter.POST("getEKUByEKID", pas.GetEKUByEKID)   // 新建Kpi
		EvaluationKpiUserRouter.PUT("updateEKU", pas.UpdateEKU)   // 新建Kpi
		
		
	}
}
