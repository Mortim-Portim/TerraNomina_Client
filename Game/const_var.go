package Game

import (
	"marvin/GraphEng/GE"
)

const (
	STANDARD_SCREEN_WIDTH = 1920
	STANDARD_SCREEN_HEIGHT = 1080
	
	FPS = 30
	RES = "./.res"
	ICON_FILES = "/Icons"
	IMAGE_FILES = "/Images"
	TITELSCREEN_FILES = IMAGE_FILES+"/GUI/Titlescreen"
	
	ICON_FORMAT = "png"
)
var (
	PARAMETER *GE.Params
	XRES, YRES float64
	
	ICON_SIZES = []int{16,32,48,64,128,256}
	
	BackImg, LoadingBar *GE.Animation
)







func (g *TerraNomina) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(XRES), int(YRES)
}