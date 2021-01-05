import React, { useState } from "react";
import i18n from "react-intl-universal";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import Link from "@material-ui/core/Link";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import Copyright from "../../components/copyright";
import { setToken, setUserInfo, removeToken } from "../../utils/dataStorage";
import { login } from "../../api/user";
const Logo = require("../../assets/img/logo.svg");

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1)
  },
  submit: {
    margin: theme.spacing(3, 0, 2)
  }
}));

export default function SignIn() {
  removeToken();
  const classes = useStyles();

  let rememberMe = localStorage.getItem("RememberMe") === "1";
  let initLoginInfo = JSON.parse(localStorage.getItem("LoginInfo")) || {};

  const [checked, setChecked] = React.useState(rememberMe);
  const [username, setUsername] = useState(initLoginInfo.username);
  const [password, setPassword] = useState(initLoginInfo.password);
  const checkedChange = event => {
    setChecked(event.target.checked);

    if (event.target.checked) {
      localStorage.setItem("RememberMe", "1");
    } else {
      localStorage.removeItem("RememberMe");
    }
  };

  const usernameChange = event => {
    setUsername(event.target.value);
  };
  const passwordChange = event => {
    setPassword(event.target.value);
  };
  const submitLogin = () => {
    //登陆成功设置
    login({ username, password })
      .then(r => {
        setToken(r.Token);
        setUserInfo(r);
        window.location.href = "./";
        localStorage.setItem(
          "LoginInfo",
          JSON.stringify({ username, password })
        );
      })
      .catch(() => {});
  };
  return (
    <Container component="main" maxWidth="xs">
      <div className={classes.paper}>
        <img style={{ width: 60, marginBottom: 20 }} src={Logo} alt="" />
        <Typography component="h1" variant="h5">
          {i18n.get("login.title")}
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            value={username}
            onChange={usernameChange}
            id="username"
            label="username"
            name="username"
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            value={password}
            required
            fullWidth
            onChange={passwordChange}
            name="password"
            label="Password"
            type="password"
            id="password"
          />
          <FormControlLabel
            control={
              <Checkbox
                checked={checked}
                onChange={checkedChange}
                value="remember"
                color="primary"
              />
            }
            label="Remember me"
          />
          <Button
            fullWidth
            variant="contained"
            color="primary"
            onClick={submitLogin}
            className={classes.submit}
          >
            {i18n.get("login.signIn")}
          </Button>
          <Grid container>
            <Grid item xs>
              {/*<Link href="#" variant="body2">*/}
              {/*    没有账户? 注册一个吧。*/}
              {/*</Link>*/}
            </Grid>
            <Grid item>
              <Link href="#/register" variant="body2">
                {i18n.get("login.tip")}
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
}
