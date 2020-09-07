import React from "react";
import { withStyles } from "@material-ui/core/styles";
import AddIcon from "@material-ui/icons/Add";
import { Fab, Paper } from "@material-ui/core";
import Edit from "./edit";
import Table from "./repositoryTable";
import Transfer from "./transfer.js";
import { getSystemInfo } from "../../utils/dataStorage";
import { repositoryFind } from "../../api/repository";
import Pagination from "@material-ui/lab/Pagination";
// import TransferWithinAStationIcon from '@material-ui/icons/TransferWithinAStation';

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
      pageNum: 1,
      transferDis: false
    };
    // this.timeout = null;
  }
  componentDidMount() {
    this.getTableData();
    window.wsFuncRepo = () => {
      this.getTableData(this.state.pageNum);
    };
  }

  componentWillUnmount() {
    window.wsFuncRepo = null;
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
    repositoryFind({ pageNum, pageSize })
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
  showTransfer(){
    this.setState({ transferDis: true });
  }
  closeDia(){
    this.setState({ transferDis: false });
  }
  render() {
    const { classes } = this.props;
    return (
      <div>
        <header className={classes.headFun}>
          {/* <Tooltip title="仓库迁移">
            <IconButton
              color="primary"
              onClick={this.showTransfer.bind(this)}
            >
              <TransferWithinAStationIcon />
            </IconButton>
          </Tooltip> */}
          {this.state.transferDis && <Transfer closeDia={this.closeDia.bind(this)} transfer={this.state.transferDis} /> }
        </header>
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
