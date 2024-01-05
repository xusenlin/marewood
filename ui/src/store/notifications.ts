import {create} from 'zustand'


export type Msg = {
  title: string
  description: string
  type: number
}


export type Notification = {
  msg: Msg[]
  push: (m:Msg) => void
}

export const useNotificationStore = create<Notification>()((set) => ({
  msg: [],
  push: (m:Msg) => set((state) => ({msg: [m,...state.msg]})),
}))
