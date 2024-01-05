import {create} from "zustand";
import {persist} from "zustand/middleware"

export enum TabEnum {
  "Dashboard",
  "Repositories",
  "Task",
  "Users",
}


export interface AppType {
  activeTabIndex: TabEnum
  showAbout:boolean

  setTabByIndex: (index: TabEnum) => void
  setShowAbout:(show: boolean)=>void
}


export const useAppStore = create<AppType>()(
    persist((set) => ({
      activeTabIndex: TabEnum.Dashboard,
      showAbout:false,
      setTabByIndex: (index: TabEnum) => set(() => ({activeTabIndex: index})),
      setShowAbout: (show: boolean) => set(() => ({showAbout: show})),
    }), {name: 'AppStore'})
)
