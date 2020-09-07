import React from "react";

import { repositories } from "../../api/repository.js";

import { withStyles } from "@material-ui/core/styles";
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import DialogTitle from '@material-ui/core/DialogTitle';
import Checkbox from '@material-ui/core/Checkbox';

const styles = theme => ({
});

class transfer extends React.Component {
        constructor(props) {
                super(props)
                this.state = {
                        list: [],
                        checked: []
                }
                this.getList()
        }
        closeDia() {
                this.props.closeDia()
        }
        isTransfer() {
                console.log(this.state.checked)
                this.closeDia()
        }
        handleChange = (id) => {
                const checkList = this.state.checked || [];
                const index = checkList.indexOf(id)

                if (checkList.indexOf(id) < 0) checkList.push(id);
                else checkList.splice(index, 1)
                this.setState({ checked: checkList })
        }
        getList = () => {
                repositories({
                        pageNum: 1, pageSize: 0
                }).then(r => {
                        this.setState({ list: r })
                })
        }
        creatcontroll = (list) => {
                let eleList = []
                for (let index = 0; index < list.length; index++) {
                        const element = list[index];
                        eleList.push((
                                <FormControlLabel
                                        key={element.ID}
                                        control={<Checkbox color="primary" value={element.ID} onChange={()=>this.handleChange(element.ID)} />}
                                        label={element.Name + element.ID}
                                />
                        ))
                }
                return eleList
        }
        render() {
                return (
                        <Dialog open={this.props.transfer || false} onClose={this.closeDia.bind(this)} aria-labelledby="form-dialog-title">
                                <DialogTitle id="form-dialog-title">请选择需要迁移的内容</DialogTitle>
                                <DialogContent>
                                        <FormGroup>
                                                {this.creatcontroll(this.state.list)}
                                        </FormGroup>
                                </DialogContent>
                                <DialogActions>
                                        <Button variant="contained" color="secondary" onClick={this.isTransfer.bind(this)}>确认迁移</Button>
                                </DialogActions>
                        </Dialog >
                )
        }
}

export default withStyles(styles)(transfer)
