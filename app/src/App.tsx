import React from "react";
import Button from "@mui/material/Button";
import "./App.css";
import { AppBar, IconButton, Toolbar, Typography } from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import Chart from "./component/chart";

export default App;

function App() {
  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" color="inherit" component="div">
            Growth of Codes
          </Typography>
        </Toolbar>
      </AppBar>
      <Chart />
    </div>
  );
}
