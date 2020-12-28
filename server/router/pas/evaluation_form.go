package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEvaluationFormRouter(Router *gin.RouterGroup) {
	EvaluationFormRouter := Router.Group("EvaluationForm").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		EvaluationFormRouter.POST("createEvaluationForm", pas.CreateEvaluationForm)   // 新建EvaluationForm
		EvaluationFormRouter.DELETE("deleteEvaluationForm", pas.DeleteEvaluationForm) // 删除EvaluationForm
		EvaluationFormRouter.DELETE("deleteEvaluationFormByIds", pas.DeleteEvaluationFormByIds) // 批量删除EvaluationForm
		EvaluationFormRouter.PUT("updateEvaluationForm", pas.UpdateEvaluationForm)    // 更新EvaluationForm
		EvaluationFormRouter.GET("findEvaluationForm", pas.FindEvaluationForm)        // 根据ID获取EvaluationForm
		EvaluationFormRouter.GET("getEvaluationFormList",pas.GetEvaluationFormList)  // 获取EvaluationForm列表
	}
}
