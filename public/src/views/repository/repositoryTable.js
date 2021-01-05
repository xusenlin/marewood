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
import RotateLeftIcon from "@material-ui/icons/RotateLeft";
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
import RepositoryTaskStatus from "./repositoryTaskStatus";
import Snackbar from "../../components/snackbar";
import {
  deleteDepend,
  destroy,
  gitPull,
  UpdateField,
  pruneBranch,
  discardChange
} from "../../api/repository";
import EditField from "../../components/editField";
import { tooltip } from "../../assets/jss/common";
import NewReleases from "@material-ui/icons/NewReleases";
import { reset } from "../../api/repository.js";
import i18n from "react-intl-universal";

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
      Snackbar.warning(i18n.get("repo.deleteDependentFailedTip"));
      return;
    }
    deleteDepend({ id: row.ID })
      .then(r => {
        Snackbar.success(r);
      })
      .catch(() => {});
  }
  discardRepoChange(row) {
    discardChange({ id: row.ID })
      .then(() => {
        Snackbar.success("git checkout . success！");
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
        Snackbar.success("git pull success");
      })
      .catch(() => {});
  }

  repositoryPruneBranch(row) {
    pruneBranch({ id: row.ID })
      .then(() => {
        Snackbar.success("git remote prune origin success");
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
  resetRepo(row) {
    if (row.JobStatus !== 0) {
      return Snackbar.warning(i18n.get("repo.resetRepoFailed"));
    }
    reset({ id: row.ID })
      .then(() => {
        Snackbar.success(i18n.get("repo.resetRepoSuccess"));
        this.props.refresh();
      })
      .catch(() => {});
  }
  render() {
    const { classes } = this.props;
    return (
      <Fragment>
        <Table className={classes.table}>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell align="center">
                {i18n.get("common.remoteUrl")}
              </TableCell>
              <TableCell align="center">{i18n.get("common.name")}</TableCell>
              <TableCell align="center">
                {i18n.get("repo.cloneStatus")}
              </TableCell>
              <TableCell align="center">
                {i18n.get("repo.workingStatus")}
                {/*<HelperTooltips help="当前工作目录正在执行其他部署任务，资源被占用" />*/}
              </TableCell>
              <TableCell align="center">
                {i18n.get("repo.auth")}
                {/*<HelperTooltips help="私密仓库是需要用户名和密码的，否则克隆失败" />*/}
              </TableCell>
              <TableCell align="center">
                {i18n.get("common.terminalOut")}
              </TableCell>
              <TableCell align="center">
                {i18n.get("common.remark")}
                <HelperTooltips help={i18n.get("common.clickEdit")} />
              </TableCell>
              <TableCell align="center">
                {i18n.get("common.dependentTools")}
              </TableCell>
              <TableCell align="center">{i18n.get("common.actions")}</TableCell>
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
                      i18n.get("common.title"),
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
                  <RepositoryTaskStatus status={row.JobStatus} />
                </TableCell>
                <TableCell align="center">
                  {row.UserName && row.Password ? (
                    <Tooltip
                      title={
                        <div>
                          {i18n.get("common.username")}：{row.UserName}
                        </div>
                      }
                      interactive
                    >
                      <IconButton color="primary">
                        <LockIcon />
                      </IconButton>
                    </Tooltip>
                  ) : (
                    <IconButton color="primary">
                      <LockOpenIcon />
                    </IconButton>
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
                        i18n.get("common.desc"),
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
                  <Tooltip title="git pull" interactive>
                    <IconButton
                      color="primary"
                      onClick={this.repositoryGitPull.bind(this, row)}
                    >
                      <SaveAlt />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title="git remote prune origin" interactive>
                    <IconButton
                      color="primary"
                      onClick={this.repositoryPruneBranch.bind(this, row)}
                    >
                      <FormatClearIcon />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title="git checkout ." interactive>
                    <IconButton
                      color="primary"
                      onClick={this.discardRepoChange.bind(this, row)}
                    >
                      <RotateLeftIcon />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title={i18n.get("repo.deleteDependent")} interactive>
                    <IconButton
                      color="primary"
                      onClick={this.deleteRepositoryDepend.bind(this, row)}
                    >
                      <RestoreFromTrash />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title={i18n.get("repo.resetRepo")}>
                    <IconButton
                      color="primary"
                      onClick={this.resetRepo.bind(this, row)}
                    >
                      <NewReleases />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title={i18n.get("common.delete")} interactive>
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
          <DialogTitle id="alert-dialog-title">
            {i18n.get("repo.deleteTitle")}
          </DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              {i18n.get("repo.deleteTip")}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button
              onClick={this.destroyDialogClose.bind(this)}
              color="primary"
            >
              {i18n.get("common.close")}
            </Button>
            <Button
              onClick={this.destroyConfirm.bind(this)}
              color="secondary"
              autoFocus
            >
              {i18n.get("common.confirm")}
            </Button>
          </DialogActions>
        </Dialog>
      </Fragment>
    );
  }
}

export default withStyles(styles)(RepositoryTable);
