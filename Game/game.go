package Game

import (
	"github.com/hajimehoshi/ebiten"
	"time"
	"fmt"
	"errors"
)

type TerraNomina struct {
	States map[int]GameState
	first bool
	
	frame, currentState int
	lastLoadingState, loadingState uint8
	initializing bool
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
	for i := 0; i < 100; i ++ {
		g.loadingState = uint8(i)
		time.Sleep(time.Millisecond*10*5)
	}
	g.initializing = false
	g.frame = 0
}
func (g *TerraNomina) ChangeState(newState int) {
	if _,ok := g.States[newState]; ok {
		g.States[g.currentState].Stop(g)
		g.currentState = newState
		g.States[g.currentState].Start(g)
	}
}
type GameState interface {
	Start(g *TerraNomina)
	Stop(g *TerraNomina)
	Update(screen *ebiten.Image) error
}