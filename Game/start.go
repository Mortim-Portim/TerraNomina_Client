package Game

import (
	"flag"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
)

var WRITELOGFILE = flag.String("log", "log.txt", "name of the logfile")
var AUTOMOVE = flag.Bool("mv", false, "presses D and A")
var WINDOWED = flag.Bool("win", true, "starts Terra Nomina in window mode")

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func StartGame(g ebiten.Game) {
	defer onUnexpectedError(g)
	icons, err := GE.InitIcons(F_ICONS, ICON_SIZES, ICON_FORMAT)
	CheckErr(err)
	ebiten.SetWindowIcon(icons)
	ebiten.SetWindowTitle("Terra Nomina")
	if !*WINDOWED {
		ebiten.SetFullscreen(true)
	} else {
		ebiten.SetWindowSize(int(XRES), int(YRES))
	}
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(int(FPS))
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(g); err != nil {
		fmt.Printf("Error: %v\n", err)
		g.(*TerraNomina).Close()
		CheckErr(err)
	}
	g.(*TerraNomina).Close()
}

func onUnexpectedError(g ebiten.Game) {
	if r := recover(); r != nil {
		fmt.Printf("unexpected Error: %v\n", r)
		g.(*TerraNomina).Close()
		panic(fmt.Sprintf("unexpected Error: %v", r))
	}
}

func Start() {
	flag.Parse()
	GE.StartProfiling(cpuprofile)
	GE.SetLogFile(RES + "/" + *WRITELOGFILE)
	GE.Init("", FPS)
	GC.InitSyncVarStandardTypes()

	x, y := ebiten.ScreenSizeInFullscreen()
	if *WINDOWED {
		x = 960
		y = 540
	}
	InitParams(F_Params, x, y)

	GE.MOVE_A_D = *AUTOMOVE

	tn := &TerraNomina{first: true, States: make(map[int]GameState)}
	tn.States[TITLESCREEN_STATE] = GetTitleScreen(tn)
	tn.States[OPTIONS_MENU_STATE] = GetOptionsMenu(tn)
	tn.States[PLAY_MENU_STATE] = GetPlayMenu(tn)
	tn.States[CONNECTING_STATE] = GetConnecting(tn)
	tn.States[INGAME_STATE] = GetInGame(tn)
	tn.States[SELRACE_STATE] = GetSelectRaceMenu(tn)
	tn.States[SELCLASS_STATE] = GetSelectClassMenu(tn)
	tn.States[SELSTATS_STATE] = GetSelectStatsMenu(tn)
	//tn.States[TEST_STATE] = getTestMenu(tn)

	Toaster = GE.GetNewToaster(XRES, YRES, 0.5, 0.04, GE.StandardFont, color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 255})

	SetupCharacterMenu()
	StartGame(tn)
}

func Println(ps ...interface{}) {
	out := fmt.Sprintln(ps...)
	GE.LogToFile(out)
	fmt.Print(out)
}
func Print(ps ...interface{}) {
	out := fmt.Sprint(ps...)
	GE.LogToFile(out)
	fmt.Print(out)
}
func Printf(s string, ps ...interface{}) {
	out := fmt.Sprintf(s, ps...)
	GE.LogToFile(out)
	fmt.Print(out)
}

func CheckErr(err error) {
	//if err != nil {
	//
	//}
	GE.ShitImDying(err)
}
