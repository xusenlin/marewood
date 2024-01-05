import {Button} from "@/components/ui/button"

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { edit } from "@/api/user.ts"
import {Input} from "@/components/ui/input"
import {Textarea} from "@/components/ui/textarea"
import {Label} from "@/components/ui/label"
import {useState, forwardRef, useImperativeHandle, ChangeEvent} from "react";
import {User} from "@/store/user.ts";
import {getRoleArr} from "@/components/ui/role.tsx";
import {toast} from "sonner";


export type UserFormType = {
  EditForm: (user: User | null) => void;
}
interface UserFormProps {
  onOk: () => void;
}
const initialForm = {
  id:0,
  username:"",
  role:10,
  status:1,
  description:""
}

const UserForm = forwardRef<UserFormType | undefined,UserFormProps>(({onOk}, ref) => {
  const [open, setOpen] = useState(false)

  const [form, setForm] = useState<Partial<User>>({...initialForm})

  const handleFormChange = (event:ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = event.target;
    setForm((prevForm) => ({ ...prevForm, [name]: value }));
  };
  const setRole = (role:string) => {
    setForm((prevForm) => ({...prevForm, role:parseInt(role) }))
  }
  const setStatus = (s:string) => {
    setForm((prevForm) => ({...prevForm, status:parseInt(s) }))
  }


  const EditForm = (user: User | null) => {
    if (user) {
      setForm({
        id:user.id,
        username:user.username,
        role:user.role,
        status:user.status,
        description:user.description
      })
    } else {
      setForm({...initialForm})
    }
    setOpen(true)
  }
  const submit = ()=>{
    edit(form).then(r=>{
      setOpen(false)
      onOk()
      toast.success(r.data.msg)
    }).catch(e => {
      toast.error(e)
    })
  }

  useImperativeHandle(ref, () => ({EditForm}))
  return (
      <Dialog open={open} onOpenChange={v => {
        setOpen(v)
      }}>
        <DialogContent className="sm:max-w-[625px]">
          <DialogHeader>
            <DialogTitle>Edit User</DialogTitle>
            <DialogDescription>
              {form.id ? `Edit User ID:${form.id}` : "Creat New User"}
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div className="grid gap-4">
              <Label htmlFor="username">Username:</Label>
              <Input
                  id="username"
                  name="username"
                  onChange={handleFormChange}
                  value={form.username}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="role">Role:</Label>
              <Select defaultValue={form.role+''} onValueChange={setRole}>
                <SelectTrigger>
                  <SelectValue  placeholder="Select a role" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    { getRoleArr().map((item) => (<SelectItem key={item.k} value={item.k+''}>{item.v}</SelectItem>))}
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="role">Status:</Label>
              <Select defaultValue={form.status+''} onValueChange={setStatus}>
                <SelectTrigger>
                  <SelectValue  placeholder="Select a status" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem value="1">Enabled</SelectItem>
                    <SelectItem value="2">Disabled</SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="description">Description:</Label>
              <Textarea
                  id="description"
                  name="description"
                  onChange={handleFormChange}
                  value={form.description}
              />
            </div>
          </div>
          <DialogFooter>
            <Button type="submit" onClick={()=>{submit()}}>Save</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
  )
})


export {UserForm}
