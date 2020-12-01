package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEvalutionFormRouter(Router *gin.RouterGroup) {
	EvalutionFormRouter := Router.Group("EvalutionForm").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		EvalutionFormRouter.POST("createEvalutionForm", pas.CreateEvalutionForm)   // 新建EvalutionForm
		EvalutionFormRouter.DELETE("deleteEvalutionForm", pas.DeleteEvalutionForm) // 删除EvalutionForm
		EvalutionFormRouter.DELETE("deleteEvalutionFormByIds", pas.DeleteEvalutionFormByIds) // 批量删除EvalutionForm
		EvalutionFormRouter.PUT("updateEvalutionForm", pas.UpdateEvalutionForm)    // 更新EvalutionForm
		EvalutionFormRouter.GET("findEvalutionForm", pas.FindEvalutionForm)        // 根据ID获取EvalutionForm
		EvalutionFormRouter.GET("getEvalutionFormList",pas.GetEvalutionFormList)  // 获取EvalutionForm列表
	}
}
