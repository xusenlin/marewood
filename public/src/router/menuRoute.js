import Dashboard from '../views/dashboard';
import Category from '../views/category';
import List from '../views/itemList';

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
        icon:HomeIcon,
        component:Dashboard

    },
    {
        title:'仓库分类',
        path:'/category',
        icon:CategoryIcon,
        component:Category
    },
    {
        title:'仓库列表',
        path:'/repository',
        icon:DashboardIcon,
        component:List
    },
    {
        title:'部署任务',
        path:'/job',
        icon:ViewListIcon,
        component:List
    },
    {
        title:'工具箱',
        path:'/folder',
        icon:FolderOpenIcon,
        component:List
    },
]
