import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {PencilRuler, Trash2, Plus, Link2, Album,Download,TimerReset} from "lucide-react";
import {forwardRef, useRef, useState, ForwardedRef, useImperativeHandle, useEffect} from "react";
import {toast} from "sonner";
import {Button} from "@/components/ui/button.tsx";
import {Badge} from "@/components/ui/badge.tsx";
import {Pagination} from "@/commonType/pagination.ts"
import {PaginationComponent} from "@/components/ui/pagination.tsx";
import {Repository, repositories, destroy, gitPull, gitCheckoutDot, RepoParams} from "@/api/repo.ts";
import {RepoFormType, RepoForm} from "@/components/pages/repository/form.tsx";
import {FloatButton} from "@/components/ui/float-button.tsx";
import {Status} from "./status.tsx"
import {TaskStatus} from "./taskStatus.tsx"
import {TooltipWrap} from "@/components/ui/tooltipWrap.tsx";
import {TerminalOut} from "@/components/ui/terminal-out.tsx"
import {ConfirmDialog} from "@/components/ui/confirm-dialog.tsx";

const columns: { id: keyof Repository, title: string }[] = [
  {
    id: "id",
    title: "ID"
  },
  {
    id: "name",
    title: "Name"
  },
  {
    id: "status",
    title: "Status"
  },
  {
    id: "taskStatus",
    title: "TaskStatus"
  },
  {
    id: "url",
    title: "Url"
  },
  {
    id: "dependTools",
    title: "DependTools"
  },
  {
    id: "terminalInfo",
    title: "TerminalInfo"
  },
  {
    id: "description",
    title: "Description"
  },

]
const cellProps = {
  className: "py-2 px-4 text-center"
}

export type ActionType = {
  reload: (resetPageIndex?: boolean) => void;
} | undefined


export type RepoTableProps = {
  prams: Partial<Repository>
  ref?: ActionType
}


const TableComponent = (props: RepoTableProps, ref: ForwardedRef<ActionType>) => {

  const [data, setData] = useState<Pagination<Repository> | null>(null)
  const [params, setParams] = useState<RepoParams>({pageNum: 1, pageSize: 7})

  const formRef = useRef<RepoFormType>();
  const isFirstRender = useRef(true);


  useEffect(() => {
    setParams({...params, ...props.prams, pageNum: 1})
  }, [props.prams])

  useEffect(() => {
    if (isFirstRender.current) {
      isFirstRender.current = false;
    } else {
      getRepoData();
    }
  }, [params])

  const getRepoData = () => {

    repositories(params).then(res => {
      setData(res.data.data);
    }).catch(e => {
      toast.error(e)
    })
  }

  const reload = (resetPageIndex = true) => {
    if (resetPageIndex) {
      setParams({...params, pageNum: 1})
    } else {
      getRepoData()
    }
  }

  const deleteRepo = (id:number)=>{
    destroy(id).then(()=>{
      toast.success("Delete success")
      reload()
    }).catch(e=>{
      toast.error(e)
    })
  }
  const pull = (id:number)=>{
    gitPull(id).then(r=>{
      toast.success(r.data.data)
    }).catch(e=>{
      toast.error(e)
    })
  }

  const checkoutDot  = (id:number)=>{
    gitCheckoutDot(id).then(r=>{
      toast.success(r.data.data)
    }).catch(e=>{
      toast.error(e)
    })
  }

  useImperativeHandle(ref, () => ({reload}))

  return (data &&
      <>
        <div className="rounded-md border">
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
                          {t.id}
                        </TableCell>
                        <TableCell {...cellProps} className="min-w-[160px]">
                          {t.name}
                        </TableCell>
                        <TableCell {...cellProps}>
                          <Status status={t.status}/>
                        </TableCell>
                        <TableCell {...cellProps}>
                          <TaskStatus status={t.taskStatus}/>
                        </TableCell>
                        <TableCell {...cellProps}>
                          <div className="flex justify-center items-center">
                            <TooltipWrap tip={<a target="_blank" href={t.url}>{t.url}</a>} child={<Link2 size={20}/>}/>
                          </div>
                        </TableCell>
                        <TableCell {...cellProps}>
                          <Badge variant="secondary">
                            {t.dependTools.toUpperCase()}
                          </Badge>
                        </TableCell>
                        <TableCell {...cellProps}>
                          <div className="flex w-full justify-center items-center">
                            <TerminalOut
                                title={`RepositoryID:${t.id}`}
                                description={<span className="flex items-center space-x-2">{t.url}</span>}
                                out={t.terminalInfo}/>
                          </div>
                        </TableCell>
                        <TableCell {...cellProps}>
                          <div className="flex justify-center items-center">
                            <TooltipWrap tip={t.description} child={<Album size={18}/>}/>
                          </div>
                        </TableCell>
                        <TableCell {...cellProps} className="min-w-[200px]">
                          <div className="flex items-center w-full justify-center">
                            <Button
                                onClick={() => {
                                  formRef.current?.EditForm(t)
                                }}
                                title="edit"
                                variant="ghost"
                                size="icon"
                                className="hover:bg-primary/0"
                            >
                              <PencilRuler size={20}/>
                            </Button>
                            <Button
                                onClick={() => {pull(t.id!)}}
                                title="git pull"
                                variant="ghost"
                                size="icon"
                                className="hover:bg-primary/0"
                            >
                              <Download size={20}/>
                            </Button>
                            <Button
                                onClick={() => {checkoutDot(t.id!)}}
                                title="git checkout ."
                                variant="ghost"
                                size="icon"
                                className="hover:bg-primary/0"
                            >
                              <TimerReset size={20}/>
                            </Button>

                            <ConfirmDialog child={
                              <Button
                                  title="delete"
                                  variant="ghost"
                                  size="icon"
                                  className="hover:bg-primary/0">
                                <Trash2 size={20}/>
                              </Button>
                            }
                                           onConfirm={()=>{deleteRepo(t.id!)}}
                            />
                          </div>
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
        <RepoForm ref={formRef} onOk={getRepoData}/>
        <FloatButton onClick={() => {
          formRef.current?.EditForm(null)
        }}>
          <Plus size={32}/>
        </FloatButton>
      </>
  )
}


export const RepoTable = forwardRef(TableComponent)
