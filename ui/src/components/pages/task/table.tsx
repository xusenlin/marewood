import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {Task, TaskParams, tasks} from "@/api/task.ts";
import {ForwardedRef, forwardRef, useEffect, useImperativeHandle, useRef, useState} from "react";
import {RowActions} from "@/components/pages/task/row-actions.tsx";
import {Pagination} from "@/commonType/pagination.ts";
import {toast} from "sonner";
import { Status } from "./status.tsx"
import {TerminalOut} from "@/components/ui/terminal-out.tsx";
import {Badge} from "@/components/ui/badge.tsx";
import {TooltipWrap} from "@/components/ui/tooltipWrap.tsx";
import {Album, Lock} from "lucide-react";
import {PaginationComponent} from "@/components/ui/pagination.tsx";
import {FormActionType, TaskForm} from "@/components/pages/task/form.tsx";

const columns:{id:keyof Task,title:string}[]= [
  {
    id:"id",
    title: "ID"
  },
  {
    id: "name",
    title: "Name"
  },
  // {
  //   id: "repositoryId",
  //   title: "Repository"
  // },
  {
    id: "commitHash",
    title: "Hash"
  },
  {
    id: "status",
    title: "Status"
  },
  {
    id: "branch",
    title: "Branch"
  },
  {
    id: "runTotal",
    title: "RunTotal"
  },
  {
    id: "terminalInfo",
    title: "TerminalOut"
  },
  {
    id: "description",
    title: "Description"
  },
  {
    id: "buildDir",
    title: "BuildDir"
  },
  {
    id: "buildCommand",
    title: "BuildCmd"
  },
]
const cellProps = {
  className: "py-2 px-4 text-center"
}
export type ActionType = {
  reload: (resetPageIndex?: boolean) => void;
  addRow: () => void;
} | undefined

export type TaskTableProps = {
  prams: Partial<Task>
  ref?: ActionType
}

const TableComponent = (props: TaskTableProps, ref: ForwardedRef<ActionType>) => {
  const [data, setData] = useState<Pagination<Task> | null>(null)
  const [params, setParams] = useState<TaskParams>({pageNum: 1, pageSize: 7})
  const formRef = useRef<FormActionType>();
  const isFirstRender = useRef(true);

  useEffect(() => {
    setParams({...params, ...props.prams, pageNum: 1})
  }, [props.prams])

  useEffect(() => {
    if (isFirstRender.current) {
      isFirstRender.current = false;
    } else {
      getTaskData();
    }
  }, [params])

  const getTaskData = () => {
    tasks(params).then(res => {
      setData(res.data.data);
    }).catch(e => {
      toast.error(e)
    })
  }

  const reload = (resetPageIndex = true) => {
    if (resetPageIndex) {
      setParams({...params, pageNum: 1})
    } else {
      getTaskData()
    }
  }
  const addRow = () => {
    formRef.current?.editForm(null);
  }
  useImperativeHandle(ref, () => ({reload,addRow}))


  return (data &&
      <>
        <div  className="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow key={1}>
                {columns.map(c => (<TableHead key={c.id} className="h-10 text-center">{c.title}</TableHead>))}
                <TableHead className="h-10 text-center">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {
                data.list.map(t => {
                  return (
                      <TableRow key={t.id}>
                        <TableCell {...cellProps}>
                          <div className="flex items-center">
                            {t.id}{t.private && <Lock size={12} color="#4caf50"/> }
                          </div>
                        </TableCell>
                        <TableCell {...cellProps} className="min-w-[200px] whitespace-nowrap">
                            {t.tag}
                            {t.name}
                        </TableCell>
                        {/*<TableCell {...cellProps}>*/}
                        {/*  {t.repositoryId}*/}
                        {/*</TableCell>*/}
                        <TableCell {...cellProps}>
                          {t.commitHash}
                        </TableCell>
                        <TableCell {...cellProps}>
                          <Status status={t.status}/>
                        </TableCell>
                        <TableCell {...cellProps}>
                          {t.branch}
                        </TableCell>
                        <TableCell {...cellProps}>
                          {t.runTotal}
                        </TableCell>
                        <TableCell {...cellProps}>
                          <div className="flex w-full justify-center items-center">
                            <TerminalOut
                                title={`TaskID:${t.id}`}
                                description={
                                  <div className="flex items-center space-x-4">
                                    <div className="flex items-center space-x-2"><span>Status:</span><Status  status={t.status}/></div>
                                    <div className="flex items-center space-x-2"><span>Hash:</span><Badge variant="secondary">{t.commitHash}</Badge></div>
                                    <div className="flex items-center space-x-2"><span>Branch:</span><Badge variant="secondary">{t.branch}</Badge></div>
                                  </div>
                                }
                                out={t.terminalInfo}/>
                          </div>
                        </TableCell>
                        <TableCell {...cellProps}>
                          <div className="flex justify-center items-center">
                            <TooltipWrap tip={t.description} child={<Album size={18}/>}/>
                          </div>
                        </TableCell>
                        <TableCell {...cellProps}>
                          {t.buildDir}
                        </TableCell>
                        <TableCell {...cellProps} className="min-w-[180px]">
                          <div className="w-full text-center">{t.buildCommand}</div>
                        </TableCell>
                        <TableCell {...cellProps} className="min-w-[160px]">
                          <RowActions
                              row={t}
                              edit={()=>{formRef.current?.editForm(t);}}
                              reload={()=>{reload()}}/>
                        </TableCell>
                      </TableRow>
                  )
                })
              }
            </TableBody>
          </Table>
        </div>
        <div className="mt-4">
          <PaginationComponent pageNum={data.pageNum} pageSize={data.pageSize} total={data.total} onPageChange={v => {
            setParams({...params, pageNum: v})
          }}/>
        </div>
        <TaskForm ref={formRef} reload={()=>{reload()}}/>
      </>

  )
}


export const TaskTable = forwardRef(TableComponent)
