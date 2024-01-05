
import {useState,forwardRef,useImperativeHandle} from "react";
import { destroy } from "@/api/user.ts"
import {
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialog
} from "@/components/ui/alert-dialog.tsx";
import {toast} from "sonner";



export type AlertDialogType = {
  Confirm: (id:number) => void;
}
interface DeleteDialogProps {
  onDelete: () => void;
}

const DeleteDialog =  forwardRef<AlertDialogType|undefined,DeleteDialogProps>(({onDelete},ref)=>{
  const [open,setOpen] = useState(false)
  const [userId,setUserId] = useState(0)

  const Confirm = (id:number) => {
    setOpen(true)
    setUserId(id)
  }
  const delOk = ()=>{
    destroy({id:userId}).then(r=>{
      onDelete();
      setOpen(false)
      toast.success(r.data.msg)
    }).catch(e => {
      toast.error(e)
    })
  }
  useImperativeHandle(ref, ()=>({Confirm}))
  return (
      <AlertDialog open={open}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirm User Deletion</AlertDialogTitle>
            <AlertDialogDescription>Are you sure you want to delete this user? This action cannot be undone.</AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel onClick={()=>{setOpen(false)}}>Cancel</AlertDialogCancel>
            <AlertDialogAction onClick={delOk}>Continue</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
  )
})


export { DeleteDialog }
