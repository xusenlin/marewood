import React from 'react';
import Snackbar from '@material-ui/core/Snackbar/index';
import SnackbarContentWrapper from './snackbarContentWrapper'


class CustomizeSnackbar extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            open: false,
            content:"",
            duration:4000,
            type:"info"
        };
    }

    openSnackbar(type, content, duration){
        this.setState({
            open: true,
            type:type,
            content:content,
            duration:duration
        })
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
                autoHideDuration={this.state.duration}
                onClose={this.handleClose.bind(this)}
            >
                <SnackbarContentWrapper
                    onClose={this.handleClose.bind(this)}
                    variant={this.state.type}
                    message={this.state.content}
                />
            </Snackbar>
        );
    }

}


export default CustomizeSnackbar

