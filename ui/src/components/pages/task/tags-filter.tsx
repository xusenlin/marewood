import { CheckIcon, BookmarkPlus } from "lucide-react"

import { cn } from "@/lib/utils"
import { Badge } from "@/components/ui/badge.tsx"
import { Button } from "@/components/ui/button.tsx"
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "@/components/ui/command"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import { Separator } from "@/components/ui/separator"
import {useEffect, useState} from "react";

export type OptionType = {
  label:string
  value:string
  count:number
}
interface TagsFilterProps {
  values: string
  title?: string
  options: OptionType[]
  onChange: (values: string) => void
}

export function TagsFilter({
  values,
  title,
  options,
  onChange
}: TagsFilterProps) {

  const [selectedValues,setSelectedValues] = useState<string[]>([])

  useEffect(() => {
    setSelectedValues(values ? values.split(","):[])
  }, [values])

  const onSelect = (value: string) => {
    const values = [...selectedValues]
    if (values.includes(value)) {
      values.splice(values.indexOf(value), 1)
    } else {
      values.push(value)
    }
    onChange(values.join(","))
  }

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button variant="outline" >
          <BookmarkPlus size={18} className="mr-1"/>
          {title}
          {selectedValues.length > 0 && (
            <>
              <Separator orientation="vertical" className="mx-2 h-4" />
              <div className="hidden space-x-1 lg:flex">
                {selectedValues.length > 2 ? (
                  <Badge
                    variant="secondary"
                    className="rounded-sm px-1"
                  >
                    <div style={{whiteSpace: "nowrap"}}>{selectedValues.length} selected</div>
                  </Badge>
                ) : (
                    selectedValues.map((option) => (
                      <Badge
                        variant="secondary"
                        key={option}
                        className="rounded-sm px-1 whitespace-nowrap"
                      >
                        {option}
                      </Badge>
                    ))
                )}
              </div>
            </>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0 mr-2" align="start">
        <Command>
          <CommandInput placeholder={title} />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup>
              {options.map((option) => {
                const isSelected = selectedValues.indexOf(option.value)!==-1
                return (
                  <CommandItem
                    key={option.value}
                    onSelect={() => {onSelect(option.value)}}
                  >
                    <div
                      className={cn(
                        "mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary",
                        isSelected
                          ? "bg-primary text-primary-foreground"
                          : "opacity-50 [&_svg]:invisible"
                      )}
                    >
                      <CheckIcon className={cn("h-4 w-4")} />
                    </div>
                    <span>{option.label}</span>
                    <span className="ml-auto flex h-4 w-4 items-center justify-center font-mono text-xs">
                        {option.count||""}
                      </span>
                  </CommandItem>
                )
              })}
            </CommandGroup>
            {selectedValues.length > 0 && (
              <>
                <CommandSeparator />
                <CommandGroup>
                  <CommandItem
                    onSelect={()=>{onChange("")}}
                    className="justify-center text-center"
                  >
                    Clear filters
                  </CommandItem>
                </CommandGroup>
              </>
            )}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
