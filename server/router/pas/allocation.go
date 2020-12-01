package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAllocationRouter(Router *gin.RouterGroup) {
	AllocationRouter := Router.Group("Allocation").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AllocationRouter.POST("createAllocation", pas.CreateAllocation)   // 新建Allocation
		AllocationRouter.DELETE("deleteAllocation", pas.DeleteAllocation) // 删除Allocation
		AllocationRouter.DELETE("deleteAllocationByIds", pas.DeleteAllocationByIds) // 批量删除Allocation
		AllocationRouter.PUT("updateAllocation", pas.UpdateAllocation)    // 更新Allocation
		AllocationRouter.GET("findAllocation", pas.FindAllocation)        // 根据ID获取Allocation
		AllocationRouter.GET("getAllocationList", pas.GetAllocationList)  // 获取Allocation列表
	}
}
