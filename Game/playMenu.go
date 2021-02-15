package Game

import (
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
	Println("Initializing PlayMenu")
	characterTabU, err := GE.LoadEbitenImg(F_BUTTONS + "/character_u.png")
	CheckErr(err)
	serverTabU, err := GE.LoadEbitenImg(F_BUTTONS + "/server_u.png")
	CheckErr(err)
	characterTabD, err := GE.LoadEbitenImg(F_BUTTONS + "/character_d.png")
	CheckErr(err)
	serverTabD, err := GE.LoadEbitenImg(F_BUTTONS + "/server_d.png")
	CheckErr(err)
	
	t.playBtn, err = GE.LoadButton(F_BUTTONS + "/play_u.png", F_BUTTONS + "/play_d.png", 0,0, 0, 0, true)
	CheckErr(err)
	t.playBtn.Img.ScaleToOriginalSize(); t.playBtn.Img.ScaleToX(XRES * PLAY_MENU_PLAY_BUTTON_WIDTH)
	t.playBtn.Img.SetBottomRight(XRES-t.playBtn.Img.H, YRES-t.playBtn.Img.H)

	TabViewUpdateAble := make([]GE.UpdateAble, 2)
	TabViewUpdateAble[0] = GE.GetGroup()
	ipAddr := GE.GetEditText("ip:port", 10, 100, 100, 20, GE.StandardFont, color.RGBA{255, 255, 255, 255}, color.RGBA{120, 120, 120, 255})
	ipAddr.RegisterOnChange(func(et *GE.EditText) {
		StandardIP_TEXT = et.GetText()
	})
	ipAddr.SetText(StandardIP_TEXT)
	TabViewUpdateAble[1] = ipAddr

	t.playBtn.RegisterOnLeftEvent(func(b *GE.Button) {
		if !b.LPressed {
			USER_INPUT_IP_ADDR = ipAddr.GetText()
			t.parent.ChangeState(CONNECTING_STATE)
		}
	})

	params := &GE.TabViewParams{Imgs: []*ebiten.Image{characterTabU, serverTabU}, Dark: []*ebiten.Image{characterTabD, serverTabD}, Scrs: TabViewUpdateAble, Y: 0, W: XRES, H: YRES}
	t.tabs = GE.GetTabView(params)

	t.playBtn.Init(nil, nil)
	t.tabs.Init(nil, nil)
}
func (t *PlayMenu) Start(oldState int) {
	Print("--------> PlayMenu   \n")
	t.oldState = oldState
	t.playBtn.Start(nil, nil)
	t.tabs.Start(nil, nil)
}
func (t *PlayMenu) Stop(newState int) {
	Print("PlayMenu    -------->")
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
