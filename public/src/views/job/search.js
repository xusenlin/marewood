import React from "react"
import { withStyles } from "@material-ui/core/styles"
import { Paper, InputBase } from "@material-ui/core"
import IconButton from "@material-ui/core/IconButton"
import SearchIcon from "@material-ui/icons/Search"
import Divider from '@material-ui/core/Divider';
import RefreshIcon from '@material-ui/icons/Refresh';

const styles = (theme) => ({
  paper: {
    display: "flex",
    alignItems: "center",
    justenAlign: "center",
    width: 500,
  },
  input: {
    marginLeft: theme.spacing(1),
    flex: 1,
  },
  iconButton: {
    padding: 10,
  },
  divider: {
    height: 28,
    margin: 4,
  },
})

class Search extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      keyword: ""
    }
  }

  inputBaseChange(event) {
    this.setState({
      keyword: event.target.value
    })
  }

  clickSearch() {
    this.props.refreshSearch(this.state.keyword);
  }

  clickReset() {
    this.setState({
      keyword: ""
    })
    this.props.refreshSearch("");
  }

  render() {
    const { classes } = this.props
    let keyword = this.state.keyword
    return (
      <div>
        <Paper component="form" className={classes.paper}>
          <InputBase
            className={classes.input}
            placeholder="请输入关键字"
            value={keyword}
            onChange={this.inputBaseChange.bind(this)}
            inputProps={{ "aria-label": "search google maps" }}
          />
          <IconButton type="submit" className={classes.iconButton} onClick={this.clickSearch.bind(this)} aria-label="search">
            <SearchIcon />
          </IconButton>
          <Divider className={classes.divider} orientation="vertical" />
          <IconButton color="primary" className={classes.iconButton} onClick={this.clickReset.bind(this)} aria-label="directions">
            <RefreshIcon />
          </IconButton>
        </Paper>
      </div>
    )
  }
}

export default withStyles(styles)(Search)
