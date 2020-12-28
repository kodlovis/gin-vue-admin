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
	}
}
