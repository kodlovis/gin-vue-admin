import service from '@/utils/request'

export const getLastPRICreatePRIU = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/getLastPRICreatePRIU",
        method: 'post',
        data
    })
}

export const getPRIUByPRIID = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/getPRIUByPRIID",
        method: 'post',
        data
    })
}
export const updatePRIU = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/updatePRIU",
        method: 'put',
        data
    })
}

export const removePRIU = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/removePRIU",
        method: 'delete',
        data
    })
}

export const createPRIU = (data) => {
   return service({
       url: "/PerformanceReviewItemUser/createPRIU",
       method: 'post',
       data
   })
}
export const getPRIUListByUser = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/getPRIUListByUser",
        method: 'post',
        data
    })
}
export const updatePRIUStatusByPRIID = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/updatePRIUStatusByPRIID",
        method: 'post',
        data
    })
}

export const getPRIUListByStatus = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/getPRIUListByStatus",
        method: 'post',
        data
    })
}

export const updatePRIUStatustByIds = (data) => {
    return service({
        url: "/PerformanceReviewItemUser/updatePRIUStatustByIds",
        method: 'post',
        data
    })
}
