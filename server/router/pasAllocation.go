package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPasAllocationRouter(Router *gin.RouterGroup) {
	PasAllocationRouter := Router.Group("pasAllocation").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PasAllocationRouter.POST("createPasAllocation", v1.CreatePasAllocation)   // 新建PasAllocation
		PasAllocationRouter.DELETE("deletePasAllocation", v1.DeletePasAllocation) // 删除PasAllocation
		PasAllocationRouter.DELETE("deletePasAllocationByIds", v1.DeletePasAllocationByIds) // 批量删除PasAllocation
		PasAllocationRouter.PUT("updatePasAllocation", v1.UpdatePasAllocation)    // 更新PasAllocation
		PasAllocationRouter.GET("findPasAllocation", v1.FindPasAllocation)        // 根据ID获取PasAllocation
		PasAllocationRouter.GET("getPasAllocationList", v1.GetPasAllocationList)  // 获取PasAllocation列表
	}
}
