import React from "react";
import { Badge, IconButton, Menu } from "@material-ui/core";
import { Alert } from "@material-ui/lab";
import NoticeIcon from "@material-ui/icons/NotificationsPaused";
import { withStyles } from "@material-ui/core/styles";
import Divider from "@material-ui/core/Divider";
import BusEvent from "../../utils/busEvents";
import i18n from "react-intl-universal";

const styles = theme => ({
  item: {
    margin: theme.spacing(1)
  },
  content:{
    display: "flex",
    maxHeight:360,
    flexDirection: "column",
  },
  notice:{
    flex:1,
    overflowY:"auto"
  },
  bar: {
    width: "100%",
    justifyContent: "space-around",
    color: "#c9c9c9",
    marginTop: 10,
    alignItems: "center",
    display: "flex"
  }
});

const Notice = props => {
  const [anchorEl, setAnchorEl] = React.useState(null);
  const [notices, setNotice] = React.useState([]);
  const [onlineUsers, setOnlineUsers] = React.useState(0);
  BusEvent.pushNotice = n => {
    setNotice([n, ...notices]);
    setOnlineUsers(n.OnlineUsers || 0);
  };
  const { classes } = props;
  return (
    <div>
      <IconButton
        edge="start"
        color="primary"
        aria-label="notice-menu"
        onClick={event => {
          setAnchorEl(event.currentTarget);
        }}
      >
        <Badge
          color="secondary"
          variant={notices.length === 0 ? "standard" : "dot"}
        >
          <NoticeIcon />
        </Badge>
      </IconButton>
      <Menu
        id="notice-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        PaperProps={{
          style: {
            padding: 0,
            width: "40ch",
          }
        }}
        onClose={() => {
          setAnchorEl(null);
        }}
      >
        <div className={classes.content}>
          <div className={classes.notice}>
            {notices.length === 0 ? (
              <Alert severity="info" className={classes.item}>
                {i18n.get("bar.noticeEmpty")}
              </Alert>
            ) : (
              notices.map((r, i) => (
                <Alert key={i} severity={r.Type} className={classes.item}>
                  {r.Message}
                </Alert>
              ))
            )}
          </div>
          <div>
            <Divider />
            <div className={classes.bar}>
          <span>
            {i18n.get("bar.noticeQuantityTip", { v: notices.length })}
          </span>
              <span>{i18n.get("bar.onlineTip", { v: onlineUsers })}</span>
            </div>
          </div>
        </div>
      </Menu>
    </div>
  );
};

export default withStyles(styles)(Notice);
