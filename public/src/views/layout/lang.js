import React from "react";
import { IconButton, Menu, MenuItem } from "@material-ui/core";
import LanguageIcon from "@material-ui/icons/Language";
import { setUserLang } from "../../utils/dataStorage";
import { withStyles } from "@material-ui/core/styles";
import i18n from "react-intl-universal";

const styles = () => ({});

const Lang = () => {
  const [anchorEl, setAnchorEl] = React.useState(null);
  // const { classes } = props;
  return (
    <div>
      <IconButton
        edge="start"
        color="primary"
        aria-label="menu"
        onClick={event => {
          setAnchorEl(event.currentTarget);
        }}
      >
        <LanguageIcon />
      </IconButton>
      <Menu
        id="lang-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        PaperProps={{
          style: {
            padding: 0,
            maxHeight: 360,
            width: "40ch"
          }
        }}
        onClose={() => {
          setAnchorEl(null);
        }}
      >
        <MenuItem
          onClick={() => {
            setUserLang("en-US");
          }}
        >
          {i18n.get("bar.switchEN")}
        </MenuItem>
        <MenuItem
          onClick={() => {
            setUserLang("zh-CN");
          }}
        >
          {i18n.get("bar.switchCN")}
        </MenuItem>
      </Menu>
    </div>
  );
};

export default withStyles(styles)(Lang);
