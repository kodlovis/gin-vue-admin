package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPerformanceReviewRouter(Router *gin.RouterGroup) {
	PerformanceReviewRouter := Router.Group("PerformanceReview").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PerformanceReviewRouter.POST("createPerformanceReview", pas.CreatePerformanceReview)   // 新建PerformanceReview
		PerformanceReviewRouter.DELETE("deletePerformanceReview", pas.DeletePerformanceReview) // 删除PerformanceReview
		PerformanceReviewRouter.DELETE("deletePerformanceReviewByIds", pas.DeletePerformanceReviewByIds) // 批量删除PerformanceReview
		PerformanceReviewRouter.PUT("updatePerformanceReview", pas.UpdatePerformanceReview)    // 更新PerformanceReview
		PerformanceReviewRouter.GET("findPerformanceReview", pas.FindPerformanceReview)        // 根据ID获取PerformanceReview
		PerformanceReviewRouter.GET("getPerformanceReviewList", pas.GetPerformanceReviewList)  // 获取PerformanceReview列表
		PerformanceReviewRouter.POST("updatePerformanceReviewByInfo", pas.UpdatePerformanceReviewByInfo)
		PerformanceReviewRouter.GET("getLastPerformanceReview", pas.GetLastPerformanceReview)  
		PerformanceReviewRouter.PUT("updatePRStatusById", pas.UpdatePRStatusById)
		PerformanceReviewRouter.POST("getPRBystatus", pas.GetPRBystatus)    
		PerformanceReviewRouter.POST("updatePRStatysByIds", pas.UpdatePRStatysByIds)
		PerformanceReviewRouter.GET("getPRByUser", pas.GetPRByUser)  
		PerformanceReviewRouter.POST("getPRListByUser", pas.GetPRListByUser) 
		PerformanceReviewRouter.PUT("updatePRResult", pas.UpdatePRResult)
		PerformanceReviewRouter.POST("getPRByID", pas.GetPRByID) 
		PerformanceReviewRouter.GET("getPRListWithoutFinishedStatus", pas.GetPRListWithoutFinishedStatus)  
		
		

	}
}
