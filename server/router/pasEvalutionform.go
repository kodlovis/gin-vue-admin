package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPasEvalutionFormRouter(Router *gin.RouterGroup) {
	PasEvalutionFormRouter := Router.Group("pasEvalutionform").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PasEvalutionFormRouter.POST("createPasEvalutionForm", v1.CreatePasEvalutionForm)   // 新建PasEvalutionForm
		PasEvalutionFormRouter.DELETE("deletePasEvalutionForm", v1.DeletePasEvalutionForm) // 删除PasEvalutionForm
		PasEvalutionFormRouter.DELETE("deletePasEvalutionFormByIds", v1.DeletePasEvalutionFormByIds) // 批量删除PasEvalutionForm
		PasEvalutionFormRouter.PUT("updatePasEvalutionForm", v1.UpdatePasEvalutionForm)    // 更新PasEvalutionForm
		PasEvalutionFormRouter.GET("findPasEvalutionForm", v1.FindPasEvalutionForm)        // 根据ID获取PasEvalutionForm
		PasEvalutionFormRouter.GET("getPasEvalutionFormList", v1.GetPasEvalutionFormList)  // 获取PasEvalutionForm列表
	}
}
