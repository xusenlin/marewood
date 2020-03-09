import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import LinkIcon from '@material-ui/icons/Link'
import Computer from '@material-ui/icons/Computer';
import Announcement from '@material-ui/icons/Announcement';
import PlayCircleFilled from '@material-ui/icons/PlayCircleFilled';
import DeviceHub from '@material-ui/icons/DeviceHub';
import {
    TableRow, TableHead, TableCell, TableBody, Table, Tooltip, IconButton
} from '@material-ui/core';
import JobStatus from "./jobStatus"
import HelperTooltips from "../../components/helperTooltips";
import SwitchBranchDialog from "./switchBranchDialog"


const styles = theme => ({
    table: {
        marginTop: theme.spacing(3),
        minWidth: 650,
    },
});


class CategoriesTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            switchBranchDialog:{
                id:0,
                show:false,
                selectedValue:"master"
            }
        };
    }

    componentDidMount() {

    }
    componentWillUnmount() {

    }
    runJob(jobId){
        alert(jobId)
    }
    openSwitchBranchDialog(row){
        this.setState({
            switchBranchDialog:{
                id:row.ID,
                show:true,
                selectedValue:row.Branch
            }
        })
    }
    closeSwitchBranchDialog(selectedValue){
        let oldVal = this.state.switchBranchDialog;
        this.setState({switchBranchDialog:{
                id:oldVal.id,
                show:false,
                selectedValue:oldVal.selectedValue
            }})
    }
    render() {
        const {classes} = this.props;
        return (
            <div>
                <Table className={classes.table}>
                    <TableHead>
                        <TableRow>
                            <TableCell>ID</TableCell>
                            <TableCell align="center">任务名称</TableCell>
                            <TableCell align="center">任务状态</TableCell>
                            <TableCell align="center">
                                当前分支
                                <HelperTooltips help="当前任务分支，如有需要请先切换分支"/>
                            </TableCell>
                            <TableCell align="center">
                                访问地址
                                <HelperTooltips help="任务状态成功(1)时可访问页面"/>
                            </TableCell>
                            <TableCell align="center">终端信息</TableCell>
                            <TableCell align="center">
                                附加脚本
                                <HelperTooltips help="打包成功运行的脚本，多个用 ; 隔开"/>
                            </TableCell>
                            <TableCell align="center">任务描述</TableCell>
                            <TableCell align="center">打包目录</TableCell>
                            <TableCell align="center">打包命令</TableCell>
                            <TableCell align="center">操作</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {
                            this.props.tableData.map(row => (
                                <TableRow key={row.ID}>
                                    <TableCell component="th" scope="row">
                                        {row.ID}
                                    </TableCell>
                                    <TableCell align="center">{row.Name}</TableCell>
                                    <TableCell align="center">
                                        <JobStatus status={row.Status} />
                                    </TableCell>
                                    <TableCell align="center">{row.Branch}</TableCell>
                                    <TableCell align="center">
                                        <Tooltip title={row.Url} interactive>
                                            <IconButton color="primary">
                                                <LinkIcon/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="center">
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
                                    <TableCell align="center">
                                        <Tooltip title={
                                            <div style={{whiteSpace: "pre-wrap"}}>
                                                {row.SuccessScript}
                                            </div>
                                        } interactive>
                                            <IconButton color="primary">
                                                <Computer/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="center">
                                        <Tooltip title={row.Desc} interactive>
                                            <IconButton color="primary">
                                                <Announcement/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell align="center">
                                        <span className="tag">{(row.BuildDir)}</span>
                                    </TableCell>
                                    <TableCell align="center">
                                        <span className="tag">{(row.BuildCommand)}</span>
                                    </TableCell>
                                    <TableCell align="center">
                                        <Tooltip title="切换分支" interactive>
                                            <IconButton color="primary" onClick={this.openSwitchBranchDialog.bind(this,row)}>
                                                <DeviceHub/>
                                            </IconButton>
                                        </Tooltip>
                                        <Tooltip title="运行任务" interactive>
                                            <IconButton color="primary"
                                                        onClick={this.runJob.bind(this, row.ID)}>
                                                <PlayCircleFilled/>
                                            </IconButton>
                                        </Tooltip>
                                    </TableCell>

                                </TableRow>
                            ))
                        }
                    </TableBody>
                </Table>
                <SwitchBranchDialog
                    jobId={this.state.switchBranchDialog.id}
                    selectedValue={this.state.switchBranchDialog.selectedValue}
                    open={this.state.switchBranchDialog.show}
                    onClose={this.closeSwitchBranchDialog.bind(this)} />
            </div>
        );
    }

}


export default withStyles(styles)(CategoriesTable)
