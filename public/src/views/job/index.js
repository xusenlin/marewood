import React from "react";
import { withStyles } from "@material-ui/core/styles";
import { Tab, Tabs, Paper, Fab } from "@material-ui/core";
import Pagination from "@material-ui/lab/Pagination";
import { jobsFind } from "../../api/job";
import { categories } from "../../api/category";
import CategoriesTable from "./categoriesTable";
import Edit from "./edit";
import AddIcon from "@material-ui/icons/Add";

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
  },
  pagination: {
    width: "100%",
    display: "flex",
    justifyContent: "center",
    textAlign: "center",
    padding: theme.spacing(3)
  }
});

class Jobs extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      category: 0,
      categories: [],
      jobs: [],
      editDialog: {
        show: false,
        categoryId: 0,
        categoryName: ""
      },
      totalPage: 1,
      pageNum: 1
    };
    this.timeoutId = null;
  }

  componentDidMount() {
    categories()
      .then(r => {
        if (r.length === 0) {
          alert("请先新建分类");
          return;
        }
        this.setState({ categories: r }, () => {
          this.setTabAndJobsByCategoryId(0);
        });
      })
      .catch(() => {});
    window.wsFunc = () => {
      this.setTabAndJobsByCategoryId(this.state.category, this.state.pageNum);
    };
  }

  changePagination(v, pageNum) {
    if (pageNum === this.state.pageNum) {
      return;
    }
    this.setTabAndJobsByCategoryId(this.state.category, pageNum);
  }

  tabsChange(event, index) {
    this.setTabAndJobsByCategoryId(index);
  }

  setTabAndJobsByCategoryId(index = 0, pageNum = 1, pageSize = 8) {
    let categoryId = this.state.categories[index].ID;
    if (this.timeoutId) {
      clearTimeout(this.timeoutId);
    }
    jobsFind({ categoryId, pageNum, pageSize })
      .then(r => {
        this.setState({
          category: index,
          jobs: r.List,
          totalPage: r.TotalPage,
          pageNum: r.PageNum
        });
        // for (let i = 0; i < r.List.length; i++) {
        //   if (r.List[i].Status === 3) {
        //     this.timeoutId = setTimeout(() => {
        //       this.setTabAndJobsByCategoryId(index, pageNum, pageSize);
        //     }, 2000);
        //     return;
        //   }
        // }
      })
      .catch(() => {});
  }

  editDialogShow() {
    let category = this.state.categories[this.state.category];
    this.setState({
      editDialog: {
        show: true,
        categoryId: category.ID,
        categoryName: category.Name
      }
    });
  }

  createSuccess() {
    this.setTabAndJobsByCategoryId(this.state.category);
    this.setState({
      editDialog: {
        show: false,
        categoryId: 0,
        categoryName: ""
      }
    });
  }

  createClose() {
    this.setState({
      editDialog: {
        show: false,
        categoryId: 0,
        categoryName: ""
      }
    });
  }

  render() {
    const { classes } = this.props;
    return (
      <div>
        <Paper className={classes.root}>
          <Tabs
            value={this.state.category}
            indicatorColor="primary"
            textColor="primary"
            centered
            onChange={this.tabsChange.bind(this)}
          >
            {this.state.categories.map(category => (
              <Tab key={category.ID} label={category.Name} />
            ))}
          </Tabs>
          <CategoriesTable
            tableData={this.state.jobs}
            refresh={this.setTabAndJobsByCategoryId.bind(
              this,
              this.state.category
            )}
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
          show={this.state.editDialog.show}
          categoryId={this.state.editDialog.categoryId}
          categoryName={this.state.editDialog.categoryName}
          handleClose={this.createClose.bind(this)}
          createSuccess={this.createSuccess.bind(this)}
        />
      </div>
    );
  }
}

export default withStyles(styles)(Jobs);
