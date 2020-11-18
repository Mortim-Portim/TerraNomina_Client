package Game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

const PLAY_MENU_PLAY_BUTTON_WIDTH = 0.1

func GetPlayMenu(g *TerraNomina) *PlayMenu {
	return &PlayMenu{parent: g}
}

type PlayMenu struct {
	tabs    *GE.TabView
	playBtn *GE.Button

	parent   *TerraNomina
	oldState int
}

func (t *PlayMenu) Init() {
	fmt.Println("Initializing PlayMenu")
	characterTab, err := GE.LoadEbitenImg(F_PLAYMENU + "/character.png")
	CheckErr(err)
	serverTab, err := GE.LoadEbitenImg(F_PLAYMENU + "/server.png")
	CheckErr(err)
	playBtnImg, err := GE.LoadEbitenImg(F_PLAYMENU + "/play.png")
	CheckErr(err)
	w, h := playBtnImg.Size()
	rel := float64(h) / float64(w)
	W := XRES * PLAY_MENU_PLAY_BUTTON_WIDTH
	H := W * rel
	t.playBtn = GE.GetImageButton(playBtnImg, XRES-W-H*0.5, YRES-H*1.5, W, H)

	TabViewUpdateAble := make([]GE.UpdateAble, 2)
	TabViewUpdateAble[0] = GE.GetGroup()
	ipAddr := GE.GetEditText("ip:port", 10, 100, 100, 20, GE.StandardFont, color.RGBA{255, 255, 255, 255}, color.RGBA{120, 120, 120, 255})
	TabViewUpdateAble[1] = ipAddr

	t.playBtn.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			USER_INPUT_IP_ADDR = ipAddr.GetText()
			t.parent.ChangeState(CONNECTING_STATE)
		}
	})

	params := &GE.TabViewParams{Imgs: []*ebiten.Image{characterTab, serverTab}, Scrs: TabViewUpdateAble, Y: 0, W: XRES, H: YRES}
	t.tabs = GE.GetTabView(params)

	t.playBtn.Init(nil, nil)
	t.tabs.Init(nil, nil)
}
func (t *PlayMenu) Start(oldState int) {
	fmt.Print("--------> PlayMenu   \n")
	t.oldState = oldState
	t.playBtn.Start(nil, nil)
	t.tabs.Start(nil, nil)
}
func (t *PlayMenu) Stop(newState int) {
	fmt.Print("PlayMenu    -------->")
	t.playBtn.Stop(nil, nil)
	t.tabs.Stop(nil, nil)
}
func (t *PlayMenu) Update(screen *ebiten.Image) error {
	down, changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if changed && !down {
		t.GetBack()
	}

	t.playBtn.Update(t.parent.frame)
	t.tabs.Update(t.parent.frame)

	t.playBtn.Draw(screen)
	t.tabs.Draw(screen)

	return nil
}

func (t *PlayMenu) GetBack() {
	t.parent.ChangeState(t.oldState)
}
