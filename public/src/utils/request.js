import Axios from 'axios'
import Config from '../config/index'
import {getToken} from '../utils/dataStorage'
import Snackbar from '../components/snackbar/index'


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
            Snackbar.error("status code is not 200");
            return Promise.reject(res)
        } else {
            if (res.data.status !== true) {
                Snackbar.warning(res.data.msg);
                return Promise.reject(res);
            }
            return res.data.data
        }
    },
    error => {
        Snackbar.error(error.message);
        return Promise.reject(error)
    }
);

export default service
