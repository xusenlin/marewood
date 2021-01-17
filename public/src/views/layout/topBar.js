import React from "react";
import {
  AppBar,
  Toolbar,
  IconButton,
  Grid,
  Menu,
  MenuItem,
  Link,
  Chip
} from "@material-ui/core";
import Typography from "@material-ui/core/Typography";
import Notice from "./notice";
import { ExitToApp, AssignmentInd } from "@material-ui/icons";

import Search from "./search";
import Lang from "./lang";
import i18n from "react-intl-universal";
import { getUserInfo } from "../../utils/dataStorage";

const Logo = require("../../assets/img/logov2.svg");
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
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            onClick={props.toggle}
          >
            {/*<MenuIcon color="primary"/>*/}
            <div style={{ width: 24, height: 24, display: "flex" }}>
              <img style={{ width: 24 }} src={Logo} alt="" />
            </div>
          </IconButton>
          <Typography variant="h6" style={{ flexGrow: 1, marginLeft: 20 }}>
            {props.title}
          </Typography>
          <div>
            <Grid container alignItems="center" spacing={2}>
              <Grid item>
                <Search />
              </Grid>
              <Grid item>
                <Notice />
              </Grid>
              <Grid item>
                <Lang />
              </Grid>
              <Grid item>
                <Chip
                  onClick={handleClick}
                  color="primary"
                  label={UserInfo.Username}
                />
                <Menu
                  id="user-menu"
                  anchorEl={anchorEl}
                  keepMounted
                  open={Boolean(anchorEl)}
                  onClose={handleClose}
                >
                  <MenuItem>
                    <AssignmentInd
                      style={{ marginRight: 10 }}
                      fontSize="small"
                    />
                    {i18n.get("role." + UserInfo.Role)}
                  </MenuItem>
                  <Link href="#/login">
                    <MenuItem>
                      <ExitToApp style={{ marginRight: 10 }} fontSize="small" />
                      {i18n.get("bar.logout")}
                    </MenuItem>
                  </Link>
                </Menu>
              </Grid>
            </Grid>
          </div>
        </Toolbar>
      </AppBar>
    </div>
  );
};

export default TopBar;
