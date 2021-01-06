import React from "react";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import i18n from "react-intl-universal";
import { create } from "../../api/category.js";

class EditCategory extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      desc: ""
    };
  }

  componentDidMount() {}

  textFieldChange(field, event) {
    this.setState({ [field]: event.target.value });
  }
  submitCategory() {
    create(this.state)
      .then(r => {
        this.props.createSuccess();
      })
      .catch(() => {});
  }

  render() {
    return (
      <div>
        <Dialog
          open={this.props.show}
          onClose={this.props.handleClose}
          aria-labelledby="form-dialog-title"
        >
          <DialogTitle id="form-dialog-title">
            {i18n.get("common.create")}
          </DialogTitle>
          <DialogContent>
            <TextField
              required
              value={this.state.name}
              margin="dense"
              id="name"
              onChange={this.textFieldChange.bind(this, "name")}
              label={i18n.get("common.name")}
              type="text"
              fullWidth
            />
            <TextField
              margin="dense"
              value={this.state.desc}
              id="desc"
              multiline
              rows={2}
              onChange={this.textFieldChange.bind(this, "desc")}
              label={i18n.get("common.desc")}
              type="text"
              fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.props.handleClose}>
              {i18n.get("common.close")}
            </Button>
            <Button onClick={this.submitCategory.bind(this)} color="primary">
              {i18n.get("common.confirm")}
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default EditCategory;
