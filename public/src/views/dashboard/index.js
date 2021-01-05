import React from "react";
import { withStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import i18n from "react-intl-universal";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import { systemInfo } from "../../api/system";
import { setSystemInfo } from "../../utils/dataStorage";
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

class Dashboard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tableData: []
    };
  }
  componentDidMount() {
    systemInfo()
      .then(r => {
        setSystemInfo(r.kvConfig);
        this.setState({ tableData: r.detailsConfig });
      })
      .catch(() => {});
  }
  componentWillUnmount() {
    BusEvent.search.setSearchKeyword("");
  }
  render() {
    const { classes } = this.props;
    return (
      <div>
        <Paper className={classes.root}>
          <Table className={classes.table}>
            <TableHead>
              <TableRow>
                <TableCell>{i18n.get("common.title")}</TableCell>
                <TableCell>{i18n.get("common.value")}</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {this.state.tableData.map(row => (
                <TableRow key={row.Title}>
                  <TableCell component="th" scope="row">
                    {row.Title}
                  </TableCell>
                  <TableCell>
                    {Array.isArray(row.Val) ? row.Val.join(" | ") : row.Val}
                  </TableCell>
                </TableRow>
              ))}
              <TableRow>
                <TableCell component="th" scope="row">
                  About
                </TableCell>
                <TableCell>{i18n.get("dashboard.about")}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </Paper>
      </div>
    );
  }
}

export default withStyles(styles)(Dashboard);
