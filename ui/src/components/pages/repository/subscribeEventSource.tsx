import {useEffect} from "react";
import {getEventUrl} from "@/api/repo.ts";
import {toast} from "sonner";
import {useNotificationStore} from "@/store/notifications.ts";


const eventTypeCloneOk = "cloneOk"
const eventTypeCloneErr = "cloneErr"
const eventTypeEditOk = "editOk"
const eventTypeDestroyOk = "destroyOk"

const eventTaskTypeRunOk = "runOk"
const eventTaskTypeBuildOk = "buildOk"

export type RepositoryEvent = {
  repositoryId: number
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

    evtSource.addEventListener(eventTypeCloneOk, m => {
      const r: RepositoryEvent = JSON.parse(m.data)
      const d = `The repository created by the ${r.userName} has been successfully cloned (Repository ID: ${r.repositoryId}).`
      toast.success(d)
      push({type:1,description:d,title:"successfully cloned"})
      props.refresh()
    })
    evtSource.addEventListener(eventTypeCloneErr, m => {
      const r: RepositoryEvent = JSON.parse(m.data)
      const d = `The repository created by ${r.userName} failed to clone (Repository ID: ${r.repositoryId}).`
      toast.error(d)
      push({type:1,description:d,title:"successfully failed"})
      props.refresh()
    })
    evtSource.addEventListener(eventTypeEditOk, m => {
      const r: RepositoryEvent = JSON.parse(m.data)
      const d = `${r.userName} successfully edited the repository (Repository ID: ${r.repositoryId}).`
      toast.success(d)
      push({type:1,description:d,title:"edited repository"})
      props.refresh()
    })

    evtSource.addEventListener(eventTypeDestroyOk, m => {
      const r: RepositoryEvent = JSON.parse(m.data)
      const d = `${r.userName} successfully deleted the repository (Repository ID: ${r.repositoryId}).`
      toast.success(d)
      push({type:1,description:d,title:"deleted repository"})
      props.refresh()
    })
    //task
    evtSource.addEventListener(eventTaskTypeRunOk, m => {
      const r: RepositoryEvent = JSON.parse(m.data)
      const d = `${r.userName} successfully initiated the compilation task (Repository ID: ${r.repositoryId}).`
      toast.success(d)
      push({type:2,description:d,title:"initiated compilation task"})
      props.refresh()
    })
    evtSource.addEventListener(eventTaskTypeBuildOk, m => {
      const r: RepositoryEvent = JSON.parse(m.data)
      const d = `Task initiated by ${r.userName} has been successfully built (Repository ID: ${r.repositoryId}).`
      toast.success(d)
      push({type:2,description:d,title:"successfully built"})
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
