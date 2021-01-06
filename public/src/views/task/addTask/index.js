import React from "react";
import { Fab } from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import Edit from "./edit";
import AddIcon from "@material-ui/icons/Add";
import PropTypes from "prop-types";

const styles = theme => ({
  fab: {
    position: "absolute",
    bottom: theme.spacing(2),
    right: theme.spacing(2)
  }
});

class AddTask extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      editDialogShow: false
    };
  }
  success() {
    this.setState({ editDialogShow: false });
    this.props.createSuccess();
  }
  render() {
    const { classes } = this.props;
    return (
      <div>
        {this.props.category ? (
          <div>
            <Fab
              color="primary"
              className={classes.fab}
              aria-label="add"
              onClick={() => {
                this.setState({ editDialogShow: true });
              }}
            >
              <AddIcon />
            </Fab>
            <Edit
              show={this.state.editDialogShow}
              categoryId={this.props.category.ID}
              categoryName={this.props.category.Name}
              handleClose={() => {
                this.setState({ editDialogShow: false });
              }}
              createSuccess={this.success.bind(this)}
            />
          </div>
        ) : (
          ""
        )}
      </div>
    );
  }
}

AddTask.propTypes = {
  category: PropTypes.object,
  createSuccess: PropTypes.func
};

export default withStyles(styles)(AddTask);
