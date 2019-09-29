import ApiUrl from './url.js'

export default {
    apiUrl: ApiUrl,
    apiPrefix: '',
    timeout: 5000,
    requestRetry:4,
    requestRetryDelay:800,
    siteName: 'MaterialUiAdmin',
    userInfoKey:'STAT_USER_INFO',//添加前缀（STAT）可以在一个域名下存在多个应用
    tokenKey: 'STAT_ACCESS_TOKEN',
    permissionsKey:'STAT_PERMISSION_ARRAY',
}
