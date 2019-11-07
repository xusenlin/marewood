import request from '../utils/request.js'

/**
 * 获取全部分类
 * @param params
 * @returns {AxiosPromise}
 */
export function repositories(params) {
    return request({
        url:  '/repositories',
        method: 'get',
        params: params
    })
}

