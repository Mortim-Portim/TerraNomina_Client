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
	generalTabU, err := GetButtonImg("general", true)
	CheckErr(err)
	recordingTabU, err := GetButtonImg("recording", true)
	CheckErr(err)
	generalTabD, err := GetButtonImg("general", false)
	CheckErr(err)
	recordingTabD, err := GetButtonImg("recording", false)
	CheckErr(err)
	
	scrollbarImg, err := GetEbitenImage(F_UI_ELEMENTS+"/scrollbar.png")
	CheckErr(err)
	scrollbarButtonImg, err := GetEbitenImage(F_UI_ELEMENTS+"/scrollbar_button.png")
	CheckErr(err)
	
	volumetext := GE.GetTextImage("Volume", XRES*0.07, YRES*0.15, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	volumeScrollbar := GE.GetImageScrollbar(XRES*0.2, YRES*0.15, XRES*0.6, YRES*0.05, scrollbarImg, scrollbarButtonImg, 0, 100, int(StandardVolume*100), GE.StandardFont)
	volumeScrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {
		pc := scrollbar.Current()
		StandardVolume = float64(pc)/100
		Soundtrack.SetVolume(StandardVolume)
	})
	
	recordText := GE.GetTextImage("Recording", XRES*0.07, YRES*0.15, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	recordButton,err := GetButton("checkbox", XRES*0.2, YRES*0.15, 0, 0 , false);CheckErr(err)
	recordButton.Img.ScaleToOriginalSize();recordButton.Img.ScaleToY(TITLESCREEN_BUTTON_HEIGHT_REL*YRES);recordButton.DrawDark = RecordAll
	recordButton.Img.SetMiddleY(YRES*0.175)
	recordButton.RegisterOnEvent(func(b *GE.Button){
		if !b.LPressed && !b.RPressed {
			b.DrawDark = !b.DrawDark
			RecordAll = b.DrawDark
		}
	})
	
	
	recordingTimeTxt := GE.GetTextImage("Time", XRES*0.07, YRES*0.3, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	recordingTimeScrollbar := GE.GetImageScrollbar(XRES*0.2, YRES*0.3, XRES*0.6, YRES*0.05, scrollbarImg, scrollbarButtonImg, 1, 30, int(RecordingLength), GE.StandardFont)
	recordingTimeScrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {
		pc := scrollbar.Current()
		RecordingLength = float64(pc)
		ResetRecorder()
	})
	recordingScaleTxt := GE.GetTextImage("Scale", XRES*0.07, YRES*0.45, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	recordingScaleScrollbar := GE.GetImageScrollbar(XRES*0.2, YRES*0.45, XRES*0.6, YRES*0.05, scrollbarImg, scrollbarButtonImg, 0, 10, int(RecordingScale*10), GE.StandardFont)
	recordingScaleScrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {
		pc := scrollbar.Current()
		RecordingScale = float64(pc)/10
		ResetRecorder()
	})
	
	TabViewUpdateAble := make([]GE.UpdateAble, 2)
	TabViewUpdateAble[0] = GE.GetGroup(volumetext, volumeScrollbar)
	TabViewUpdateAble[1] = GE.GetGroup(recordText, recordButton, recordingTimeTxt, recordingTimeScrollbar, recordingScaleTxt, recordingScaleScrollbar)

	tabPs := &GE.TabViewParams{
		X:0,
		Y:0,
		W:XRES,
		H:YRES,
		Imgs:[]*ebiten.Image{generalTabU, recordingTabU},
		Dark:[]*ebiten.Image{generalTabD, recordingTabD},
		Scrs: TabViewUpdateAble,
		TabH: TITLESCREEN_BUTTON_HEIGHT_REL*YRES,
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
