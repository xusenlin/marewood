import {Button} from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Checkbox } from "@/components/ui/checkbox.tsx"
import {Input} from "@/components/ui/input"
import {Textarea} from "@/components/ui/textarea"
import {Label} from "@/components/ui/label"
import {useState, forwardRef, useImperativeHandle, ChangeEvent, ForwardedRef, useEffect} from "react";
import {edit, Task} from "@/api/task.ts";
import {branch, packageJson, RepositoryLess, repositoryList} from "@/api/repo.ts";
import {toast} from "sonner";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";
import {SelectTag} from "@/components/pages/task/select-tag.tsx";


export type FormActionType = {
  editForm: (task: Task | null) => void;
}|undefined

const initialForm = {
  id: 0,
  name: "",
  repositoryId:0,
  branch: "",
  alias:"",
  buildDir: "dist",
  buildCommand:"",
  private:false,
  tag:"",
  description:""
}

const TaskFormComponent =  (props: { reload: () => void; }, ref: ForwardedRef<FormActionType>) => {
  const [open, setOpen] = useState(false)
  const [repositories, setRepositories] = useState<RepositoryLess[]>([])
  const [branches, setBranches] = useState<string[]>([])
  const [buildCmd, setBuildCmd] = useState<{k:string,v:string}[]>([])

  const [form, setForm] = useState<Partial<Task>>({...initialForm})

  const handleFormChange = (event:ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = event.target;
    setForm((prevForm) => ({ ...prevForm, [name]: value }));
  };
  useEffect(() => {
    repositoryList().then(r => {
      setRepositories(r.data.data)
    }).catch(e=>{
      toast.error(e)
    })

  }, [])

  const editForm = (task: Task | null) => {
    if (task) {
      const newForm: Partial<Task> = {...initialForm}
      for (const key in newForm) {
        newForm[key]= task[key]
      }
      initialRepositoryData(task.repositoryId)
      setForm(newForm)
    } else {
      setForm({...initialForm})
    }
    setOpen(true)
  }
  const onChangeRepository = (repositoryId:number) => {
    initialRepositoryData(repositoryId)
    setForm((prevForm) => ({...prevForm, repositoryId,branch:"" }));

    const name = repositories.find(r=>r.id===repositoryId)?.tip
    setForm((prevForm) => ({...prevForm, alias:name }));
  }
  const initialRepositoryData = (repositoryId:number)=>{
    branch(repositoryId).then(r=>{
      setBranches(r.data.data)
    }).catch(e=>{
      toast.error(e)
    })
    packageJson(repositoryId).then(r=>{
      const scripts = r.data.data.scripts
      const command:{k:string,v:string}[] = []
      for (const s in scripts){
        if (s.includes("build")){
          command.push({k:s,v:scripts[s]})
        }
      }
      setBuildCmd(command)
    }).catch(e=>{
      toast.error(e)
    })
  }
  const submit = () => {
    edit(form).then(()=>{
      toast.success("Edit Task Success")
      props.reload()
      setOpen(false)
    }).catch(e=>{
      toast.error(e)
    })
  }

  useImperativeHandle(ref, () => ({editForm}))
  return (
      <Dialog open={open} onOpenChange={v => {
        setOpen(v)
      }}>
        <DialogContent className="sm:max-w-[625px]">
          <DialogHeader>
            <DialogTitle>Edit Task</DialogTitle>
            <DialogDescription>
              {form.id ? `Edit Task ID:${form.id}` : "Creat New Task"}
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4 dialog-form">
            <div className="grid gap-4">
              <Label htmlFor="name">
                Name
              </Label>
              <Input
                  id="name"
                  name="name"
                  onChange={handleFormChange}
                  value={form.name}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="repository">
                Repository
              </Label>
              <Select onValueChange={id=>{onChangeRepository(parseInt(id))}}>
                <SelectTrigger className="w-full">
                  { form.repositoryId ?
                    repositories.map(r=>{
                      if (r.id === form.repositoryId) {
                        return <SelectValue key={r.id} placeholder={r.name} />
                      }
                    }) :
                    <SelectValue placeholder="" />
                  }
                  {/*<SelectValue placeholder={form.repositoryId} />*/}
                </SelectTrigger>
                <SelectContent>
                  { repositories.map(r => <SelectItem key={r.id} value={r.id.toString()}>{r.name}({r.tip})</SelectItem>) }
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="branch">
                Branch
              </Label>
              <Select onValueChange={b=>{setForm((prevForm) => ({...prevForm, branch:b }))}}>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder={form.branch} />
                </SelectTrigger>
                <SelectContent>
                  { branches.length === 0 && <div className="p-2 text-center text-sm text-muted-foreground">Please select a repository first.</div>}
                  { branches.map(b => <SelectItem key={b} value={b}>{b}</SelectItem>) }
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="alias">
                Alias(Aliases are used for URL access to directories.)
              </Label>
              <Input
                  id="alias"
                  name="alias"
                  onChange={handleFormChange}
                  value={form.alias}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="buildDir">
                BuildDir
              </Label>
              <Input
                  id="buildDir"
                  name="buildDir"
                  onChange={handleFormChange}
                  value={form.buildDir}
              />
            </div>
            <div className="grid gap-4">
              <Label htmlFor="buildCommand">
                BuildCommand(Npm command with build keyword)
              </Label>
              <Select onValueChange={b=>{setForm((prevForm) => ({...prevForm, buildCommand:b }))}}>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder={form.buildCommand} />
                </SelectTrigger>
                <SelectContent>
                  { buildCmd.length === 0 && <div className="p-2 text-center text-sm text-muted-foreground">Please select a repository first.</div>}
                  { buildCmd.map(c => <SelectItem key={c.k} value={c.k}>{c.k}({c.v})</SelectItem>) }
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="private">
                Private
              </Label>
              <div className="flex items-center ">
                <Checkbox id="terms" checked={form.private} onCheckedChange={v => setForm((prevForm) => ({...prevForm, private:v as boolean }))} />
                <label
                    htmlFor="terms"
                    className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 ml-2"
                >
                  This task is visible only to itself.
                </label>
              </div>
            </div>
            <div className="grid gap-4">
              <Label htmlFor="tag">
                Tag
              </Label>
              <div className="flex">
                <Input
                    id="tag"
                    name="tag"
                    className="mr-2"
                    onChange={handleFormChange}
                    value={form.tag}
                />
                <SelectTag onChange={v=>{setForm((prevForm) => ({...prevForm, tag:v }))}}/>
              </div>

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
            <Button type="submit" onClick={submit}>Save</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
  )
}



export const TaskForm = forwardRef(TaskFormComponent)
