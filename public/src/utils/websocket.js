import Snackbar from "../components/snackbar/index";
// let host = window.location.host;
let host = "localhost:8081";

export function connWebsocket(token) {
  window.ws = new WebSocket(`ws://${host}/websocket?token=${token}`);

  window.ws.addEventListener("message", r => {
    let data = JSON.parse(r.data);
    if (data.Type === 0) {
      Snackbar.success(data.Message);
    }
  });
  window.ws.addEventListener("error", e => {
    Snackbar.error("WebSocket 链接出错！");
  });
}
