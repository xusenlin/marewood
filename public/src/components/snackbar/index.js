import snackbar from './snackbarInstance'


let snackbarInstance;
const notice = (type, content, duration ) => {
    if (!snackbarInstance) snackbarInstance = snackbar;
    return snackbarInstance.openSnackbar(type, content, duration)
};

export default {
    info(content, duration= 2000) {
        return notice('info', content, duration)
    },
    success(content, duration) {
        return notice('success', content, duration)
    },
    warning(content, duration) {
        return notice('warning', content, duration)
    },
    error(content, duration) {
        return notice('error', content, duration)
    },
}
