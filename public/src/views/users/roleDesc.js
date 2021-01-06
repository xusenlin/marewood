import React from "react";
import i18n from "react-intl-universal";
const RLOECOLOR = ["#f44336", "#f44336", "#ff9800", "#3f51b5", "#4caf50"];
class RoleDesc extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div>
        <span
          className="tag"
          style={{ backgroundColor: RLOECOLOR[this.props.role] }}
        >
          {i18n.get("role." + this.props.role)}
        </span>
      </div>
    );
  }
}

export default RoleDesc;
