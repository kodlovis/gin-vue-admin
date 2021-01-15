package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitKpiTagRouter(Router *gin.RouterGroup) {
	KpiTagRouter := Router.Group("KpiTag").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		KpiTagRouter.POST("createKpiTag", pas.CreateKpiTag)
	}
}
