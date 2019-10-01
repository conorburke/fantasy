package main

import (
	"fmt"
	"strings"
)

func selectOffensePlayers(s string) (players []offensePlayer, err error) {
	s = strings.ToUpper(s)
	if s == "" {
		s = "%"
	}
	rows, err := db.Query(fmt.Sprintf(`select first_name, last_name, team,
	bye, pos, pass_yards, interceptions, rush_yards, catches, rec_yards,
	reg_tds, bonus, dynst, points_espn, vbd_reg, points_ppr, ppr_vbd,
	points_td, vbd_td, points_two_qb, vbd_two_qb, points_custom, vbd_custom from offense_players
	where pos like '%s'`, s))
	if err != nil {
		return
	}
	for rows.Next() {
		p := offensePlayer{}
		err = rows.Scan(&p.FirstName, &p.LastName, &p.Team,
			&p.Bye, &p.Pos, &p.PassYards, &p.Int, &p.RushYds,
			&p.Catch, &p.RecYds, &p.RegTDs, &p.Bonus,
			&p.Dynst, &p.PointsESPN, &p.VBDReg, &p.PointsPPR,
			&p.PPRVBD, &p.PointsTD, &p.VBDTD, &p.PointsTwoQB,
			&p.VBDTwoQB, &p.PointsCustom, &p.VBDCustom)
		if err != nil {
			return
		}
		players = append(players, p)
	}
	rows.Close()
	return
}
