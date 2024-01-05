import { BellIcon, BellRing,PackageCheck } from "lucide-react"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import {Button} from "@/components/ui/button.tsx";
import {useNotificationStore} from "@/store/notifications.ts";

type ItemType = {
  icon: JSX.Element,
  title: string,
  description: string,
}

const NotificationItem = ({icon, title, description}:ItemType) => {
  return (
      <div
          className="flex items-start space-x-4 rounded-md p-2 transition-all hover:bg-accent hover:text-accent-foreground">
        {icon}
        <div className="space-y-1">
          <p className="text-sm font-medium leading-none">{title}</p>
          <p className="text-sm text-muted-foreground">{description}</p>
        </div>
      </div>
  )
}



export const Notifications = ()=>{
  const {msg} = useNotificationStore()

  return (
      <Popover>
        <PopoverTrigger asChild>
          <Button
              variant="ghost"
              size="icon"
              className="relative hover:bg-primary/0 hover:text-primary text-muted-foreground focus-visible:ring-0 focus-visible:ring-offset-0"
          >
            <BellRing className="h-4 w-4" />
            { msg.length > 0 && <div className="right-1 top-1 w-1 h-1 rounded absolute animate-ping bg-red-600"></div> }
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[340px] mr-2">
          <div className="flex justify-center items-center mb-2 text-muted-foreground">Notifications({ msg.length })</div>
          <div className="-mx-1 my-1 h-px bg-muted"></div>
          <div className="max-h-[300px] overflow-y-auto overflow-x-hidden">
            { msg.map((m,idx)=><NotificationItem key={idx} title={m.title} description={m.description} icon={m.type ===2?<PackageCheck className="mt-px h-5 w-5" />:<BellIcon className="mt-px h-5 w-5" />}/>) }
            { msg.length === 0 && <div className="flex justify-center items-center text-muted-foreground h-20">No notifications</div>}
          </div>
        </PopoverContent>
      </Popover>
  )
}
