import React from "react";
import { withStyles } from "@material-ui/core/styles";
import { Tab, Tabs, Paper } from "@material-ui/core";
import Pagination from "@material-ui/lab/Pagination";
import { taskFind } from "../../api/task";
import { categories } from "../../api/category";
import CategoriesTable from "./categoriesTable";
import AddTask from "./addTask/index";
import BusEvent from "../../utils/busEvents";
import i18n from "react-intl-universal";

const defPageSize = 8;

const styles = theme => ({
  root: {
    width: "100%",
    marginTop: theme.spacing(1),
    overflowX: "auto"
  },
  pagination: {
    width: "100%",
    display: "flex",
    justifyContent: "center",
    textAlign: "center",
    padding: theme.spacing(3)
  }
});

class Task extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tabIndex: 0, //tab索引
      categories: [], //tab的分类数据
      tasks: [], //表格数据
      totalPage: 1, //分页总数
      pageNum: 1 //当前页码
    };
  }
  componentDidMount() {
    categories()
      .then(r => {
        if (r.length === 0) {
          alert(i18n.get("category.createTip"));
          return;
        }
        this.setState({ categories: r }, () => {
          let { ID } = this.state.categories[this.state.tabIndex];
          this.getTask(ID, 1);
        });
        BusEvent.updateCurrentTable = () => {
          let { ID } = this.state.categories[this.state.tabIndex];
          this.getTask(ID, 1);
        };
      })
      .catch(() => {});
  }
  componentWillUnmount() {
    BusEvent.updateCurrentTable = () => {};
    BusEvent.search.setSearchKeyword("");
  }
  changePagination(v, pageNum) {
    if (pageNum === this.state.pageNum) {
      return;
    }
    let { ID } = this.state.categories[this.state.tabIndex];
    this.getTask(ID, pageNum, () => {});
  }
  tabsChange(event, index) {
    let { ID } = this.state.categories[index];
    this.getTask(ID, 1, () => {
      this.setState({ tabIndex: index });
    });
  }
  getTask(
    categoryId,
    pageNum = this.state.pageNum,
    cb = () => {},
    name = BusEvent.search.getSearchKeyword(),
    pageSize = defPageSize
  ) {
    taskFind({ categoryId, pageNum, pageSize, name })
      .then(r => {
        this.setState(
          {
            tasks: r.List,
            totalPage: r.TotalPage,
            pageNum: r.PageNum
          },
          () => {
            cb();
          }
        );
      })
      .catch(() => {});
  }
  render() {
    const { classes } = this.props;
    return (
      <div>
        <Paper className={classes.root}>
          <Tabs
            value={this.state.tabIndex}
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
            tableData={this.state.tasks}
            refresh={() => {
              let { ID } = this.state.categories[this.state.tabIndex];
              this.getTask(ID, 1);
            }}
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
        <AddTask
          category={this.state.categories[this.state.tabIndex]}
          createSuccess={() => {
            let { ID } = this.state.categories[this.state.tabIndex];
            this.getTask(ID, 1);
          }}
        />
      </div>
    );
  }
}

export default withStyles(styles)(Task);
