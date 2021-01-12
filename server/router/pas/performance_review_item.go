package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPerformanceReviewItemRouter(Router *gin.RouterGroup) {
	PerformanceReviewItemRouter := Router.Group("PerformanceReviewItem").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PerformanceReviewItemRouter.POST("createPerformanceReviewItem", pas.CreatePerformanceReviewItem)
		PerformanceReviewItemRouter.DELETE("deletePerformanceReviewItem", pas.DeletePerformanceReviewItem)
		PerformanceReviewItemRouter.DELETE("deletePerformanceReviewItemByIds", pas.DeletePerformanceReviewItemByIds)
		PerformanceReviewItemRouter.POST("getPerformanceReviewListById", pas.GetPerformanceReviewListById) 
		
		
	}
}
