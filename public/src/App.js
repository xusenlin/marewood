import React from "react";
import {HashRouter, Switch, Route,Redirect} from "react-router-dom";
import CssBaseline from '@material-ui/core/CssBaseline';
import Admin from './views/layout/index.js'
import Login from './views/login/index.js'
import Register from './views/register/index.js'
import NotFound from "./views/error/404.js";
import { getToken } from "./utils/dataStorage.js"
function App() {
    return (
        <div>
            <CssBaseline/>
            <HashRouter>
                <Switch>
                    <Route exact path="/register" component={Register} />
                    <Route exact path="/login" component={Login} />
                    <Route exact path="/404" component={NotFound}/>
                    <Route path="/">
                      {
                        getToken() ?  <Admin/> : <Redirect  to="/login"/>
                      }
                    </Route>
                </Switch>
            </HashRouter>
        </div>
    );
}

export default App;
