import React from "react";
import {HashRouter, Switch, Route} from "react-router-dom";
import CssBaseline from '@material-ui/core/CssBaseline';
import Admin from './views/layout/index.js'


function App() {
    return (
        <React.Fragment>
            <CssBaseline/>
            <HashRouter>
                <Switch>
                    {/*login page*/}
                    <Route path="/" component={Admin} />
                </Switch>
            </HashRouter>
        </React.Fragment>
    );
}

export default App;
