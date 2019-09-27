CREATE TABLE IF NOT EXISTS offense_players (
	first_name    character varying DEFAULT ''::character varying NOT NULL,
	last_name     character varying DEFAULT ''::character varying NOT NULL,
	team         character varying DEFAULT ''::character varying NOT NULL,
	bye          character varying DEFAULT ''::character varying NOT NULL,
	pos          character varying DEFAULT ''::character varying NOT NULL,
	pass_yards    integer DEFAULT '0'::integer NOT NULL,
	interceptions integer DEFAULT '0'::integer NOT NULL,
	rush_yards      integer DEFAULT '0'::integer NOT NULL,
	catches        integer DEFAULT '0'::integer NOT NULL,
	rec_yards       integer DEFAULT '0'::integer NOT NULL,
	reg_tds       integer DEFAULT '0'::integer NOT NULL,
	bonus        integer DEFAULT '0'::integer NOT NULL,
	dynst        integer DEFAULT '0'::integer NOT NULL,
	points_espn   integer DEFAULT '0'::integer NOT NULL,
	vbd_reg       integer DEFAULT '0'::integer NOT NULL,
	points_ppr    integer DEFAULT '0'::integer NOT NULL,
	ppr_vbd       integer DEFAULT '0'::integer NOT NULL,
	points_td     integer DEFAULT '0'::integer NOT NULL,
	vbd_td        integer DEFAULT '0'::integer NOT NULL,
	points_two_qb  integer DEFAULT '0'::integer NOT NULL,
	vbd_two_qb     integer DEFAULT '0'::integer NOT NULL,
	points_custom integer DEFAULT '0'::integer NOT NULL,
	vbd_custom    integer DEFAULT '0'::integer NOT NULL
);

CREATE TABLE IF NOT EXISTS kickers (
	first_name    character varying DEFAULT ''::character varying NOT NULL,
	last_name     character varying DEFAULT ''::character varying NOT NULL,
	team         character varying DEFAULT ''::character varying NOT NULL,
	bye          character varying DEFAULT ''::character varying NOT NULL,
	pos          character varying DEFAULT ''::character varying NOT NULL,
	fg1_to_39    integer DEFAULT '0'::integer NOT NULL,
	fg40_to_49 integer DEFAULT '0'::integer NOT NULL,
	fg50_plus      integer DEFAULT '0'::integer NOT NULL,
	xp        integer DEFAULT '0'::integer NOT NULL,
	points_espn   integer DEFAULT '0'::integer NOT NULL,
	vbd_reg       integer DEFAULT '0'::integer NOT NULL,
	points_ppr    integer DEFAULT '0'::integer NOT NULL,
	ppr_vbd       integer DEFAULT '0'::integer NOT NULL,
	points_td     integer DEFAULT '0'::integer NOT NULL,
	vbd_td        integer DEFAULT '0'::integer NOT NULL,
	points_two_qb  integer DEFAULT '0'::integer NOT NULL,
	vbd_two_qb     integer DEFAULT '0'::integer NOT NULL,
	points_custom integer DEFAULT '0'::integer NOT NULL,
	vbd_custom    integer DEFAULT '0'::integer NOT NULL
);

CREATE TABLE IF NOT EXISTS defenses (
	team         character varying DEFAULT ''::character varying NOT NULL,
	bye          character varying DEFAULT ''::character varying NOT NULL,
	reg_score   integer DEFAULT '0'::integer NOT NULL,
	vbd_reg       integer DEFAULT '0'::integer NOT NULL,
	points_ppr    integer DEFAULT '0'::integer NOT NULL,
	ppr_vbd       integer DEFAULT '0'::integer NOT NULL,
	points_td     integer DEFAULT '0'::integer NOT NULL,
	vbd_td        integer DEFAULT '0'::integer NOT NULL,
	points_two_qb  integer DEFAULT '0'::integer NOT NULL,
	vbd_two_qb     integer DEFAULT '0'::integer NOT NULL,
	points_custom integer DEFAULT '0'::integer NOT NULL,
	vbd_custom    integer DEFAULT '0'::integer NOT NULL
);