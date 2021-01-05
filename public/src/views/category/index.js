import React from "react";
import { withStyles } from "@material-ui/core/styles";
import AddIcon from "@material-ui/icons/Add";
import EditIcon from "@material-ui/icons/Edit";
import DeleteIcon from "@material-ui/icons/Delete";
import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableCell,
  Paper,
  Fab,
  IconButton,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Tooltip
} from "@material-ui/core";
import EditCategory from "./edit.js";
import { categories, destroy, UpdateField } from "../../api/category.js";
import EditField from "../../components/editField.js";
import Snackbar from "../../components/snackbar";
import Announcement from "@material-ui/icons/Announcement";
import BusEvent from "../../utils/busEvents";
import i18n from "react-intl-universal";

const styles = theme => ({
  root: {
    width: "100%",
    marginTop: theme.spacing(1),
    overflowX: "auto"
  },
  table: {
    minWidth: 650
  },
  fab: {
    position: "absolute",
    bottom: theme.spacing(2),
    right: theme.spacing(2)
  }
});

class CategoryTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tableData: [],
      destroyDialogShow: false,
      editCategoryShow: false,
      editField: {
        id: 0,
        open: false,
        rows: 1,
        desc: "",
        field: "",
        fieldContent: ""
      }
    };
    this.destroyId = 0;
  }

  componentDidMount() {
    this.getTableData();
    BusEvent.updateCurrentTable = () => {
      this.getTableData();
    };
  }
  componentWillUnmount() {
    BusEvent.updateCurrentTable = () => {};
    BusEvent.search.setSearchKeyword("");
  }
  getTableData() {
    let name = BusEvent.search.getSearchKeyword();
    categories({ name })
      .then(r => {
        this.setState({ tableData: r });
      })
      .catch(() => {});
  }

  destroyDialogOpen(id) {
    this.destroyId = id;
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
        this.getTableData();
      })
      .catch(() => {});
  }

  editDialogShow() {
    this.setState({ editCategoryShow: true });
  }
  editDialogClose() {
    this.setState({ editCategoryShow: false });
  }
  createCategorySuccess() {
    this.setState({ editCategoryShow: false });
    this.getTableData();
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
        this.getTableData();
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
      <div>
        <Paper className={classes.root}>
          <Table className={classes.table}>
            <TableHead>
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell align="left">{i18n.get("common.name")}</TableCell>
                <TableCell align="left">{i18n.get("common.total")}</TableCell>
                <TableCell align="left">{i18n.get("common.remark")}</TableCell>
                <TableCell align="left">
                  {i18n.get("common.createdAt")}
                </TableCell>
                <TableCell align="left">{i18n.get("common.actions")}</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {this.state.tableData.map(row => (
                <TableRow key={row.ID} hover>
                  <TableCell component="th" scope="row">
                    {row.ID}
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
                  <TableCell align="left">{row.JobQuantity}</TableCell>
                  <TableCell align="left">
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
                  <TableCell align="left">{row.CreatedAt}</TableCell>
                  <TableCell align="left">
                    <Tooltip title={i18n.get("common.delete")}>
                      <IconButton
                        edge="start"
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
        <Fab
          color="primary"
          className={classes.fab}
          aria-label="add"
          onClick={this.editDialogShow.bind(this)}
        >
          <AddIcon />
        </Fab>
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
        <EditCategory
          show={this.state.editCategoryShow}
          handleClose={this.editDialogClose.bind(this)}
          createSuccess={this.createCategorySuccess.bind(this)}
        />
      </div>
    );
  }
}

export default withStyles(styles)(CategoryTable);
