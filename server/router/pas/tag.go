package pas

import (
	"gin-vue-admin/api/v1/pas"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(Router *gin.RouterGroup) {
	TagRouter := Router.Group("Tag").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		TagRouter.POST("createTag", pas.CreateTag)   // 新建Tag
		TagRouter.DELETE("deleteTag", pas.DeleteTag) // 删除Tag
		TagRouter.DELETE("deleteTagByIds", pas.DeleteTagByIds) // 批量删除Tag
		TagRouter.PUT("updateTag", pas.UpdateTag)    // 更新Tag
		TagRouter.GET("findTag", pas.FindTag)        // 根据ID获取Tag
		TagRouter.GET("getTagList", pas.GetTagList)  // 获取Tag列表
	}
}
