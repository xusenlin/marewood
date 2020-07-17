import React from "react";
import { CircularProgress, IconButton, Tooltip } from "@material-ui/core";

const STATUS = ["依赖处理中", "依赖处理完成", "依赖处理失败"];

class DependStatus extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  renderCircularProgress(status) {
    if (0 === status) return <CircularProgress size={20} />;
    if (1 === status)
      return <CircularProgress variant="static" value={100} size={20} />;
    return (
      <CircularProgress
        variant="static"
        value={100}
        size={20}
        color="secondary"
      />
    );
  }
  render() {
    return (
      <div>
        <Tooltip title={STATUS[this.props.status]}>
          <IconButton color="primary">
            {this.renderCircularProgress(this.props.status)}
          </IconButton>
        </Tooltip>
      </div>
    );
  }
}

export default DependStatus;
