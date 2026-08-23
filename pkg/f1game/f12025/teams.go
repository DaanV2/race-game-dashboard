package f12025

import "github.com/daanv2/race-game-dashboard/pkg/f1game/common"

var (
	TEAM_ID_AIX_RACING_24            = common.NewTeam(161, "AIX Racing '24")
	TEAM_ID_ALPINE                   = common.NewTeam(5, "Alpine")
	TEAM_ID_ALPINE_24                = common.NewTeam(190, "Alpine '24")
	TEAM_ID_APXGP_24                 = common.NewTeam(142, "APXGP '24")
	TEAM_ID_APXGP_25                 = common.NewTeam(154, "APXGP '25")
	TEAM_ID_ART_GP_24                = common.NewTeam(158, "Art GP '24")
	TEAM_ID_ASTON_MARTIN             = common.NewTeam(4, "Aston Martin")
	TEAM_ID_ASTON_MARTIN_24          = common.NewTeam(189, "Aston Martin '24")
	TEAM_ID_CAMPOS_24                = common.NewTeam(159, "Campos '24")
	TEAM_ID_DAMS_24                  = common.NewTeam(162, "DAMS '24")
	TEAM_ID_F1_CUSTOM_TEAM           = common.NewTeam(104, "F1 Custom Team")
	TEAM_ID_F1_GENERIC               = common.NewTeam(41, "F1 Generic")
	TEAM_ID_FERRARI                  = common.NewTeam(1, "Ferrari")
	TEAM_ID_FERRARI_24               = common.NewTeam(186, "Ferrari '24")
	TEAM_ID_HAAS                     = common.NewTeam(7, "Haas")
	TEAM_ID_HAAS_24                  = common.NewTeam(192, "Haas '24")
	TEAM_ID_HITECH_24                = common.NewTeam(163, "Hitech '24")
	TEAM_ID_INVICTA_24               = common.NewTeam(168, "Invicta '24")
	TEAM_ID_KONNERSPORT              = common.NewTeam(129, "Konnersport")
	TEAM_ID_KONNERSPORT_24           = common.NewTeam(155, "Konnersport '24")
	TEAM_ID_MCLAREN                  = common.NewTeam(8, "McLaren")
	TEAM_ID_MCLAREN_24               = common.NewTeam(193, "McLaren '24")
	TEAM_ID_MERCEDES                 = common.NewTeam(0, "Mercedes")
	TEAM_ID_MERCEDES_24              = common.NewTeam(185, "Mercedes '24")
	TEAM_ID_MP_MOTORSPORT_24         = common.NewTeam(164, "MP Motorsport '24")
	TEAM_ID_PREMA_24                 = common.NewTeam(165, "Prema '24")
	TEAM_ID_RB                       = common.NewTeam(6, "RB")
	TEAM_ID_RB_24                    = common.NewTeam(191, "RB '24")
	TEAM_ID_RED_BULL_RACING          = common.NewTeam(2, "Red Bull Racing")
	TEAM_ID_RED_BULL_RACING_24       = common.NewTeam(187, "Red Bull Racing '24")
	TEAM_ID_RODIN_MOTORSPORT_24      = common.NewTeam(160, "Rodin Motorsport '24")
	TEAM_ID_SAUBER                   = common.NewTeam(9, "Sauber")
	TEAM_ID_SAUBER_24                = common.NewTeam(194, "Sauber '24")
	TEAM_ID_TRIDENT_24               = common.NewTeam(166, "Trident '24")
	TEAM_ID_VAN_AMERSFOORT_RACING_24 = common.NewTeam(167, "Van Amersfoort Racing '24")
	TEAM_ID_WILLIAMS                 = common.NewTeam(3, "Williams")
	TEAM_ID_WILLIAMS_24              = common.NewTeam(188, "Williams '24")
)

func TeamsData() []common.Team {
	return []common.Team{
		TEAM_ID_AIX_RACING_24,
		TEAM_ID_ALPINE,
		TEAM_ID_ALPINE_24,
		TEAM_ID_APXGP_24,
		TEAM_ID_APXGP_25,
		TEAM_ID_ART_GP_24,
		TEAM_ID_ASTON_MARTIN,
		TEAM_ID_ASTON_MARTIN_24,
		TEAM_ID_CAMPOS_24,
		TEAM_ID_DAMS_24,
		TEAM_ID_F1_CUSTOM_TEAM,
		TEAM_ID_F1_GENERIC,
		TEAM_ID_FERRARI,
		TEAM_ID_FERRARI_24,
		TEAM_ID_HAAS,
		TEAM_ID_HAAS_24,
		TEAM_ID_HITECH_24,
		TEAM_ID_INVICTA_24,
		TEAM_ID_KONNERSPORT,
		TEAM_ID_KONNERSPORT_24,
		TEAM_ID_MCLAREN,
		TEAM_ID_MCLAREN_24,
		TEAM_ID_MERCEDES,
		TEAM_ID_MERCEDES_24,
		TEAM_ID_MP_MOTORSPORT_24,
		TEAM_ID_PREMA_24,
		TEAM_ID_RB,
		TEAM_ID_RB_24,
		TEAM_ID_RED_BULL_RACING,
		TEAM_ID_RED_BULL_RACING_24,
		TEAM_ID_RODIN_MOTORSPORT_24,
		TEAM_ID_SAUBER,
		TEAM_ID_SAUBER_24,
		TEAM_ID_TRIDENT_24,
		TEAM_ID_VAN_AMERSFOORT_RACING_24,
		TEAM_ID_WILLIAMS,
		TEAM_ID_WILLIAMS_24,
	}
}

func TeamsMap() map[int]common.Team {
	result := map[int]common.Team{}

	for _, t := range TeamsData() {
		result[t.ID] = t
	}

	return result
}
