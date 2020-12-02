import request from "../utils/request.js";

/**
 * 获取全部任务
 * @param params
 * @returns {AxiosPromise}
 */
export function jobs(params) {
  return request({
    url: "/jobs",
    method: "get",
    params: params
  });
}

/**
 * 通过分类ID获取任务
 * @param params
 * @returns {AxiosPromise}
 */
export function jobsFind(params) {
  return request({
    url: "/jobs_find",
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
    url: "/job/create",
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
    url: "/job/update_branch",
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
    url: "/job/delete",
    method: "get",
    params: params
  });
}

/**
 * 运行任务
 * @param params
 * @returns {AxiosPromise}
 */
export function RunJob(params) {
  return request({
    url: "/job/run",
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
    url: "/job/update_field",
    method: "get",
    params: params
  });
}

/**
 * 更新任务描述
 * @param params
 * @returns {AxiosPromise}
 */
export function jobLock(params) {
  return request({
    url: "/job/lock",
    method: "get",
    params: params
  });
}

/**
 * 搜索任务描述
 * @param params
 * @returns {AxiosPromise}
 */
export function jobSerach(params) {
  return request({
    url: "/search_job",
    method: "get",
    params: params
  });
}
