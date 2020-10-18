package Game

import (
	"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
	"time"
	"fmt"
	"log"
	"errors"
	"os/signal"
	"os"
)

type TerraNomina struct {
	States map[int]GameState
	first bool
	
	frame, currentState int
	lastLoadingState, loadingState uint8
	initializing bool
	interrupt chan os.Signal
}
func (g *TerraNomina) Update(screen *ebiten.Image) error {
	if g.first {
		g.loadingState = 0
		g.lastLoadingState = 0
		g.frame = 0
		g.first = false
		g.initializing = true
		go g.Init()
	}
	if g.initializing {
		return g.Initializing(screen)
	}
	state, ok := g.States[g.currentState]
	if ok {
		return state.Update(screen)
	}
	return errors.New(fmt.Sprintf("Cannot update state %v, does not exist", g.currentState))
}
func (g *TerraNomina) Close() {
	state, ok := g.States[g.currentState]
	if ok {
		state.Stop(g)
	}
}






func (g *TerraNomina) Initializing(screen *ebiten.Image) error {
	TITLE_BackImg.Update(g.frame)
	TITLE_BackImg.DrawImageObj(screen)
	if g.lastLoadingState != g.loadingState {
		TITLE_LoadingBar.Update(int(g.loadingState))
		TITLE_Name.Update(int(g.loadingState))
		g.lastLoadingState = g.loadingState
	}
	TITLE_LoadingBar.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)
	return nil
}
func (g *TerraNomina) Init() {
	g.interrupt = make(chan os.Signal, 1)
	signal.Notify(g.interrupt, os.Interrupt)
	go func(){
		<-g.interrupt
		g.Close()
		log.Fatal("User Termination")
		return
	}()
	
	mt, err := GE.LoadSounds(RES+SOUNDTRACK_FILES+"/main")
	CheckErr(err)
	bt, err := GE.LoadSounds(RES+SOUNDTRACK_FILES+"/battle")
	CheckErr(err)
	MainTheme = mt; BattleTheme = bt
	
	for i := 0; i < 30; i ++ {
		g.loadingState = uint8(i)
		time.Sleep(time.Millisecond*100)
	}
	g.frame = 0
	newState := g.currentState
	g.currentState = -1
	g.ChangeState(newState)
	g.initializing = false
}
func (g *TerraNomina) ChangeState(newState int) {
	if _,ok := g.States[newState]; ok {
		if _,ok := g.States[g.currentState]; ok {
			g.States[g.currentState].Stop(g)
		}
		g.currentState = newState
		g.States[g.currentState].Start(g)
	}
}
type GameState interface {
	Start(g *TerraNomina)
	Stop(g *TerraNomina)
	Update(screen *ebiten.Image) error
}