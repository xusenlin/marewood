import React from "react";
import ReactDOM from "react-dom";
import "nprogress/nprogress.css";
import "./assets/style/style.scss";
import { getToken } from "./utils/dataStorage.js";
import { connWebsocket } from "./utils/websocket.js";
import App from "./App";
import * as serviceWorker from "./serviceWorker";

ReactDOM.render(<App />, document.getElementById("root"));

let token = getToken();
if (token && !window.ws && window.location.href.indexOf("/login") === -1) {
  connWebsocket(token);
}

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
