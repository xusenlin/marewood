import Dashboard from '../views/dashboard';
import Category from '../views/category';
import Repository from '../views/repository';
import WebHookRecord from '../views/webHookRecord/index'
import Job from '../views/job';

import CategoryIcon from '@material-ui/icons/Category';
import HomeIcon from '@material-ui/icons/Home';
import ViewListIcon from '@material-ui/icons/ViewList';
import DashboardIcon from '@material-ui/icons/Dashboard';
import FolderOpenIcon from '@material-ui/icons/FolderOpen';



//https://material.io/resources/icons/?style=baseline


export default [
    {
        title:'仪表盘',
        path:'/dashboard',
        isMenu:true,
        icon:HomeIcon,
        component:Dashboard

    },
    {
        title:'仓库列表',
        path:'/repository',
        isMenu:true,
        icon:DashboardIcon,
        component:Repository
    },
    {
        title:'任务分类',
        path:'/category',
        isMenu:true,
        icon:CategoryIcon,
        component:Category
    },
    {
        title:'部署任务',
        path:'/job',
        isMenu:true,
        icon:ViewListIcon,
        component:Job
    },
    {
        title:'工具箱',
        path:'/folder',
        isMenu:true,
        icon:FolderOpenIcon,
        component:Job
    },
    {
        title:'仓库更新记录',
        path:'/webHookRecord',
        component:WebHookRecord
    },
]
