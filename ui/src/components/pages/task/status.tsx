import { Loader2,Smile,Ban,CheckCircle2 } from "lucide-react"
import {TooltipWrap} from "@/components/ui/tooltipWrap.tsx";

export const Status = (props:{status:number})=>{
  switch (props.status){
    case 0:
      return <div className="flex items-center justify-center">
        <TooltipWrap tip="Leisured" child={<Smile size={18} color="#9E9E9E"/>}/>
      </div>
    case 1:
      return <div className="flex items-center justify-center">
        <TooltipWrap tip="Busy" child={<Loader2  size={18} className="animate-spin" />}/>
      </div>
    case 2:
      return <div className="flex items-center justify-center">
        <TooltipWrap tip="Failed" child={<Ban   size={18} color="#f44336" />}/>
      </div>
    case 3:
      return <div className="flex items-center justify-center">
        <TooltipWrap tip="Success" child={<CheckCircle2   size={18} color="#4caf50" />}/>
      </div>
  }
}
