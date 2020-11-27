package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPasEvaluationRouter(Router *gin.RouterGroup) {
	PasEvaluationRouter := Router.Group("pasEvaluation").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PasEvaluationRouter.POST("createPasEvaluation", v1.CreatePasEvaluation)   // 新建PasEvaluation
		PasEvaluationRouter.DELETE("deletePasEvaluation", v1.DeletePasEvaluation) // 删除PasEvaluation
		PasEvaluationRouter.DELETE("deletePasEvaluationByIds", v1.DeletePasEvaluationByIds) // 批量删除PasEvaluation
		PasEvaluationRouter.PUT("updatePasEvaluation", v1.UpdatePasEvaluation)    // 更新PasEvaluation
		PasEvaluationRouter.GET("findPasEvaluation", v1.FindPasEvaluation)        // 根据ID获取PasEvaluation
		PasEvaluationRouter.GET("getPasEvaluationList", v1.GetPasEvaluationList)  // 获取PasEvaluation列表
	}
}
