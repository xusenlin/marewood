import React from "react";
import { CircularProgress, IconButton, Tooltip } from "@material-ui/core";
import i18n from "react-intl-universal";
const STATUS = ["repo.repoLeisured", "repo.repoBusy"];

class RepositoryTaskStatus extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  renderCircularProgress(status) {
    if (1 === status) return <CircularProgress size={20} color="secondary" />;
    return <CircularProgress variant="static" value={100} size={20} />;
  }
  render() {
    return (
      (this.props.status === 0 || this.props.status === 1)  ?
        (<div>
          <Tooltip title={i18n.get(STATUS[this.props.status])}>
            <IconButton color="primary">
              {this.renderCircularProgress(this.props.status)}
            </IconButton>
          </Tooltip>
        </div>):""
    );
  }
}

export default RepositoryTaskStatus;
