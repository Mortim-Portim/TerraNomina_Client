package Game

import (
	"marvin/GraphEng/GE"
)

const (
	STANDARD_SCREEN_WIDTH = 1920
	STANDARD_SCREEN_HEIGHT = 1080
	
	FPS = 30
	RES = "./.res"
	AUDIO_FILES = "/Audio"
	ICON_FILES = "/Icons"
	IMAGE_FILES = "/Images"
	TITELSCREEN_FILES = IMAGE_FILES+"/GUI/Titlescreen"
	SOUNDTRACK_FILES = AUDIO_FILES+"/Soundtrack"
	
	ICON_FORMAT = "png"
)
var (
	PARAMETER *GE.Params
	XRES, YRES float64
	
	ICON_SIZES = []int{16,32,48,64,128,256}
	
	TITLE_BackImg, TITLE_LoadingBar, TITLE_Name *GE.Animation
	
	MainTheme, BattleTheme *GE.Sounds
)







func (g *TerraNomina) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(XRES), int(YRES)
}