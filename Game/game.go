package Game

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
)

type TerraNomina struct {
	States map[int]GameState
	first  bool

	doneSavingSC chan bool
	currentSaveFileSC string
	
	frame, currentState            int
	loadingState uint8
	initializing                   bool
	interrupt                      chan os.Signal
}

func (g *TerraNomina) Update(screen *ebiten.Image) error {
	defer func() {
		g.frame++ 
		Toaster.Update(g.frame)
		Toaster.Draw(screen)
	}()
	if Keyli != nil {
		Keyli.UpdateMapped()
	}
	if g.first {
		g.loadingState = 0
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
		err := state.Update(screen)
		if RecordAll {
			RecorderLock.Lock()
			Recorder.NextFrame(screen)
			RecorderLock.Unlock()
			down,chng := Keyli.GetMappedKeyState(record_key_id)
			if down && chng && !Recorder.IsSaving() {
				g.currentSaveFileSC = fmt.Sprintf("%s_%s", RecordingFile, GE.GetTime())
				go g.OnDoneSaving()
				RecorderLock.Lock()
				Recorder.Save(g.currentSaveFileSC, g.doneSavingSC)
				RecorderLock.Unlock()
			}
		}
		return err
	}
	return errors.New(fmt.Sprintf("Cannot update state %v, does not exist", g.currentState))
}
func (g *TerraNomina) Close() {
	state, ok := g.States[g.currentState]
	if ok {
		state.Stop(-1)
	}
	Keyli.SaveConfig(F_KEYLI_MAPPER)
	Soundtrack.FadeOut()
	VarsToParams()
	PARAMETER.SaveToFile(F_Params)
	
	time.Sleep(time.Duration(float64(time.Second) * (GE.STANDARD_FADE_TIME + 0.5)))
	fmt.Println()
	GE.CloseLogFile()
}

func (g *TerraNomina) Initializing(screen *ebiten.Image) error {
	TITLE_BackImg.Update(g.frame)
	TITLE_BackImg.DrawImageObj(screen)
	
	TITLE_LoadingBar.SetTo(int(g.loadingState))
	
	TITLE_LoadingBar.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)
	return nil
}
func (g *TerraNomina) Init() {
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			g.loadingState = uint8(i)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		<-done
		g.frame = 0
		newState := g.currentState
		g.currentState = -1
		g.ChangeState(newState)
		g.initializing = false
	}()

	g.interrupt = make(chan os.Signal, 1)
	signal.Notify(g.interrupt, os.Interrupt)
	go func() {
		<-g.interrupt
		g.Close()
		log.Fatal("User Termination")
		return
	}()
	
	g.doneSavingSC = make(chan bool)
	
	Toaster.New("Loading soundtrack", FPS*2)
	st, err := GE.LoadSoundTrack(F_SOUNDTRACK, 1)
	CheckErr(err)
	Soundtrack = st
	Soundtrack.OnFinished = func() {
		Soundtrack.NextTrack = Soundtrack.GetRandomTrack(time.Now().UnixNano())
	}
	Soundtrack.Play(SOUNDTRACK_MAIN)
	Soundtrack.SetVolume(StandardVolume)
	
	Toaster.New("Loading keyboard layout", FPS*2)
	Keyli = &GE.KeyLi{}
	Keyli.Reset()
	
	left_key_id = 		Keyli.MappKey(ebiten.KeyA)
	right_key_id = 		Keyli.MappKey(ebiten.KeyD)
	up_key_id = 		Keyli.MappKey(ebiten.KeyW)
	down_key_id = 		Keyli.MappKey(ebiten.KeyS)
	ESC_KEY_ID = 		Keyli.MappKey(ebiten.KeyEscape)
	record_key_id = 	Keyli.MappKey(ebiten.KeyR)
	
	Keyli.LoadConfig(F_KEYLI_MAPPER)
	//Keyli.RegisterKeyEventListener(ESC_KEY_ID, func(l *GE.KeyLi, state bool){fmt.Printf("Esc is %v\n", state)})
	
	Toaster.New("Creating client", FPS*2)
	Client = GC.GetNewClient()
	ClientManager = GC.GetClientManager(Client)
	
	Toaster.New("Creating recorder", FPS*2)
	RecordAll = true
	RecordingFile = "./screencapture"
	ResetRecorder()

	Toaster.New("Initializing gamestates", FPS*2)
	for _, state := range g.States {
		state.Init()
	}
	
	Toaster.New("Finished", FPS*1)
	close(done)
}
func (g *TerraNomina) OnDoneSaving() {
	<-g.doneSavingSC
	Toaster.New(fmt.Sprintf("Saved to %s", g.currentSaveFileSC), FPS*3)
}
func (g *TerraNomina) ChangeState(newState int) error {
	if _, ok := g.States[newState]; ok {
		if _, ok := g.States[g.currentState]; ok {
			g.States[g.currentState].Stop(newState)
		}
		g.States[newState].Start(g.currentState)
		g.currentState = newState
		return nil
	}
	return errors.New(fmt.Sprintf("Cannot change to state %v, does not exist", g.currentState))
}
func (g *TerraNomina) GetCurrentFrame() int {
	return g.frame
}

type GameState interface {
	Init()
	Start(lastState int)
	Stop(nextState int)
	Update(screen *ebiten.Image) error
}
