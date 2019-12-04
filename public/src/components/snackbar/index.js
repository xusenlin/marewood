import snackbar from './snackbar'


let snackbarInstance;
const notice = (type, content, duration = 2000, onClose) => {
    if (!snackbarInstance) snackbarInstance = snackbar;
    return snackbarInstance.openSnackbar(type, content, duration, onClose)
};

export default {
    info(content, duration, onClose) {
        return notice('info', content, duration, onClose)
    },
    success(content, duration, onClose) {
        return notice('success', content, duration, onClose)
    },
    warning(content, duration, onClose) {
        return notice('warning', content, duration, onClose)
    },
    error(content, duration, onClose) {
        return notice('error', content, duration, onClose)
    },
}
