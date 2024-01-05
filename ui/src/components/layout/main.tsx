import {Nav} from "@/components/layout/nav.tsx"
import {Account} from "@/components/layout/account.tsx"
import {Dashboard} from "@/components/pages/dashboard";
import {TaskIndex} from "@/components/pages/task";
import {User} from "@/components/pages/user";
import {Notifications} from "./notifications.tsx"
import {Repository} from "@/components/pages/repository";

import {useAppStore, TabEnum} from "@/store/app.ts";


function Main() {
  const {activeTabIndex} = useAppStore()
  return (
      <div className="main">
        <div className="border-b">
          <div className="flex h-16 items-center px-4">
            <Nav/>
            <div className="ml-auto flex items-center space-x-1 mx-6">
              <Notifications/>
              <Account/>
            </div>
          </div>
        </div>
        <div className="content text-foregroun">
          {activeTabIndex === TabEnum.Dashboard && <Dashboard/>}
          {activeTabIndex === TabEnum.Repositories && <Repository/>}
          {activeTabIndex === TabEnum.Task && <TaskIndex/>}
          {activeTabIndex === TabEnum.Users && <User/>}
        </div>
      </div>
  )
}

export default Main
