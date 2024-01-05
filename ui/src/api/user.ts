import request, {Response, } from "@/lib/request.ts"
import {User} from "@/store/user.ts";
import {PaginationParams, RespPagination} from "@/commonType/pagination.ts";

type Account = {
  username:string
  password:string
}

export const register = (data: Account) => {
  return request<Response<number>>({
    url: '/v1/register',
    method: 'post',
    data
  })
}

export const login = (data: Account) => {
  return request<Response<User>>({
    url: '/v1/login',
    method: 'post',
    data
  })
}

export type UserParams = Partial<User>&PaginationParams

export const getUsers = (params:UserParams) => {
  return request<RespPagination<User>>({
    url: '/v1/users',
    method: 'get',
    params: params
  })
}

export const destroy = (params: {id: number}) => {
  return request<Response<number>>({
    url: `/v1/users/${params.id}`,
    method: 'delete',
  })
}

export const edit = (data: Partial<User>) => {
  return request<Response<number>>({
    url: '/v1/users',
    method: 'post',
    data
  })
}
