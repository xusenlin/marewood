import React from "react";
import i18n from "react-intl-universal";
import { CircularProgress, IconButton, Tooltip } from "@material-ui/core";

const STATUS = [
  "task.statusLeisured",
  "task.statusSuccess",
  "task.statusFail",
  "task.statusProcessing"
];

class RepositoryStatus extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  renderCircularProgress(status) {
    if (0 === status)
      return (
        <CircularProgress
          variant="static"
          style={{ color: "#9E9E9E" }}
          value={100}
          size={20}
        />
      );
    if (1 === status)
      return <CircularProgress variant="static" value={100} size={20} />;
    if (2 === status)
      return (
        <CircularProgress
          variant="static"
          value={100}
          size={20}
          color="secondary"
        />
      );
    return <CircularProgress size={20} />;
  }
  render() {
    return (
      <div>
        <Tooltip title={i18n.get(STATUS[this.props.status])}>
          <IconButton color="primary">
            {this.renderCircularProgress(this.props.status)}
          </IconButton>
        </Tooltip>
      </div>
    );
  }
}

export default RepositoryStatus;
