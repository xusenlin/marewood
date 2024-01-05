import {Button} from "@/components/ui/button.tsx"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {Input} from "@/components/ui/input"
import {Label} from "@/components/ui/label"
import {Logo} from "@/components/layout/logo.tsx"
import {useUserStore} from "@/store/user.ts";
import { toast } from "sonner"
import {useState} from "react";
import { login } from "@/api/user.ts"

export function Login({create}:{create:()=>void}) {
  const {setUserInfo} = useUserStore()
  const [username, setUsername ] = useState("")
  const [password, setPassword ] = useState("")

  const submit = ()=>{
    login({username,password}).then(res=>{
      toast.success(res.data.msg)
      setUserInfo(res.data.data)
    }).catch(msg=>{
      toast.error(msg)
    })

  }

  return (
      <Card className="w-[430px]">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl">Login system</CardTitle>
          <CardDescription>
            Enter your username and password to log in to the system.
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-4">
          <div className="flex justify-center items-center">
            <Logo/>
          </div>
          <div className="relative">
            <div className="absolute inset-0 flex items-center">
              <span className="w-full border-t"/>
            </div>
            <div className="relative flex justify-center text-xs uppercase">
            <span className="bg-background px-2 text-muted-foreground">
              MareWood
            </span>
            </div>
          </div>
          <div className="grid gap-2">
            <Label htmlFor="name">Name</Label>
            <Input onChange={r=>setUsername(r.target.value)} id="name" type="text" placeholder="Admin"/>
          </div>
          <div className="grid gap-2">
            <Label htmlFor="password">Password</Label>
            <Input onChange={r=>setPassword(r.target.value)} id="password" type="password"/>
          </div>
          <Button variant="link" onClick={create}>Create an account</Button>
        </CardContent>
        <CardFooter>
          <Button className="w-full" onClick={submit}>Login account</Button>
        </CardFooter>
      </Card>
  )
}
