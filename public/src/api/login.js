import request from '../utils/request.js'

/**
 * 用户登陆
 * @param params
 * @returns {AxiosPromise}
 */
export function login(params) {
    return request({
        url:  '/userService/api/manager/login',
        method: 'post',
        data: params
    })
}
