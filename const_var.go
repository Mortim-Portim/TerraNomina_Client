package main

import (
	"runtime"
	"sync"

	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
)

const (
	STANDARD_SCREEN_WIDTH  = 1920
	STANDARD_SCREEN_HEIGHT = 1080
	RES                    = "./res"
	F_Params               = RES + "/params.txt"
	F_KEYLI_MAPPER         = RES + "/keyli.txt"
	F_ICONS                = RES + "/Icons"
	F_MAPS                 = RES + "/Maps"
	F_CHARACTER            = RES + "/Character"
	F_STRUCTURES           = F_MAPS + "/structures"
	F_TILES                = F_MAPS + "/tiles"

	F_AUDIO      = RES + "/Audio"
	F_SOUNDTRACK = F_AUDIO + "/Soundtrack"

	F_IMAGES        = RES + "/Images"
	F_SKILLS        = F_IMAGES + "/Skills"
	F_GUI           = F_IMAGES + "/GUI"
	F_InGame        = F_GUI + "/InGame"
	F_NPCICON       = F_GUI + "/NPCIcons"
	F_TITLESCREEN   = F_GUI + "/Titlescreen"
	F_CHARACTERMENU = F_GUI + "/CharacterMenu"
	F_BUTTONS       = F_GUI + "/Buttons"
	F_CONNECTING    = F_GUI + "/Connecting"
	F_UI_ELEMENTS   = F_GUI + "/Elements"

	F_ENTITY = RES + "/Entities"

	ICON_FORMAT = "png"

	TITLESCREEN_STATE = 0
	//CHARACTER_MENU_STATE = 1
	OPTIONS_MENU_STATE = 2
	PLAY_MENU_STATE    = 3
	CONNECTING_STATE   = 4
	INGAME_STATE       = 5
	SELSTATS_STATE     = 6
	//TEST_STATE         = 9

	SOUNDTRACK_MAIN   = "Main"
	SOUNDTRACK_ORK    = "Ork"
	SOUNDTRACK_BATTLE = "Battle"

	MAP_REQUEST = GC.MESSAGE_TYPES + 0
	CHAR_SEND   = GC.MESSAGE_TYPES + 1

	PLAYER_MODELL_HEIGHT = 2
)

var (
	PARAMETER       *GE.Params
	StandardIP_TEXT string
	XRES, YRES      float64

	ICON_SIZES = []int{16, 32, 48, 64, 128, 256}

	TITLE_BackImg, TITLE_LoadingBar, TITLE_Name *GE.Animation

	Soundtrack     *GE.SoundTrack
	StandardVolume float64

	Keyli *GE.KeyLi

	USER_INPUT_IP_ADDR string

	Client        *GC.Client
	ClientManager *GC.ClientManager

	SmallWorld *TNE.SmallWorld
	OwnPlayer  *TNE.Player

	RecorderLock                      sync.Mutex
	Recorder                          *GE.Recorder
	RecordAll                         bool
	RecordingLength, RecordingScale   float64
	ScreenCaptureFile, ScreenShotFile string

	ServerClosing chan bool

	Toaster *GE.Toaster
)

//Should be saved to a file
var (
	FPS = TNE.FPS
	// MOVEMENT_SPEED         = 0.5
	// MOVEMENT_UPDATE_PERIOD = int(1.0 / MOVEMENT_SPEED)
)

var (
	ESC_KEY_ID          int
	left_key_id         int
	right_key_id        int
	up_key_id           int
	down_key_id         int
	record_key_1_id     int
	record_key_2_id     int
	screenshot_key_2_id int
	interaction_key     int

	attack_key_1, attack_key_2, attack_key_3 int
)

func (g *TerraNomina) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(XRES), int(YRES)
}
func ResetRecorder() {
	RecorderLock.Lock()
	Recorder.Delete()
	Recorder = GE.GetNewRecorder(int(FPS*RecordingLength), int(XRES*RecordingScale), int(YRES*RecordingScale), int(FPS))
	RecorderLock.Unlock()
	runtime.GC()
}
