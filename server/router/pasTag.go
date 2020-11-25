package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPasTagRouter(Router *gin.RouterGroup) {
	PasTagRouter := Router.Group("pasTag").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PasTagRouter.POST("createPasTag", v1.CreatePasTag)   // 新建PasTag
		PasTagRouter.DELETE("deletePasTag", v1.DeletePasTag) // 删除PasTag
		PasTagRouter.DELETE("deletePasTagByIds", v1.DeletePasTagByIds) // 批量删除PasTag
		PasTagRouter.PUT("updatePasTag", v1.UpdatePasTag)    // 更新PasTag
		PasTagRouter.GET("findPasTag", v1.FindPasTag)        // 根据ID获取PasTag
		PasTagRouter.GET("getPasTagList", v1.GetPasTagList)  // 获取PasTag列表
	}
}
