import {Logo} from "./logo.tsx"
import {Button} from "@/components/ui/button"
import {useAppStore, TabEnum} from "@/store/app.ts";

export const Nav = () => {
    const {setTabByIndex, activeTabIndex} = useAppStore()
    const classList = (active: boolean) => {
        const classes = "text-sm hover:bg-primary/0 font-medium transition-colors "
        return active ? classes + "hover:text-primary" : classes + "text-muted-foreground hover:text-primary"
    }

    return (
        <nav className="mx-6 flex items-center space-x-4 lg:space-x-6">
            <Button
                onClick={()=>window.open("https://github.com/xusenlin/marewood", "_blank")}
                variant="ghost"
                size="icon"
                className="hover:bg-primary/0"
            >
                <Logo/>
            </Button>

            <Button
                onClick={() => setTabByIndex(TabEnum.Dashboard)}
                variant="ghost"
                className={classList(activeTabIndex === TabEnum.Dashboard)}
            >
                Dashboard
            </Button>

            <Button
                onClick={() => setTabByIndex(TabEnum.Repositories)}
                variant="ghost"
                className={classList(activeTabIndex === TabEnum.Repositories)}
            >
                Repositories
            </Button>

            <Button
                onClick={() => setTabByIndex(TabEnum.Task)}
                variant="ghost"
                className={classList(activeTabIndex === TabEnum.Task)}
            >
                Task
            </Button>
            <Button
                onClick={() => setTabByIndex(TabEnum.Users)}
                variant="ghost"
                className={classList(activeTabIndex === TabEnum.Users)}
            >
                Users
            </Button>
        </nav>
    )
}
