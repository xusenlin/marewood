import Dashboard from '../views/dashboard';
import Category from '../views/category';
import Repository from '../views/repository';
import Job from '../views/job';

import CategoryIcon from '@material-ui/icons/Category';
import HomeIcon from '@material-ui/icons/Home';
import ViewListIcon from '@material-ui/icons/ViewList';
import DashboardIcon from '@material-ui/icons/Dashboard';
import PeopleIcon from '@material-ui/icons/People';

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
        title:'用户管理',
        path:'/folder',
        isMenu:true,
        icon:PeopleIcon,
        component:Dashboard
    },

]
