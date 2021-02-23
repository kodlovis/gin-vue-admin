import service from '@/utils/request'

// @Tags PerformanceReview
// @Summary 创建PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "创建PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PerformanceReview/createPerformanceReview [post]
export const createPerformanceReview = (data) => {
     return service({
         url: "/PerformanceReview/createPerformanceReview",
         method: 'post',
         data
     })
 }


// @Tags PerformanceReview
// @Summary 删除PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "删除PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PerformanceReview/deletePerformanceReview [delete]
 export const deletePerformanceReview = (data) => {
     return service({
         url: "/PerformanceReview/deletePerformanceReview",
         method: 'delete',
         data
     })
 }

// @Tags PerformanceReview
// @Summary 删除PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PerformanceReview/deletePerformanceReview [delete]
 export const deletePerformanceReviewByIds = (data) => {
     return service({
         url: "/PerformanceReview/deletePerformanceReviewByIds",
         method: 'delete',
         data
     })
 }

// @Tags PerformanceReview
// @Summary 更新PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "更新PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PerformanceReview/updatePerformanceReview [put]
 export const updatePerformanceReview = (data) => {
     return service({
         url: "/PerformanceReview/updatePerformanceReview",
         method: 'put',
         data
     })
 }


// @Tags PerformanceReview
// @Summary 用id查询PerformanceReview
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PerformanceReview true "用id查询PerformanceReview"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PerformanceReview/findPerformanceReview [get]
 export const findPerformanceReview = (params) => {
     return service({
         url: "/PerformanceReview/findPerformanceReview",
         method: 'get',
         params
     })
 }


// @Tags PerformanceReview
// @Summary 分页获取PerformanceReview列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PerformanceReview列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PerformanceReview/getPerformanceReviewList [get]
 export const getPerformanceReviewList = (params) => {
     return service({
         url: "/PerformanceReview/getPerformanceReviewList",
         method: 'get',
         params
     })
 }

 export const updatePerformanceReviewByInfo = (data) => {
    return service({
        url: "/PerformanceReview/updatePerformanceReviewByInfo",
        method: 'post',
        data
    })
}

export const getLastPerformanceReview = (params) => {
    return service({
        url: "/PerformanceReview/getLastPerformanceReview",
        method: 'get',
        params
    })
}
export const updatePRStatusById = (data) => {
    return service({
        url: "/PerformanceReview/updatePRStatusById",
        method: 'put',
        data
    })
}
export const getPRBystatus = (data) => {
    return service({
        url: "/PerformanceReview/getPRBystatus",
        method: 'post',
        data
    })
}

export const updatePRStatysByIds = (data) => {
    return service({
        url: "/PerformanceReview/updatePRStatysByIds",
        method: 'post',
        data
    })
}

export const getPRByUser = (params) => {
    return service({
        url: "/PerformanceReview/getPRByUser",
        method: 'get',
        params
    })
}
export const getPRListByUser = (data) => {
    return service({
        url: "/PerformanceReview/getPRListByUser",
        method: 'post',
        data
    })
}

export const updatePRResult = (data) => {
    return service({
        url: "/PerformanceReview/updatePRResult",
        method: 'put',
        data
    })
}

export const getPRByID = (data) => {
    return service({
        url: "/PerformanceReview/getPRByID",
        method: 'post',
        data
    })
}

export const getPRListWithoutFinishedStatus = (params) => {
    return service({
        url: "/PerformanceReview/getPRListWithoutFinishedStatus",
        method: 'get',
        params
    })
}