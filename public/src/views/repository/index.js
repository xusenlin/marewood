import React from "react";
import { withStyles } from "@material-ui/core/styles";
import AddIcon from "@material-ui/icons/Add";
import { Fab, Paper } from "@material-ui/core";
import Edit from "./edit";
import Table from "./repositoryTable";
import { getSystemInfo } from "../../utils/dataStorage";
import { repositoryFind } from "../../api/repository";
import Pagination from "@material-ui/lab/Pagination";
import SearchBox from "../../components/searchBox";


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
  },
  search: {
    width: "400px",
    height: "64px",
    position: "absolute",
    top: 0,
    left: "50%",
    transform: 'translateX(-50%)',
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
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
      pageNum: 1,
      name: ""
    };
  }
  componentDidMount() {
    this.getTableData();
    window.wsUpdateDataFunc = () => {
      this.getTableData(this.state.pageNum);
    };
  }

  componentWillUnmount() {
    window.wsUpdateDataFunc = null;
    // if (this.timeout) clearTimeout(this.timeout);
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
    // if (this.timeout) clearTimeout(this.timeout);
    let name = this.state.name;
    repositoryFind({ pageNum, pageSize, name })
      .then(r => {
        this.setState({
          tableData: r.List,
          totalPage: r.TotalPage,
          pageNum: r.PageNum
        });

        // for (let i = 0; i < r.length; i++) {
        //   if (
        //     r[i].Status === 0 ||
        //     (r[i].JobStatus === 1 && r[i].Status === 1)
        //   ) {
        //     //仓库正在克隆当中  或者 （一个正常的仓库很繁忙）的情况就会刷新
        //     this.timeout = setTimeout(() => {
        //       this.getTableData();
        //     }, 5000);
        //     return;
        //   }
        // }
      })
      .catch(() => {});
  }

  changePagination(v, pageNum) {
    if (pageNum === this.state.pageNum) {
      return;
    }
    this.getTableData(pageNum);
  }

  refreshSearch(val) {
    this.setState({name: val}, () => {
      this.getTableData()
    })
  }

  render() {
    const { classes } = this.props;
    return (
      <div>
        <div className={classes.search}>
          <SearchBox refreshSearch={this.refreshSearch.bind(this)} />
        </div>
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
