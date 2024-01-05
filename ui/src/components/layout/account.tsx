import {Button} from "@/components/ui/button.tsx";
import {Role} from "@/components/ui/role.tsx";


import {
  Github,
  LogOut,
  UserCircle2
} from "lucide-react"

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

import {useUserStore,outLogin} from "@/store/user.ts";


export const Account = () => {
  const {user } = useUserStore()

  return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
              variant="ghost"
              className="hover:bg-primary/0 hover:text-primary text-muted-foreground focus-visible:ring-0 focus-visible:ring-offset-0"
          >
            <UserCircle2 className="mr-2 h-4 w-4" />
            {user?.username}
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-50 mr-2">
          <DropdownMenuLabel>My Account</DropdownMenuLabel>
          <div className="text-muted-foreground p-2  text-sm outline-none flex flex-col  justify-center items-center">
            <div className="mb-2">{user?.username} (ID: {user?.id})</div>
            <Role id={user!.role}/>
          </div>
          <DropdownMenuSeparator />
          <DropdownMenuItem onClick={()=>window.open("https://github.com/xusenlin/marewood", "_blank")}>
            <Github className="mr-2 h-4 w-4" />
            <span>GitHub</span>
          </DropdownMenuItem>
          <DropdownMenuSeparator />
          <DropdownMenuItem onClick={outLogin}>
            <LogOut className="mr-2 h-4 w-4" />
            <span>Log out</span>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>

  )
}
