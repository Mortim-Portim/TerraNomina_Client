package Game

import (
	"fmt"
	"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
)
const (
	TITLESCREEN_BUTTON_WIDTH = 1.0/8.0
	TITLESCREEN_BUTTON_HEIGHT_REL = 3.0/10.0
)

func GetTitleScreen(g *TerraNomina) *TitleScreen {
	return &TitleScreen{parent:g}
}
type TitleScreen struct {
	Buttons *GE.Group
	parent *TerraNomina
}

func (t *TitleScreen) Init(g *TerraNomina) {
	fmt.Println("Initializing TitleScreen")
	Play_i, err 	 := GE.LoadEbitenImg(RES+BUTTON_FILES+"/play.png"); CheckErr(err)
	Character_i, err := GE.LoadEbitenImg(RES+BUTTON_FILES+"/character.png"); CheckErr(err)
	Options_i, err	 := GE.LoadEbitenImg(RES+BUTTON_FILES+"/options.png"); CheckErr(err)
	
	w := TITLESCREEN_BUTTON_WIDTH*XRES; h := w*TITLESCREEN_BUTTON_HEIGHT_REL
	Play_B      := GE.GetImageButton(Play_i     , XRES-w*1.5, YRES/3    , w, h)
	Character_B := GE.GetImageButton(Character_i, XRES-w*1.5, YRES/3+h  , w, h)
	Options_B   := GE.GetImageButton(Options_i  , XRES-w*1.5, YRES/3+h*2, w, h)
	Play_B.RegisterOnLeftEvent(func(b *GE.Button) {if !b.LPressed {t.parent.ChangeState(PLAY_MENU_STATE)}})
	Character_B.RegisterOnLeftEvent(func(b *GE.Button) {if !b.LPressed {t.parent.ChangeState(CHARACTER_MENU_STATE)}})
	Options_B.RegisterOnLeftEvent(func(b *GE.Button) {if !b.LPressed {t.parent.ChangeState(OPTIONS_MENU_STATE)}})
	
	
	t.Buttons = GE.GetGroup(Play_B, Character_B, Options_B)
	t.Buttons.Init(nil,nil)
}
func (t *TitleScreen) Start(g *TerraNomina, oldState int) {
	fmt.Print("--------> TitleScreen\n")
	t.Buttons.Start(nil,nil)
}
func (t *TitleScreen) Stop(g *TerraNomina, newState int) {
	fmt.Print("TitleScreen -------->")
	t.Buttons.Stop(nil,nil)
}
func (t *TitleScreen) Update(screen *ebiten.Image) error {
	TITLE_BackImg.Update(t.parent.frame)
	TITLE_BackImg.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)
	
	t.Buttons.Update(t.parent.frame)
	t.Buttons.Draw(screen)
	return nil
}