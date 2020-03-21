import request from '../utils/request.js'

/**
 * 用户注册
 * @param params
 * @returns {AxiosPromise}
 */
export function register(params) {
    return request({
        url:  '/register',
        method: 'post',
        data: params
    })
}

/**
 * 用户登陆
 * @param params
 * @returns {AxiosPromise}
 */
export function login(params) {
    return request({
        url:  '/login',
        method: 'post',
        data: params
    })
}
