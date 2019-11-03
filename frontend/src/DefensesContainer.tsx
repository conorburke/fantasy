import React from 'react';
import './App.css';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';

import { defense } from './types';
import { Order, getSorting, stableSort } from './sorting'

const useStyles = makeStyles({
  root: {
    width: '100%',
    overflowX: 'auto',
  },
  table: {
    minWidth: 650,
  },
});

const DefensesContainer: React.FC<{ defenses: Array<defense>; }> = (props: {defenses: Array<defense>}) => {
  const [order, setOrder] = React.useState<Order>('desc');
  const [orderBy, setOrderBy] = React.useState<keyof defense>('regScore');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);

  const classes = useStyles();

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  return (
    <div className="App">
      <Paper>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Team</TableCell>
              <TableCell>Bye</TableCell>
              <TableCell align="right">Points Regular</TableCell>
              <TableCell align="right">VBD Regular</TableCell>
              <TableCell align="right">Points PPR</TableCell>
              <TableCell align="right">VBD PPR</TableCell>
              <TableCell align="right">Points TD</TableCell>
              <TableCell align="right">VBD TD</TableCell>
              <TableCell align="right">Points Two QB</TableCell>
              <TableCell align="right">VBD Two QB</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {stableSort(props.defenses, getSorting(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map(row => (
              <TableRow key={row.team}>
                <TableCell component="th" scope="row">
                  {row.team}
                </TableCell>
                <TableCell>{row.bye}</TableCell>
                <TableCell align="right">{row.regScore}</TableCell>
                <TableCell align="right">{row.vbdReg}</TableCell>
                <TableCell align="right">{row.pprScore}</TableCell>
                <TableCell align="right">{row.vbdPpr}</TableCell>
                <TableCell align="right">{row.tdScore}</TableCell>
                <TableCell align="right">{row.vbdTd}</TableCell>
                <TableCell align="right">{row.twoQbScore}</TableCell>
                <TableCell align="right">{row.vbdTwoQb}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25, props.defenses.length]}
          component="div"
          count={props.defenses.length}
          rowsPerPage={rowsPerPage}
          page={page}
          backIconButtonProps={{
            'aria-label': 'previous page',
          }}
          nextIconButtonProps={{
            'aria-label': 'next page',
          }}
          onChangePage={handleChangePage}
          onChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Paper>
      {/* <ul>
        { allPlayers.map((item:player) => <li key={item.lastName + item.firstName + item.team}>{item.lastName}</li>)}
      </ul>
      <ul>
        { kickers.map((item:kicker) => <li key={item.lastName + item.firstName + item.team}>{item.firstName}</li>)}
      </ul>
      <ul>
        { defenses.map((item:defense) => <li key={item.team}>{item.team}</li>)}
      </ul> */}
      {/* <Table numRows={allPlayers.length} numFrozenColumns={allPlayers.length > 0 ? Object.keys(allPlayers[0]).length -1 : 0}>
        <Column name="Player" cellRenderer={cellRenderer} />
      </Table> */}
    </div>
  );
}

export default DefensesContainer;
