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

import { kicker } from './types';
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

const KickersContainer: React.FC<{ kickers: Array<kicker> }> = (props: {kickers: Array<kicker>}) => {
  const [order, setOrder] = React.useState<Order>('desc');
  const [orderBy, setOrderBy] = React.useState<keyof kicker>('espn');
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
              <TableCell>Name</TableCell>
              <TableCell>Team</TableCell>
              <TableCell>Bye</TableCell>
              <TableCell>FG 1 to 39</TableCell>
              <TableCell>FG 40 to 49</TableCell>
              <TableCell>FG 50+</TableCell>
              <TableCell>XP</TableCell>
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
            {stableSort(props.kickers, getSorting(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map(row => (
              <TableRow key={row.lastName + row.firstName + row.team}>
                <TableCell component="th" scope="row">
                  {row.firstName + " " + row.lastName}
                </TableCell>
                <TableCell>{row.team}</TableCell>
                <TableCell>{row.bye}</TableCell>
                <TableCell>{row.fg1to39}</TableCell>
                <TableCell>{row.fg40to49}</TableCell>
                <TableCell>{row.fg50Plus}</TableCell>
                <TableCell>{row.xp}</TableCell>
                <TableCell align="right">{row.espn}</TableCell>
                <TableCell align="right">{row.vbdReg}</TableCell>
                <TableCell align="right">{row.pointsPpr}</TableCell>
                <TableCell align="right">{row.pprVbd}</TableCell>
                <TableCell align="right">{row.pointsTd}</TableCell>
                <TableCell align="right">{row.vbdTd}</TableCell>
                <TableCell align="right">{row.pointsTwoQb}</TableCell>
                <TableCell align="right">{row.vbdTwoQb}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25, props.kickers.length]}
          component="div"
          count={props.kickers.length}
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
    </div>
  );
}

export default KickersContainer;