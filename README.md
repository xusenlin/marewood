# MareWood
>MareWood 是一个轻量级的前端部署工具，使用了 GOLANG、GIN、GORM、JWT、SQLITE、REACE、MATERIAL-UI 开发，
不同于 Jenkins 的大而全，它很简单且只针对前端，你可以很灵活的配置各种部署环境。
如果你愿意，线上发布也可以是点击一下按钮这么简单的事情,当然也可以配置 WEBHOOK，提交 GIT 代码既自动发布。


# MareWood 在做什么
MareWood 可以帮你克隆前端仓库并使用你选择的工具安装依赖，并能在同一个仓库下通过不同的任务不同的打包命令打包出多个环境下的前端资源并放在 resources\webs 下，通过任务 ID 一一对应，并提供快捷访问。



# 使用

1. 服务器请先安装 git、node、npm，建议同时安装 cnpm 和 yarn 以供选择。
2. 修改 config.json 来配置 MareWood 启动的端口号、域名、和支持的前端工具。
3. 编译 MareWood 并放入后台运行。
4. 客户端在 public 目录，请安装依赖并修改正式请求地址然后运行 npm run build 即可。（请求地址=> /public/src/config/url.js）

其他：MareWood 提供 http 访问，建议 nginx 提供 https 静态文件服务器（指向resources\webs）这样可以灵活的选择 http 或者 https 去访问打包的前端项目。


那么，一个 React 或者 Vue 项目如何配置多个打包命令呢？
请移步http://xusenlin.com/article?path=%2f%e5%89%8d%e7%ab%af%e5%bc%80%e5%8f%91%2fREACT%e5%92%8cVUE%e7%9a%84%e5%a4%9a%e5%a5%97%e6%89%93%e5%8c%85%e7%8e%af%e5%a2%83.md

如果不明白也可以参考实际的项目

- https://github.com/xusenlin/VueMultiplePages.git
- https://github.com/xusenlin/VueElementUiAdmin.git

# 角色




# 预览

### Repository
![Repository](./preview/仓库.png)
### Category
![Category](./preview/分类.png)
### Job
![Job](./preview/任务.png)


# 番外

MareWood 翻译过来是 玛瑞·伍德 看上去像是一个人的名字，其实它是代表两个人，一个咒语。



# License

MIT © Richard McRichface
