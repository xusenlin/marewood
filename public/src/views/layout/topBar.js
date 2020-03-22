import React from 'react';
import { AppBar,Toolbar,IconButton} from '@material-ui/core'
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import Person from '@material-ui/icons/Person';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import Link from "@material-ui/core/Link";
import Role from "../../config/role"
import { getUserInfo } from "../../utils/dataStorage"
import Avatar from '@material-ui/core/Avatar';
import Chip from '@material-ui/core/Chip';

let UserInfo = getUserInfo();


const TopBar = props => {

    const [anchorEl, setAnchorEl] = React.useState(null);

    const handleClick = event => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

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
                        <Chip onClick={handleClick} color="primary" label={UserInfo.Username} />
                        <Menu
                            id="user-menu"
                            anchorEl={anchorEl}
                            keepMounted
                            open={Boolean(anchorEl)}
                            onClose={handleClose}
                        >
                            <MenuItem >{ Role[UserInfo.Role] }</MenuItem>
                            <Link href="#/login">
                                <MenuItem >退出</MenuItem>
                            </Link>

                        </Menu>
                    </div>
                </Toolbar>
            </AppBar>
        </div>
    )
};

export default TopBar
