import request from '../utils/request.js'

/**
 * 获取记录
 * @param params
 * @returns {AxiosPromise}
 */
export function webHookRecord(params) {
    return request({
        url:  '/webHookRecord',
        method: 'get',
        params: params
    })
}
