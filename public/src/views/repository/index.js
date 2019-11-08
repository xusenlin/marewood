import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import AddIcon from '@material-ui/icons/Add';
import DeleteIcon from '@material-ui/icons/Delete';
import Fingerprint from '@material-ui/icons/Fingerprint';
import Computer from '@material-ui/icons/Computer';
import Announcement from '@material-ui/icons/Announcement';
import LinkIcon from '@material-ui/icons/Link'
import Comment from '@material-ui/icons/Comment'
import {
    Dialog, DialogContent, DialogTitle, DialogContentText, Tooltip,
    DialogActions, Button, IconButton, Fab, Paper, TableRow, TableHead, TableCell, TableBody, Table
} from '@material-ui/core';
import HelperTooltips from '../../components/helperTooltips'
import ApiUrl from '../../config/url.js'
import RepositoryStatus from './repositoryStatus'
import Edit from './edit'
import {repositories, destroy} from '../../api/repository'

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
            editShow: false
        };
        this.destroyId = 0; //记录当前要删除的id
        this.timeout = null;
    }

    componentDidMount() {
        this.getTableData()
    }
    componentWillUnmount() {
        if(this.timeout)clearTimeout(this.timeout);
    }
    getTableData() {
        if(this.timeout)clearTimeout(this.timeout);

        repositories().then(r => {
            this.setState({tableData: r});

            for (let i = 0; i < r.length; i++) {
                if (r[i].Status === 0) {
                    this.timeout = setTimeout(()=>{
                        this.getTableData()
                    },5000);
                    return
                }
            }

        }).catch(() => {})
    }

    destroyDialogOpen(id) {
        this.destroyId = id;
        this.setState({destroyDialogShow: true})
    }

    destroyDialogClose() {
        this.setState({destroyDialogShow: false})
    }

    destroyConfirm() {
        destroy({id: this.destroyId}).then(r => {
            this.setState({destroyDialogShow: false});
            this.getTableData()
        }).catch(() => {
        })
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
                                <TableCell align="left">操作 <HelperTooltips help="删除/提交记录"/> </TableCell>
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
                                                {row.TerminalInfo}
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
                                        <IconButton color="primary"
                                                    onClick={this.destroyDialogOpen.bind(this, row.ID)}>
                                            <DeleteIcon/>
                                        </IconButton>
                                        <IconButton color="primary"
                                                    onClick={this.destroyDialogOpen.bind(this, row.ID)}>
                                            <Comment/>
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
                    <DialogTitle id="alert-dialog-title">{"确认删除仓库?"}</DialogTitle>
                    <DialogContent>
                        <DialogContentText id="alert-dialog-description">
                            你确认要删除这个仓库？为了不影响以前的任务，本地的代码不会被移除。
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
                <Edit
                    show={this.state.editShow}
                    handleClose={this.editDialogClose.bind(this)}
                    createSuccess={this.createSuccess.bind(this)}
                />
            </div>
        );
    }

}


export default withStyles(styles)(RepositoryTable)
