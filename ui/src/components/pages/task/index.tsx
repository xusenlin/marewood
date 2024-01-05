import {ActionType, TaskTable} from "./table.tsx"
import { FloatButton  } from "@/components/ui/float-button.tsx"
import {useEffect, useRef, useState} from "react";
import {Eraser, Plus} from "lucide-react";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {OptionType, TagsFilter} from "./tags-filter.tsx"
import {allTags} from "@/api/task.ts";
import {toast} from "sonner";
import {SubscribeEventSource} from "./subscribeEventSource.tsx";

const initialPrams: {name: string, id: number|null, tags:string } = {
  name: "",
  id: null,
  // private:false,
  tags:"",
}

export const TaskIndex = () => {
  const [prams, setPrams] = useState(initialPrams)
  const tableRef = useRef<ActionType>()
  const [tags, setTags] = useState<OptionType[]>([])


  useEffect(()=>{
    allTags().then(res=>{
      const tags = res.data.data.map(item=>({label:item.tag,value:item.tag,count:item.count}))
      setTags(tags)
    }).catch(e=>{
      toast.error(e)
    })
  },[])

  return (
      <>
        <div className="flex items-center justify-between px-6">
          <div className="flex flex-col space-y-1.5">
            <div className="text-2xl font-semibold">Task</div>
            <div className="text-sm text-muted-foreground">🌼 Different branches and npm commands can be used to create multiple tasks in the same repository.</div>
          </div>
          <div className="flex items-center">

            <Input
                value={prams.id+''}
                onChange={e => {
                  setPrams({...prams, id: e.target.value ? parseInt(e.target.value):null})
                }}
                className="border-2 border-gray-200 w-[60px]  mr-2"
                placeholder="ID" type="number"/>
            <Input
                value={prams.name}
                onChange={e => {
                  setPrams({...prams, name: e.target.value})
                }}
                className="border-2 border-gray-200  mr-2"
                placeholder="Search task name" type="text"/>
            <div className="mr-2">
              <TagsFilter
                  onChange={v=>{setPrams({...prams,tags:v})}}
                  values={prams.tags}
                  title="Tags"
                  options={tags}
              />
            </div>
            <Button style={{flexShrink:0}} variant="outline" size="icon" onClick={()=>setPrams(initialPrams)}>
              <Eraser size={16} />
            </Button>
          </div>
        </div>

        <div className="p-6">
          <TaskTable prams={prams} ref={tableRef} />
        </div>

        <FloatButton onClick={()=>{ tableRef.current?.addRow()}}>
            <Plus size={32} />
        </FloatButton>
        <SubscribeEventSource refresh={()=>{tableRef.current?.reload()}}/>
      </>
  )
}
