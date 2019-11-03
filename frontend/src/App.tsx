import React from 'react';
import axios from 'axios';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

import './App.css';
import DefensesContainer from './DefensesContainer';
import KickersContainer from './KickersContainer';
import PlayersContainer from './PlayersContainer';
import { defense, kicker, player } from './types';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
  }),
);

const App: React.FC = () => {
  const [entity, setEntity] = React.useState('player');
  const [allPlayers, setAllPlayers] = React.useState<Array<player>>([]);
  const [kickers, setKickers] = React.useState<Array<kicker>>([]);
  const [defenses, setDefenses] = React.useState<Array<defense>>([]);

  const classes = useStyles();

  const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEntity(event.target.value as string);
  };

  React.useEffect(() => {
    axios.get('http://localhost:7000/offenseplayers')
      .then(res => {
        setAllPlayers(res.data);
      });
    axios.get('http://localhost:7000/kickers')
      .then(res => {
        setKickers(res.data);
      }); 
    axios.get('http://localhost:7000/defenses')
      .then(res => {
        setDefenses(res.data);
      });   
  }, [])

  return (
    <div className="App">
      <h1 style={{fontFamily: 'Roboto', fontWeight: 400}}>Walter's Notebook</h1>
      <FormControl className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">Entity</InputLabel>
        <Select
          labelId="demo-simple-select-label"
          id="demo-simple-select"
          value={entity}
          onChange={handleChange}
        >
          <MenuItem value={'player'}>Offense</MenuItem>
          <MenuItem value={'kicker'}>Kicker</MenuItem>
          <MenuItem value={'defense'}>Defense</MenuItem>
        </Select>
      </FormControl>
      { entity === 'player' ? <PlayersContainer players={allPlayers}/> : entity === 'kicker' ? <KickersContainer kickers={kickers} /> : <DefensesContainer defenses={defenses} /> }
      <footer>Rankings are pulled hourly from <a href="https://walterfootball.com/fantasy.php" target="_blank" rel="noopener noreferrer">Walter Football</a></footer>
    </div>
  );
}

export default App;
