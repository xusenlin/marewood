import {Response} from "@/lib/request.ts";

export type PaginationParams = {
  pageNum: number
  pageSize: number
}


export type Pagination<T> = {
  total: number
  list: T[]
  totalPage: number
} & PaginationParams


export type RespPagination<T> = Response<Pagination<T>>
