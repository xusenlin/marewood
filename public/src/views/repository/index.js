import React from "react";
import { withStyles } from "@material-ui/core/styles";
import AddIcon from "@material-ui/icons/Add";
import { Fab, Paper } from "@material-ui/core";
import Edit from "./edit";
import Table from "./repositoryTable";
import { getSystemInfo } from "../../utils/dataStorage";
import { repositoryFind } from "../../api/repository";
import Pagination from "@material-ui/lab/Pagination";
import BusEvent from "../../utils/busEvents";
const styles = theme => ({
  root: {
    width: "100%",
    marginTop: theme.spacing(1),
    overflowX: "auto"
  },
  fab: {
    position: "absolute",
    bottom: theme.spacing(2),
    right: theme.spacing(2)
  },
  pagination: {
    width: "100%",
    display: "flex",
    justifyContent: "center",
    textAlign: "center",
    padding: theme.spacing(3)
  },
  headFun: {
    textAlign: "right"
  }
});

class RepositoryTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tableData: [],
      editShow: false,
      dependentSupport: getSystemInfo("DependTools") || [],
      totalPage: 1,
      pageNum: 1
    };
  }
  componentDidMount() {
    let { query } = this.props.location;
    if (query && query.id) {
      BusEvent.search.setSearchKeyword("id:" + query.id);
    }
    setTimeout(() => {
      this.getTableData();
    }, 150);
    BusEvent.updateCurrentTable = () => {
      this.getTableData();
    };
  }

  componentWillUnmount() {
    BusEvent.updateCurrentTable = () => {};
    BusEvent.search.setSearchKeyword("");
  }
  editDialogShow() {
    this.setState({ editShow: true });
  }

  editDialogClose() {
    this.setState({ editShow: false });
  }

  createSuccess() {
    this.setState({ editShow: false });
    this.getTableData();
  }
  getTableData(pageNum = 1, pageSize = 8) {
    let name = BusEvent.search.getSearchKeyword();
    repositoryFind({ pageNum, pageSize, name })
      .then(r => {
        this.setState({
          tableData: r.List,
          totalPage: r.TotalPage,
          pageNum: r.PageNum
        });
      })
      .catch(() => {});
  }

  changePagination(v, pageNum) {
    if (pageNum === this.state.pageNum) {
      return;
    }
    this.getTableData(pageNum);
  }
  render() {
    const { classes } = this.props;
    return (
      <div>
        <Paper className={classes.root}>
          <Table
            tableData={this.state.tableData}
            refresh={this.getTableData.bind(this)}
          />
          <div className={classes.pagination}>
            <Pagination
              count={this.state.totalPage}
              page={this.state.pageNum}
              onChange={this.changePagination.bind(this)}
              color="primary"
              shape="rounded"
            />
          </div>
        </Paper>
        <Fab
          color="primary"
          className={classes.fab}
          aria-label="add"
          onClick={this.editDialogShow.bind(this)}
        >
          <AddIcon />
        </Fab>
        <Edit
          dependentSupport={this.state.dependentSupport}
          show={this.state.editShow}
          handleClose={this.editDialogClose.bind(this)}
          createSuccess={this.createSuccess.bind(this)}
        />
      </div>
    );
  }
}

export default withStyles(styles)(RepositoryTable);
