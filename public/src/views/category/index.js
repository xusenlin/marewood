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
import DeleteIcon from '@material-ui/icons/Delete';
import {IconButton} from "@material-ui/core";
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

import {categories,destroy} from '../../api/category.js'

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
            tableData:[],
            destroyDialogShow:false,

        };
        this.destroyId = 0;
    }

    componentDidMount() {
        this.getTableData()
    }

    getTableData(){
        categories().then(r=>{
            this.setState({tableData:r})
        }).catch(()=>{})
    }

    destroyDialogOpen(id){
        this.destroyId = id;
        this.setState({destroyDialogShow:true})
    };

    destroyDialogClose () {
        this.setState({destroyDialogShow:false})
    }
    destroyConfirm () {
        destroy({id:this.destroyId}).then(r=>{
            this.setState({destroyDialogShow:false});
            this.getTableData()
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
                                <TableCell align="left">任务数量</TableCell>
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
                                    <TableCell align="left">{ 10 }</TableCell>
                                    <TableCell align="left">{row.Desc}</TableCell>
                                    <TableCell align="left">{row.CreatedAt}</TableCell>
                                    <TableCell align="left">
                                        <IconButton edge="start" color="primary" onClick={this.destroyDialogOpen.bind(this,row.ID)}>
                                            <DeleteIcon />
                                        </IconButton>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>
                <Dialog
                    open={this.state.destroyDialogShow}
                    onClose={this.destroyDialogClose.bind(this)}
                    aria-labelledby="alert-dialog-title"
                    aria-describedby="alert-dialog-description"
                >
                    <DialogTitle id="alert-dialog-title">{"确认删除分类?"}</DialogTitle>
                    <DialogContent>
                        <DialogContentText id="alert-dialog-description">
                            如果你确认要删除这个分类，请确保在这个分类下面已经没有任何任务了。
                        </DialogContentText>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.destroyDialogClose.bind(this)} color="primary">
                            关闭
                        </Button>
                        <Button onClick={this.destroyConfirm.bind(this)} color="secondary" autoFocus>
                            确认
                        </Button>
                    </DialogActions>
                </Dialog>
                <Fab color="primary" className={classes.fab} aria-label="add">
                    <AddIcon />
                </Fab>
            </div>
        );
    }

}


export default withStyles(styles)(CategoryTable)