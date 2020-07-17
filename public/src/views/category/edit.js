import React from "react";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";

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
          <DialogTitle id="form-dialog-title">新增分类</DialogTitle>
          <DialogContent>
            <TextField
              required
              value={this.state.name}
              margin="dense"
              id="name"
              onChange={this.textFieldChange.bind(this, "name")}
              label="分类名称"
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
              label="分类描述"
              type="text"
              fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.props.handleClose}>关闭</Button>
            <Button onClick={this.submitCategory.bind(this)} color="primary">
              确认
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default EditCategory;
