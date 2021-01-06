import request from "../utils/request.js";

/**
 * 系统信息
 * @param params
 * @returns {AxiosPromise}
 */
export function systemInfo(params) {
  return request({
    url: "/system/info",
    method: "get",
    params: params
  });
}
