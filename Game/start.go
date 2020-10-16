package Game

import (
	"github.com/hajimehoshi/ebiten"
	"marvin/GraphEng/GE"
	fonts "marvin/TerraNomina_Client/.res/Fonts"
)

func StartGame(g ebiten.Game) {
	icons, err := GE.InitIcons(RES+ICON_FILES, ICON_SIZES, ICON_FORMAT)
	CheckErr(err)
	ebiten.SetWindowIcon(icons)
	ebiten.SetWindowTitle("Terra Nomina")
	ebiten.SetFullscreen(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(FPS)
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(g); err != nil {
		CheckErr(err)
	}
	GE.CloseLogFile()
}

func Start() {
	GE.Init("")
	GE.StandardFont = GE.ParseFontFromBytes(fonts.MONO_TTF)
	GE.SetLogFile(RES+"/log.txt")
	
	InitParams(RES+"/params.txt")
	
	tn := &TerraNomina{first:true, States:make(map[int]GameState)}
	
	StartGame(tn)
}

func CheckErr(err error) {
	//if err != nil {
	//	
	//}
	GE.ShitImDying(err)
}