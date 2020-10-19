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
	AUDIO_FILES = 								"/Audio"
	ICON_FILES = 								"/Icons"
	IMAGE_FILES = 								"/Images"
	TITELSCREEN_FILES = 	IMAGE_FILES+"/GUI/Titlescreen"
	PLAYMENU_FILES = 		IMAGE_FILES+"/GUI/PlayMenu"
	BUTTON_FILES =			IMAGE_FILES+"/GUI/Buttons"
	CONNECTING_FILES =		IMAGE_FILES+"/GUI/Connecting"
	SOUNDTRACK_FILES = 		AUDIO_FILES+"/Soundtrack"
	KEYLI_MAPPER_FILE =							"/keyli.txt"
	
	STRUCTURE_FILES = 							"/Maps/structures"
	TILE_FILES = 								"/Maps/tiles"
	
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







func (g *TerraNomina) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(XRES), int(YRES)
}