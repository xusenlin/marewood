import request from "../utils/request.js";

/**
 * 用户注册
 * @param params
 * @returns {AxiosPromise}
 */
export function register(params) {
  return request({
    url: "/register",
    method: "post",
    data: params
  });
}

/**
 * 用户登陆
 * @param params
 * @returns {AxiosPromise}
 */
export function login(params) {
  return request({
    url: "/login",
    method: "post",
    data: params
  });
}

/**
 * 用户列表
 * @param params
 * @returns {AxiosPromise}
 */
export function users(params) {
  return request({
    url: "/users",
    method: "get",
    params: params
  });
}

/**
 * 删除用户
 * @param params
 * @returns {AxiosPromise}
 */
export function deleteUsers(params) {
  return request({
    url: "/user/delete",
    method: "get",
    params: params
  });
}

/**
 * 删除用户
 * @param params
 * @returns {AxiosPromise}
 */
export function roleEdit(params) {
  return request({
    url: "/user/role_edit",
    method: "get",
    params: params
  });
}
