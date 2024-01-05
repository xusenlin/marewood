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
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

import {Input} from "@/components/ui/input"
import {Textarea} from "@/components/ui/textarea"
import {Label} from "@/components/ui/label"
import {useState, forwardRef, useImperativeHandle, ChangeEvent} from "react";
import {toast} from "sonner";
import {Repository,edit} from "@/api/repo.ts";


export type RepoFormType = {
  EditForm: (repo: Repository | null) => void;
}
interface RepoFormProps {
  onOk: () => void;
}
const initialForm:Partial<Repository> = {
  id:0,
  name:"",
  url:"",
  userName:"",
  password:"",
  dependTools:"",
  description:"",
  defaultBranch:"master",
}

const RepoForm = forwardRef<RepoFormType | undefined,RepoFormProps>(({onOk}, ref) => {
  const [open, setOpen] = useState(false)

  const [form, setForm] = useState<Partial<Repository>>({...initialForm})

  const handleFormChange = (event:ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = event.target;
    setForm((prevForm) => ({ ...prevForm, [name]: value }));
  };

  const EditForm = (repo: Repository | null) => {
    if (repo) {
      setForm({
        id:repo.id,
        name:repo.name,
        url:repo.url,
        userName:repo.userName,
        password:repo.password,
        dependTools:repo.dependTools,
        description:repo.description,
        defaultBranch:repo.defaultBranch,
      })
    } else {
      setForm({...initialForm})
    }
    setOpen(true)
  }
  const onChangeTools = (value:string)=>{
    setForm((prevForm) => ({ ...prevForm, dependTools: value }));
  }

  const onChangeBranch = (value:string)=>{
    setForm((prevForm) => ({...prevForm, defaultBranch: value }));
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
            <DialogTitle>Edit Repo</DialogTitle>
            <DialogDescription>
              {form.id ? `Edit Repo ID:${form.id}` : "Creat New Repo"}
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4 dialog-form">
            <div className="grid gap-4">
              <Label htmlFor="name">Name:</Label>
              <Input
                  id="name"
                  name="name"
                  onChange={handleFormChange}
                  value={form.name}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="url">Url(http[s]):</Label>
              <Input
                  id="url"
                  name="url"
                  onChange={handleFormChange}
                  value={form.url}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="dependTools">DefaultBranch:</Label>
              <Select onValueChange={onChangeBranch} defaultValue={form.defaultBranch}>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder={form.defaultBranch} />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="main">main</SelectItem>
                  <SelectItem value="master">master</SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="userName">UserName:</Label>
              <Input
                  id="userName"
                  name="userName"
                  onChange={handleFormChange}
                  value={form.userName}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="password">Password:</Label>
              <Input
                  id="password"
                  name="password"
                  onChange={handleFormChange}
                  value={form.password}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="dependTools">DependTools:</Label>
              {/*<Input*/}
              {/*    id="dependTools"*/}
              {/*    name="dependTools"*/}
              {/*    onChange={handleFormChange}*/}
              {/*    value={form.dependTools}*/}
              {/*/>*/}
              <Select onValueChange={onChangeTools} defaultValue={form.dependTools}>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder="NPM" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="npm">NPM</SelectItem>
                  <SelectItem value="yarn">YARN</SelectItem>
                  <SelectItem value="pnpm">PNPM</SelectItem>
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


export {RepoForm}
