import service from '@/utils/request'

export const getLastPRICreatePRIU = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/getLastPRICreatePRIU",
        method: 'post',
        data
    })
}