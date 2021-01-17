import React, { useState } from "react";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Link from "@material-ui/core/Link";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import Copyright from "../../components/copyright";
import { register } from "../../api/user";
import snackbar from "../../components/snackbar/index";
import i18n from "react-intl-universal";
const Logo = require("../../assets/img/logov2.svg");

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
    marginTop: theme.spacing(3)
  },
  submit: {
    margin: theme.spacing(3, 0, 2)
  }
}));

export default function SignUp() {
  const classes = useStyles();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const usernameChange = event => {
    setUsername(event.target.value);
  };
  const passwordChange = event => {
    setPassword(event.target.value);
  };
  const submitRegister = () => {
    register({ username, password })
      .then(() => {
        snackbar.success(i18n.get("register.registerSuccess"));
      })
      .catch(() => {});
  };
  return (
    <Container component="main" maxWidth="xs">
      <div className={classes.paper}>
        <img style={{ width: 60, marginBottom: 20 }} src={Logo} alt="" />
        <Typography component="h1" variant="h5">
          {i18n.get("register.title")}
        </Typography>
        <form className={classes.form} noValidate>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                required
                fullWidth
                onChange={usernameChange}
                id="username"
                label="Username"
                name="username"
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                required
                fullWidth
                onChange={passwordChange}
                name="password"
                label="Password"
                type="password"
                id="password"
              />
            </Grid>
            <Grid item xs={12}>
              <span>{i18n.get("register.roleTip")} </span>
            </Grid>
          </Grid>
          <Button
            fullWidth
            onClick={submitRegister}
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            {i18n.get("register.createAccount")}
          </Button>
          <Grid container justify="flex-end">
            <Grid item>
              <Link href="#/login" variant="body2">
                {i18n.get("register.tip")}
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={5}>
        <Copyright />
      </Box>
    </Container>
  );
}
