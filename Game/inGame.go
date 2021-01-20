package Game

import (
	"github.com/mortim-portim/TN_Engine/TNE"
	"fmt"
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
	Println("Initializing InGame")
}
func (i *InGame) Start(oldState int) {
	Print("--------> InGame     \n")
	i.sm = SmallWorld
	i.ef = i.sm.Ef
	
	//fmt.Println(i.sm.Print())
	
	Soundtrack.Play(SOUNDTRACK_MAIN)
}
func (i *InGame) Stop(newState int) {
	Print("InGame      -------->")
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
	
	i.sm.UpdateAll()
	
	if !i.sending {
		i.sending = true
		go func(){
			i.sm.ActivePlayer.UpdateVarsFromPlayer()
			i.sm.ActivePlayer.UpdateSyncVars(ClientManager)
			Client.WaitForConfirmation()
			i.sending = false
		}()
	}
	
//	x,y := OwnPlayer.IntPos()
//	Printf("%v; %v; ", x, y)
//	for _,pl := range(i.sm.Plys) {
//		if pl.HasPlayer() {
//			xp,yp := pl.Se.Entity.IntPos()
//			Printf("%v; %v; ", xp, yp)
//		}
//	}
//	Println()
	//Println(i.sm.Struct.ObjMat.Print())
	i.sm.Draw(screen)
	
	msg := fmt.Sprintf("TPS: %0.1f, Ping: %v", ebiten.CurrentTPS(), Client.Ping)
	ebitenutil.DebugPrint(screen, msg)
	return nil
}
