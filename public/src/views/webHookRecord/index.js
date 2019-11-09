import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import Computer from '@material-ui/icons/Computer';
import {Tooltip, IconButton,CircularProgress, Paper, TableRow, TableHead, TableCell, TableBody, Table} from '@material-ui/core';

import {webHookRecord} from '../../api/webHookRecord'


const styles = theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(3),
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


class RecordTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            tableData: [],
        };
    }

    componentDidMount() {
        if (this.props.location && this.props.location.query) {
            webHookRecord({id:this.props.location.query.id}).then(r => {
                this.setState({tableData: r});
            }).catch(() => {})
        }
    }
    render() {
        const {classes} = this.props;
        return (
            <div>
                <Paper className={classes.root}>
                    <Table className={classes.table}>
                        <TableHead>
                            <TableRow>
                                <TableCell>ID</TableCell>
                                <TableCell align="left">状态</TableCell>
                                <TableCell align="left">终端信息</TableCell>
                                <TableCell align="left">创建时间</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.tableData.map(row => (
                                <TableRow key={row.ID}>
                                    <TableCell component="th" scope="row">
                                        {row.ID}
                                    </TableCell>
                                    <TableCell component="th" scope="row">
                                        {row.Status === 1 ? (
                                            <CircularProgress variant="static" value={100} size={20}/>
                                        ):(
                                            <CircularProgress variant="static" value={100} size={20} color="secondary"/>
                                            )}
                                    </TableCell>
                                    <TableCell align="left">
                                        <Tooltip title={
                                            <div style={{whiteSpace: "pre-wrap"}}>
                                                {row.TerminalInfo}
                                            </div>
                                        } interactive>
                                            <IconButton color="primary">
                                                <Computer/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell component="th" scope="row">
                                        {row.CreatedAt}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>
            </div>
        );
    }
}


export default withStyles(styles)(RecordTable)
