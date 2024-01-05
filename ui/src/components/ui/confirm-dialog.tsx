
import {
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
  AlertDialog
} from "@/components/ui/alert-dialog.tsx";

import React from "react";



export type ConfirmDialogProps =  {
  child: React.ReactNode
  title?: string
  description?: string
  onConfirm: ()=>void
}

export const ConfirmDialog =  (props:ConfirmDialogProps)=>{

  return (
      <AlertDialog>
        <AlertDialogTrigger asChild>
          {props.child}
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>{props.title||"Confirm operation"}</AlertDialogTitle>
            <AlertDialogDescription>{props.title||"Are you sure you want to perform this operation?"}</AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <AlertDialogAction onClick={props.onConfirm}>Continue</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
  )
}



