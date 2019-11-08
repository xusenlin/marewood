import React from 'react';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';
import {Tooltip} from '@material-ui/core';


class HelperTooltips extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    componentDidMount() {

    }
    render() {
        return (
            <Tooltip title={this.props.help}>
                <HelpOutlineIcon fontSize="small"/>
            </Tooltip>
        );
    }
}

export default HelperTooltips
