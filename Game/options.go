package Game

import (
	"fmt"
	"image/color"

	//"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetOptionsMenu(g *TerraNomina) *OptionsMenu {
	return &OptionsMenu{parent: g}
}

type OptionsMenu struct {
	parent   *TerraNomina
	oldState int

	drawobject *GE.Group
}

func (t *OptionsMenu) Init() {
	fmt.Println("Initializing OptionsMenu")

	volumetext := GE.GetTextImage("Volume", XRES*0.07, YRES*0.1, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)

	scrollbar := GE.GetStandardScrollbar(XRES*0.2, YRES*0.1, XRES*0.6, YRES*0.05, 0, 100, 100, GE.StandardFont)
	scrollbar.RegisterOnChange(func(scrollbar *GE.ScrollBar) {

	})

	t.drawobject = GE.GetGroup(volumetext, scrollbar)
	t.drawobject.Init(nil, nil)
}
func (t *OptionsMenu) Start(oldState int) {
	fmt.Print("--------> OptionsMenu\n")
	t.oldState = oldState
}
func (t *OptionsMenu) Stop(newState int) {
	fmt.Print("OptionsMenu -------->")
}
func (t *OptionsMenu) Update(screen *ebiten.Image) error {
	down, changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if changed && !down {
		t.GetBack()
	}

	screen.Fill(color.RGBA{168, 255, 68, 255})

	t.drawobject.Update(t.parent.GetCurrentFrame())
	t.drawobject.Draw(screen)

	return nil
}

func (t *OptionsMenu) GetBack() {
	t.parent.ChangeState(t.oldState)
}
