import React from "react";
import {HashRouter, Switch, Route} from "react-router-dom";
import { SnackbarProvider } from 'notistack';
import CssBaseline from '@material-ui/core/CssBaseline';
import Admin from './views/layout/index.js'


function App() {
    return (

        <SnackbarProvider maxSnack={1}>
            <CssBaseline/>
            <HashRouter>
                <Switch>
                    {/*login page*/}
                    <Route path="/" component={Admin} />
                </Switch>
            </HashRouter>
        </SnackbarProvider>
    );
}

export default App;
