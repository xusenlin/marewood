import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import AddIcon from '@material-ui/icons/Add';
import { Fab, Paper} from '@material-ui/core';
import Edit from './edit'
import Table from "./repositoryTable"
import { getSystemInfo } from "../../utils/dataStorage"

const styles = theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(1),
        overflowX: 'auto',
    },
    fab: {
        position: 'absolute',
        bottom: theme.spacing(2),
        right: theme.spacing(2),
    },
});


class RepositoryTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            editShow: false,
            dependentSupport: getSystemInfo("DependTools") ||[]
        };

    }

    editDialogShow() {
        this.setState({editShow: true})
    }

    editDialogClose() {
        this.setState({editShow: false})
    }

    createSuccess() {
        this.setState({editShow: false});
        this.getTableData()
    }

    render() {
        const {classes} = this.props;
        return (
            <div>
                <Paper className={classes.root}>
                    <Table/>
                </Paper>
                <Fab color="primary" className={classes.fab} aria-label="add" onClick={this.editDialogShow.bind(this)}>
                    <AddIcon/>
                </Fab>
                <Edit
                    dependentSupport={this.state.dependentSupport}
                    show={this.state.editShow}
                    handleClose={this.editDialogClose.bind(this)}
                    createSuccess={this.createSuccess.bind(this)}
                />
            </div>
        );
    }

}


export default withStyles(styles)(RepositoryTable)
