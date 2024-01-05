import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog"
import {useState,forwardRef,useImperativeHandle} from "react";
import {Share2} from "lucide-react"

import {
  Command,
  CommandEmpty,
  CommandInput,
  CommandItem,
  CommandGroup,
  CommandList,
} from "@/components/ui/command"
import {branch} from "@/api/repo.ts";
import {toast} from "sonner";
import {updateBranch} from "@/api/task.ts";


export type SwitchBranchType = {
  Show: (id:number,branch:string,repositoryId:number) => void;
}

type SwitchBranchProps = {
  reload:(resetPageIndex:boolean)=>void
}

const SwitchBranch =  forwardRef<SwitchBranchType|undefined,SwitchBranchProps>(({reload}, ref)=>{
  const [open,setOpen] = useState(false)
  const [name,setName] = useState("")
  const [taskId,setTaskId] = useState(0)
  const [branches,setBranches] = useState<string[]>([])
  const selectBranch = (b:string)=>{
    updateBranch({id:taskId,branch:b}).then(()=>{
      toast.success("Switch branch success.")
      setOpen(false)
      reload(false)
    }).catch(e=>{
      toast.error(e)
    })
  }

  const Show = (id:number,name:string,repositoryId:number) => {
    setName(name)
    setTaskId(id)
    branch(repositoryId).then(r=>{
      setBranches(r.data.data)
      setOpen(true)
    }).catch(e=>{
      toast.error(e)
    })
  }
  useImperativeHandle(ref, ()=>({Show}))
  return (
      <Dialog open={open} onOpenChange={v=>{setOpen(v)}}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>SwitchBranch ( taskId:{taskId} )</DialogTitle>
            <DialogDescription>The current branch of the task is {name}.</DialogDescription>
          </DialogHeader>
          <Command className="rounded">
            <CommandInput placeholder="Search branch ..." />
            <CommandList>
              <CommandEmpty>No results found.</CommandEmpty>
              <CommandGroup>
                {
                  branches.map(b=>{
                    const isCurrent = name===b
                    return (
                        <CommandItem
                            key={b}
                            onSelect={()=>{selectBranch(b)}}
                            disabled={isCurrent}
                            className={isCurrent?"text-muted-foreground":""}>
                          <Share2 className="mr-2 h-4 w-4" />
                          <span>{b}</span>
                        </CommandItem>
                    )
                  })
                }
              </CommandGroup>

            </CommandList>
          </Command>
        </DialogContent>
      </Dialog>
  )
})


export { SwitchBranch }
