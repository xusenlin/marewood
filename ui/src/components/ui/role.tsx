import { Badge } from "@/components/ui/badge"



const description:{[key:number]:string} =  {
  1:"SuperAdministrator",
  3:"Administer",
  6:"Developer",
  10:"Reporter"
}

export const getRoleArr = ():{k:number,v:string}[] => {
  const res:{k:number,v:string}[] = []
  for(const i in description){
    res.push({
      k:Number(i),
      v:description[i]
    })
  }
  return res
}

export const Role = (props:{id:number})=> (<Badge className="hover:bg-gray-100/50 dark:hover:bg-gray-800/50" variant="secondary">{description[props.id]}</Badge>)
