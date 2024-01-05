import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet"
import {TerminalSquare} from "lucide-react";
import { JSX } from "react"

type TerminalOutTProps = {
  title:string|JSX.Element
  out:string|JSX.Element
  description:string|JSX.Element
}


export const TerminalOut = (props:TerminalOutTProps)=>{

  return <Sheet>
    <SheetTrigger>
      <TerminalSquare size={18}/>
    </SheetTrigger>
    <SheetContent className="terminal-out sm:max-w-[40%]">
      <SheetHeader>
        <SheetTitle>{props.title}</SheetTitle>
        <div className="text-sm text-muted-foreground">{props.description}</div>
      </SheetHeader>
      <div className="out-content">
        {props.out}
      </div>
    </SheetContent>
  </Sheet>
}
