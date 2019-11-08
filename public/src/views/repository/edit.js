import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import {create} from '../../api/repository'



class Edit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name:'',
            url:'',
            desc:'',
            userName:'',
            password:'',
            needAuth:false
        };
    }

    componentDidMount() {

    }

    textFieldChange(field, event){
        this.setState({[field]:event.target.value})
    }
    submitForm(){
        create(this.state).then(r=>{
            this.props.createSuccess()
        }).catch(()=>{})
    }

    render() {
        return (
            <div>
                <Dialog open={this.props.show} onClose={this.props.handleClose} aria-labelledby="form-dialog-title">
                    <DialogTitle id="form-dialog-title">编辑仓库</DialogTitle>
                    <DialogContent>
                        <TextField
                            required autoFocus value={this.state.name} margin="dense" id="name"
                            onChange={this.textFieldChange.bind(this,'name')} label="仓库名称" type="text" fullWidth
                        />
                        <TextField
                            required autoFocus value={this.state.url} margin="dense" id="url"
                            onChange={this.textFieldChange.bind(this,'url')} label="仓库Url" type="text" fullWidth
                        />
                        <TextField
                            autoFocus value={this.state.userName} margin="dense" id="userName"
                            onChange={this.textFieldChange.bind(this,'userName')} label="用户名" type="text" fullWidth
                        />
                        <TextField
                            autoFocus value={this.state.password} margin="dense" id="password"
                            onChange={this.textFieldChange.bind(this,'password')} label="密码" type="password" fullWidth
                        />

                        <TextField
                            required
                            margin="dense" value={this.state.desc} id="desc" multiline rows={2}
                            onChange={this.textFieldChange.bind(this,'desc')} label="仓库描述" type="text" fullWidth
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.props.handleClose}>
                            关闭
                        </Button>
                        <Button onClick={this.submitForm.bind(this)} color="primary">
                            确认
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>
        );
    }

}


export default Edit
