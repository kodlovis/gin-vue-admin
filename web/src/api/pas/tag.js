import service from '@/utils/request'

// @Tags Tag
// @Summary 创建Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "创建Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Tag/createTag [post]
export const createTag = (data) => {
     return service({
         url: "/Tag/createTag",
         method: 'post',
         data
     })
 }


// @Tags Tag
// @Summary 删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Tag/deleteTag [delete]
 export const deleteTag = (data) => {
     return service({
         url: "/Tag/deleteTag",
         method: 'delete',
         data
     })
 }

// @Tags Tag
// @Summary 删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Tag/deleteTag [delete]
 export const deleteTagByIds = (data) => {
     return service({
         url: "/Tag/deleteTagByIds",
         method: 'delete',
         data
     })
 }

// @Tags Tag
// @Summary 更新Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "更新Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Tag/updateTag [put]
 export const updateTag = (data) => {
     return service({
         url: "/Tag/updateTag",
         method: 'put',
         data
     })
 }


// @Tags Tag
// @Summary 用id查询Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "用id查询Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Tag/findTag [get]
 export const findTag = (params) => {
     return service({
         url: "/Tag/findTag",
         method: 'get',
         params
     })
 }


// @Tags Tag
// @Summary 分页获取Tag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Tag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Tag/getTagList [get]
 export const getTagList = (params) => {
     return service({
         url: "/Tag/getTagList",
         method: 'get',
         params
     })
 }