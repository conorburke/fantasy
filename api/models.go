package main

type fantasyEntity interface {
	getTeam() string
}

type offensePlayer struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Team         string `json:"team"`
	Bye          string `json:"bye"`
	Pos          string `json:"position"`
	PassYards    int `json:"passYards"`
	Int          int `json:"interceptions"`
	RushYds      int `json:"rushYards"`
	Catch        int `json:"catches"`
	RecYds       int `json:"receivingYards"`
	RegTDs       int `json:"regYards"`
	Bonus        int `json:"bonus"`
	Dynst        int `json:"dynst"`
	PointsESPN   int `json:"pointsEspn"`
	VBDReg       int `json:"vbdReg"`
	PointsPPR    int `json:"pointsPpr"`
	PPRVBD       int `json:"pprVbd"`
	PointsTD     int `json:"pointsTd"`
	VBDTD        int `json:"vbdTd"`
	PointsTwoQB  int `json:"pointsTwoQb"`
	VBDTwoQB     int `json:"vbdTwoQb"`
	PointsCustom int `json:"pointsCustom"`
	VBDCustom    int `json:"vbdCustom"`
}

type kicker struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Team         string `json:"team"`
	Bye          string `json:"bye"`
	Pos          string `json:"position"`
	FG1to39      int `json:"fg1to39"`
	FG40to49     int `json:"fg40to49"`
	FG50Plus     int `json:"fg50Plus"`
	XP           int `json:"xp"`
	ESPN         int `json:"espn"`
	VBDReg       int `json:"vbdReg"`
	PointsPPR    int `json:"pointsPpr"`
	PPRVBD       int `json:"pprVbd"`
	PointsTD     int `json:"pointsTd"`
	VBDTD        int `json:"vbdTd"`
	PointsTwoQB  int `json:"pointsTwoQb"`
	VBDTwoQB     int `json:"vbdTwoQb"`
	PointsCustom int `json:"pointsCustom"`
	VBDCustom    int `json:"vbdCustom"`
}

type defense struct {
	Team        string `json:"team"`
	Bye         string `json:"bye"`
	RegScore    int `json:"regScore"`
	VBDReg      int `json:"vbdReg"`
	PPRScore    int `json:"pprScore"`
	VBDPPR      int `json:"vbdPpr"`
	TDScore     int `json:"tdScore"`
	VBDTD       int `json:"vbdTd"`
	TwoQBScore  int `json:"twoQbScore"`
	VBDTwoQB    int `json:"vbdTwoQb"`
	CustomScore int `json:"customScore"`
	VBDCustom   int `json:"vbdCustom"`
}

func (o offensePlayer) getTeam() string {
	return o.Team
}

func (k kicker) getTeam() string {
	return k.Team
}

func (d defense) getTeam() string {
	return d.Team
}