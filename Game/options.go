package Game

import (
	"fmt"
	//"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
)

func GetOptionsMenu(g *TerraNomina) *OptionsMenu {
	return &OptionsMenu{parent: g}
}

type OptionsMenu struct {
	parent   *TerraNomina
	oldState int
}

func (t *OptionsMenu) Init() {
	fmt.Println("Initializing OptionsMenu")

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

	return nil
}

func (t *OptionsMenu) GetBack() {
	t.parent.ChangeState(t.oldState)
}
