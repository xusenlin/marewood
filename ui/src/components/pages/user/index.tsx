import { Input } from "@/components/ui/input"
import {UserTable} from "./table.tsx"
import {useState} from "react";
import {Button} from "@/components/ui/button.tsx";
import { Eraser} from "lucide-react"

export const User = () => {
  const [username, setName] = useState("")

  return (
      <>
        <div className="flex items-center justify-between px-6">
          <div className="flex flex-col space-y-1.5">
            <div className="text-2xl font-semibold">Users</div>
            <div className="text-sm text-muted-foreground"> 🌼  Role list: SuperAdministrator、Administer、Developer、Reporter.</div>
          </div>

          <div className="flex max-w-sm items-center">
            <Input
                value={username}
                onChange={e=>{setName(e.target.value)}}
                className="border-2 border-gray-200 mr-2"
                placeholder="Search username" type="text" />
            <Button variant="outline" size="icon" onClick={()=>setName("")}>
              <Eraser size={16} />
            </Button>
          </div>

        </div>
        <div className="p-6">
          <UserTable prams={{username}}/>
        </div>
      </>
  )
}
