import request from "../utils/request.js";

/**
 * 获取全部任务
 * @param params
 * @returns {AxiosPromise}
 */
export function tasks(params) {
  return request({
    url: "/tasks",
    method: "get",
    params: params
  });
}

/**
 * 通过分类ID获取任务
 * @param params
 * @returns {AxiosPromise}
 */
export function taskFind(params) {
  return request({
    url: "/task_find",
    method: "get",
    params: params
  });
}

/**
 * 任务创建
 * @param params
 * @returns {AxiosPromise}
 */
export function create(params) {
  return request({
    url: "/task/create",
    method: "post",
    data: params
  });
}

/**
 * 修改任务分支
 * @param params
 * @returns {AxiosPromise}
 */
export function updateBranch(params) {
  return request({
    url: "/task/update_branch",
    method: "get",
    params: params
  });
}

/**
 * 删除
 * @param params
 * @returns {AxiosPromise}
 */
export function destroy(params) {
  return request({
    url: "/task/delete",
    method: "get",
    params: params
  });
}

/**
 * 运行任务
 * @param params
 * @returns {AxiosPromise}
 */
export function RunTask(params) {
  return request({
    url: "/task/run",
    method: "get",
    params: params
  });
}

/**
 * 更新任务白名单字段
 * @param params
 * @returns {AxiosPromise}
 */
export function UpdateField(params) {
  return request({
    url: "/task/update_field",
    method: "get",
    params: params
  });
}

/**
 * 更新任务描述
 * @param params
 * @returns {AxiosPromise}
 */
export function TaskLock(params) {
  return request({
    url: "/task/lock",
    method: "get",
    params: params
  });
}
