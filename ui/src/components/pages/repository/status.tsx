import {Loader2, CheckCircle2, Ban} from "lucide-react"
import {TooltipWrap} from "@/components/ui/tooltipWrap.tsx";

export const Status = (props:{status:number})=>{
  switch (props.status){
    case 0:
      return  <div className="flex items-center justify-center">
        <TooltipWrap tip="In Progress" child={<Loader2   size={18} className="animate-spin" />}/>
      </div>
    case 1:
      return  <div className="flex items-center justify-center">
        <TooltipWrap tip="Done" child={<CheckCircle2  size={18} color="#4caf50"/>}/>
      </div>
    case 2:
      return  <div className="flex items-center justify-center">
        <TooltipWrap tip="Error" child={<Ban  size={18} color="#f44336"/>}/>
      </div>

  }
}
