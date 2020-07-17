import React from "react";
import PropTypes from "prop-types";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import DialogTitle from "@material-ui/core/DialogTitle";
import Dialog from "@material-ui/core/Dialog";
import { gitBranch } from "../../api/repository";
import { updateBranch } from "../../api/job";
import { DialogContentText } from "@material-ui/core";

class SwitchBranchDialog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      branchList: []
    };
  }

  componentDidMount() {}
  onEnteredDialog() {
    this.setState({ branchList: [] });
    gitBranch({ id: this.props.repositoryId })
      .then(r => {
        this.setState({ branchList: r });
      })
      .catch(() => {});
  }
  selectBranch(branch) {
    updateBranch({ branch: branch, id: this.props.jobId })
      .then(() => {
        this.props.switchSuccess();
      })
      .catch(() => {});
  }

  render() {
    return (
      <Dialog
        open={this.props.open}
        onEntering={this.onEnteredDialog.bind(this)}
        onClose={this.props.onClose}
        aria-labelledby="dialog-title"
      >
        <DialogTitle id="dialog-title">切换当前任务的分支</DialogTitle>
        <DialogContentText>
          {this.state.branchList.length === 0 ? "分支无法使用" : ""}
        </DialogContentText>
        <List>
          {this.state.branchList.map(branch => (
            <ListItem
              button
              onClick={this.selectBranch.bind(this, branch)}
              key={branch}
            >
              <ListItemText primary={branch} />
            </ListItem>
          ))}
        </List>
      </Dialog>
    );
  }
}

SwitchBranchDialog.propTypes = {
  onClose: PropTypes.func,
  open: PropTypes.bool,
  jobId: PropTypes.number,
  repositoryId: PropTypes.number,
  switchSuccess: PropTypes.func
};

export default SwitchBranchDialog;
