package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEvaluationRouter(Router *gin.RouterGroup) {
	EvaluationRouter := Router.Group("Evaluation").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		EvaluationRouter.POST("createEvaluation", pas.CreateEvaluation)   // 新建Evaluation
		EvaluationRouter.DELETE("deleteEvaluation", pas.DeleteEvaluation) // 删除Evaluation
		EvaluationRouter.DELETE("deleteEvaluationByIds", pas.DeleteEvaluationByIds) // 批量删除Evaluation
		EvaluationRouter.PUT("updateEvaluation", pas.UpdateEvaluation)    // 更新Evaluation
		EvaluationRouter.GET("findEvaluation", pas.FindEvaluation)        // 根据ID获取Evaluation
		EvaluationRouter.GET("getEvaluationList",pas.GetEvaluationList)  // 获取Evaluation列表
		
		
	}
}
