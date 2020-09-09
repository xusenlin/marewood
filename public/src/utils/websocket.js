import { host } from "../config/url.js";
import Snackbar, { notice } from "../components/snackbar/index.js";

export function connWebsocket(token) {
  window.ws = new WebSocket(`ws://${host}/websocket?token=${token}`);
  window.ws.onopen = () => {
    Snackbar.success("WebSocket已连接");
  };
  window.ws.onmessage = r => {
    let data = JSON.parse(r.data);
    notice(data.Type, data.Message);
    if (window.wsUpdateDataFunc) {
      window.wsUpdateDataFunc();
    }
  };

  window.ws.onerror = () => {
    Snackbar.error("WebSocket出错");
  };
  window.ws.onclose = () => {
    Snackbar.warning("WebSocket已关闭");
  };
}
