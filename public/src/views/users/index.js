import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { users,deleteUsers,roleEdit } from "../../api/user";
import {IconButton, Tooltip} from "@material-ui/core";
import DeleteIcon from '@material-ui/icons/Delete';
import ThumbDown from '@material-ui/icons/ThumbDown';
import ThumbUp from '@material-ui/icons/ThumbUp';
import Snackbar from "../../components/snackbar";
import RoleDesc from './roleDesc'



const styles = theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(1),
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
    sept: {
        marginTop:20,
        padding: theme.spacing(3, 2),
    },
});



class User extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            tableData: [],
        };
    }
    componentDidMount() {
        this.getUserList()
    }
    getUserList(){
        users().then(r=>{
            this.setState({tableData:r})
        }).catch(()=>{})
    }
    destroyUser(row){
        deleteUsers({id:row.ID}).then(r=>{
            Snackbar.success("删除成功");
            this.getUserList();
        }).catch(()=>{})
    }
    role(row,isUp = 0){
        roleEdit({id:row.ID,isUp:isUp}).then(r=>{
            Snackbar.success("修改成功");
            this.getUserList();
        }).catch(()=>{})
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
                                <TableCell align="left">用户名字</TableCell>
                                <TableCell align="left">角色</TableCell>
                                <TableCell align="left">备注</TableCell>
                                <TableCell align="center">操作</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.tableData.map(row => (
                                <TableRow key={row.ID} hover>
                                    <TableCell component="th" scope="row">
                                        {row.ID}
                                    </TableCell>
                                    <TableCell align="left">{ row.Username }</TableCell>
                                    <TableCell align="left"><RoleDesc role={row.Role}/></TableCell>
                                    <TableCell align="left">{ row.Desc }</TableCell>
                                    <TableCell align="center">
                                        <Tooltip title="角色提升">
                                            <IconButton color="primary" onClick={this.role.bind(this,row,1)}>
                                                <ThumbUp />
                                            </IconButton>
                                        </Tooltip>
                                        <Tooltip title="角色降级">
                                            <IconButton color="primary" onClick={this.role.bind(this,row,0)}>
                                                <ThumbDown />
                                            </IconButton>
                                        </Tooltip>
                                        <Tooltip title="删除用户">
                                            <IconButton color="primary" onClick={this.destroyUser.bind(this,row)}>
                                                <DeleteIcon />
                                            </IconButton>
                                        </Tooltip>
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

export default withStyles(styles)(User)

