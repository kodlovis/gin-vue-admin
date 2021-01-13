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

