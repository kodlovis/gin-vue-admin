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