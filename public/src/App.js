import React, { useEffect } from "react";
import { HashRouter, Switch, Route, Redirect } from "react-router-dom";
import i18n from "react-intl-universal";
import CssBaseline from "@material-ui/core/CssBaseline";
import Admin from "./views/layout/index.js";
import Login from "./views/login/index.js";
import Register from "./views/register/index.js";
import NotFound from "./views/error/404.js";
import { getToken, getUserLang } from "./utils/dataStorage.js";

const locales = {
  "en-US": require("./i18n/en-US.json"),
  "zh-CN": require("./i18n/zh-CN.json")
};

const App = () => {
  const [initI18nDone, setInitI18nDone] = React.useState(false);

  const loadLang = () => {
    i18n
      .init({
        currentLocale: getUserLang(),
        locales
      })
      .then(() => {
        setInitI18nDone(true);
      });
  };
  useEffect(() => {
    if (initI18nDone) return;
    loadLang();
  });

  return (
    initI18nDone && (
      <div>
        <CssBaseline />
        <HashRouter>
          <Switch>
            <Route exact path="/register" component={Register} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/404" component={NotFound} />
            <Route path="/">
              {getToken() ? <Admin /> : <Redirect to="/login" />}
            </Route>
          </Switch>
        </HashRouter>
      </div>
    )
  );
};

export default App;
