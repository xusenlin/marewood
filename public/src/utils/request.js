import Axios from 'axios'
import Config from '../config/index.js'
import {getToken, removeToken} from '../utils/dataStorage.js'


const service = Axios.create({
    baseURL: Config.apiUrl + '/' + Config.apiPrefix,
    headers: {
        'Accept': '*/*'
    },
    timeout: Config.timeout
});

service.defaults.retry = Config.requestRetry;
service.defaults.retryDelay = Config.requestRetryDelay;


service.interceptors.request.use(
    config => {
        if (!config.closeLoading) {
            //加载提示
        }
        // config.headers['businessSource'] = 3;
        // config.headers['terminalPlatform'] = 4;
        // config.headers['requestId'] = new Date().getTime();
        config.headers['Authorization'] = getToken() || '';
        return config
    },
    error => {
        Promise.reject(error)
    }
);


service.interceptors.response.use(
    res => {

        if (!res.config.closeLoading) {
            //关闭loading
        }

        if (res.status !== 200) {
            // notification.warn({
            //     message: '数据返回出错',
            //     description: "请稍后重试"
            // });
            return Promise.reject('error')
        } else {
            if (res.config.closeInterceptors) {
                return res.data
            }

            // if (res.data.code !== 100000) {
            //     notification.warn({
            //         description: res.data.msg
            //     });
            //     if (res.data.code === 100004) {//登录状态失效
            //         removeToken();
            //         setTimeout(_ => {
            //             window.location.href = './#/login'
            //             // Router.push('/login')
            //         }, 1000)
            //     }
            //     return Promise.reject('error');
            // }

            return res.data.data
        }
    },
    error => {
        return Promise.reject(error)
    }
);

export default service
