import {useState, useEffect} from "react"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card.tsx";
import {Hexagon, FolderKanban, FolderHeart, Blocks, Tag, Settings, Server, GitBranch, Folder, Database, Github, Laptop} from "lucide-react"
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
                <div className="flex flex-wrap items-center gap-x-8 gap-y-4 mt-8">
                  <div className="flex items-center">
                    <span className="font-semibold text-base mr-2">
                      <Tag className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                      Version:
                    </span>
                    <Badge variant="secondary" className="text-base px-3 py-1">{sysInfo.version}</Badge>
                  </div>

                  <div className="flex items-center">
                    <span className="font-semibold text-base mr-2"><Settings className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>GinMode:</span>
                    <Badge variant="secondary" className="text-base px-3 py-1">{sysInfo.ginMode}</Badge>
                  </div>
                  <div className="flex items-center">
                    <span className="font-semibold text-base mr-2"><Server className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>HttpPort:</span>
                    <Badge variant="secondary" className="text-base px-3 py-1">{sysInfo.httpPort}</Badge>
                  </div>
                </div>
                <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4 mt-8">
                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Laptop className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        System Info
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.systemInfo}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <GitBranch className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        Git version
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.dependTools.git}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Folder className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        ClientDir
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.clientDir}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Folder className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        ResourcesDir
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.resourcesDir}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Database className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        DBDns
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.dbDns}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Folder className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        RepositoryDir
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.repositoryDir}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Folder className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        WebRootDir
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">{sysInfo.webRootDir}</code>
                    </CardContent>
                  </Card>

                  <Card>
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                      <CardTitle className="text-sm font-medium flex items-center">
                        <Github className="h-4 w-4 mr-1 text-muted-foreground inline-block align-text-bottom"/>
                        Github
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <code className="break-all text-muted-foreground text-sm">https://github.com/xusenlin/marewood</code>
                    </CardContent>
                  </Card>
                </div>
              </CardContent>
            </Card>
            }
          </div>
        </div>
      </>
  )
}
