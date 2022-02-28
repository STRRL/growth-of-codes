import React from "react";
import "./App.css";
import { AppBar, Toolbar, Typography } from "@mui/material";
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
