import { host } from "../config/url.js";
import Snackbar, { notice } from "../components/snackbar/index.js";

const UpdateDataTypeIsNotice = 0,
  UpdateDataTypeIsRepoAction = 1,
  UpdateDataTypeIsCategoryAction = 2,
  UpdateDataTypeIsJobAction = 3,
  UpdateDataTypeIsUserAction = 4;

export function connWebsocket(token) {
  window.ws = new WebSocket(`ws://${host}/websocket?token=${token}`);

  window.ws.addEventListener("message", r => {
    let data = JSON.parse(r.data);
    notice(data.Type, data.Message);

    if (data.UpdateDataType === UpdateDataTypeIsNotice) {
      return;
    }
    if (data.UpdateDataType === UpdateDataTypeIsJobAction && window.wsFuncJob) {
      window.wsFuncJob();
      return;
    }
    if (
      data.UpdateDataType === UpdateDataTypeIsRepoAction &&
      window.wsFuncRepo
    ) {
      window.wsFuncRepo();
      return;
    }
    if (
      data.UpdateDataType === UpdateDataTypeIsCategoryAction &&
      window.wsFuncCategory
    ) {
      window.wsFuncCategory();
      return;
    }
    if (
      data.UpdateDataType === UpdateDataTypeIsUserAction &&
      window.wsFuncUser
    ) {
      window.wsFuncUser();
      return;
    }
  });
  window.ws.addEventListener("error", e => {
    Snackbar.error("WebSocket 链接出错！");
  });
}
