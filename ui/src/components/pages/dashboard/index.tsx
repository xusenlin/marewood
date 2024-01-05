import {useState, useEffect} from "react"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card.tsx";
import {Hexagon, FolderKanban, FolderHeart,Blocks} from "lucide-react"
import {Sys, getSysInfo} from "@/api/sys.ts";
import {toast} from "sonner";
import {Badge} from "@/components/ui/badge.tsx";


export const Dashboard = () => {
  const [sysInfo, setSysInfo] = useState<Sys | null>();

  useEffect(() => {
    getSysInfo().then(res => {
      setSysInfo(res.data.data);
    }).catch(e => {
      toast.error(e)
    })
  }, [])

  return (
      <>
        <div className="flex flex-col space-y-1.5 px-6">
          <div className="text-2xl font-semibold">Dashboard</div>
          <div className="text-sm text-muted-foreground flex items-center">
            💐 Please restart the service if you want to update the server environment.
          </div>
        </div>
        <div className="p-6">
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            <Card className="flex flex-col justify-between">
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  Node version
                </CardTitle>
                <Hexagon className="h-4 w-4 text-muted-foreground"/>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">{sysInfo && sysInfo.dependTools.node}</div>
              </CardContent>
            </Card>
            <Card className="flex flex-col justify-between">
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  Pnpm version
                </CardTitle>
                <Blocks className="h-4 w-4 text-muted-foreground"/>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">{sysInfo && sysInfo.dependTools.pnpm}</div>
              </CardContent>
            </Card>
            <Card className="flex flex-col justify-between">
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">Yarn version</CardTitle>
                <FolderHeart className="h-4 w-4 text-muted-foreground"/>

              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">{sysInfo && sysInfo.dependTools.yarn}</div>
              </CardContent>
            </Card>
            <Card className="flex flex-col justify-between">
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  Npm version
                </CardTitle>
                <FolderKanban className="h-4 w-4 text-muted-foreground"/>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">{sysInfo && sysInfo.dependTools.npm}</div>
              </CardContent>
            </Card>
          </div>
          <div className="mt-4">

            {sysInfo && <Card className="bg-white dark:bg-gray-900 shadow rounded-lg p-6">
              <CardHeader>
                <CardTitle className="text-lg font-semibold">MareWood</CardTitle>
                <CardDescription>
                  🌻 Marewood is an open source automatic packaging deployment tool for front-end
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="flex justify-between items-center mt-8">
                  <div>
                    <span className="font-medium mr-2">
                      Version:
                    </span>
                    <Badge className="text-sm">{sysInfo.version}</Badge>
                  </div>

                  <div>
                    <span className="font-medium mr-2">GinMode:</span>
                    <Badge className="text-sm">{sysInfo.ginMode}</Badge>
                  </div>
                  <div>
                    <span className="font-medium mr-2">HttpPort:</span>
                    <Badge className="text-sm">{sysInfo.httpPort}</Badge>
                  </div>
                </div>
                <div className="grid grid-cols-2 gap-4 mt-8">
                  <div>
                    <span className="font-medium">Git version:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>{sysInfo.dependTools.git}</code>
                  </div>
                  <div>
                    <span className="font-medium">ClientDir:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>{sysInfo.clientDir}</code>
                  </div>
                  <div>
                    <span className="font-medium">ResourcesDir:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>{sysInfo.resourcesDir}</code>
                  </div>
                  <div>
                    <span className="font-medium">DBDns:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>{sysInfo.dbDns}</code>
                  </div>
                  <div>
                    <span className="font-medium">RepositoryDir:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>{sysInfo.repositoryDir}</code>
                  </div>
                  <div>
                    <span className="font-medium">WebRootDir:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>{sysInfo.webRootDir}</code>
                  </div>
                  <div>
                    <span className="font-medium">Github:</span>
                  </div>
                  <div className="text-muted-foreground">
                    <code>https://github.com/xusenlin/marewood</code>
                  </div>
                </div>
              </CardContent>
            </Card>
            }
          </div>
        </div>
      </>
  )
}
