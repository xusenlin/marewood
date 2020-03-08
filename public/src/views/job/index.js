import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import {
    Tab,Tabs,Paper
} from '@material-ui/core';
import {jobsFind} from '../../api/job'
import {categories} from '../../api/category'
import CategoriesTable from "./categoriesTable"


const styles = theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(3),
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
    fab: {
        position: 'absolute',
        bottom: theme.spacing(2),
        right: theme.spacing(2),
    },
});


class Jobs extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            category:0,
            categories:[],
            jobs:[]
        };
    }

    componentDidMount() {

        categories().then(r=>{
            this.setState({categories:r},()=>{
                this.setTabAndJobsByCategoryId(0)
            });
        }).catch(()=>{})
    }
    tabsChange(event, index){
        this.setTabAndJobsByCategoryId(index)
    }
    setTabAndJobsByCategoryId(index = 0){
        let categoryId = this.state.categories[index].ID;
        jobsFind({id:categoryId}).then(r=>{
            this.setState({category:index,jobs:r});
        }).catch(()=>{})
    }
    render() {
        const {classes} = this.props;
        return (
            <div>
                <Paper className={classes.root}>
                    <Tabs
                        value={this.state.category}
                        indicatorColor="primary"
                        textColor="primary"
                        centered
                        onChange={this.tabsChange.bind(this)}
                    >
                        {
                            this.state.categories.map(category=>(
                                <Tab key={category.ID} label={category.Name} />
                            ))
                        }
                    </Tabs>
                    <CategoriesTable tableData={this.state.jobs} />
                </Paper>
            </div>
        );
    }

}


export default withStyles(styles)(Jobs)
