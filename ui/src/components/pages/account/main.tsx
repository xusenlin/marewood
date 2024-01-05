import {useState} from "react";
import { Login } from "./login.tsx"
import { CreateAccount } from "./create.tsx"

export const Account = ()=>{
  const [isLogin,setIsLogin ] = useState(true)
  // setIsLogin(true)

  return (
      <div className="w-full h-[100vh] flex justify-center items-center">
        { isLogin ? <Login create={()=>{setIsLogin(false)}}/> : <CreateAccount login={()=>{setIsLogin(true)}}/> }
      </div>
  )
}
