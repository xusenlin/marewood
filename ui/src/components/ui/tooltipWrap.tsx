import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip.tsx";


type TooltipWrapType = {
  tip:string|JSX.Element
  child: JSX.Element
}
export const TooltipWrap = (props:TooltipWrapType)=> <TooltipProvider>
  <Tooltip >
    <TooltipTrigger asChild>
        { props.child }
    </TooltipTrigger>
    <TooltipContent side="bottom">
      {props.tip}
    </TooltipContent>
  </Tooltip>
</TooltipProvider>
