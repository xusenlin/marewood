import request, {Response} from "@/lib/request.ts"

export type Sys = {
  [key: string]: any
}

export const getSysInfo = () => {
  return request<Response<Sys>>({
    url: '/v1/sys',
    method: 'get',
  })
}
