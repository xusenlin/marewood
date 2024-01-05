import request, {baseURL, Response} from "@/lib/request.ts"
import {PaginationParams, RespPagination} from "@/commonType/pagination.ts";
import {getToken} from "@/lib/utils.ts";

export type Repository = {
  id:number|null
  name:string
  description:string
  url:string
  userName:string
  password:string
  status:number
  taskStatus:number
  dependTools:string
  defaultBranch:string
  terminalInfo:string
}

export type RepoParams = Partial<Repository>&Partial<PaginationParams>

export const repositories = (params:RepoParams) => {
  return request<RespPagination<Repository>>({
    url: '/v1/repositories',
    method: 'get',
    params
  })
}

export type RepositoryLess = {
  id:number
  name:string
  url:string
  tip:string
}

export const repositoryList = () => {
  return request<Response<RepositoryLess[]>>({
    url: '/v1/repository/list',
    method: 'get',
  })
}



export const edit = (data: Partial<Repository>) => {
  return request<Response<number>>({
    url: '/v1/repository',
    method: 'put',
    data
  })
}

export const destroy = (id: number) => {
  return request<Response<number>>({
    url: `/v1/repository/${id}`,
    method: 'delete',
  })
}

export const gitPull = (id: number) => {
  return request<Response<string>>({
    url: `/v1/repository/git_pull/${id}`,
    method: 'get',
  })
}
export const gitCheckoutDot = (id: number) => {
  return request<Response<string>>({
    url: `/v1/repository/git_checkout_dot/${id}`,
    method: 'get',
  })
}

export const branch = (repositoryId: number) => {
  return request<Response<string[]>>({
    url: `/v1/repository/branch/${repositoryId}`,
    method: 'get',
  })
}
export type PageJson  ={
  name:string
  version:string
  private:string
  dependencies:{[key:string]:string}
  devDependencies:{[key:string]:string}
  scripts:{[key:string]:string}
}

export const packageJson = (repositoryId: number) => {
  return request<Response<PageJson>>({
    url: `/v1/repository/package_json/${repositoryId}`,
    method: 'get',
  })
}


export const getEventUrl = ()=>`${baseURL}/v1/event/repository?token=${getToken()}`
