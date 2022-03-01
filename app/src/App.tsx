import React, { useEffect, useState } from "react";
import "./App.css";
import { AppBar, Toolbar, Typography, Container, Box } from "@mui/material";
import Chart, { Point } from "./component/chart";

export default App;

function App() {
  useEffect(() => {
    fetch("/api/random/20")
      .then((response) => response.json())
      .then((response) => setTimeSeries(response))
      .catch((error) => console.error(error));
  }, []);

  const [timeSeries, setTimeSeries] = useState<Point[]>([]);
  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" color="inherit" component="div">
            Growth of Codes
          </Typography>
        </Toolbar>
      </AppBar>
      <Container>
        <Box sx={{ height: "400px" }}>
          <Chart timeSeries={timeSeries} />
        </Box>
      </Container>
    </div>
  );
}
