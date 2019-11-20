import React from 'react';
import {CircularProgress, IconButton, Tooltip} from "@material-ui/core";

const STATUS = ["正在更新","更新成功","更新失败"];

class RecordStatus extends React.Component {
    constructor(props) {
        super(props);
        this.state ={}
    }
    renderCircularProgress(status){
        if (1 === status ) return (<CircularProgress variant="static" value={100} size={20}/>);
        return (<CircularProgress variant="static" value={100} size={20} color="secondary"/>);
    }
    render() {
        return (
            <div>
                <Tooltip title={STATUS[this.props.status]}>
                    <IconButton  color="primary">
                        {this.renderCircularProgress(this.props.status)}
                    </IconButton>
                </Tooltip>
            </div>
        );
    }
}


export default RecordStatus
