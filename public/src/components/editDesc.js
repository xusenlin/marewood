import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import MuiDialogTitle from '@material-ui/core/DialogTitle';
import MuiDialogContent from '@material-ui/core/DialogContent';
import MuiDialogActions from '@material-ui/core/DialogActions';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';
import Typography from '@material-ui/core/Typography';
import {TextField} from "@material-ui/core";


const styles = (theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(2),
  },
  closeButton: {
    position: 'absolute',
    right: theme.spacing(1),
    top: theme.spacing(1),
    color: theme.palette.grey[500],
  },
});

const DialogTitle = withStyles(styles)((props) => {
  const {children, classes, onClose, ...other} = props;
  return (
    <MuiDialogTitle disableTypography className={classes.root} {...other}>
      <Typography variant="h6">{children}</Typography>
      {onClose ? (
        <IconButton aria-label="close" className={classes.closeButton} onClick={onClose}>
          <CloseIcon/>
        </IconButton>
      ) : null}
    </MuiDialogTitle>
  );
});

const DialogContent = withStyles((theme) => ({
  root: {
    padding: theme.spacing(2),
  },
}))(MuiDialogContent);

const DialogActions = withStyles((theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(1),
  },
}))(MuiDialogActions);


class EditDesc extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      desc: ""
    };
  }

  saveDesc() {
    this.props.editSuccess(this.props.id, this.state.desc)
  }

  onEnteredDialog() {
    this.setState({desc: this.props.desc});
  }

  textFieldChange(event) {
    this.setState({desc: event.target.value});
  }

  render() {
    return (
      <div>
        <Dialog onClose={this.props.onClose} onEntering={this.onEnteredDialog.bind(this)}
                aria-labelledby="customized-dialog-title" open={this.props.open}>
          <DialogTitle id="customized-dialog-title" onClose={this.props.onClose}>
            编辑描述
          </DialogTitle>
          <DialogContent dividers style={{width: 600}}>
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              margin="dense" value={this.state.desc} id="desc" multiline rows={8}
              onChange={this.textFieldChange.bind(this)} label="描述" type="text" fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button autoFocus onClick={this.saveDesc.bind(this)} color="primary">
              保存
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default EditDesc
