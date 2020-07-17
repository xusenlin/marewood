import React from "react";
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import MuiDialogActions from "@material-ui/core/DialogActions";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import { TextField } from "@material-ui/core";
import PropTypes from "prop-types";

const styles = theme => ({
  root: {
    margin: 0,
    padding: theme.spacing(2)
  },
  closeButton: {
    position: "absolute",
    right: theme.spacing(1),
    top: theme.spacing(1),
    color: theme.palette.grey[500]
  }
});

const DialogTitle = withStyles(styles)(props => {
  const { children, classes, onClose, ...other } = props;
  return (
    <MuiDialogTitle disableTypography className={classes.root} {...other}>
      <Typography variant="h6">{children}</Typography>
      {onClose ? (
        <IconButton
          aria-label="close"
          className={classes.closeButton}
          onClick={onClose}
        >
          <CloseIcon />
        </IconButton>
      ) : null}
    </MuiDialogTitle>
  );
});

const DialogContent = withStyles(theme => ({
  root: {
    padding: theme.spacing(2)
  }
}))(MuiDialogContent);

const DialogActions = withStyles(theme => ({
  root: {
    margin: 0,
    padding: theme.spacing(1)
  }
}))(MuiDialogActions);

class EditField extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      fieldContent: ""
    };
  }

  saveFieldContent() {
    this.props.editSuccess(
      this.props.id,
      this.props.field,
      this.state.fieldContent
    );
  }

  onEnteredDialog() {
    this.setState({ fieldContent: this.props.fieldContent });
  }

  textFieldChange(event) {
    this.setState({ fieldContent: event.target.value });
  }

  render() {
    return (
      <div>
        <Dialog
          onClose={this.props.onClose}
          onEntering={this.onEnteredDialog.bind(this)}
          aria-labelledby="customized-dialog-title"
          open={this.props.open}
        >
          <DialogTitle
            id="customized-dialog-title"
            onClose={this.props.onClose}
          >
            修改{this.props.desc}
          </DialogTitle>
          <DialogContent dividers style={{ width: 600 }}>
            <TextField
              InputLabelProps={{
                shrink: true
              }}
              margin="dense"
              value={this.state.fieldContent}
              id="desc"
              multiline
              rows={this.props.rows}
              onChange={this.textFieldChange.bind(this)}
              label={this.props.desc}
              type="text"
              fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button
              autoFocus
              onClick={this.saveFieldContent.bind(this)}
              color="primary"
            >
              保存
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

EditField.propTypes = {
  id: PropTypes.number.isRequired,
  open: PropTypes.bool.isRequired,
  row: PropTypes.number,
  field: PropTypes.string.isRequired,
  desc: PropTypes.string.isRequired,
  fieldContent: PropTypes.string.isRequired,
  onClose: PropTypes.func.isRequired,
  editSuccess: PropTypes.func.isRequired
};

EditField.defaultProps = {
  row: 1
};
export default EditField;
