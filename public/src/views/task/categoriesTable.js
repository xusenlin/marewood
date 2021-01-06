import React from "react";
import { withStyles } from "@material-ui/core/styles";
import LinkIcon from "@material-ui/icons/Link";
import EditIcon from "@material-ui/icons/Edit";
import LinkOffIcon from "@material-ui/icons/LinkOff";
import UsbIcon from "@material-ui/icons/Usb";
import Computer from "@material-ui/icons/Computer";
import Announcement from "@material-ui/icons/Announcement";
import PlayCircleFilled from "@material-ui/icons/PlayCircleFilled";
import GpsFixedIcon from "@material-ui/icons/GpsFixed";
import DeviceHub from "@material-ui/icons/DeviceHub";
import {
  TextField,
  TableRow,
  TableHead,
  TableCell,
  TableBody,
  Table,
  Tooltip,
  IconButton,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
  Dialog
} from "@material-ui/core";
import { withRouter } from "react-router-dom";
import TaskStatus from "./taskStatus";
import HelperTooltips from "../../components/helperTooltips";
import EditField from "../../components/editField.js";
import DetailsPanel from "../../components/detailsPanel.js";
import SwitchBranchDialog from "./switchBranchDialog";
import Snackbar from "../../components/snackbar/index";
import LockIcon from "@material-ui/icons/Lock";
import LockOpenIcon from "@material-ui/icons/LockOpen";
import DeleteIcon from "@material-ui/icons/Delete";
import { destroy, RunTask, UpdateField, TaskLock } from "../../api/task";
import { tooltip } from "../../assets/jss/common";
import { useApiUrl } from "../../config/url";
import { getSystemInfo } from "../../utils/dataStorage";
import i18n from "react-intl-universal";

const styles = theme => ({
  table: {
    marginTop: theme.spacing(3),
    minWidth: 650
  },
  tooltip
});

class CategoriesTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      destroyDialogShow: false,
      switchBranchDialog: {
        id: 0,
        show: false,
        repositoryId: 0
      },
      lockPassword: {
        id: 0,
        password: "",
        show: false
      },
      editField: {
        id: 0,
        open: false,
        desc: "",
        rows: 1,
        field: "",
        fieldContent: ""
      },
      detailsPanel: {
        title: "",
        open: false,
        content: ""
      }
    };
    this.destroyId = 0;
    this.addressUrl = getSystemInfo("OtherAddressUrl") || [];
  }

  toggleLock(row) {
    this.setState({
      lockPassword: {
        id: row.ID,
        password: "",
        show: true
      }
    });
  }

  runTask(row) {
    if (row.Status === 3) {
      Snackbar.warning(i18n.get("task.waitBuild"));
      return;
    }
    if (row.LockPassword !== "") {
      Snackbar.warning(i18n.get("task.unlockTip"));
      return;
    }

    RunTask({ id: row.ID })
      .then(() => {
        Snackbar.success(i18n.get("task.buildSuccessTip"));
        this.props.refresh();
      })
      .catch(() => {});
  }

  openSwitchBranchDialog(row) {
    if (row.LockPassword !== "") {
      Snackbar.warning(i18n.get("task.unlockTip"));
      return;
    }
    if (row.Status === 3) {
      Snackbar.warning(i18n.get("task.waitBuild"));
      return;
    }
    this.setState({
      switchBranchDialog: {
        id: row.ID,
        show: true,
        repositoryId: row.RepositoryId
      }
    });
  }

  switchSuccess() {
    this.setState({
      switchBranchDialog: {
        id: 0,
        show: false,
        repositoryId: 0
      }
    });
    this.props.refresh();
  }

  closeSwitchBranchDialog() {
    this.setState({
      switchBranchDialog: {
        id: 0,
        show: false,
        repositoryId: 0
      }
    });
  }

  destroyDialogOpen(row) {
    if (row.LockPassword !== "") {
      Snackbar.warning(i18n.get("task.unlockTip"));
      return;
    }
    this.destroyId = row.ID;
    this.setState({ destroyDialogShow: true });
  }

  destroyDialogClose() {
    this.setState({ destroyDialogShow: false });
  }

  destroyConfirm() {
    destroy({ id: this.destroyId })
      .then(r => {
        Snackbar.success(i18n.get("common.deleteSuccess"));
        this.setState({ destroyDialogShow: false });
        this.props.refresh();
      })
      .catch(() => {});
  }

  openTaskUrl(row, url) {
    if (row.Status !== 1) {
      Snackbar.error(i18n.get("task.waitBuildSuccess"));
      return;
    }
    window.open(url);
  }

  closeLockTaskDialog() {
    this.setState({
      lockPassword: {
        id: 0,
        password: "",
        show: false
      }
    });
  }

  lockTaskDialogConfirm() {
    let p = {
      id: this.state.lockPassword.id,
      password: this.state.lockPassword.password
    };
    TaskLock(p)
      .then(() => {
        Snackbar.success(i18n.get("common.success"));
        this.props.refresh();
        this.setState({
          lockPassword: {
            id: 0,
            password: "",
            show: false
          }
        });
      })
      .catch(() => {});
  }

  passwordFieldChange(event) {
    let p = this.state.lockPassword;
    p.password = event.target.value;
    this.setState({ lockPassword: p });
  }

  closeCommitRecord() {
    this.setState({
      detailsPanel: { ...this.state.detailsPanel, open: false }
    });
  }
  positioningRepo(row) {
    // row.RepositoryId
    this.props.history.push({
      pathname: "/repository",
      query: { id: row.RepositoryId }
    });
    // commitRecord({ id: row.RepositoryId })
    //   .then(r => {
    //     this.setState({
    //       detailsPanel: {
    //         open: true,
    //         title: "最近4次的代码提交详情",
    //         content: r
    //       }
    //     });
    //   })
    //   .catch(() => {});
  }
  clickEditField(row, inputRows, desc, field) {
    if (row.LockPassword !== "") {
      Snackbar.warning(i18n.get("task.unlockTip"));
      return;
    }
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
        this.props.refresh();
      })
      .catch(() => {});
  }

  closeEditFieldDialog() {
    this.setState({
      editField: { ...this.state.editField, open: false }
    });
  }

  render() {
    const { classes } = this.props;
    return (
      <div>
        <Table className={classes.table}>
          <TableHead>
            <TableRow>
              <TableCell align="center">ID</TableCell>
              <TableCell align="center">
                {i18n.get("common.lock")}
                <HelperTooltips help={i18n.get("task.lockTip")} />
              </TableCell>
              <TableCell align="center">{i18n.get("common.name")}</TableCell>
              <TableCell align="center">{i18n.get("common.status")}</TableCell>
              <TableCell align="center">{i18n.get("common.branch")}</TableCell>
              <TableCell align="center">
                {i18n.get("task.runQuantity")}
              </TableCell>
              <TableCell align="center">{i18n.get("common.address")}</TableCell>
              <TableCell align="center">WebHook</TableCell>
              <TableCell align="center">
                {i18n.get("common.terminalOut")}
              </TableCell>
              <TableCell align="center">
                {i18n.get("common.remark")}
                <HelperTooltips help={i18n.get("common.clickEdit")} />
              </TableCell>
              <TableCell align="center">{i18n.get("task.buildDir")}</TableCell>
              <TableCell align="center">{i18n.get("task.npmScript")}</TableCell>
              <TableCell align="center">{i18n.get("common.actions")}</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {this.props.tableData.map(row => (
              <TableRow key={row.ID} hover>
                <TableCell align="center">{row.ID}</TableCell>
                <TableCell align="center">
                  {row.LockPassword ? (
                    <Tooltip
                      title={i18n.get("common.actionsUser") + "：" + row.User}
                      interactive
                    >
                      <IconButton
                        color="primary"
                        onClick={this.toggleLock.bind(this, row)}
                      >
                        <LockIcon />
                      </IconButton>
                    </Tooltip>
                  ) : (
                    <Tooltip title={i18n.get("task.notLock")} interactive>
                      <IconButton
                        color="primary"
                        onClick={this.toggleLock.bind(this, row)}
                      >
                        <LockOpenIcon />
                      </IconButton>
                    </Tooltip>
                  )}
                </TableCell>
                <TableCell align="left" style={{ fontSize: "12px" }}>
                  <EditIcon
                    style={{ fontSize: 14, marginRight: 10, cursor: "pointer" }}
                    onClick={this.clickEditField.bind(
                      this,
                      row,
                      1,
                      i18n.get("common.name"),
                      "Name"
                    )}
                    color="primary"
                  />
                  {row.Name}
                </TableCell>
                <TableCell align="center">
                  <TaskStatus status={row.Status} />
                </TableCell>
                <TableCell align="center">{row.Branch}</TableCell>
                <TableCell align="center">{row.RunQuantity}</TableCell>
                <TableCell align="center">
                  <Tooltip
                    style={{ padding: 10 }}
                    title={
                      row.Status !== 1 ? (
                        i18n.get("task.accessTip")
                      ) : (
                        <React.Fragment>
                          <a
                            target="_blank"
                            rel="noopener noreferrer"
                            style={{ color: "#fff" }}
                            href={useApiUrl + row.Url}
                          >
                            {useApiUrl + row.Url}
                          </a>
                          {this.addressUrl.map(url => {
                            return (
                              <React.Fragment key={url}>
                                <br />
                                <a
                                  target="_blank"
                                  rel="noopener noreferrer"
                                  style={{ color: "#fff" }}
                                  href={url + "/" + row.ID}
                                >
                                  {url + "/" + row.ID}
                                </a>
                              </React.Fragment>
                            );
                          })}
                        </React.Fragment>
                      )
                    }
                    interactive
                  >
                    <IconButton
                      color="primary"
                      onClick={this.openTaskUrl.bind(
                        this,
                        row,
                        useApiUrl + row.Url
                      )}
                    >
                      {row.Status !== 1 ? <LinkOffIcon /> : <LinkIcon />}
                    </IconButton>
                  </Tooltip>
                </TableCell>
                <TableCell align="center">
                  <Tooltip title={useApiUrl + row.WebHookUrl} interactive>
                    <IconButton color="primary">
                      <UsbIcon />
                    </IconButton>
                  </Tooltip>
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
                  <span className="tag">{row.BuildDir}</span>
                </TableCell>
                <TableCell align="center">
                  <span className="tag">{row.BuildCommand}</span>
                </TableCell>
                <TableCell align="center">
                  <Tooltip title={i18n.get("task.switchBranch")} interactive>
                    <IconButton
                      color="primary"
                      onClick={this.openSwitchBranchDialog.bind(this, row)}
                    >
                      <DeviceHub />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title={i18n.get("task.runTask")} interactive>
                    <IconButton
                      color="primary"
                      onClick={this.runTask.bind(this, row)}
                    >
                      <PlayCircleFilled />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title={i18n.get("task.positioning")} interactive>
                    <IconButton
                      color="primary"
                      onClick={this.positioningRepo.bind(this, row)}
                    >
                      <GpsFixedIcon />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title={i18n.get("common.delete")} interactive>
                    <IconButton
                      color="primary"
                      onClick={this.destroyDialogOpen.bind(this, row)}
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
          onClose={this.closeEditFieldDialog.bind(this)}
          editSuccess={this.editFieldSuccess.bind(this)}
        />
        <DetailsPanel
          open={this.state.detailsPanel.open}
          content={this.state.detailsPanel.content}
          onClose={this.closeCommitRecord.bind(this)}
          title={this.state.detailsPanel.title}
        />
        <SwitchBranchDialog
          taskId={this.state.switchBranchDialog.id}
          repositoryId={this.state.switchBranchDialog.repositoryId}
          open={this.state.switchBranchDialog.show}
          onClose={this.closeSwitchBranchDialog.bind(this)}
          switchSuccess={this.switchSuccess.bind(this)}
        />
        <Dialog
          open={this.state.destroyDialogShow}
          onClose={this.destroyDialogClose.bind(this)}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <DialogTitle id="alert-dialog-title">
            {i18n.get("task.deleteTitle")}
          </DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              {i18n.get("task.deleteTip")}
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
        {/*任务加锁密码*/}
        <Dialog
          open={this.state.lockPassword.show}
          onClose={this.closeLockTaskDialog.bind(this)}
        >
          <DialogTitle id="form-dialog-title">
            {i18n.get("task.lockTitle")}
          </DialogTitle>
          <DialogContent>
            <DialogContentText>
              {i18n.get("task.lockInputTip")}
            </DialogContentText>
            <TextField
              onChange={this.passwordFieldChange.bind(this)}
              autoFocus
              margin="dense"
              id="password"
              label="Lock Password"
              type="text"
              fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button
              onClick={this.closeLockTaskDialog.bind(this)}
              color="primary"
            >
              {i18n.get("common.close")}
            </Button>
            <Button
              onClick={this.lockTaskDialogConfirm.bind(this)}
              color="primary"
            >
              {i18n.get("common.confirm")}
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default withStyles(styles)(withRouter(CategoriesTable));
