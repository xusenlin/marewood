import Dashboard from "../views/dashboard";
import Category from "../views/category";
import Repository from "../views/repository";
import Task from "../views/task";
import Users from "../views/users";

import CategoryIcon from "@material-ui/icons/Category";
import HomeIcon from "@material-ui/icons/Home";
import ViewListIcon from "@material-ui/icons/ViewList";
import DashboardIcon from "@material-ui/icons/Dashboard";
import PeopleIcon from "@material-ui/icons/People";

export default [
  {
    title: "bar.dashboard",
    path: "/dashboard",
    isMenu: true,
    icon: HomeIcon,
    component: Dashboard
  },
  {
    title: "bar.repositories",
    path: "/repository",
    isMenu: true,
    icon: DashboardIcon,
    component: Repository
  },
  {
    title: "bar.category",
    path: "/category",
    isMenu: true,
    icon: CategoryIcon,
    component: Category
  },
  {
    title: "bar.task",
    path: "/task",
    isMenu: true,
    icon: ViewListIcon,
    component: Task
  },
  {
    title: "bar.users",
    path: "/users",
    isMenu: true,
    icon: PeopleIcon,
    component: Users
  }
];
