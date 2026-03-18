import { Sheet, SheetContent, SheetHeader, SheetTitle } from "@/components/ui/sheet.tsx";
import { forwardRef, useImperativeHandle, useState } from "react";
import { getHistoryByTaskId, History, restoreHistory } from "@/api/history.ts";
import { toast } from "sonner";
import { Button } from "@/components/ui/button.tsx";
import { GitCommit, User, Clock, RotateCcw } from "lucide-react";
import { cn } from "@/lib/utils";

export type HistoryVersionType = {
    Show: (id: number) => void;
};

const restore = (id: number) => {
    restoreHistory(id)
        .then(() => toast.success("Restored successfully"))
        .catch((err) => toast.error(err));
};

export const HistoryVersion = forwardRef<HistoryVersionType>((_, ref) => {
    const [isOpened, setIsOpened] = useState(false);
    const [history, setHistory] = useState<History[]>([]);

    useImperativeHandle(ref, () => ({
        Show: (id) => {
            getHistoryByTaskId(id)
                .then((res) => setHistory(res.data.data))
                .catch((err) => toast.error(err));
            setIsOpened(true);
        },
    }));

    return (
        <Sheet open={isOpened} onOpenChange={setIsOpened}>
            <SheetContent className="sm:max-w-[680px] flex flex-col gap-0 p-0">

                {/* ── Header ── */}
                <SheetHeader className="px-6 pt-6 pb-4 border-b border-border/60 pr-14">
                    <div className="flex items-center justify-between">
                        <SheetTitle className="text-base font-semibold tracking-tight">
                            History
                        </SheetTitle>
                        {history.length > 0 && (
                            <span className="text-xs font-medium tabular-nums px-2 py-0.5 rounded-full bg-muted text-muted-foreground">
                                {history.length} versions
                            </span>
                        )}
                    </div>
                </SheetHeader>

                {/* ── List ── */}
                <div className="flex-1 overflow-y-auto px-6 py-5">
                    {history.length === 0 ? (
                        <div className="flex flex-col items-center justify-center h-48 text-muted-foreground gap-2">
                            <GitCommit className="h-8 w-8 opacity-30" />
                            <p className="text-sm">No history versions yet</p>
                        </div>
                    ) : (
                        <ol className="relative">
                            {/* vertical timeline line */}
                            <div className="absolute left-[7px] top-2 bottom-2 w-px bg-border/70" />

                            {history.map((h, i) => (
                                <li
                                    key={h.id}
                                    className={cn(
                                        "group relative pl-7 pb-6 last:pb-0",
                                    )}
                                >
                                    {/* timeline dot */}
                                    <span
                                        className={cn(
                                            "absolute left-0 top-1.5 flex h-3.5 w-3.5 items-center justify-center rounded-full border-2 bg-background transition-colors duration-150",
                                            i === 0
                                                ? "border-foreground"
                                                : "border-border group-hover:border-foreground/50",
                                        )}
                                    >
                                        {i === 0 && (
                                            <span className="h-1.5 w-1.5 rounded-full bg-foreground" />
                                        )}
                                    </span>

                                    {/* card */}
                                    <div className="flex items-center justify-between gap-4 rounded-lg border border-border/60 bg-card px-4 py-3.5 transition-all duration-150 hover:border-border hover:shadow-sm">
                                        <div className="min-w-0 flex-1">

                                            {/* name + latest badge */}
                                            <div className="flex items-center gap-2 mb-2">
                                                <p className="text-sm font-medium leading-none truncate">
                                                    {h.name}
                                                </p>
                                                {i === 0 && (
                                                    <span className="shrink-0 text-[10px] font-medium px-1.5 py-0.5 rounded bg-foreground text-background leading-none">
                                                        latest
                                                    </span>
                                                )}
                                            </div>

                                            {/* metadata row */}
                                            <div className="flex flex-wrap items-center gap-x-3 gap-y-1 text-xs text-muted-foreground">
                                                {/* git hash pill */}
                                                <span className="inline-flex items-center gap-1 font-mono px-1.5 py-0.5 rounded bg-muted text-muted-foreground">
                                                    <GitCommit className="h-3 w-3 shrink-0" />
                                                    {h.gitHash.substring(0, 7)}
                                                </span>

                                                <span className="flex items-center gap-1">
                                                    <User className="h-3 w-3 shrink-0" />
                                                    {h.userName}
                                                </span>

                                                <span className="flex items-center gap-1">
                                                    <Clock className="h-3 w-3 shrink-0" />
                                                    {new Date(h.createdAt!).toLocaleString()}
                                                </span>
                                            </div>
                                        </div>

                                        {/* restore button */}
                                        <Button
                                            size="sm"
                                            variant="outline"
                                            className="shrink-0 h-7 px-2.5 text-xs gap-1.5 opacity-0 group-hover:opacity-100 transition-opacity duration-150"
                                            onClick={() => restore(h.id!)}
                                        >
                                            <RotateCcw className="h-3 w-3" />
                                            Restore
                                        </Button>
                                    </div>
                                </li>
                            ))}
                        </ol>
                    )}
                </div>
            </SheetContent>
        </Sheet>
    );
});