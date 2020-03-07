import ApiUrl from './url.js'

export default {
    apiUrl: ApiUrl,
    apiPrefix: 'v1',
    timeout: 10000,
    requestRetry:4,
    requestRetryDelay:800,
    siteName: "FEAutoDeploy",
    systemInfoKey: "FEAutoDeploy_SYSTEM_INFO",
    userInfoKey:"FEAutoDeploy_USER_INFO",//添加前缀（FEAutoDeploy）可以在一个域名下存在多个应用
    tokenKey: "FEAutoDeploy_ACCESS_TOKEN",
}
