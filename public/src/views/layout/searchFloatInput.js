import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import InputBase from "@material-ui/core/InputBase";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import RotateLeft from "@material-ui/icons/RotateLeft";
import DirectionsIcon from "@material-ui/icons/Directions";
import BusEvent from "../../utils/busEvents";
import i18n from "react-intl-universal";
const useStyles = makeStyles(theme => ({
  root: {
    position: "fixed",
    left: 0,
    top: 0,
    padding: "2px 4px",
    display: "flex",
    alignItems: "center",
    width: "100%",
    height: 64,
    background: "#fff",
    zIndex: 100
  },
  input: {
    marginLeft: theme.spacing(1),
    flex: 1
  },
  iconButton: {
    padding: 10
  },
  divider: {
    height: 28,
    margin: 4
  }
}));

const SearchFloatInput = props => {
  const classes = useStyles();

  return (
    <div component="form" className={classes.root}>
      <IconButton
        className={classes.iconButton}
        aria-label="close"
        onClick={() => {
          props.closeInput();
        }}
      >
        <CloseIcon />
      </IconButton>
      <InputBase
        value={props.keyword}
        onChange={event => {
          props.setKeyword(event.target.value);
        }}
        onKeyDown={e => {
          if (e.keyCode !== 13) {
            return;
          }
          BusEvent.updateCurrentTable();
        }}
        className={classes.input}
        placeholder={i18n.get("bar.searchTip")}
        inputProps={{ "aria-label": "search keyword" }}
      />
      <IconButton
        color="primary"
        className={classes.iconButton}
        aria-label="search"
        onClick={() => {
          BusEvent.updateCurrentTable();
        }}
      >
        <DirectionsIcon />
      </IconButton>
      <Divider className={classes.divider} orientation="vertical" />
      <IconButton
        color="primary"
        className={classes.iconButton}
        aria-label="directions"
        onClick={() => {
          props.setKeyword("");
          setTimeout(() => {
            BusEvent.updateCurrentTable();
          }, 100);
        }}
      >
        <RotateLeft />
      </IconButton>
    </div>
  );
};
export default SearchFloatInput;
