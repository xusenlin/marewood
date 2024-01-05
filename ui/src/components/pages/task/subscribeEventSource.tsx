import {useEffect} from "react";
import {getEventUrl} from "@/api/task.ts";
import {toast} from "sonner";
import {useNotificationStore} from "@/store/notifications.ts";

const eventTypeEditOk = "editOk"
const eventTypeDestroyOk = "destroyOk"

const eventTaskTypeRunOk = "runOk"
const eventTaskTypeBuildOk = "buildOk"
const eventTaskTypeBuildFail = "buildFail"



export type TaskEvent = {
  taskId: number
  userId: number
  userName: string
  msg: string
}

type SubscribeEventSourceProps = {
  refresh: () => void
}

export const SubscribeEventSource = (props: SubscribeEventSourceProps) => {
  const {push} = useNotificationStore()

  let evtSource:EventSource

  const eventSourceInit = ()=>{
    evtSource = new EventSource(getEventUrl())

    evtSource.addEventListener(eventTypeEditOk, m => {
      const r: TaskEvent = JSON.parse(m.data)
      const d = `${r.userName} successfully edited the task (Task ID: ${r.taskId}).`
      toast.success(d)
      push({type:1,description:d,title:"edited task"})
      props.refresh()
    })

    evtSource.addEventListener(eventTypeDestroyOk, m => {
      const r: TaskEvent = JSON.parse(m.data)
      const d = `${r.userName} successfully deleted the task (Task ID: ${r.taskId}).`
      toast.success(d)
      push({type:1,description:d,title:"deleted task"})
      props.refresh()
    })
    //task
    evtSource.addEventListener(eventTaskTypeRunOk, m => {
      const r: TaskEvent = JSON.parse(m.data)
      const d = `${r.userName} successfully initiated the compilation task (Task ID: ${r.taskId}).`
      toast.success(d)
      push({type:2,description:d,title:"compilation task"})
      props.refresh()
    })
    evtSource.addEventListener(eventTaskTypeBuildOk, m => {
      const r: TaskEvent = JSON.parse(m.data)
      const d = `Task initiated by ${r.userName} has been successfully built (Task ID: ${r.taskId}).`
      toast.success(d)
      push({type:2,description:d,title:"successfully built"})
      props.refresh()
    })
    evtSource.addEventListener(eventTaskTypeBuildFail, m => {
      const r: TaskEvent = JSON.parse(m.data)
      const d = `The task initiated by the ${r.userName} has failed to build (Task ID: ${r.taskId}).`
      toast.error(d)
      push({type:2,description:d,title:"failed built"})
      props.refresh()
    })



    evtSource.onerror = (error) => {
      // @ts-ignore
      const {data} = error
      evtSource.close()
      toast("Network Error", {
        description:data || "Repository status subscription failed.",
        action: {
          label: 'Retry',
          onClick: () => eventSourceInit()
        },
      })
    };
  }


  useEffect(() => {
    eventSourceInit()
    return () => {
      evtSource.close()
    }
  }, [])

  return <></>
}
