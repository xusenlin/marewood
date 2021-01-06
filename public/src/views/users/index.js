import React from "react";
import { withStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import { users, deleteUsers, roleEdit } from "../../api/user";
import { IconButton, Tooltip } from "@material-ui/core";
import i18n from "react-intl-universal";
import DeleteIcon from "@material-ui/icons/Delete";
import ThumbDown from "@material-ui/icons/ThumbDown";
import ThumbUp from "@material-ui/icons/ThumbUp";
import Snackbar from "../../components/snackbar";
import RoleDesc from "./roleDesc";
import BusEvent from "../../utils/busEvents";

const styles = theme => ({
  root: {
    width: "100%",
    marginTop: theme.spacing(1),
    overflowX: "auto"
  },
  table: {
    minWidth: 650
  },
  sept: {
    marginTop: 20,
    padding: theme.spacing(3, 2)
  }
});

class User extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tableData: []
    };
  }
  componentDidMount() {
    this.getUserList();
    BusEvent.updateCurrentTable = () => {
      this.getUserList();
    };
  }
  componentWillUnmount() {
    BusEvent.updateCurrentTable = () => {};
    BusEvent.search.setSearchKeyword("");
  }
  getUserList() {
    let name = BusEvent.search.getSearchKeyword();
    users({ name })
      .then(r => {
        this.setState({ tableData: r });
      })
      .catch(() => {});
  }
  destroyUser(row) {
    deleteUsers({ id: row.ID })
      .then(() => {
        Snackbar.success(i18n.get("common.deleteSuccess"));
        this.getUserList();
      })
      .catch(() => {});
  }
  role(row, isUp = 0) {
    roleEdit({ id: row.ID, isUp: isUp })
      .then(() => {
        Snackbar.success(i18n.get("common.modifySuccess"));
        this.getUserList();
      })
      .catch(() => {});
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
                <TableCell align="left">{i18n.get("common.role")}</TableCell>
                <TableCell align="left">{i18n.get("common.remark")}</TableCell>
                <TableCell align="center">
                  {i18n.get("common.actions")}
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {this.state.tableData.map(row => (
                <TableRow key={row.ID} hover>
                  <TableCell component="th" scope="row">
                    {row.ID}
                  </TableCell>
                  <TableCell align="left">{row.Username}</TableCell>
                  <TableCell align="left">
                    <RoleDesc role={row.Role} />
                  </TableCell>
                  <TableCell align="left">{row.Desc}</TableCell>
                  <TableCell align="center">
                    <Tooltip title={i18n.get("users.roleUpgrade")}>
                      <IconButton
                        color="primary"
                        onClick={this.role.bind(this, row, 1)}
                      >
                        <ThumbUp />
                      </IconButton>
                    </Tooltip>
                    <Tooltip title={i18n.get("users.roleDowngrade")}>
                      <IconButton
                        color="primary"
                        onClick={this.role.bind(this, row, 0)}
                      >
                        <ThumbDown />
                      </IconButton>
                    </Tooltip>
                    <Tooltip title={i18n.get("common.delete")}>
                      <IconButton
                        color="primary"
                        onClick={this.destroyUser.bind(this, row)}
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
      </div>
    );
  }
}

export default withStyles(styles)(User);
