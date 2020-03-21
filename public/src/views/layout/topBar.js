import React from 'react';
import { AppBar,Toolbar,IconButton} from '@material-ui/core'
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Link from "@material-ui/core/Link";


const TopBar = props => {
    return (
        <div className="topBar">
            <AppBar position="static" color="default">
                <Toolbar>
                    <IconButton edge="start" color="inherit" aria-label="menu" onClick={ props.toggle }>
                        <MenuIcon/>
                    </IconButton>
                    <Typography variant="h6" style={{flexGrow: 1,marginLeft:20}}>
                        { props.title }
                    </Typography>
                    <div>
                        <Link href="#/login">
                            <AccountCircle />
                        </Link>
                    </div>
                </Toolbar>
            </AppBar>
        </div>
    )
};

export default TopBar
