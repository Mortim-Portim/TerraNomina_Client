package Game

import (
	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
	"runtime"
	"sync"
)

const (
	STANDARD_SCREEN_WIDTH  = 1920
	STANDARD_SCREEN_HEIGHT = 1080

	FPS            = 30
	RES            = "./res"
	F_Params	   = RES + "/params.txt"
	F_KEYLI_MAPPER = RES + "/keyli.txt"
	F_ICONS        = RES + "/Icons"
	F_MAPS         = RES + "/Maps"
	F_CHARACTER    = RES + "/Character"
	F_STRUCTURES   = F_MAPS + "/structures"
	F_TILES        = F_MAPS + "/tiles"

	F_AUDIO      = RES + "/Audio"
	F_SOUNDTRACK = F_AUDIO + "/Soundtrack"

	F_IMAGES        = RES + "/Images"
	F_GUI           = F_IMAGES + "/GUI"
	F_TITLESCREEN   = F_GUI + "/Titlescreen"
	F_PLAYMENU      = F_GUI + "/PlayMenu"
	F_CHARACTERMENU = F_GUI + "/CharacterMenu"
	F_BUTTONS       = F_GUI + "/Buttons"
	F_CONNECTING    = F_GUI + "/Connecting"

	F_ENTITY   = RES + "/Entities"

	ICON_FORMAT = "png"

	TITLESCREEN_STATE = 0
	//CHARACTER_MENU_STATE = 1
	OPTIONS_MENU_STATE = 2
	PLAY_MENU_STATE    = 3
	CONNECTING_STATE   = 4
	INGAME_STATE       = 5
	SELRACE_STATE      = 6
	SELCLASS_STATE     = 7
	SELSTATS_STATE     = 8
	TEST_STATE         = 9

	SOUNDTRACK_MAIN         = "Main"
	SOUNDTRACK_ORK          = "Ork"
	SOUNDTRACK_BATTLE 		= "Battle"

	MAP_REQUEST = GC.MESSAGE_TYPES + 0
	CHAR_SEND   = GC.MESSAGE_TYPES + 1

	PLAYER_MODELL_HEIGHT = 2
)

var (
	PARAMETER  *GE.Params
	XRES, YRES float64

	ICON_SIZES = []int{16, 32, 48, 64, 128, 256}

	TITLE_BackImg, TITLE_LoadingBar, TITLE_Name *GE.Animation

	Soundtrack *GE.SoundTrack

	Keyli      *GE.KeyLi
	
	USER_INPUT_IP_ADDR string

	Client        *GC.Client
	ClientManager *GC.ClientManager

	SmallWorld *TNE.SmallWorld
	OwnPlayer *TNE.Player
	
	RecorderLock sync.Mutex
	Recorder *GE.Recorder
	RecordAll bool
	RecordingLength, RecordingScale float64
	RecordingFile string
)

//Should be saved to a file
var (
	MOVEMENT_SPEED                                    = 0.5
	MOVEMENT_UPDATE_PERIOD                            = int(1.0 / MOVEMENT_SPEED)
)

var (
	ESC_KEY_ID int
	left_key_id int
	right_key_id int
	up_key_id int
	down_key_id int
	record_key_id int
)

func (g *TerraNomina) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(XRES), int(YRES)
}
func ResetRecorder() {
	RecorderLock.Lock()
	Recorder.Delete()
	Recorder = GE.GetNewRecorder(int(FPS*RecordingLength), int(XRES*RecordingScale), int(YRES*RecordingScale), FPS)
	RecorderLock.Unlock()
	runtime.GC()
}