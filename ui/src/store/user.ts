import {create} from "zustand";
import {persist} from "zustand/middleware"

export interface User  {
  id:number
  username:string
  role:number
  status:number
  token:string
  description?:string
}

export interface UserType {
  user: User | null
  setUserInfo: (user: User) => void
}


export const useUserStore = create<UserType>()(
    persist((set) => ({
      user: null,
      setUserInfo: (user: User) => {
        localStorage.setItem("MareWoodToken",user.token)
        set(() => ({user}))
      },
    }), {name: 'UserStore'})
)

export const outLogin = () => {
  localStorage.removeItem("MareWoodToken")
  useUserStore.setState({user: null})
}
