package f12024

import "github.com/daanv2/race-game-dashboard/pkg/f1game/common"

type TrackID int

var (
	TRACK_ID_Alpine                   = common.NewTrack(5, "Alpine")
	TRACK_ID_Art_GP_23                = common.NewTrack(143, "Art GP '23")
	TRACK_ID_Aston_Martin             = common.NewTrack(4, "Aston Martin")
	TRACK_ID_Campos_23                = common.NewTrack(144, "Campos '23")
	TRACK_ID_Carlin_23                = common.NewTrack(145, "Carlin '23 ")
	TRACK_ID_Dams_23                  = common.NewTrack(147, "Dams '23")
	TRACK_ID_F1_Custom_Team           = common.NewTrack(104, "F1 Custom Team")
	TRACK_ID_F1_Generic               = common.NewTrack(41, "F1 Generic")
	TRACK_ID_Ferrari                  = common.NewTrack(1, "Ferrari")
	TRACK_ID_Haas                     = common.NewTrack(7, "Haas")
	TRACK_ID_Hitech_23                = common.NewTrack(148, "Hitech '23")
	TRACK_ID_McLaren                  = common.NewTrack(8, "McLaren")
	TRACK_ID_Mercedes                 = common.NewTrack(0, "Mercedes")
	TRACK_ID_MP_Motorsport_23         = common.NewTrack(149, "MP Motorsport '23")
	TRACK_ID_PHM_23                   = common.NewTrack(146, "PHM '23")
	TRACK_ID_Prema_23                 = common.NewTrack(150, "Prema '23")
	TRACK_ID_RB                       = common.NewTrack(6, "RB")
	TRACK_ID_Red_Bull_Racing          = common.NewTrack(2, "Red Bull Racing")
	TRACK_ID_Sauber                   = common.NewTrack(9, "Sauber")
	TRACK_ID_Trident_23               = common.NewTrack(151, "Trident '23")
	TRACK_ID_Van_Amersfoort_Racing_23 = common.NewTrack(152, "Van Amersfoort Racing '23")
	TRACK_ID_Virtuosi_23              = common.NewTrack(153, "Virtuosi '23")
	TRACK_ID_Williams                 = common.NewTrack(3, "Williams")
)

func TrackData() []common.Track {
	return []common.Track{
		TRACK_ID_Alpine,
		TRACK_ID_Art_GP_23,
		TRACK_ID_Aston_Martin,
		TRACK_ID_Campos_23,
		TRACK_ID_Carlin_23,
		TRACK_ID_Dams_23,
		TRACK_ID_F1_Custom_Team,
		TRACK_ID_F1_Generic,
		TRACK_ID_Ferrari,
		TRACK_ID_Haas,
		TRACK_ID_Hitech_23,
		TRACK_ID_McLaren,
		TRACK_ID_Mercedes,
		TRACK_ID_MP_Motorsport_23,
		TRACK_ID_PHM_23,
		TRACK_ID_Prema_23,
		TRACK_ID_RB,
		TRACK_ID_Red_Bull_Racing,
		TRACK_ID_Sauber,
		TRACK_ID_Trident_23,
		TRACK_ID_Van_Amersfoort_Racing_23,
		TRACK_ID_Virtuosi_23,
		TRACK_ID_Williams,
	}
}

func TrackMap() map[int]common.Track {
	result := map[int]common.Track{}

	for _, t := range TrackData() {
		result[t.ID] = t
	}

	return result
}
