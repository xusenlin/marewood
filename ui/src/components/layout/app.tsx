import  Main from "./main.tsx"
import {Account} from '@/components/pages/account/main.tsx'
import {useUserStore} from "@/store/user.ts";


function App() {
  const {user} = useUserStore()
  return user ? <Main/> : <Account/>
}

export default App
