package Game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

const (
	TITLESCREEN_BUTTON_HEIGHT_REL = 0.08
	TITLESCREEN_BUTTON_XCOORD_REL = 0.9
)

func GetTitleScreen(g *TerraNomina) *TitleScreen {
	return &TitleScreen{parent: g}
}

type TitleScreen struct {
	//TimeDrawer *GE.TimeDrawer
	Buttons *GE.Group
	parent  *TerraNomina
}

func (t *TitleScreen) Init() {
	Println("Initializing TitleScreen")
	H := TITLESCREEN_BUTTON_HEIGHT_REL * YRES
	X := TITLESCREEN_BUTTON_XCOORD_REL * XRES

	Play_B, err := GetButton("play", 0, TITLE_Name.H+H, H, H, true)
	CheckErr(err)
	Play_B.Img.ScaleToOriginalSize()
	Play_B.Img.ScaleToY(H)
	Play_B.Img.SetMiddleX(X)
	Character_B, err := GetButton("character", 0, TITLE_Name.H+H*2, H, H, true)
	CheckErr(err)
	Character_B.Img.ScaleToOriginalSize()
	Character_B.Img.ScaleToY(H)
	Character_B.Img.SetMiddleX(X)
	Options_B, err := GetButton("options", 0, TITLE_Name.H+H*3, H, H, true)
	CheckErr(err)
	Options_B.Img.ScaleToOriginalSize()
	Options_B.Img.ScaleToY(H)
	Options_B.Img.SetMiddleX(X)

	Play_B.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			t.parent.ChangeState(PLAY_MENU_STATE)
		}
	})
	Character_B.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			t.parent.ChangeState(SELSTATS_STATE)
		}
	})
	Options_B.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			t.parent.ChangeState(OPTIONS_MENU_STATE)
		}
	})

	// back, err := GetEbitenImage(F_UI_ELEMENTS + "/timeDrawerBack.png")
	// CheckErr(err)
	// front, err := GetEbitenImage(F_UI_ELEMENTS + "/timeDrawerFront.png")
	// CheckErr(err)
	// sun, err := GetEbitenImage(F_UI_ELEMENTS + "/timeDrawerSun.png")
	// CheckErr(err)
	// moon, err := GetEbitenImage(F_UI_ELEMENTS + "/timeDrawerMoon.png")
	// CheckErr(err)
	//t.TimeDrawer = GE.GetTimeDrawer(back, front, sun, moon, 0, TITLE_Name.H, XRES*0.2, XRES*0.2*(120.0/220.0))
	//t.TimeDrawer.Percent = 0.0

	t.Buttons = GE.GetGroup(Play_B, Character_B, Options_B)
	t.Buttons.Init(nil, nil)
}
func (t *TitleScreen) Start(oldState int) {
	Print("--------> TitleScreen\n")
	t.Buttons.Start(nil, nil)
}
func (t *TitleScreen) Stop(newState int) {
	Print("TitleScreen -------->")
	t.Buttons.Stop(nil, nil)
}
func (t *TitleScreen) Update() error {
	TITLE_BackImg.Update(t.parent.frame)
	TITLE_Name.Update(t.parent.frame)
	t.Buttons.Update(t.parent.frame)

	// t.TimeDrawer.Percent += 0.01
	// if t.TimeDrawer.Percent > 2.0 {
	// 	t.TimeDrawer.Percent = 0.0
	// }
	// t.TimeDrawer.Update(0)
	return nil
}
func (t *TitleScreen) Draw(screen *ebiten.Image) {
	TITLE_BackImg.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)
	t.Buttons.Draw(screen)
	//t.TimeDrawer.Draw(screen)
}
