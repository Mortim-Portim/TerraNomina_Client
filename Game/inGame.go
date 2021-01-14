package Game

import (
	"fmt"

	//"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func GetInGame(g *TerraNomina) *InGame {
	return &InGame{parent: g}
}

type InGame struct {
	parent *TerraNomina

	left, right, up, down bool
	
	sm *TNE.SmallWorld
	ef *TNE.EntityFactory
}

func (i *InGame) Init() {
	fmt.Println("Initializing InGame")
	Keyli.MappIDToKey(left_key_id, ebiten.KeyLeft)
	Keyli.MappIDToKey(right_key_id, ebiten.KeyRight)
	Keyli.MappIDToKey(up_key_id, ebiten.KeyUp)
	Keyli.MappIDToKey(down_key_id, ebiten.KeyDown)
}
func (i *InGame) Start(oldState int) {
	fmt.Print("--------> InGame     \n")
	i.sm = SmallWorld
	i.ef = i.sm.Ef
	
	Soundtrack.Play(SOUNDTRACK_MAIN)
}
func (i *InGame) Stop(newState int) {
	fmt.Print("InGame      -------->")
}
func (i *InGame) Update(screen *ebiten.Image) error {
	i.left, _ = Keyli.GetMappedKeyState(left_key_id)
	i.right, _ = Keyli.GetMappedKeyState(right_key_id)
	i.up, _ = Keyli.GetMappedKeyState(up_key_id)
	i.down, _ = Keyli.GetMappedKeyState(down_key_id)
	
	moving := false
	if i.left || i.right || i.up || i.down {
		moving = true
	}
	
	if i.left {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_CHANGE_ORIENTATION_LEFT)
	}else if i.right {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_CHANGE_ORIENTATION_RIGHT)
	}else if i.up {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_CHANGE_ORIENTATION_UP)
	}else if i.down {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_CHANGE_ORIENTATION_DOWN)
	}
	
	if moving && !OwnPlayer.IsMoving() {
		OwnPlayer.Move()
	}
	OwnPlayer.UpdateAll(nil)
	i.sm.Draw(screen)
	
	msg := fmt.Sprintf(`TPS: %0.2f`, ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)
	return nil
}
