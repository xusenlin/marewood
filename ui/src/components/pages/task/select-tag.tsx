import {
  Command,
  CommandGroup,
  CommandItem,
} from "@/components/ui/command.tsx"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover.tsx"
import {useEffect, useState} from "react";
import {PanelBottomClose} from "lucide-react";
import {toast} from "sonner";
import {allTags} from "@/api/task.ts";
export type TagType = {
  tag:string
  count:number
}
export type AntoInputProps = {
  onChange: (value: string) => void,
}



export const SelectTag = ({onChange}:AntoInputProps) => {
  const [open, setOpen] = useState(false)
  const [tags, setTags] = useState<TagType[]>([])

  useEffect(() => {

    allTags().then(res=>{
      setTags(res.data.data)
    }).catch(e=>{
      toast.error(e)
    })

  }, [])
  return (
      <div>
        <Popover open={open} onOpenChange={setOpen}>
          <PopoverTrigger>
            <div className="h-10 px-4 py-2 flex items-center justify-center rounded-md text-sm font-medium border border-input bg-background">
              <PanelBottomClose size={16} />
            </div>
          </PopoverTrigger>
          <PopoverContent className="w-[200px] p-0">
            <Command>
              <CommandGroup>
                {tags.map((t) => (
                    <CommandItem
                        key={t.tag}
                        value={t.tag}
                        onSelect={(v) => {
                          onChange(v)
                          setOpen(false)
                        }}
                    >
                      <div className="flex w-[200px] items-center justify-between">
                        <span>{t.tag}</span>
                        <span>{t.count}</span>
                      </div>
                    </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </PopoverContent>
        </Popover>
      </div>

  )
}
