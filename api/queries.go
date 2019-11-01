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

func selectKickers() (kickers []kicker, err error) {
	rows, err := db.Query(fmt.Sprintf(`select first_name, last_name, team,
	bye, pos, fg1_to_39, fg40_to_49, fg50_plus, xp, points_espn, vbd_reg, points_ppr, ppr_vbd,
	points_td, vbd_td, points_two_qb, vbd_two_qb, points_custom, vbd_custom from kickers`))
	if err != nil {
		return
	}
	for rows.Next() {
		k := kicker{}
		err = rows.Scan(&k.FirstName, &k.LastName, &k.Team,
			&k.Bye, &k.Pos, &k.FG1to39, &k.FG40to49, &k.FG50Plus,
			&k.XP, &k.ESPN, &k.VBDReg, &k.PointsPPR,
			&k.PPRVBD, &k.PointsTD, &k.VBDTD, &k.PointsTwoQB,
			&k.VBDTwoQB, &k.PointsCustom, &k.VBDCustom)
		if err != nil {
			return
		}
		kickers = append(kickers, k)
	}
	rows.Close()
	return
}

func selectDefenses() (defenses []defense, err error) {
	rows, err := db.Query(fmt.Sprintf(`select team,
	bye, reg_score, vbd_reg, points_ppr, ppr_vbd,
	points_td, vbd_td, points_two_qb, vbd_two_qb, points_custom, vbd_custom from defenses`))
	if err != nil {
		return
	}
	for rows.Next() {
		d := defense{}
		err = rows.Scan(&d.Team, &d.Bye, &d.RegScore,
		  &d.VBDReg, &d.PPRScore,
			&d.VBDPPR, &d.TDScore, &d.VBDTD, &d.TwoQBScore,
			&d.VBDTwoQB, &d.CustomScore, &d.VBDCustom)
		if err != nil {
			return
		}
		defenses = append(defenses, d)
	}
	rows.Close()
	return
}
