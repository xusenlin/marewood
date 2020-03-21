import React from "react";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import Config from "../config/index"

export default function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://github.com/xusenlin/MareWood">
        { Config.siteName }
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}
