package Game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/mortim-portim/TN_Engine/TNE"
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

	lastUpdate time.Time
	meanDelay  int
	isDelayed  bool
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
	i.lastUpdate = time.Now()
	i.meanDelay = 33288
}
func (i *InGame) Stop(newState int) {
	Print("InGame      -------->")
}
func (i *InGame) Update() error {
	left, lC := Keyli.GetMappedKeyState(left_key_id)
	right, rC := Keyli.GetMappedKeyState(right_key_id)
	up, uC := Keyli.GetMappedKeyState(up_key_id)
	down, dC := Keyli.GetMappedKeyState(down_key_id)
	esc, esc_changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if esc_changed && !esc {
		i.OpenOptions()
	}

	moving := false
	if left || right || up || down {
		moving = true
	}
	dir := TNE.GetNewDirection()
	if lC || rC || uC || dC {
		dir.L = left
		dir.R = right
		dir.U = up
		dir.D = down
		dir.FromKeys()
		OwnPlayer.ChangeOrientation(dir)
	}

	if moving && !OwnPlayer.IsMoving() {
		OwnPlayer.Move()
	}
	OwnPlayer.KeepMoving(moving)

	i.sm.ActivePlayer.UpdateChanFromPlayer()
	i.sm.ActivePlayer.UpdateSyncVars(ClientManager)
	i.sm.UpdateAll(false)
	OwnPlayer.CheckNearbyDialogs(i.sm.Ents...)

	smMsg, num := i.sm.Print(false)
	if num > 0 {
		Println(smMsg)
	}
	i.sm.ResetActions()

	delay := time.Now().Sub(i.lastUpdate).Microseconds()
	i.lastUpdate = time.Now()
	i.isDelayed = float64(delay)/float64(i.meanDelay) > 1.2
	i.meanDelay = int(float64(i.meanDelay)*(9.0/10.0) + float64(delay)*(1.0/10.0))
	return nil
}
func (i *InGame) Draw(screen *ebiten.Image) {
	i.sm.Draw(screen)
	msg := fmt.Sprintf("Time: %v, TPS: %0.1f, Ping: %v", i.sm.Struct.CurrentTime, ebiten.CurrentTPS(), Client.Ping)
	if i.isDelayed {
		msg += fmt.Sprintf(", meanDelay: %v", i.meanDelay)
		//Toaster.New(fmt.Sprintf("%v/%v", i.meanDelay, delay), 6)
	}
	ebitenutil.DebugPrint(screen, msg)
}
func (i *InGame) Close() {
	i.parent.ChangeState(TITLESCREEN_STATE)
}
func (i *InGame) OpenOptions() {

}
