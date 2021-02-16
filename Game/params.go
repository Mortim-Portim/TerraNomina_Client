package Game

import (
	"github.com/mortim-portim/GraphEng/GE"
)
const(
	VOLUME_PARAM = "volume"
	RECORDING_TIME_PARAM = "rec_time"
	RECORDING_SCALE_PARAM = "rec_scale"
	IP_ADDR_PARAM = "ip_addr"
)
/**
volume:0.5
rec_time:5
rec_scale:0.1
ip_addr:ip:port
**/
func VarsToParams() {
	PARAMETER.Set(VOLUME_PARAM, StandardVolume)
	PARAMETER.Set(RECORDING_TIME_PARAM, RecordingLength)
	PARAMETER.Set(RECORDING_SCALE_PARAM, RecordingScale)
	PARAMETER.SetS(IP_ADDR_PARAM, StandardIP_TEXT)
}
func ParamsToVars() {
	StandardVolume = PARAMETER.Get(VOLUME_PARAM)
	RecordingLength = PARAMETER.Get(RECORDING_TIME_PARAM)
	RecordingScale = PARAMETER.Get(RECORDING_SCALE_PARAM)
	StandardIP_TEXT = PARAMETER.GetS(IP_ADDR_PARAM)
}
func InitParams(path string, x, y int) {
	GE.InitParams(nil)
	InitImages()
	InitButtons()

	PARAMETER = &GE.Params{}
	err := PARAMETER.LoadFromFile(path)
	CheckErr(err)
	ParamsToVars()

	//ebiten.ScreenSizeInFullscreen()
	XRES = float64(x)
	YRES = float64(y)

	backParams := &GE.Params{}
	backParams.LoadFromFile(F_TITLESCREEN + "/back.txt")
	back, err2 := GetEbitenImage(F_TITLESCREEN + "/back.png")
	CheckErr(err2)
	TITLE_BackImg = GE.GetAnimationFromParams(0, 0, XRES, YRES, backParams, back)

	loadingParams := &GE.Params{}
	loadingParams.LoadFromFile(F_TITLESCREEN + "/loading.txt")
	loading, err2 := GetEbitenImage(F_TITLESCREEN + "/loading.png")
	CheckErr(err2)
	TITLE_LoadingBar = GE.GetAnimationFromParams(0, 0, XRES, YRES, loadingParams, loading)

	nameParams := &GE.Params{}
	nameParams.LoadFromFile(F_TITLESCREEN + "/name.txt")
	name, err3 := GetEbitenImage(F_TITLESCREEN + "/name.png")
	CheckErr(err3)
	TITLE_Name = GE.GetAnimationFromParams(0, 0, XRES, XRES, nameParams, name)
	TITLE_Name.ScaleToOriginalSize()
	TITLE_Name.ScaleToX(XRES)
}

