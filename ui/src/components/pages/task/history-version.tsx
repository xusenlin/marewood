import {Sheet, SheetContent, SheetHeader, SheetTitle} from "@/components/ui/sheet.tsx";
import {forwardRef, useImperativeHandle, useState} from "react";
import {getHistoryByTaskId, History, restoreHistory} from "@/api/history.ts";
import {toast} from "sonner";
import {Button} from "@/components/ui/button.tsx";
import {GitCommit, User, Calendar} from "lucide-react";

export type HistoryVersionType = {
    Show: (id: number) => void
}

const restore = (id: number) => {
    restoreHistory(id).then(() => {
        toast.success("Restore success")
    }).catch(err => {
        toast.error(err)
    })
}

export const HistoryVersion = forwardRef<HistoryVersionType>((_, ref) => {

    const [isOpened, setIsOpened] = useState(false)
    const [history, setHistory] = useState<History[]>([])

    useImperativeHandle(ref, () => ({
        Show: (id) => {
            getHistoryByTaskId(id).then(res => {
                setHistory(res.data.data)
            }).catch(err => {
                toast.error(err)
            })
            setIsOpened(true)
        }
    }));

    return (
        <Sheet open={isOpened} onOpenChange={setIsOpened}>
            <SheetContent className="sm:max-w-[700px]">
                <SheetHeader>
                    <SheetTitle>History Version</SheetTitle>
                </SheetHeader>
                <div className="py-4 h-full overflow-y-auto">
                    {history.map(h => (
                        <div key={h.id} className="flex items-center justify-between py-3 border-b last:border-0">
                            <div>
                                <p className="text-sm font-medium leading-none">{h.name}</p>
                                <div className="flex items-center text-xs text-muted-foreground mt-2 space-x-3">
                                    <div className="flex items-center">
                                        <GitCommit className="mr-1 h-3 w-3" />
                                        <span>{h.gitHash.substring(0, 7)}</span>
                                    </div>
                                    <div className="flex items-center">
                                        <User className="mr-1 h-3 w-3" />
                                        <span>{h.userName}</span>
                                    </div>
                                    <div className="flex items-center">
                                        <Calendar className="mr-1 h-3 w-3" />
                                        <span>{new Date(h.createdAt!).toLocaleString()}</span>
                                    </div>
                                </div>
                            </div>
                            <Button size="sm" onClick={() => restore(h.id!)}>Restore</Button>
                        </div>
                    ))}
                </div>
            </SheetContent>
        </Sheet>
    )
})