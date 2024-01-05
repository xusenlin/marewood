import axios from 'axios'
import {outLogin} from "@/store/user.ts";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import nprogress from "nprogress"
import {getToken} from "@/lib/utils.ts";

export const baseURL = window.location.origin
//export const baseURL = "http://localhost:8088"

const service = axios.create({baseURL})

export type Response<T> = {
  status: boolean
  data: T
  msg: string
}


service.interceptors.request.use(
    (config) => {
      nprogress.start();
      config.headers.Authorization = getToken()
      return config
    }
)

service.interceptors.response.use(
    (response) => {
      nprogress.done();
      if (response.status !== 200) {
        return Promise.reject(response.statusText)
      }
      if (response.data.status === false) {
        const msg = response.data.msg
        if (msg.indexOf("token is expired") !== -1) {
          outLogin()
        }
        if (msg.indexOf("please log in first") !== -1) {
          outLogin()
        }

        return Promise.reject(msg)
      }
      return response
    },
    (error) => {
      nprogress.done();
      return Promise.reject(error.toString())
    }
)

export default service
