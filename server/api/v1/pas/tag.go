package pas

import (
	"gin-vue-admin/global"
    mp "gin-vue-admin/model/pas"
    rp "gin-vue-admin/model/request/pas"
    "gin-vue-admin/model/response"
    sp "gin-vue-admin/service/pas"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Tag
// @Summary 创建Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "创建Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Tag/createTag [post]
func CreateTag(c *gin.Context) {
	var Tag mp.Tag
	_ = c.ShouldBindJSON(&Tag)
	if err := sp.CreateTag(Tag); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Tag
// @Summary 删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Tag/deleteTag [delete]
func DeleteTag(c *gin.Context) {
	var Tag mp.Tag
	_ = c.ShouldBindJSON(&Tag)
	if err := sp.DeleteTag(Tag); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Tag
// @Summary 批量删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Tag/deleteTagByIds [delete]
func DeleteTagByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteTagByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Tag
// @Summary 更新Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "更新Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Tag/updateTag [put]
func UpdateTag(c *gin.Context) {
	var Tag mp.Tag
	_ = c.ShouldBindJSON(&Tag)
	if err := sp.UpdateTag(&Tag); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Tag
// @Summary 用id查询Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "用id查询Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Tag/findTag [get]
func FindTag(c *gin.Context) {
	var Tag mp.Tag
	_ = c.ShouldBindQuery(&Tag)
	if err, reTag := sp.GetTag(Tag.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTag": reTag}, c)
	}
}

// @Tags Tag
// @Summary 分页获取Tag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TagSearch true "分页获取Tag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Tag/getTagList [get]
func GetTagList(c *gin.Context) {
	var pageInfo rp.TagSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetTagInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
