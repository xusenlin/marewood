import React from 'react';
import Role from "../../config/role"
const RLOECOLOR = ["#f44336","#f44336","#ff9800","#3f51b5","#4caf50"];
class RoleDesc extends React.Component {
    constructor(props) {
        super(props);
        this.state ={}
    }
    render() {
        return (
            <div>
                <span className="tag"
                      style={{backgroundColor:RLOECOLOR[this.props.role]}}>
                    {Role[this.props.role]}
                </span>
            </div>
        );
    }
}

export default RoleDesc
