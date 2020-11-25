package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags PasTag
// @Summary 创建PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "创建PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasTag/createPasTag [post]
func CreatePasTag(c *gin.Context) {
	var pasTag model.PasTag
	_ = c.ShouldBindJSON(&pasTag)
	if err := service.CreatePasTag(pasTag); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PasTag
// @Summary 删除PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "删除PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pasTag/deletePasTag [delete]
func DeletePasTag(c *gin.Context) {
	var pasTag model.PasTag
	_ = c.ShouldBindJSON(&pasTag)
	if err := service.DeletePasTag(pasTag); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PasTag
// @Summary 批量删除PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pasTag/deletePasTagByIds [delete]
func DeletePasTagByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePasTagByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PasTag
// @Summary 更新PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "更新PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pasTag/updatePasTag [put]
func UpdatePasTag(c *gin.Context) {
	var pasTag model.PasTag
	_ = c.ShouldBindJSON(&pasTag)
	if err := service.UpdatePasTag(&pasTag); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PasTag
// @Summary 用id查询PasTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasTag true "用id查询PasTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pasTag/findPasTag [get]
func FindPasTag(c *gin.Context) {
	var pasTag model.PasTag
	_ = c.ShouldBindQuery(&pasTag)
	if err, repasTag := service.GetPasTag(pasTag.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repasTag": repasTag}, c)
	}
}

// @Tags PasTag
// @Summary 分页获取PasTag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PasTagSearch true "分页获取PasTag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pasTag/getPasTagList [get]
func GetPasTagList(c *gin.Context) {
	var pageInfo request.PasTagSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPasTagInfoList(pageInfo); err != nil {
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
