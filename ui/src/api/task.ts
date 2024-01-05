import {PaginationParams, RespPagination} from "@/commonType/pagination.ts";
import request, {baseURL, Response} from "@/lib/request.ts";
import {getToken} from "@/lib/utils.ts";


export type Task = {
  id:number|null
  name:string
  description:string
  status:number
  branch:string
  alias:string
  runTotal:number
  repositoryId:number
  buildDir:string
  buildCommand:string
  createdId:number
  private:boolean
  tag:string
  commitHash:string
  terminalInfo:string
  [key: string]: string|number|boolean|null;
}
export type TaskParams = Partial<Task>&Partial<PaginationParams>

export const tasks = (params:TaskParams) => {
  return request<RespPagination<Task>>({
    url: '/v1/tasks',
    method: 'get',
    params
  })
}
export const run = (id:number) => {
  return request<Response<number>>({
    url: '/v1/task/run',
    method: 'get',
    params:{id}
  })
}


export const edit = (data: Partial<Task>) => {
  return request<Response<number>>({
    url: '/v1/task',
    method: 'put',
    data
  })
}


export type TagType = {
  tag:string
  count:number
}

export const allTags = () => {
  return request<Response<TagType[]>>({
    url: '/v1/task/tags',
    method: 'get',
  })
}

export const updateBranch = (params: {id:number,branch:string }) => {
  return request<Response<TagType[]>>({
    url: '/v1/task/update_branch',
    method: 'get',
    params
  })
}

export const destroy = (id: number) => {
  return request<Response<number>>({
    url: `/v1/task/${id}`,
    method: 'delete',
  })
}


export const getEventUrl = ()=>`${baseURL}/v1/event/task?token=${getToken()}`


export const getTaskTarUrl = (id:number)=>`${baseURL}/v1/task/tar?id=${id}`


