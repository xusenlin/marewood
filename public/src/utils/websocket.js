import { host } from "../config/url.js";
import Snackbar from "../components/snackbar/index.js";

export function connWebsocket(token) {
  window.ws = new WebSocket(`ws://${host}/websocket?token=${token}`);

  window.ws.addEventListener("message", r => {
    let data = JSON.parse(r.data);
    if (data.Type === 0) {
      Snackbar.success(data.Message);
      return;
    }
    if (window.wsFunc) {
      window.wsFunc();
    }
    Snackbar.info(data.Message);
  });
  window.ws.addEventListener("error", e => {
    Snackbar.error("WebSocket 链接出错！");
  });
}
