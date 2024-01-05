import { Loader2,Smile } from "lucide-react"
import {TooltipWrap} from "@/components/ui/tooltipWrap.tsx";

export const TaskStatus = (props:{status:number})=>{
  switch (props.status){
    case 0:
      return <div className="flex items-center justify-center">
        <TooltipWrap tip="Leisured" child={<Smile size={18} color="#4caf50"/>}/>
      </div>
    case 1:
      return <div className="flex items-center justify-center">
        <TooltipWrap tip="Busy" child={<Loader2   size={18} className="animate-spin" />}/>
      </div>
  }
}
