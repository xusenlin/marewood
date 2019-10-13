import React from 'react';
import {withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Fab from '@material-ui/core/Fab';
import AddIcon from '@material-ui/icons/Add';
import {categories} from '../../api/category.js'

const styles = theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(3),
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
    fab: {
        position: 'absolute',
        bottom: theme.spacing(2),
        right: theme.spacing(2),
    },
});


class CategoryTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            tableData:[]
        };
    }

    componentDidMount() {
        categories().then(r=>{
            this.setState({tableData:r})
        }).catch(()=>{})
    }
    render() {
        const { classes } = this.props;
        return (
            <div>
                <Paper className={classes.root}>
                    <Table className={classes.table}>
                        <TableHead>
                            <TableRow>
                                <TableCell>ID</TableCell>
                                <TableCell align="left">分类名字</TableCell>
                                <TableCell align="left">项目数量</TableCell>
                                <TableCell align="left">备注</TableCell>
                                <TableCell align="left">创建时间</TableCell>
                                <TableCell align="left">操作</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.tableData.map(row => (
                                <TableRow key={row.ID}>
                                    <TableCell component="th" scope="row">
                                        {row.ID}
                                    </TableCell>
                                    <TableCell align="left">{row.Name}</TableCell>
                                    <TableCell align="left">{row.fat}</TableCell>
                                    <TableCell align="left">{row.carbs}</TableCell>
                                    <TableCell align="left">{row.protein}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>
                <Fab color="primary" className={classes.fab} aria-label="add">
                    <AddIcon />
                </Fab>
            </div>
        );
    }

}


export default withStyles(styles)(CategoryTable)