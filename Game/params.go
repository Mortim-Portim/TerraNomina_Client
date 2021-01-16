package Game

import (
	"github.com/mortim-portim/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

func InitParams(path string) {
	GE.InitParams(nil)

	PARAMETER = &GE.Params{}
	err := PARAMETER.LoadFromFile(path)
	CheckErr(err)

	xres, yres := ebiten.ScreenSizeInFullscreen()
	XRES = float64(xres)
	YRES = float64(yres)

	backParams := &GE.Params{}
	backParams.LoadFromFile(F_TITLESCREEN + "/back.txt")
	back, err2 := GE.LoadEbitenImg(F_TITLESCREEN + "/back.png")
	CheckErr(err2)
	TITLE_BackImg = GE.GetAnimationFromParams(0, 0, XRES, YRES, backParams, back)

	loadingParams := &GE.Params{}
	loadingParams.LoadFromFile(F_TITLESCREEN + "/loading.txt")
	loading, err2 := GE.LoadEbitenImg(F_TITLESCREEN + "/loading.png")
	CheckErr(err2)
	TITLE_LoadingBar = GE.GetAnimationFromParams(0, 0, XRES, YRES, loadingParams, loading)

	nameParams := &GE.Params{}
	nameParams.LoadFromFile(F_TITLESCREEN + "/name.txt")
	name, err3 := GE.LoadEbitenImg(F_TITLESCREEN + "/name.png")
	CheckErr(err3)
	TITLE_Name = GE.GetAnimationFromParams(0, 0, XRES, XRES*0.19, nameParams, name)
}
