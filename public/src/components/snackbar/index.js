import snackbar from "./snackbarInstance";

let snackbarInstance;
export const notice = (type, content, duration = 4000) => {
  if (!snackbarInstance) snackbarInstance = snackbar;
  return snackbarInstance.openSnackbar(type, content, duration);
};

export default {
  info(content, duration = 4000) {
    return notice("info", content, duration);
  },
  success(content, duration = 4000) {
    return notice("success", content, duration);
  },
  warning(content, duration = 4000) {
    return notice("warning", content, duration);
  },
  error(content, duration = 4000) {
    return notice("error", content, duration);
  }
};
