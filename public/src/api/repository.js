import request from "../utils/request.js";

/**
 * 更新仓库代码
 * @param params
 * @returns {AxiosPromise}
 */
export function gitPull(params) {
  return request({
    url: "/repository/git_pull",
    method: "get",
    params: params
  });
}

/**
 * 通过仓库 ID 获取分支
 * @param params
 * @returns {AxiosPromise}
 */
export function gitBranch(params) {
  return request({
    url: "/repository/git_branch",
    method: "get",
    params: params
  });
}

/**
 * 通过仓库 ID 裁剪支
 * @param params
 * @returns {AxiosPromise}
 */
export function pruneBranch(params) {
  return request({
    url: "/repository/prune_branch",
    method: "get",
    params: params
  });
}

/**
 * 删除仓库依赖
 * @param params
 * @returns {AxiosPromise}
 */
export function deleteDepend(params) {
  return request({
    url: "/repository/delete_depend",
    method: "get",
    params: params
  });
}

/**
 * 获取全部仓库
 * @param params
 * @returns {AxiosPromise}
 */
export function repositories(params) {
  return request({
    url: "/repositories",
    method: "get",
    params: params
  });
}

/**
 * 获取仓库分页
 * @param params
 * @returns {AxiosPromise}
 */
export function repositoryFind(params) {
  return request({
    url: "/repository_find",
    method: "get",
    params: params
  });
}

/**
 * 删除仓库
 * @param params
 * @returns {AxiosPromise}
 */
export function destroy(params) {
  return request({
    url: "/repository/delete",
    method: "get",
    params: params
  });
}

/**
 * 编辑
 * @param params
 * @returns {AxiosPromise}
 */
export function create(params) {
  return request({
    url: "/repository/create",
    method: "post",
    data: params
  });
}

/**
 * 获取仓库可执行脚本
 * @param params
 * @returns {AxiosPromise}
 */
export function getScript(params) {
  return request({
    url: "/repository/get_script",
    method: "get",
    params: params
  });
}

/**
 * 更新仓库白名单字段
 * @param params
 * @returns {AxiosPromise}
 */
export function UpdateField(params) {
  return request({
    url: "/repository/update_field",
    method: "get",
    params: params
  });
}
