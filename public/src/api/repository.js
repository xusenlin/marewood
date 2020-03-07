import request from '../utils/request.js'



/**
 * 更新仓库代码
 * @param params
 * @returns {AxiosPromise}
 */
export function gitPull(params) {
    return request({
        url:  '/repository/git_pull',
        method: 'get',
        params: params
    })
}


/**
 * 删除仓库依赖
 * @param params
 * @returns {AxiosPromise}
 */
export function deleteDepend(params) {
    return request({
        url:  '/repository/delete_depend',
        method: 'get',
        params: params
    })
}




/**
 * 获取全部仓库
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


/**
 * 删除仓库
 * @param params
 * @returns {AxiosPromise}
 */
export function destroy(params) {
    return request({
        url:  '/repository/delete',
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
        url:  '/repository/create',
        method: 'post',
        data: params
    })
}
