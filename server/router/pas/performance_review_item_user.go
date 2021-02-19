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
		PerformanceReviewItemUserRouter.POST("getPRIUByPRIID", pas.GetPRIUByPRIID) 
		PerformanceReviewItemUserRouter.PUT("updatePRIU", pas.UpdatePRIU)  
		PerformanceReviewItemUserRouter.DELETE("removePRIU", pas.RemovePRIU)  
		PerformanceReviewItemUserRouter.POST("createPRIU", pas.CreatePRIU)   
		PerformanceReviewItemUserRouter.POST("getPRIUListByUser", pas.GetPRIUListByUser)
		PerformanceReviewItemUserRouter.POST("updatePRIUStatusByPRIID", pas.UpdatePRIUStatusByPRIID)
		PerformanceReviewItemUserRouter.POST("getPRIUListByStatus", pas.GetPRIUListByStatus)
		PerformanceReviewItemUserRouter.POST("updatePRIUStatustByIds", pas.UpdatePRIUStatusByIds)
		
	}
}
