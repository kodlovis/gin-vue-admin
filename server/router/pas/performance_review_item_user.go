package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPerformanceReviewItemUserRouter(Router *gin.RouterGroup) {
	PerformanceReviewItemUserRouter := Router.Group("PerformanceReviewItemUser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PerformanceReviewItemUserRouter.POST("getLastPRICreatePRIU", pas.GetLastPRICreatePRIU)
	}
}
