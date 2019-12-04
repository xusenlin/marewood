import Axios from 'axios'
import Config from '../config/index'
import {getToken} from '../utils/dataStorage'
import Toast from '../components/snackbar/index'


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
        config.headers['Authorization'] = getToken() || '';
        return config
    },
    error => {
        Promise.reject(error)
    }
);


service.interceptors.response.use(
    res => {
        if (res.status !== 200) {
            return Promise.reject(res)
        } else {
            if (res.data.status !== true) {
                Toast.error(res.data.msg);
                return Promise.reject(res);
            }
            return res.data.data
        }
    },
    error => {
        return Promise.reject(error)
    }
);

export default service
