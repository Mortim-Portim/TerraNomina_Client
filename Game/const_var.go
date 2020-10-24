package Game

import (
	"marvin/GraphEng/GE"
	"marvin/GameConn/GC"
)

const (
	STANDARD_SCREEN_WIDTH = 1920
	STANDARD_SCREEN_HEIGHT = 1080
	
	FPS = 30
	RES = 									"./.res"
	F_KEYLI_MAPPER = 						RES + 	"/keyli.txt"
	F_ICONS =								RES +	"/Icons"
	F_MAPS =								RES +	"/Maps"
	F_STRUCTURES = 									F_MAPS + 	"/structures"
	F_TILES = 										F_MAPS + 	"/tiles"
	
	F_AUDIO = 								RES + 	"/Audio"
	F_SOUNDTRACK = 									F_AUDIO + 	"/Soundtrack"
	
	F_IMAGES =								RES +	"/Images"
	F_GUI = 										F_IMAGES +	"/GUI"
	F_TITLESCREEN =												F_GUI + 		"/Titlescreen"
	F_PLAYMENU = 												F_GUI +			"/PlayMenu"
	F_BUTTONS = 												F_GUI +			"/Buttons"
	F_CONNECTING = 												F_GUI +			"/Connecting"
	F_ANIMS = 										F_IMAGES +	"/Anims"
	F_DAYNIGHT = 												F_ANIMS + 		"/DayNight"
	
	
	/**
	RES = 									"./.res"
	AUDIO_FILES = 								"/Audio"
	ICON_FILES = 								"/Icons"
	IMAGE_FILES = 								"/Images"
	TITELSCREEN_FILES = 	IMAGE_FILES+"/GUI/Titlescreen"
	PLAYMENU_FILES = 		IMAGE_FILES+"/GUI/PlayMenu"
	BUTTON_FILES =			IMAGE_FILES+"/GUI/Buttons"
	CONNECTING_FILES =		IMAGE_FILES+"/GUI/Connecting"
	ANIMATION_FILES = 		IMAGE_FILES+"/Anims"
	SOUNDTRACK_FILES = 		AUDIO_FILES+"/Soundtrack"
	KEYLI_MAPPER_FILE =							"/keyli.txt"
	
	STRUCTURE_FILES = 							"/Maps/structures"
	TILE_FILES = 								"/Maps/tiles"
	**/
	
	ICON_FORMAT = "png"
	
	TITLESCREEN_STATE = 		0
	CHARACTER_MENU_STATE = 		1
	OPTIONS_MENU_STATE = 		2
	PLAY_MENU_STATE = 			3
	CONNECTING_STATE = 			4
	INGAME_STATE = 				5
	
	SOUNDTRACK_MAIN =			"main"
	SOUNDTRACK_ORK =			"ork"
	SOUNDTRACK_BATTLE_INTRO =	"battle_intro"
	SOUNDTRACK_BATTLE_CYCLE =	"battle_cycle"
	
	MAP_REQUEST = 				GC.MESSAGE_TYPES+0
	
	PLAYER_MODELL_HEIGHT = 		2
)
var (
	PARAMETER *GE.Params
	XRES, YRES float64
	
	ICON_SIZES = []int{16,32,48,64,128,256}
	
	TITLE_BackImg, TITLE_LoadingBar, TITLE_Name *GE.Animation
	
	Soundtrack *GE.SoundTrack
	
	Keyli *GE.KeyLi
	ESC_KEY_ID int
	
	USER_INPUT_IP_ADDR string
	
	Client *GC.Client
	ClientManager *GC.ClientManager
	
	WorldStructure *GE.WorldStructure
)

//Should be saved to a file
var (
	MOVEMENT_SPEED = 0.5
	MOVEMENT_UPDATE_PERIOD = int(1.0/MOVEMENT_SPEED)
	left_key_id, right_key_id, up_key_id, down_key_id int
)





func (g *TerraNomina) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(XRES), int(YRES)
}