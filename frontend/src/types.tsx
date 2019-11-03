export interface defense {
  team: string
	bye: string
	regScore: number
	vbdReg: number
	pprScore: number
	vbdPpr: number
	tdScore: number
	vbdTd: number
	twoQbScore: number
	vbdTwoQb: number
	customScore: number
	vbdCustom: number
}

export interface kicker {
	firstName: string
	lastName: string
	team: string
	bye: string
	position: string
	fg1to39: string
	fg40to49: string
	fg50Plus: string
	xp: string
	espn: string
	vbdReg: string
	pointsPpr: string
	pprVbd: string
	pointsTd: string
	vbdTd: string
	pointsTwoQb: string
	vbdTwoQb: string
	pointsCustom: string
	vbdCustom: string
}

export interface player {
  firstName: string
	lastName: string
	team: string
	bye: string
	position: string
	passYards: number
	interceptions: number
	rushYards: number
	catches: number
	receivingYards: number
	regYards: number
	bonus: number
	dynst: number
	pointsEspn: number
	vbdReg: number
	pointsPpr: number
	pprVbd: number
	pointsTd: number
	vbdTd: number
	pointsTwoQb: number
	vbdTwoQb: number
	pointsCustom: number
	vbdCustom: number
}
