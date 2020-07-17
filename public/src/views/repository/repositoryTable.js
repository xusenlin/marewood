import React, { Fragment } from "react";
import {
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Tooltip,
  IconButton,
  Table,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
  Dialog
} from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";
import EditIcon from "@material-ui/icons/Edit";
import FormatClearIcon from "@material-ui/icons/FormatClear";
import RestoreFromTrash from "@material-ui/icons/RestoreFromTrash";
import LockIcon from "@material-ui/icons/Lock";
import LockOpenIcon from "@material-ui/icons/LockOpen";
import SaveAlt from "@material-ui/icons/SaveAlt";
import Computer from "@material-ui/icons/Computer";
import Announcement from "@material-ui/icons/Announcement";
import LinkIcon from "@material-ui/icons/Link";
import { withStyles } from "@material-ui/core/styles";
import HelperTooltips from "../../components/helperTooltips";
import RepositoryStatus from "./repositoryStatus";
import RepositoryJobStatus from "./repositoryJobStatus";
import Snackbar from "../../components/snackbar";
import {
  deleteDepend,
  destroy,
  gitPull,
  UpdateField,
  pruneBranch
} from "../../api/repository";
import EditField from "../../components/editField";
import { tooltip } from "../../assets/jss/common";

const styles = () => ({
  table: {
    minWidth: 400
  },
  tooltip
});

class RepositoryTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      destroyDialogShow: false,
      editField: {
        id: 0,
        open: false,
        rows: 1,
        desc: "",
        field: "",
        fieldContent: ""
      }
    };
    this.destroyId = 0; //记录当前要删除的id
  }

  destroyDialogOpen(id) {
    this.destroyId = id;
    this.setState({ destroyDialogShow: true });
  }

  destroyDialogClose() {
    this.setState({ destroyDialogShow: false });
  }

  deleteRepositoryDepend(row) {
    if (row.Status !== 1) {
      Snackbar.warning("仓库状态不正常，无法删除依赖");
      return;
    }
    deleteDepend({ id: row.ID })
      .then(r => {
        Snackbar.success(r);
      })
      .catch(() => {});
  }

  destroyConfirm() {
    destroy({ id: this.destroyId })
      .then(r => {
        this.setState({ destroyDialogShow: false });
        this.props.refresh();
      })
      .catch(() => {});
  }

  repositoryGitPull(row) {
    gitPull({ id: row.ID })
      .then(() => {
        Snackbar.success("已经更新代码");
      })
      .catch(() => {});
  }

  repositoryPruneBranch(row) {
    pruneBranch({ id: row.ID })
      .then(() => {
        Snackbar.success("分支已经裁剪");
      })
      .catch(() => {});
  }

  clickEditField(row, inputRows, desc, field) {
    this.setState({
      editField: {
        id: row.ID,
        open: true,
        rows: inputRows,
        desc: desc,
        field: field,
        fieldContent: row[field]
      }
    });
  }
  editFieldSuccess(id, field, fieldContent) {
    UpdateField({ id, field, fieldContent })
      .then(() => {
        this.closeResetEditFieldDialog();
        this.props.refresh();
      })
      .catch(() => {});
  }
  closeResetEditFieldDialog() {
    this.setState({
      editField: {
        id: 0,
        open: false,
        rows: 1,
        desc: "",
        field: "",
        fieldContent: ""
      }
    });
  }

  render() {
    const { classes } = this.props;
    return (
      <Fragment>
        <Table className={classes.table}>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell align="center">仓库地址</TableCell>
              <TableCell align="center">仓库名字</TableCell>
              <TableCell align="center">克隆状态</TableCell>
              <TableCell align="center">
                工作状态
                <HelperTooltips help="当前工作目录正在执行其他部署任务，资源被占用" />
              </TableCell>
              <TableCell align="center">
                仓库权限
                <HelperTooltips help="私密仓库是需要用户名和密码的，否则克隆失败" />
              </TableCell>
              <TableCell align="center">终端信息</TableCell>
              <TableCell align="center">
                备注
                <HelperTooltips help="点击可以修改" />
              </TableCell>
              <TableCell align="center">依赖工具</TableCell>
              <TableCell align="center">操作</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {this.props.tableData.map(row => (
              <TableRow key={row.ID} hover>
                <TableCell component="th" scope="row">
                  {row.ID}
                </TableCell>
                <TableCell align="center">
                  <Tooltip title={row.Url} interactive>
                    <IconButton color="primary">
                      <LinkIcon />
                    </IconButton>
                  </Tooltip>
                </TableCell>
                <TableCell align="left">
                  <EditIcon
                    style={{
                      fontSize: 14,
                      marginRight: 10,
                      cursor: "pointer"
                    }}
                    onClick={this.clickEditField.bind(
                      this,
                      row,
                      1,
                      "标题",
                      "Name"
                    )}
                    color="primary"
                  />
                  {row.Name}
                </TableCell>
                <TableCell align="center">
                  <RepositoryStatus status={row.Status} />
                </TableCell>
                <TableCell align="center">
                  <RepositoryJobStatus status={row.JobStatus} />
                </TableCell>
                <TableCell align="center">
                  {row.UserName && row.Password ? (
                    <Tooltip
                      title={<div>用户名：{row.UserName}</div>}
                      interactive
                    >
                      <IconButton color="primary">
                        <LockIcon />
                      </IconButton>
                    </Tooltip>
                  ) : (
                    <Tooltip title="仓库非私密" interactive>
                      <IconButton color="primary">
                        <LockOpenIcon />
                      </IconButton>
                    </Tooltip>
                  )}
                </TableCell>
                <TableCell align="center">
                  <Tooltip
                    title={row.TerminalInfo}
                    classes={{ tooltip: classes.tooltip }}
                    interactive
                  >
                    <IconButton color="primary">
                      <Computer />
                    </IconButton>
                  </Tooltip>
                </TableCell>
                <TableCell align="center">
                  <Tooltip
                    title={row.Desc}
                    classes={{ tooltip: classes.tooltip }}
                    interactive
                  >
                    <IconButton
                      color="primary"
                      onClick={this.clickEditField.bind(
                        this,
                        row,
                        8,
                        "描述",
                        "Desc"
                      )}
                    >
                      <Announcement />
                    </IconButton>
                  </Tooltip>
                </TableCell>
                <TableCell align="center">
                  <span className="tag">{row.DependTools.toUpperCase()}</span>
                </TableCell>
                <TableCell align="center">
                  <Tooltip title="执行 GIT PULL 命令" interactive>
                    <IconButton
                      color="primary"
                      onClick={this.repositoryGitPull.bind(this, row)}
                    >
                      <SaveAlt />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title="裁剪分支" interactive>
                    <IconButton
                      color="primary"
                      onClick={this.repositoryPruneBranch.bind(this, row)}
                    >
                      <FormatClearIcon />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title="删除依赖" interactive>
                    <IconButton
                      color="primary"
                      onClick={this.deleteRepositoryDepend.bind(this, row)}
                    >
                      <RestoreFromTrash />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title="删除仓库" interactive>
                    <IconButton
                      color="primary"
                      onClick={this.destroyDialogOpen.bind(this, row.ID)}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </Tooltip>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <EditField
          id={this.state.editField.id}
          open={this.state.editField.open}
          desc={this.state.editField.desc}
          rows={this.state.editField.rows}
          field={this.state.editField.field}
          fieldContent={this.state.editField.fieldContent}
          onClose={this.closeResetEditFieldDialog.bind(this)}
          editSuccess={this.editFieldSuccess.bind(this)}
        />
        <Dialog
          open={this.state.destroyDialogShow}
          onClose={this.destroyDialogClose.bind(this)}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <DialogTitle id="alert-dialog-title">{"确认删除仓库?"}</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              你确认要删除这个仓库？没有任务使用此仓库才能删除。
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button
              onClick={this.destroyDialogClose.bind(this)}
              color="primary"
            >
              关闭
            </Button>
            <Button
              onClick={this.destroyConfirm.bind(this)}
              color="secondary"
              autoFocus
            >
              确认
            </Button>
          </DialogActions>
        </Dialog>
      </Fragment>
    );
  }
}

export default withStyles(styles)(RepositoryTable);
