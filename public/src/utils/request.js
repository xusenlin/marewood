import Axios from "axios";
import NProgress from "nprogress";
import { useApiUrl } from "../config/url.js";
import Config from "../config/index";
import { getToken } from "../utils/dataStorage";
import Snackbar from "../components/snackbar/index";

const service = Axios.create({
  baseURL: useApiUrl + "/" + Config.apiPrefix,
  headers: {
    Accept: "*/*"
  },
  timeout: Config.timeout
});

service.defaults.retry = Config.requestRetry;
service.defaults.retryDelay = Config.requestRetryDelay;

service.interceptors.request.use(
  config => {
    NProgress.start();
    config.headers["Authorization"] = getToken() || "";
    return config;
  },
  error => {
    Promise.reject(error);
  }
);

service.interceptors.response.use(
  res => {
    setTimeout(() => {
      NProgress.done();
    }, 300);
    if (res.status !== 200) {
      Snackbar.error("Status Code Is Not 200");
      return Promise.reject(res);
    } else {
      if (res.data.status !== true) {
        if (res.data.msg.indexOf("token is expired") !== -1) {
          Snackbar.warning("登陆已经失效，请重新登陆");
          window.location.href = "./#/login";
          return Promise.reject(res);
        } else {
          Snackbar.warning(res.data.msg);
          return Promise.reject(res);
        }
      }
      return res.data.data;
    }
  },
  error => {
    Snackbar.error(error.message);
    NProgress.done();
    return Promise.reject(error);
  }
);

export default service;
