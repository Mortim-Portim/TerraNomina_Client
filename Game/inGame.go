package Game

import (
	"fmt"
	//"time"
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
	
	sm *TNE.SmallWorld
	ef *TNE.EntityFactory
}

func (i *InGame) Init() {
	fmt.Println("Initializing InGame")
}
func (i *InGame) Start(oldState int) {
	fmt.Print("--------> InGame     \n")
	i.sm = SmallWorld
	i.ef = i.sm.Ef
	
	//fmt.Println(i.sm.Print())
	
	Soundtrack.Play(SOUNDTRACK_MAIN)
}
func (i *InGame) Stop(newState int) {
	fmt.Print("InGame      -------->")
}
func (i *InGame) Update(screen *ebiten.Image) error {
	left, lC := Keyli.GetMappedKeyState(left_key_id)
	right, rC := Keyli.GetMappedKeyState(right_key_id)
	up, uC := Keyli.GetMappedKeyState(up_key_id)
	down, dC := Keyli.GetMappedKeyState(down_key_id)
	
	moving := false
	if left || right || up || down {
		moving = true
	}
	if left && (lC || (rC && !right) || (uC && !up) || (dC && !down)) {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_ORIENTATION_LEFT)
	}else if right && (rC || (lC && !left) || (uC && !up) || (dC && !down)) {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_ORIENTATION_RIGHT)
	}else if up && (uC || (rC && !right) || (lC && !left) || (dC && !down)) {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_ORIENTATION_UP)
	}else if down && (dC || (rC && !right) || (uC && !up) || (lC && !left)) {
		OwnPlayer.ChangeOrientation(TNE.ENTITY_ORIENTATION_DOWN)
	}
	
	if moving && !OwnPlayer.IsMoving() {
		OwnPlayer.Move()
	}
	OwnPlayer.KeepMoving(moving)
	OwnPlayer.UpdateAll(nil)
	
	SmallWorld.ActivePlayer.UpdateVarsFromPlayer()
	//st := time.Now()
	SmallWorld.ActivePlayer.UpdateSyncVars(ClientManager)
	Client.WaitForConfirmation()
	//fmt.Printf("Updating Vars took: %v\n", time.Now().Sub(st))
	if moving {
		x,y := OwnPlayer.IntPos()
		fmt.Println("OwnPl: ", x, ":", y)
	}
	i.sm.Draw(screen)
	
	msg := fmt.Sprintf(`TPS: %0.2f`, ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)
	return nil
}
