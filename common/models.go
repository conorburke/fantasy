package common

type OffensePlayer struct {
	FirstName    string
	LastName     string
	Team         string
	Bye          string
	Pos          string
	PassYards    int
	Int          int
	RushYds      int
	Catch        int
	RecYds       int
	RegTDs       int
	Bonus        int
	Dynst        int
	PointsESPN   int
	VBDReg       int
	PointsPPR    int
	PPRVBD       int
	PointsTD     int
	VBDTD        int
	PointsTwoQB  int
	VBDTwoQB     int
	PointsCustom int
	VBDCustom    int
}

type Kicker struct {
	FirstName    string
	LastName     string
	Team         string
	Bye          string
	Pos          string
	FG1to39      int
	FG40to49     int
	FG50Plus     int
	XP           int
	ESPN         int
	VBDReg       int
	PointsPPR    int
	PPRVBD       int
	PointsTD     int
	VBDTD        int
	PointsTwoQB  int
	VBDTwoQB     int
	PointsCustom int
	VBDCustom    int
}

type Defense struct {
	Team        string
	Bye         string
	RegScore    int
	VBDReg      int
	PPRScore    int
	VBDPPR      int
	TDScore     int
	VBDTD       int
	TwoQBScore  int
	VBDTwoQB    int
	CustomScore int
	VBDCustom   int
}
