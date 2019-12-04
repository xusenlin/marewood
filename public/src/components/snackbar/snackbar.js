import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import ErrorIcon from '@material-ui/icons/Error';
import InfoIcon from '@material-ui/icons/Info';
import CloseIcon from '@material-ui/icons/Close';
import {amber, green} from '@material-ui/core/colors/index';
import IconButton from '@material-ui/core/IconButton/index';
import Snackbar from '@material-ui/core/Snackbar/index';
import SnackbarContent from '@material-ui/core/SnackbarContent/index';
import WarningIcon from '@material-ui/icons/Warning';
import {makeStyles} from '@material-ui/core/styles/index';
import ReactDOM from "react-dom";


const variantIcon = {
    success: CheckCircleIcon,
    warning: WarningIcon,
    error: ErrorIcon,
    info: InfoIcon,
};
const useStyles1 = makeStyles(theme => ({
    success: {
        backgroundColor: green[600],
    },
    error: {
        backgroundColor: theme.palette.error.dark,
    },
    info: {
        backgroundColor: theme.palette.primary.main,
    },
    warning: {
        backgroundColor: amber[700],
    },
    icon: {
        fontSize: 20,
    },
    iconVariant: {
        opacity: 0.9,
        marginRight: theme.spacing(1),
    },
    message: {
        display: 'flex',
        alignItems: 'center',
    },
}));

function SnackbarContentWrapper(props) {
    const classes = useStyles1();
    const {message, onClose, variant, ...other} = props;
    const Icon = variantIcon[variant];

    return (
        <SnackbarContent
            className={clsx(classes[variant])}
            aria-describedby="client-snackbar"
            message={
                <span id="client-snackbar" className={classes.message} style={{whiteSpace: "pre-wrap"}}>
          <Icon className={clsx(classes.icon, classes.iconVariant)}/>
                    {message}
        </span>
            }
            action={[
                <IconButton key="close" aria-label="close" color="inherit" onClick={onClose}>
                    <CloseIcon className={classes.icon}/>
                </IconButton>,
            ]}
            {...other}
        />
    );
}

SnackbarContentWrapper.propTypes = {
    message: PropTypes.string,
    onClose: PropTypes.func,
    variant: PropTypes.oneOf(['error', 'info', 'success', 'warning']).isRequired,
};

class Notification extends React.Component {
    constructor(props) {
        super(props);
        this.state = {open: true};
    }

    // componentWillReceiveProps(nextProps) {
    //
    // }
    openSnackbar(type, content, duration, onClose){
        console.log(type, content, duration, onClose);
        this.setState({open: true})
    }
    handleClose() {
        this.setState({open: false})
    }

    render() {

        return (
            <Snackbar
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                open={this.state.open}
                autoHideDuration={6000}
                onClose={this.handleClose.bind(this)}
            >
                <SnackbarContentWrapper
                    onClose={this.handleClose.bind(this)}
                    variant={this.props.type}
                    message={this.props.msg}
                />
            </Snackbar>
        );
    }

}


export default function createNotification() {
    const snackbarDom = document.getElementById('snackbar');
    return ReactDOM.render(<Notification type={"info"} msg={"saaaaaaaaa"} />, snackbarDom)
}

