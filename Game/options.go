package Game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetOptionsMenu(g *TerraNomina) *OptionsMenu {
	return &OptionsMenu{parent: g}
}

type OptionsMenu struct {
	parent   *TerraNomina
	oldState int

	tabs *GE.TabView
}

func (t *OptionsMenu) Init() {
	Println("Initializing OptionsMenu")
	generalTabU, err := GE.LoadEbitenImg(F_BUTTONS + "/general_u.png")
	CheckErr(err)
	recordingTabU, err := GE.LoadEbitenImg(F_BUTTONS + "/recording_u.png")
	CheckErr(err)
	generalTabD, err := GE.LoadEbitenImg(F_BUTTONS + "/general_d.png")
	CheckErr(err)
	recordingTabD, err := GE.LoadEbitenImg(F_BUTTONS + "/recording_d.png")
	CheckErr(err)
	
	volumetext := GE.GetTextImage("Volume", XRES*0.07, YRES*0.15, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	volumeScrollbar, err := GE.GetImageScrollbarFromPath(XRES*0.2, YRES*0.15, XRES*0.6, YRES*0.05, F_UI_ELEMENTS+"/scrollbar.png", F_UI_ELEMENTS+"/scrollbar_button.png", 0, 100, int(StandardVolume*100), GE.StandardFont)
	CheckErr(err)
	volumeScrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {
		pc := scrollbar.Current()
		StandardVolume = float64(pc)/100
		Soundtrack.SetVolume(StandardVolume)
	})
	
	recordingTimeTxt := GE.GetTextImage("Time", XRES*0.07, YRES*0.15, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	recordingTimeScrollbar, err := GE.GetImageScrollbarFromPath(XRES*0.2, YRES*0.15, XRES*0.6, YRES*0.05, F_UI_ELEMENTS+"/scrollbar.png", F_UI_ELEMENTS+"/scrollbar_button.png", 1, 30, int(RecordingLength), GE.StandardFont)
	CheckErr(err)
	recordingTimeScrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {
		pc := scrollbar.Current()
		RecordingLength = float64(pc)
		ResetRecorder()
	})
	recordingScaleTxt := GE.GetTextImage("Scale", XRES*0.07, YRES*0.3, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	recordingScaleScrollbar, err := GE.GetImageScrollbarFromPath(XRES*0.2, YRES*0.3, XRES*0.6, YRES*0.05, F_UI_ELEMENTS+"/scrollbar.png", F_UI_ELEMENTS+"/scrollbar_button.png", 0, 10, int(RecordingScale*10), GE.StandardFont)
	CheckErr(err)
	recordingScaleScrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {
		pc := scrollbar.Current()
		RecordingScale = float64(pc)/10
		ResetRecorder()
	})
	
	TabViewUpdateAble := make([]GE.UpdateAble, 2)
	TabViewUpdateAble[0] = GE.GetGroup(volumetext, volumeScrollbar)
	TabViewUpdateAble[1] = GE.GetGroup(recordingTimeTxt, recordingTimeScrollbar, recordingScaleTxt, recordingScaleScrollbar)

	tabPs := &GE.TabViewParams{
		X:0,
		Y:0,
		W:XRES,
		H:YRES,
		Imgs:[]*ebiten.Image{generalTabU, recordingTabU},
		Dark:[]*ebiten.Image{generalTabD, recordingTabD},
		Scrs: TabViewUpdateAble,
	}
	t.tabs = GE.GetTabView(tabPs)
	t.tabs.Init(nil, nil)
}
func (t *OptionsMenu) Start(oldState int) {
	Print("--------> OptionsMenu\n")
	t.oldState = oldState
}
func (t *OptionsMenu) Stop(newState int) {
	Print("OptionsMenu -------->")
}
func (t *OptionsMenu) Update(screen *ebiten.Image) error {
	down, changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if changed && !down {
		t.GetBack()
	}
	screen.Fill(color.RGBA{168, 255, 68, 255})
	
	t.tabs.Update(t.parent.GetCurrentFrame())
	t.tabs.Draw(screen)
	return nil
}

func (t *OptionsMenu) GetBack() {
	t.parent.ChangeState(t.oldState)
}
