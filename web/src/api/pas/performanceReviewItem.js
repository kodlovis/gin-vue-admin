import service from '@/utils/request'

export const createPerformanceReviewItem = (data) => {
    return service({
        url: "/PerformanceReviewItem/createPerformanceReviewItem",
        method: 'post',
        data
    })
}

export const deletePerformanceReviewItem = (data) => {
    return service({
        url: "/PerformanceReviewItem/deletePerformanceReviewItem",
        method: 'delete',
        data
    })
}
export const deletePRItemById = (data) => {
    return service({
        url: "/PerformanceReviewItem/deletePRItemById",
        method: 'delete',
        data
    })
}

export const deletePerformanceReviewItemByIds = (data) => {
    return service({
        url: "/PerformanceReviewItem/deletePerformanceReviewItemByIds",
        method: 'delete',
        data
    })
}

export const getPerformanceReviewItemListById = (data) => {
    return service({
        url: "/PerformanceReviewItem/getPerformanceReviewItemListById",
        method: 'post',
        data
    })
}


export const updatePerformanceReviewItemByInfo = (data) => {
    return service({
        url: "/PerformanceReviewItem/updatePerformanceReviewItemByInfo",
        method: 'put',
        data
    })
}


export const getPRItemListByUser = (data) => {
    return service({
        url: "/PerformanceReviewItem/getPRItemListByUser",
        method: 'post',
        data
    })
}

export const updatePRItemStatusById = (data) => {
    return service({
        url: "/PerformanceReviewItem/updatePRItemStatusById",
        method: 'put',
        data
    })
}
export const updatePRItemStatusByPrId = (data) => {
    return service({
        url: "/PerformanceReviewItem/updatePRItemStatusByPrId",
        method: 'put',
        data
    })
}
export const getPRItemCount = (data) => {
    return service({
        url: "/PerformanceReviewItem/getPRItemCount",
        method: 'post',
        data
    })
}
export const updatePRItemStatysByIds = (data) => {
    return service({
        url: "/PerformanceReviewItem/updatePRItemStatysByIds",
        method: 'post',
        data
    })
}

export const getPRItemListByStatusPrid = (data) => {
    return service({
        url: "/PerformanceReviewItem/getPRItemListByStatusPrid",
        method: 'post',
        data
    })
}

export const getPRItemListByStatus = (data) => {
    return service({
        url: "/PerformanceReviewItem/getPRItemListByStatus",
        method: 'post',
        data
    })
}
export const getPRItemListByPrids = (data) => {
    return service({
        url: "/PerformanceReviewItem/getPRItemListByPrids",
        method: 'post',
        data
    })
}
export const getLastPRI = (params) => {
    return service({
        url: "/PerformanceReviewItem/getLastPRI",
        method: 'get',
        params
    })
}
