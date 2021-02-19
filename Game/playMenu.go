package Game

import (
	"fmt"
	"io/ioutil"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
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
	characterTabU, err := GetButtonImg("character", true)
	CheckErr(err)
	serverTabU, err := GetButtonImg("server", true)
	CheckErr(err)
	characterTabD, err := GetButtonImg("character", false)
	CheckErr(err)
	serverTabD, err := GetButtonImg("server", false)
	CheckErr(err)
	
	t.playBtn, err = GetButton("play", 0,0, 0, 0, true)
	CheckErr(err)
	t.playBtn.Img.ScaleToOriginalSize(); t.playBtn.Img.ScaleToX(XRES * PLAY_MENU_PLAY_BUTTON_WIDTH)
	t.playBtn.Img.SetBottomRight(XRES-t.playBtn.Img.H, YRES-t.playBtn.Img.H)

	TabViewUpdateAble := make([]GE.UpdateAble, 2)
	TabViewUpdateAble[0] = t.GetCharacterNameList()
	
	ipAddr := GE.GetEditText("ip:port", XRES/200, YRES*TITLESCREEN_BUTTON_HEIGHT_REL, YRES*TITLESCREEN_BUTTON_HEIGHT_REL, 25, GE.StandardFont, color.RGBA{255, 255, 255, 255}, color.RGBA{120, 120, 120, 255})
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

	params := &GE.TabViewParams{Imgs: []*ebiten.Image{characterTabU, serverTabU}, Dark: []*ebiten.Image{characterTabD, serverTabD}, Scrs: TabViewUpdateAble, Y: 0, W: XRES, H: YRES, TabH: YRES*TITLESCREEN_BUTTON_HEIGHT_REL}
	t.tabs = GE.GetTabView(params)

	t.playBtn.Init(nil, nil)
	t.tabs.Init(nil, nil)
}
func (t *PlayMenu) GetCharacterNameList() *GE.ScrollPanel {
	charFiles, err := GE.OSReadDir(F_CHARACTER);CheckErr(err)
	return GE.GetScrollPanelFromStrings(XRES/200, YRES*TITLESCREEN_BUTTON_HEIGHT_REL, XRES/2, YRES*(1-TITLESCREEN_BUTTON_HEIGHT_REL), YRES*TITLESCREEN_BUTTON_HEIGHT_REL, XRES/200, color.RGBA{0,0,0,255}, color.RGBA{255,255,255,255}, charFiles...)
}
func (t *PlayMenu) Start(oldState int) {
	Print("--------> PlayMenu   \n")
	t.oldState = oldState
	
	sm, err := TNE.GetSmallWorld(0, 0, XRES, YRES, F_TILES, F_STRUCTURES, F_ENTITY)
	CheckErr(err)
	sm.RegisterOnEntityChangeListeners()
	SmallWorld = sm	
	
	ple, err := sm.Ef.GetByName("Goblin")
	CheckErr(err)
	OwnPlayer = &TNE.Player{Entity:ple}
	
	scrollpannel := t.GetCharacterNameList()
	t.tabs.Screens.Set(scrollpannel, 0)
	for _,btn := range scrollpannel.Content() {
		btn.ChangeDrawDarkOnLeft = true
		btn.RegisterOnEvent(func(b *GE.Button){
				if !b.LPressed && !b.RPressed {
					name := b.Data.(string)
					data, err := ioutil.ReadFile(F_CHARACTER+"/"+name)
					if err == nil {
						char, err := TNE.LoadChar(data)
						if err == nil {
							ent, err := SmallWorld.Ef.GetFromCharacter(char)
							if err == nil {
								OwnPlayer.Entity = ent
								Toaster.New(fmt.Sprintf("Player %s loaded", OwnPlayer.Char.Name), FPS*1.5)
							}
						}
					}
				}
		})
	}
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
