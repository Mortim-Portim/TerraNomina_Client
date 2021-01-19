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
	
	sending bool
	
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
	if left || right || up || down {moving = true}
	dir := TNE.GetNewDirection()
	if lC || rC || uC || dC {
		dir.L = left; dir.R = right; dir.U = up; dir.D = down
		dir.FromKeys()
		OwnPlayer.ChangeOrientation(dir)
	}
	
	if moving && !OwnPlayer.IsMoving() {
		OwnPlayer.Move()
	}
	OwnPlayer.KeepMoving(moving)
	
	SmallWorld.UpdateAll()
	
	if !i.sending {
		i.sending = true
		go func(){
			SmallWorld.ActivePlayer.UpdateVarsFromPlayer()
			SmallWorld.ActivePlayer.UpdateSyncVars(ClientManager)
			Client.WaitForConfirmation()
			i.sending = false
		}()
	}
	
	x,y := OwnPlayer.IntPos()
	fmt.Printf("%p: %v, %v\n", OwnPlayer, x, y)
	for _,pl := range(i.sm.Plys) {
		if pl.HasPlayer() {
			xp,yp := pl.Se.Entity.IntPos()
			fmt.Printf("%p: %v, %v\n", pl.Se.Entity, xp, yp)
		}
	}
	i.sm.Draw(screen)
	
	msg := fmt.Sprintf(`TPS: %0.2f`, ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)
	return nil
}
