package Game

import (
	"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
)

func InitParams(path string) {
	GE.InitParams(nil)
	
	PARAMETER = &GE.Params{}
	err := PARAMETER.LoadFromFile(path)
	CheckErr(err)
	
	xres,yres := ebiten.ScreenSizeInFullscreen(); XRES = float64(xres); YRES = float64(yres)
	
	backParams := &GE.Params{}; backParams.LoadFromFile(RES+TITELSCREEN_FILES+"/back.txt")
	back, err2 := GE.LoadEbitenImg(RES+TITELSCREEN_FILES+"/back.png")
	CheckErr(err2)
	TITLE_BackImg = GE.GetAnimationFromParams(0, 0, XRES, YRES, backParams, back)
	
	loadingParams := &GE.Params{}; loadingParams.LoadFromFile(RES+TITELSCREEN_FILES+"/loading.txt")
	loading, err2 := GE.LoadEbitenImg(RES+TITELSCREEN_FILES+"/loading.png")
	CheckErr(err2)
	TITLE_LoadingBar = GE.GetAnimationFromParams(0, 0, XRES, YRES, loadingParams, loading)
	
	nameParams := &GE.Params{}; nameParams.LoadFromFile(RES+TITELSCREEN_FILES+"/name.txt")
	name, err3 := GE.LoadEbitenImg(RES+TITELSCREEN_FILES+"/name.png")
	CheckErr(err3)
	TITLE_Name = GE.GetAnimationFromParams(0, 0, XRES, YRES/3.32, nameParams, name)
}

