import React from 'react';
import PropTypes from "prop-types";

import {
    Button,TextField,Dialog,DialogActions,InputLabel,
    DialogContent,DialogTitle,Select,MenuItem
} from '@material-ui/core';
import {create} from '../../api/job'
import { repositories,getScript } from '../../api/repository'




class Edit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name:'',
            desc:'',
            repositoryId:'',
            buildDir:'dist',
            buildCommand:'build',
            password:'',
            successScript:'',
            repositories:[],
            buildCommandList:{}
        };
    }

    componentDidMount() {
        repositories({isNormal:1}).then(r=>{
            this.setState({repositories:r})
        }).catch(()=>{})
    }
    textFieldChange(field, event){
        this.setState({[field]:event.target.value})
    }
    onChangeRepository(event){
        getScript({id:event.target.value}).then(r=>{
            this.setState({buildCommandList:r});
        }).catch(()=>{});
        this.setState({repositoryId:event.target.value});
    }
    submitForm(){
        create({...this.state,categoryId:this.props.categoryId}).then(r=>{
            this.props.createSuccess()
        }).catch(()=>{})
    }

    render() {
        return (
            <div>
                <Dialog open={this.props.show} onClose={this.props.handleClose} aria-labelledby="form-dialog-title">
                    <DialogTitle id="form-dialog-title">在《{this.props.categoryName}》分类下添加任务</DialogTitle>
                    <DialogContent>
                        <TextField
                            required autoFocus value={this.state.name} margin="dense" id="name"
                            onChange={this.textFieldChange.bind(this,'name')} label="任务名称" type="text" fullWidth
                        />
                        <InputLabel style={{marginTop:10,fontSize:"12px"}}>选择仓库</InputLabel>
                        <Select
                            style={{width:"100%"}}
                            value={this.state.repositoryId}
                            onChange={this.onChangeRepository.bind(this)}
                        >
                            {
                                this.state.repositories.map(r=>{
                                    return <MenuItem key={r.ID} value={r.ID}>{r.Name}</MenuItem>
                                })
                            }
                        </Select>

                        <TextField
                            required
                            autoFocus value={this.state.buildDir} margin="dense" id="buildDir"
                            onChange={this.textFieldChange.bind(this,'buildDir')} label="打包目录" type="text" fullWidth
                        />

                        <InputLabel style={{marginTop:10,fontSize:"12px"}}>选择打包的命令</InputLabel>
                        <Select
                            style={{width:"100%"}}
                            value={this.state.buildCommand}
                            onChange={this.textFieldChange.bind(this,'buildCommand')}
                        >
                            {
                                Object.keys(this.state.buildCommandList).map(key=>{
                                    return <MenuItem key={key} value={key}>{key} => { this.state.buildCommandList[key] } </MenuItem>
                                })
                            }
                        </Select>

                        <TextField
                            autoFocus value={this.state.password} margin="dense" id="password"
                            onChange={this.textFieldChange.bind(this,'password')} label="任务加密" type="password" fullWidth
                        />
                        <TextField
                            margin="dense" value={this.state.successScript} id="desc" multiline rows={6}
                            onChange={this.textFieldChange.bind(this,'successScript')} label="打包成功执行的命令（用户权限模块没有完成，现在不会运行）" type="text" fullWidth
                        />
                        <TextField
                            margin="dense" value={this.state.desc} id="desc" multiline rows={6}
                            onChange={this.textFieldChange.bind(this,'desc')} label="任务描述" type="text" fullWidth
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

Edit.propTypes = {
    show:PropTypes.bool.isRequired,
    categoryId: PropTypes.number.isRequired,
    categoryName: PropTypes.string.isRequired,
};

export default Edit
