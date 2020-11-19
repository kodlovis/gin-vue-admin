package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPasKpiRouter(Router *gin.RouterGroup) {
	PasKpiRouter := Router.Group("pasKpi").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PasKpiRouter.POST("createPasKpi", v1.CreatePasKpi)   // 新建PasKpi
		PasKpiRouter.DELETE("deletePasKpi", v1.DeletePasKpi) // 删除PasKpi
		PasKpiRouter.DELETE("deletePasKpiByIds", v1.DeletePasKpiByIds) // 批量删除PasKpi
		PasKpiRouter.PUT("updatePasKpi", v1.UpdatePasKpi)    // 更新PasKpi
		//PasKpiRouter.PUT("updatePasKpiByIds", v1.UpdatePasKpiByIds) // 樊批量归档PasKpi
		PasKpiRouter.GET("findPasKpi", v1.FindPasKpi)        // 根据ID获取PasKpi
		PasKpiRouter.GET("getPasKpiList", v1.GetPasKpiList)  // 获取PasKpi列表
	}
}
