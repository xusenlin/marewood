import * as React from "react";
import {cn} from "@/lib/utils.ts";


export const FloatButton = React.forwardRef<
    HTMLDivElement,
    React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
    <div
        ref={ref}
        className={cn("fixed hover:shadow-lg float-button",  className)}
        {...props}
    />
))
