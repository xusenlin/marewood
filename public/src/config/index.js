import ApiUrl from './url.js'

export default {
    apiUrl: ApiUrl,
    apiPrefix: 'api_v1',
    timeout: 5000,
    requestRetry:4,
    requestRetryDelay:800,
    siteName: 'FEAutoDeploy',
    userInfoKey:'FEAutoDeploy_USER_INFO',//添加前缀（FEAutoDeploy）可以在一个域名下存在多个应用
    tokenKey: 'FEAutoDeploy_ACCESS_TOKEN',
}
