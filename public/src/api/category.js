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


/**
 * 删除分类
 * @param params
 * @returns {AxiosPromise}
 */
export function destroy(params) {
    return request({
        url:  '/category/delete',
        method: 'get',
        params: params
    })
}

/**
 * 编辑
 * @param params
 * @returns {AxiosPromise}
 */
export function create(params) {
    return request({
        url:  '/category/create',
        method: 'post',
        data: params
    })
}

/**
 * 更新分类描述
 * @param params
 * @returns {AxiosPromise}
 */
export function UpdateDesc(params) {
    return request({
        url:  '/category/update_desc',
        method: 'get',
        params: params
    })
}
