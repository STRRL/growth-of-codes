import { useEffect, useState } from "react";
import "./App.css";
import { AppBar, Toolbar, Typography, Container, Box } from "@mui/material";
import Chart, { Point } from "./component/chart";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
export default App;

function App() {
  const [repos, setRepos] = useState<string[]>([]);
  const [languages, setLanguages] = useState<string[]>([]);

  const [repo, setRepo] = useState("");
  const [language, setLanguage] = useState("");

  useEffect(() => {
    fetch(`/api/repo/list`)
      .then((response) => response.json())
      .then((response: string[]) => {
        if (response.length > 0) {
          setRepos(response);
        }
      });
  }, []);

  useEffect(() => {
    if (repo !== "") {
      fetch(`/api/language/list?repo=${repo}`)
        .then((response) => response.json())
        .then((response: string[]) => {
          if (response.length > 0) {
            setLanguages(response);
          }
        });
    }
  }, [repo]);

  useEffect(() => {
    if (language !== "" && repo !== "") {
      fetch(`/api/complexity?repo=${repo}&language=${language}`)
        .then((response) => response.json())
        .then((response) => setTimeSeries(response))
        .catch((error) => console.error(error));
    }
  }, [repo, language]);

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
        <RepoSelect repos={repos} onRepoChange={(repo) => setRepo(repo)} />
        <LanguageSelect
          languages={languages}
          onLanguageChange={(value) => setLanguage(value)}
        />
        <Box sx={{ height: "400px" }}>
          <Chart timeSeries={timeSeries} />
        </Box>
      </Container>
    </div>
  );
}

function RepoSelect(props: {
  repos: string[];
  onRepoChange: (event: string) => void;
}) {
  const [repo, setRepo] = useState("");
  const handleChange = (event: any) => {
    setRepo(event.target.value);
    props.onRepoChange(event.target.value);
  };
  return (
    <Box sx={{ minWidth: 120, padding: "20px" }}>
      <FormControl fullWidth>
        <InputLabel id="repo-select-label">Repository</InputLabel>
        <Select
          labelId="repo-select-label"
          id="repo-select"
          value={repo}
          label="Language"
          onChange={handleChange}
        >
          {props.repos.sort().map((repo) => (
            <MenuItem value={repo}>{repo}</MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
}

function LanguageSelect(props: {
  languages: string[];
  onLanguageChange: (event: string) => void;
}) {
  const [language, setLanguage] = useState("");

  const handleChange = (event: any) => {
    setLanguage(event.target.value);
    props.onLanguageChange(event.target.value);
  };

  return (
    <Box sx={{ minWidth: 120, padding: "20px" }}>
      <FormControl fullWidth>
        <InputLabel id="language-select-label">Language</InputLabel>
        <Select
          labelId="language-select-label"
          id="language-select"
          value={language}
          label="Language"
          onChange={handleChange}
        >
          {props.languages.sort().map((language) => (
            <MenuItem value={language}>{language}</MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
}
