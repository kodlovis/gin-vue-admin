import service from '@/utils/request'

export const createKpiTag = (data) => {
    return service({
        url: "/KpiTag/createKpiTag",
        method: 'post',
        data
    })
}