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
		PerformanceReviewItemRouter.POST("getPerformanceReviewItemListById", pas.GetPerformanceReviewItemListById)
		PerformanceReviewItemRouter.PUT("updatePerformanceReviewItemByInfo", pas.UpdatePerformanceReviewItemByInfo)
		PerformanceReviewItemRouter.POST("getPRItemListByUser", pas.GetPRItemListByUser)
		PerformanceReviewItemRouter.PUT("updatePRItemStatusById", pas.UpdatePRItemStatusById)
		PerformanceReviewItemRouter.POST("getPRItemCount", pas.GetPRItemCount)
		PerformanceReviewItemRouter.POST("updatePRItemStatysByIds", pas.UpdatePRItemStatysByIds)
		PerformanceReviewItemRouter.PUT("updatePRItemStatusByPrId", pas.UpdatePRItemStatusByPrId)
		PerformanceReviewItemRouter.POST("getPRItemListByStatusPrid", pas.GetPRItemListByStatusPrid)
		PerformanceReviewItemRouter.POST("getPRItemListByStatus", pas.GetPRItemListByStatus)
		PerformanceReviewItemRouter.DELETE("deletePRItemById", pas.DeletePRItemById)

	}
}
