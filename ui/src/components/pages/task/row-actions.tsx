import {Button} from "@/components/ui/button.tsx";
import {ArrowDownToLine, MoreHorizontal,PencilRuler, PlayCircle, Link, Share2, Trash2} from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {destroy, getTaskArchiverUrl, run, Task} from "@/api/task.ts";
import {useRef} from "react";
import {SwitchBranch, SwitchBranchType} from "@/components/pages/task/switch-branch.tsx";
import {toast} from "sonner";

type RowActionsProps = {
  row:Task
  reload:(resetPageIndex:boolean)=>void
  edit:()=>void
}

export const RowActions = ({row,reload,edit}:RowActionsProps) => {

  const switchBranchRef = useRef<SwitchBranchType>();

  const openUrl = () => {
    window.open(`${window.location.protocol}//${window.location.host}/webs/${row.alias}`, "_blank");
  }
  const del = (id:number) =>{
    destroy(id).then(()=>{
      toast.success("delete success")
      reload(true)
    }).catch(e=>{
      toast.error(e)
    })
  }
  const runTask = (id:number)=>{
    run(id).then(()=>{
      toast.success("Compiling in the background...")
      reload(true)
    }).catch(e=>{
      toast.error(e)
    })
  }
  return (
      <div className="flex items-center w-full justify-center">
        <Button
            variant="ghost"
            size="icon"
            onClick={()=>{runTask(row.id!)}}
            className="hover:bg-primary/0"
        >
          <PlayCircle size={20}/>
        </Button>
        <Button
            variant="ghost"
            onClick={openUrl}
            size="icon"
            className="hover:bg-primary/0"
        >
          <Link size={20}/>
        </Button>
        <Button
            variant="ghost"
            size="icon"
            onClick={()=>switchBranchRef.current?.Show(row.id!,row.branch,row.repositoryId)}
            className="hover:bg-primary/0"
        >
          <Share2 size={20}/>
        </Button>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button
                onClick={()=>{}}
                variant="ghost"
                size="icon"
                className="hover:bg-primary/0"
            >
              <MoreHorizontal size={20}/>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-50  mr-2">
            <DropdownMenuLabel>More Actions</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              <DropdownMenuItem onClick={edit}>
                <PencilRuler className="mr-2 h-4 w-4" />
                <span>Edit Task</span>
              </DropdownMenuItem>
              <DropdownMenuItem onClick={()=>{window.open(getTaskArchiverUrl(row.id!,1), "_blank");}}>
                <ArrowDownToLine className="mr-2 h-4 w-4" />
                <span>Download Tar</span>
              </DropdownMenuItem>
              <DropdownMenuItem onClick={()=>{window.open(getTaskArchiverUrl(row.id!,2), "_blank");}}>
                <ArrowDownToLine className="mr-2 h-4 w-4" />
                <span>Download Zip</span>
              </DropdownMenuItem>
              <DropdownMenuItem onClick={()=>{del(row.id!)}}>
                <Trash2 className="mr-2 h-4 w-4" />
                <span>Delete Task</span>
              </DropdownMenuItem>
            </DropdownMenuGroup>
          </DropdownMenuContent>
        </DropdownMenu>
        <SwitchBranch ref={switchBranchRef} reload={reload} />
      </div>
  )
}
