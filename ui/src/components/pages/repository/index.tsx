import {useRef, useState} from "react";
import {Input} from "@/components/ui/input.tsx";
import {ActionType, RepoTable} from "./table.tsx"
import { SubscribeEventSource } from "./subscribeEventSource.tsx"
import {Eraser} from "lucide-react";
import {Button} from "@/components/ui/button.tsx";

const initialPrams:{name:string, id:number|null} = {
  name: "",
  id: null,
}

export const Repository = () => {
  const [prams, setPrams] = useState(initialPrams)
  const tableRef = useRef<ActionType>()

  return (
      <>
        <div className="flex items-center justify-between px-6">
          <div className="flex flex-col space-y-1.5">
            <div className="text-2xl font-semibold">Repositories</div>
            <div className="text-sm text-muted-foreground">🌴 Any modifications to the repository will result in the source code being re-cloned.</div>
          </div>
          <div className="flex max-w-sm items-center">
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
                placeholder="Search repository name" type="text"/>
            <Button variant="outline" size="icon" onClick={()=>setPrams(initialPrams)}>
              <Eraser size={16} />
            </Button>
          </div>
        </div>
        <div className="p-6">
          <RepoTable prams={prams} ref={tableRef} />
        </div>
        <SubscribeEventSource refresh={()=>{tableRef.current?.reload()}}/>
      </>
  )
}
