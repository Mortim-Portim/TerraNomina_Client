package Game

import (
	"github.com/mortim-portim/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

const (
	TITLESCREEN_BUTTON_HEIGHT_REL = 0.08
	TITLESCREEN_BUTTON_XCOORD_REL = 0.9
)

func GetTitleScreen(g *TerraNomina) *TitleScreen {
	return &TitleScreen{parent: g}
}

type TitleScreen struct {
	Buttons *GE.Group
	parent  *TerraNomina
}

func (t *TitleScreen) Init() {
	Println("Initializing TitleScreen")
	H := TITLESCREEN_BUTTON_HEIGHT_REL * YRES
	X := TITLESCREEN_BUTTON_XCOORD_REL*XRES
	
	Play_B, err := GE.LoadButton(F_BUTTONS + "/play_u.png", F_BUTTONS + "/play_d.png", 0, TITLE_Name.H, H, H, true)
	CheckErr(err)
	Play_B.Img.ScaleToOriginalSize();Play_B.Img.ScaleToY(H);Play_B.Img.SetMiddleX(X)
	Character_B, err := GE.LoadButton(F_BUTTONS + "/character_u.png", F_BUTTONS + "/character_d.png", 0, TITLE_Name.H+H, H, H, true)
	CheckErr(err)
	Character_B.Img.ScaleToOriginalSize();Character_B.Img.ScaleToY(H);Character_B.Img.SetMiddleX(X)
	Options_B, err := GE.LoadButton(F_BUTTONS + "/options_u.png", F_BUTTONS + "/options_d.png", 0, TITLE_Name.H+H*2, H, H, true)
	CheckErr(err)
	Options_B.Img.ScaleToOriginalSize();Options_B.Img.ScaleToY(H);Options_B.Img.SetMiddleX(X)
	
	Play_B.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			t.parent.ChangeState(PLAY_MENU_STATE)
		}
	})
	Character_B.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			t.parent.ChangeState(SELRACE_STATE)
		}
	})
	Options_B.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			t.parent.ChangeState(OPTIONS_MENU_STATE)
		}
	})

	t.Buttons = GE.GetGroup(Play_B, Character_B, Options_B)
	t.Buttons.Init(nil, nil)
}
func (t *TitleScreen) Start(oldState int) {
	Print("--------> TitleScreen\n")
	t.Buttons.Start(nil, nil)
	Soundtrack.Play(SOUNDTRACK_MAIN)
}
func (t *TitleScreen) Stop(newState int) {
	Print("TitleScreen -------->")
	t.Buttons.Stop(nil, nil)
}
func (t *TitleScreen) Update(screen *ebiten.Image) error {
	TITLE_BackImg.Update(t.parent.frame)
	TITLE_Name.Update(t.parent.frame)
	TITLE_BackImg.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)

	t.Buttons.Update(t.parent.frame)
	t.Buttons.Draw(screen)
	return nil
}
