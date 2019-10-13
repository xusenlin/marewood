import request from '../utils/request.js'

/**
 * 获取全部分类
 * @param params
 * @returns {AxiosPromise}
 */
export function categories(params) {
    return request({
        url:  '/categories',
        method: 'get',
        params: params
    })
}
