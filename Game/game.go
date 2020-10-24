package Game

import (
	"marvin/GraphEng/GE"
	"marvin/GameConn/GC"
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
	defer func(){g.frame ++}()
	if Keyli != nil {
		Keyli.UpdateMapped()
	}
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
		state.Stop(g, -1)
	}
	Keyli.SaveConfig(F_KEYLI_MAPPER)
	Soundtrack.FadeOut()
	time.Sleep(time.Duration(float64(time.Second)*(GE.STANDARD_FADE_TIME+0.5)))
	fmt.Println()
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
	done := make(chan struct{})
	go func(){
		for i := 0; i <= 30; i ++ {
			g.loadingState = uint8(i)
			time.Sleep(time.Millisecond*100)
		}
		<-done
		g.frame = 0
		newState := g.currentState
		g.currentState = -1
		g.ChangeState(newState)
		g.initializing = false
	}()
	
	g.interrupt = make(chan os.Signal, 1)
	signal.Notify(g.interrupt, os.Interrupt)
	go func(){
		<-g.interrupt
		g.Close()
		log.Fatal("User Termination")
		return
	}()
	
	st, err := GE.LoadSoundTrack(F_SOUNDTRACK)
	CheckErr(err)
	Soundtrack = st
	st.Play(SOUNDTRACK_MAIN)
	
	Keyli = &GE.KeyLi{}
	Keyli.Reset()
	Keyli.LoadConfig(F_KEYLI_MAPPER)
	ESC_KEY_ID = Keyli.MappKey(ebiten.KeyEscape)
	ESC_KEY_ID = Keyli.MappKey(ebiten.KeyEscape)
	//Keyli.RegisterKeyEventListener(ESC_KEY_ID, func(l *GE.KeyLi, state bool){fmt.Printf("Esc is %v\n", state)})
	
	Client = GC.GetNewClient()
	ClientManager = GC.GetClientManager(Client)
	
	for _,state := range(g.States) {
		state.Init(g)
	}
	
	close(done)
}
func (g *TerraNomina) ChangeState(newState int) error {
	if _,ok := g.States[newState]; ok {
		if _,ok := g.States[g.currentState]; ok {
			g.States[g.currentState].Stop(g, newState)
		}
		g.States[newState].Start(g, g.currentState)
		g.currentState = newState
		return nil
	}
	return errors.New(fmt.Sprintf("Cannot change to state %v, does not exist", g.currentState))
}
type GameState interface {
	Init(g *TerraNomina)
	Start(g *TerraNomina, lastState int)
	Stop(g *TerraNomina, nextState int)
	Update(screen *ebiten.Image) error
}