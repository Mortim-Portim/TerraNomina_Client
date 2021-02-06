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
	
	oldState int
	
	sending bool
	
	sm *TNE.SmallWorld
	ef *TNE.EntityFactory
}

func (i *InGame) Init() {
	Println("Initializing InGame")
}
func (i *InGame) Start(oldState int) {
	Print("--------> InGame     \n")
	i.oldState = oldState
	i.sm = SmallWorld
	i.ef = i.sm.Ef
	
	go func() {
		<-ServerClosing
		i.Close()
	}()
	
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
	esc, esc_changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if esc_changed && !esc {
		i.OpenOptions()
	}
	
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
	
	i.sm.UpdateAll(false)
	
	if !i.sending {
		i.sending = true
		go func(){
			i.sm.ActivePlayer.UpdateChanFromPlayer()
			i.sm.ActivePlayer.UpdateSyncVars(ClientManager)
			Client.WaitForConfirmation()
			i.sending = false
		}()
	}
	
	x,y := OwnPlayer.IntPos()
	Printf("%v; %v; ", x, y)
	for _,ent := range(i.sm.Ents) {
		if ent.HasEntity() {
			xp,yp := ent.Entity.IntPos()
			Printf("%v; %v; ", xp, yp)
		}
	}
	Println()
	i.sm.Draw(screen)
	
	msg := fmt.Sprintf("TPS: %0.1f, Ping: %v", ebiten.CurrentTPS(), Client.Ping)
	ebitenutil.DebugPrint(screen, msg)
	return nil
}

func (i *InGame) Close() {
	i.parent.ChangeState(TITLESCREEN_STATE)
}
func (i *InGame) OpenOptions() {
	
}
