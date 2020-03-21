import React from 'react';
import {Switch, Route, Redirect,withRouter} from "react-router-dom";
import Drawer from '@material-ui/core/Drawer';
import TopBar from './topBar.js'
import RouteList from '../../router/route'
import {IconButton} from "@material-ui/core";
import './admin.scss'


const Layout = props => {

    const [menuState,setMenuState] = React.useState(false);
    const [title,setTitle] = React.useState('');

    React.useEffect(() => {
        for (let i=0;i<RouteList.length;i++){
            let r = RouteList[i];
            if(r.path === props.location.pathname){
                setTitle(r.title);
                document.title = r.title;
                return;
            }
        }
    },[props.location.pathname]);

    const openMenu = (menu)=>{
        if(menu.path === props.location.pathname) return;
        setTitle(menu.title);
        setMenuState(false);
        props.history.push(menu.path);
    };

    return (
        <div id="admin" className="admin">
            <TopBar title={title} toggle={()=>{setMenuState(true)}} clickRoute={openMenu} />
            <Drawer anchor="top" open={menuState} onClose={()=>{setMenuState(false)}} style={{flexDirection:"column"}}>
                <div className="menuList">
                    {
                        RouteList.map(menu => {
                            if(menu.isMenu)
                            return (
                                <div className="menuBox" key={menu.path}>
                                    <IconButton edge="start" color="primary" aria-label="menu" onClick={()=>{openMenu(menu)}}>
                                        <menu.icon/>
                                    </IconButton>
                                    <p>{ menu.title }</p>
                                </div>
                            );
                            return false;
                        })
                    }
                </div>
            </Drawer>
            <div className="container">
                <Switch>
                    {
                        RouteList.map(route => {
                            return (
                                <Route exact key={route.path} path={route.path} component={route.component}/>
                            )
                        })
                    }
                    <Redirect exact from="/" to="/dashboard"/>
                    <Redirect  to="/404" />
                </Switch>
            </div>
        </div>
    );

};


export default withRouter(Layout);
