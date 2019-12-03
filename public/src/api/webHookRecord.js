import request from '../utils/request.js'

/**
 * 获取记录
 * @param params
 * @returns {AxiosPromise}
 */
export function webHookRecord(params) {
    return request({
        url:  '/webHook_record',
        method: 'get',
        params: params
    })
}
