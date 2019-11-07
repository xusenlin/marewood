import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import AddIcon from '@material-ui/icons/Add';
import DeleteIcon from '@material-ui/icons/Delete';
import Fingerprint from '@material-ui/icons/Fingerprint';
import Computer from '@material-ui/icons/Computer';
import Announcement from '@material-ui/icons/Announcement';
import LinkIcon from '@material-ui/icons/Link'
import {
    Dialog, DialogContent, DialogTitle, DialogContentText, Tooltip,
    DialogActions, Button, IconButton, Fab, Paper, TableRow, TableHead, TableCell, TableBody, Table
} from '@material-ui/core';
import ApiUrl from '../../config/url.js'
import RepositoryStatus from './repositoryStatus.js'
import EditCategory from './edit.js'
import {repositories} from '../../api/repository.js'

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


class RepositoryTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            tableData: [],
            destroyDialogShow: false,
            editCategoryShow: false
        };
        this.destroyId = 0;
    }

    componentDidMount() {
        this.getTableData()
    }

    getTableData() {
        repositories().then(r => {
            this.setState({tableData: r})
        }).catch(() => {
        })
    }

    destroyDialogOpen(id) {
        this.destroyId = id;
        this.setState({destroyDialogShow: true})
    }

    destroyDialogClose() {
        this.setState({destroyDialogShow: false})
    }

    destroyConfirm() {
        // destroy({id:this.destroyId}).then(r=>{
        //     this.setState({destroyDialogShow:false});
        //     this.getTableData()
        // }).catch(()=>{})
    }

    editDialogShow() {
        this.setState({editCategoryShow: true})
    }

    editDialogClose() {
        this.setState({editCategoryShow: false})
    }

    createCategorySuccess() {
        this.setState({editCategoryShow: false});
        this.getTableData()
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
                                <TableCell align="left">仓库名字</TableCell>
                                <TableCell align="left">状态</TableCell>
                                <TableCell align="left">仓库权限</TableCell>
                                <TableCell align="left">终端信息</TableCell>
                                <TableCell align="left">备注</TableCell>
                                <TableCell align="left">仓库地址</TableCell>
                                <TableCell align="left">webHook密钥</TableCell>
                                <TableCell align="left">webHook地址</TableCell>
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
                                    <TableCell align="left">
                                        <RepositoryStatus status={row.Status}/>
                                    </TableCell>
                                    <TableCell align="left">
                                        <Tooltip title={
                                            <div>
                                                用户名：{row.UserName}
                                                <br/>
                                                密码：{row.Password}
                                            </div>
                                        } interactive>
                                            <IconButton color="primary">
                                                <Fingerprint/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="left">
                                        <Tooltip title={
                                            <div style={{whiteSpace: "pre-wrap"}}>
                                                {row.ConsoleOutput}
                                            </div>
                                        } interactive>
                                            <IconButton color="primary">
                                                <Computer/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="left">
                                        <Tooltip title={row.Desc} interactive>
                                            <IconButton color="primary">
                                                <Announcement/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="left">
                                        <Tooltip title={row.Url} interactive>
                                            <IconButton color="primary">
                                                <LinkIcon/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="left">{row.WebHookSecret}</TableCell>
                                    <TableCell align="left">{ApiUrl}/web_hook?id={row.ID}</TableCell>
                                    <TableCell align="left">{row.CreatedAt}</TableCell>
                                    <TableCell align="left">
                                        <IconButton edge="start" color="primary"
                                                    onClick={this.destroyDialogOpen.bind(this, row.ID)}>
                                            <DeleteIcon/>
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
                <Fab color="primary" className={classes.fab} aria-label="add" onClick={this.editDialogShow.bind(this)}>
                    <AddIcon/>
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


export default withStyles(styles)(RepositoryTable)
