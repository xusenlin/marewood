import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {PencilRuler, Trash2,ShieldCheck,ShieldBan} from "lucide-react";
import {ForwardedRef, forwardRef, useEffect, useImperativeHandle, useRef, useState} from "react";
import {User} from "@/store/user.ts";
import {toast} from "sonner";
import {getUsers, UserParams} from "@/api/user.ts";
import { DeleteDialog ,AlertDialogType} from "./delete-dialog.tsx"
import {Button} from "@/components/ui/button.tsx";
import {UserForm, UserFormType} from "@/components/pages/user/form.tsx";
import {Pagination} from "@/commonType/pagination.ts"
import {PaginationComponent} from "@/components/ui/pagination.tsx";
import {Role} from "@/components/ui/role.tsx";


const columns: { id: keyof User, title: string }[] = [
  {
    id: "id",
    title: "ID"
  },
  {
    id: "username",
    title: "Username"
  },
  {
    id: "role",
    title: "Role"
  },
  {
    id: "status",
    title: "Status"
  }
]
const cellProps = {
  className: "py-2 px-4 text-center"
}

export type ActionType ={
  reload: (resetPageIndex?: boolean) => void;
} | undefined


export type UserTableProps = {
  prams:Partial<User>
  ref?:ActionType
}


const TableComponent = (props:UserTableProps,ref:ForwardedRef<ActionType>) => {

  const [data, setData] = useState<Pagination<User>|null>(null)
  const [params, setParams] = useState<UserParams>({pageNum:1,pageSize:7})

  const userFormRef = useRef<UserFormType>();
  const deleteDialogRef = useRef<AlertDialogType>();

  const isFirstRender = useRef(true);

  useEffect(() => {
    setParams({...params,...props.prams,pageNum:1})
  }, [props.prams])

  useEffect(() => {
    if (isFirstRender.current) {
      isFirstRender.current = false;
    } else {
      getUserData();
    }
  }, [params])

  const getUserData = ()=>{
    getUsers(params).then(res => {
      setData(res.data.data);
    }).catch(e => {
      toast.error(e)
    })
  }


  const reload = (resetPageIndex = true) => {
    if(resetPageIndex){
      setParams({...params,pageNum:1})
    }else {
      getUserData()
    }
  }

  useImperativeHandle(ref, () => ({reload}))

  return ( data &&
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
                        {
                          columns.map(c => {
                            switch (c.id) {
                              case "status":
                                return (
                                    <TableCell key={t.id + c.id} {...cellProps}>
                                      <div className="flex item-center justify-center">
                                        {t.status === 1 ? <ShieldCheck size={18} color="#4caf50"/> : <ShieldBan size={18} color="#ff9800" />}
                                      </div>
                                    </TableCell>
                                )
                              case "role":
                                return (
                                    <TableCell key={t.id + c.id} {...cellProps}>
                                      <Role id={t.role}/>
                                    </TableCell>
                                )
                              default:
                                return (
                                    <TableCell key={t.id + c.id} {...cellProps}>
                                      {t[c.id]}
                                    </TableCell>
                                );
                            }
                          })
                        }
                        <TableCell {...cellProps}>
                          <Button
                              onClick={() => {
                                userFormRef.current?.EditForm(t)
                              }}
                              variant="ghost"
                              size="icon"
                              className="hover:bg-primary/0"
                          >
                            <PencilRuler size={20}/>
                          </Button>
                          <Button
                              onClick={() => {
                                deleteDialogRef.current?.Confirm(t.id)
                              }}
                              variant="ghost"
                              size="icon"
                              className="hover:bg-primary/0"
                          >
                            <Trash2 size={20}/>
                          </Button>
                        </TableCell>
                      </TableRow>
                  )
                })
              }
            </TableBody>
          </Table>
        </div>

        <div className="mt-4">
          <PaginationComponent pageNum={data.pageNum} pageSize={data.pageSize} total={data.total} onPageChange={v=>{setParams({...params,pageNum:v})}}/>
        </div>
        <DeleteDialog ref={deleteDialogRef} onDelete={getUserData}/>
        <UserForm ref={userFormRef} onOk={getUserData}/>
      </>
  )
}

export const UserTable = forwardRef(TableComponent)
