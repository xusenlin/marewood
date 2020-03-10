import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import AddIcon from '@material-ui/icons/Add';
import DeleteIcon from '@material-ui/icons/Delete';
import {
    Table, TableRow, TableBody, TableHead, TableCell, Paper, Fab, IconButton, Button,
    Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions, Tooltip
} from '@material-ui/core';
import EditCategory from './edit.js'
import {categories,destroy} from '../../api/category.js'
import Snackbar from "../../components/snackbar";


const styles = theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(1),
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
            editCategoryShow:false
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
    }

    destroyDialogClose () {
        this.setState({destroyDialogShow:false})
    }
    destroyConfirm () {
        destroy({id:this.destroyId}).then(r=>{
            Snackbar.success("删除成功！");
            this.setState({destroyDialogShow:false});
            this.getTableData()
        }).catch(()=>{})
    }
    editDialogShow () {
        this.setState({editCategoryShow:true})
    }
    editDialogClose () {
        this.setState({editCategoryShow:false})
    }

    createCategorySuccess(){
        this.setState({editCategoryShow:false});
        this.getTableData()
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
                                    <TableCell align="left">{ row.JobQuantity }</TableCell>
                                    <TableCell align="left">{row.Desc}</TableCell>
                                    <TableCell align="left">{row.CreatedAt}</TableCell>
                                    <TableCell align="left">
                                        <Tooltip title="删除分类">
                                            <IconButton edge="start" color="primary" onClick={this.destroyDialogOpen.bind(this,row.ID)}>
                                                <DeleteIcon />
                                            </IconButton>
                                        </Tooltip>
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
                <Fab color="primary" className={classes.fab} aria-label="add" onClick={this.editDialogShow.bind(this)}>
                    <AddIcon />
                </Fab>
                <EditCategory
                    show={this.state.editCategoryShow}
                    handleClose={this.editDialogClose.bind(this)}
                    createSuccess={this.createCategorySuccess.bind(this)}
                />
            </div>
        );
    }

}


export default withStyles(styles)(CategoryTable)
