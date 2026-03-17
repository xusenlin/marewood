import request, {Response} from "@/lib/request.ts";

export interface History {
    id?: number;
    taskId: number;
    name: string;
    gitHash: string;
    userId: number;
    userName: string;
    createdAt?: string;
}


export const getHistoryByTaskId = (taskId: number) => {
    return request<Response<History[]>>({
        url: "/v1/histories",
        method: "get",
        params: {taskId}
    });
}

export const restoreHistory = (id: number) => {
    return request<Response<number>>({
        url: `/v1/history/restore/${id}`,
        method: "post"
    });
}
