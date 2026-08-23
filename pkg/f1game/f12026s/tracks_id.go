package f12026s

import "github.com/daanv2/race-game-dashboard/pkg/f1game/common"

type TrackID int

var (
	TRACK_ID_ABU_DHABI           = common.NewTrack(14, "Abu Dhabi")
	TRACK_ID_AUSTRIA             = common.NewTrack(17, "Austria")
	TRACK_ID_AUSTRIA_REVERSE     = common.NewTrack(40, "Austria (Reverse)")
	TRACK_ID_BAKU                = common.NewTrack(20, "Baku (Azerbaijan)")
	TRACK_ID_BRAZIL              = common.NewTrack(16, "Brazil")
	TRACK_ID_CATALUNYA           = common.NewTrack(4, "Catalunya")
	TRACK_ID_HUNGARORING         = common.NewTrack(9, "Hungaroring")
	TRACK_ID_IMOLA               = common.NewTrack(27, "Imola")
	TRACK_ID_JEDDAH              = common.NewTrack(29, "Jeddah")
	TRACK_ID_LAS_VEGAS           = common.NewTrack(31, "Las Vegas")
	TRACK_ID_LOSAIL              = common.NewTrack(32, "Losail")
	TRACK_ID_MADRID              = common.NewTrack(42, "Madrid")
	TRACK_ID_MELBOURNE           = common.NewTrack(0, "Melbourne")
	TRACK_ID_MEXICO              = common.NewTrack(19, "Mexico")
	TRACK_ID_MIAMI               = common.NewTrack(30, "Miami")
	TRACK_ID_MONACO              = common.NewTrack(5, "Monaco")
	TRACK_ID_MONTREAL            = common.NewTrack(6, "Montreal")
	TRACK_ID_MONZA               = common.NewTrack(11, "Monza")
	TRACK_ID_SAKHIR              = common.NewTrack(3, "Sakhir (Bahrain)")
	TRACK_ID_SHANGHAI            = common.NewTrack(2, "Shanghai")
	TRACK_ID_SILVERSTONE         = common.NewTrack(7, "Silverstone")
	TRACK_ID_SILVERSTONE_REVERSE = common.NewTrack(39, "Silverstone (Reverse)")
	TRACK_ID_SINGAPORE           = common.NewTrack(12, "Singapore")
	TRACK_ID_SPA                 = common.NewTrack(10, "Spa")
	TRACK_ID_SUZUKA              = common.NewTrack(13, "Suzuka")
	TRACK_ID_TEXAS               = common.NewTrack(15, "Texas")
	TRACK_ID_ZANDVOORT           = common.NewTrack(26, "Zandvoort")
	TRACK_ID_ZANDVOORT_REVERSE   = common.NewTrack(41, "Zandvoort (Reverse)")
)

func TrackData() []common.Track {
	return []common.Track{
		TRACK_ID_ABU_DHABI,
		TRACK_ID_AUSTRIA,
		TRACK_ID_AUSTRIA_REVERSE,
		TRACK_ID_BAKU,
		TRACK_ID_BRAZIL,
		TRACK_ID_CATALUNYA,
		TRACK_ID_HUNGARORING,
		TRACK_ID_IMOLA,
		TRACK_ID_JEDDAH,
		TRACK_ID_LAS_VEGAS,
		TRACK_ID_LOSAIL,
		TRACK_ID_MADRID,
		TRACK_ID_MELBOURNE,
		TRACK_ID_MEXICO,
		TRACK_ID_MIAMI,
		TRACK_ID_MONACO,
		TRACK_ID_MONTREAL,
		TRACK_ID_MONZA,
		TRACK_ID_SAKHIR,
		TRACK_ID_SHANGHAI,
		TRACK_ID_SILVERSTONE,
		TRACK_ID_SILVERSTONE_REVERSE,
		TRACK_ID_SINGAPORE,
		TRACK_ID_SPA,
		TRACK_ID_SUZUKA,
		TRACK_ID_TEXAS,
		TRACK_ID_ZANDVOORT,
		TRACK_ID_ZANDVOORT_REVERSE,
	}
}

func TrackMap() map[int]common.Track {
	result := map[int]common.Track{}

	for _, t := range TrackData() {
		result[t.ID] = t
	}

	return result
}
