import { host } from "../config/url.js";
import Snackbar, { notice } from "../components/snackbar/index.js";
import BusEvent from "./busEvents";

const websocketRetryQuantity = 5;
let websocketQuantity = 0;
export function connWebsocket(token) {
  if (!token) return;
  websocketQuantity++;
  window.ws = new WebSocket(`ws://${host}/websocket?token=${token}`);
  window.ws.onopen = () => {
    Snackbar.success("WebSocket已连接");
  };
  window.ws.onmessage = r => {
    let data = JSON.parse(r.data);
    notice(data.Type, data.Message);
    BusEvent.updateCurrentTable();
    BusEvent.pushNotice(data);
  };
  window.ws.onerror = () => {
    if (websocketQuantity > websocketRetryQuantity) {
      return;
    }
    Snackbar.error("WebSocket出错,2s后尝试链接...");
    setTimeout(() => {
      connWebsocket(token);
    }, 2000);
  };
  window.ws.onclose = () => {
    if (websocketQuantity > websocketRetryQuantity) {
      return;
    }
    Snackbar.warning("WebSocket已关闭,2s后尝试链接...");
    setTimeout(() => {
      connWebsocket(token);
    }, 2000);
  };
}
