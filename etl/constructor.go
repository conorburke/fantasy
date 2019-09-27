package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readToStruct(s string) {
	// Open CSV file
	f, err := os.Open(fmt.Sprintf("/tmp/%s", s))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// fmt.Println("lines")
	// fmt.Println(lines)

	for i, line := range lines {
		column := strings.Split(line[0], "|")
		if i > 0 {
			if s == "Ks" {
				data := kicker{
					FirstName:    column[0],
					LastName:     column[1],
					Team:         column[2],
					Bye:          column[3],
					Pos:          column[4],
					FG1to39:      sToI(column[5]),
					FG40to49:     sToI(column[6]),
					FG50Plus:     sToI(column[7]),
					XP:           sToI(column[8]),
					ESPN:         sToI(column[9]),
					VBDReg:       sToI(column[10]),
					PointsPPR:    sToI(column[11]),
					PPRVBD:       sToI(column[12]),
					PointsTD:     sToI(column[13]),
					VBDTD:        sToI(column[14]),
					PointsTwoQB:  sToI(column[15]),
					VBDTwoQB:     sToI(column[16]),
					PointsCustom: sToI(column[17]),
					VBDCustom:    sToI(column[18]),
				}

				statement := `insert into kickers (first_name, last_name, team,
					bye, pos, fg1_to_39, fg40_to_49, fg50_plus, xp, points_espn,
					vbd_reg, points_ppr, ppr_vbd, points_td, vbd_td, points_two_qb, 
					vbd_two_qb, points_custom, vbd_custom) 
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
					$16, $17, $18, $19)`
				stmt, err := db.Prepare(statement)
				fmt.Println(stmt)
				if err != nil {
					panic(err)
				}
				defer stmt.Close()
				stmt.QueryRow(data.FirstName, data.LastName, data.Team,
					data.Bye, data.Pos, data.FG1to39, data.FG40to49, data.FG50Plus,
					data.XP, data.ESPN, data.VBDReg, data.PointsPPR,
					data.PPRVBD, data.PointsTD, data.VBDTD, data.PointsTwoQB,
					data.VBDTwoQB, data.PointsCustom, data.VBDCustom).Scan(&data.FirstName)
			} else if s == "DEFs" {
				data := defense{
					Team:        column[0],
					Bye:         column[1],
					RegScore:    sToI(column[2]),
					VBDReg:      sToI(column[3]),
					PPRScore:    sToI(column[4]),
					VBDPPR:      sToI(column[5]),
					TDScore:     sToI(column[6]),
					VBDTD:       sToI(column[7]),
					TwoQBScore:  sToI(column[8]),
					VBDTwoQB:    sToI(column[9]),
					CustomScore: sToI(column[10]),
					VBDCustom:   sToI(column[11]),
				}
				statement := `insert into defenses (team, bye, reg_score,
					vbd_reg, points_ppr, ppr_vbd, points_td, vbd_td, points_two_qb, 
					vbd_two_qb, points_custom, vbd_custom) 
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
				stmt, err := db.Prepare(statement)
				fmt.Println(stmt)
				if err != nil {
					panic(err)
				}
				defer stmt.Close()
				stmt.QueryRow(data.Team, data.Bye, data.RegScore, data.VBDReg, data.PPRScore,
					data.VBDPPR, data.TDScore, data.VBDTD, data.TwoQBScore,
					data.VBDTwoQB, data.CustomScore, data.VBDCustom).Scan(&data.Team)
			} else {
				data := offensePlayer{
					FirstName:    column[0],
					LastName:     column[1],
					Team:         column[2],
					Bye:          column[3],
					Pos:          column[4],
					PassYards:    sToI(column[5]),
					Int:          sToI(column[6]),
					RushYds:      sToI(column[7]),
					Catch:        sToI(column[8]),
					RecYds:       sToI(column[9]),
					RegTDs:       sToI(column[10]),
					Bonus:        sToI(column[11]),
					Dynst:        sToI(column[12]),
					PointsESPN:   sToI(column[13]),
					VBDReg:       sToI(column[14]),
					PointsPPR:    sToI(column[15]),
					PPRVBD:       sToI(column[16]),
					PointsTD:     sToI(column[17]),
					VBDTD:        sToI(column[18]),
					PointsTwoQB:  sToI(column[19]),
					VBDTwoQB:     sToI(column[20]),
					PointsCustom: sToI(column[21]),
					VBDCustom:    sToI(column[22]),
				}

				statement := `insert into offense_players (first_name, last_name, team,
					bye, pos, pass_yards, interceptions, rush_yards, catches, rec_yards,
					reg_tds, bonus, dynst, points_espn, vbd_reg, points_ppr, ppr_vbd,
					points_td, vbd_td, points_two_qb, vbd_two_qb, points_custom, vbd_custom) 
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
					$16, $17, $18, $19, $20, $21, $22, $23)`
				stmt, err := db.Prepare(statement)
				fmt.Println(stmt)
				if err != nil {
					panic(err)
				}
				defer stmt.Close()
				stmt.QueryRow(data.FirstName, data.LastName, data.Team,
					data.Bye, data.Pos, data.PassYards, data.Int, data.RushYds,
					data.Catch, data.RecYds, data.RegTDs, data.Bonus,
					data.Dynst, data.PointsESPN, data.VBDReg, data.PointsPPR,
					data.PPRVBD, data.PointsTD, data.VBDTD, data.PointsTwoQB,
					data.VBDTwoQB, data.PointsCustom, data.VBDCustom).Scan(&data.FirstName)
			}

		}
	}
}

func sToI(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
