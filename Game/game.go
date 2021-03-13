package Game

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
)

type TerraNomina struct {
	States map[int]GameState
	first  bool

	doneSavingSC, StopPrintSavingTo chan bool
	currentSaveFileSC               string
	HasUpdatedLastFrame             bool

	frame, currentState int
	loadingState        uint8
	initializing        bool
	interrupt           chan os.Signal
}

func (g *TerraNomina) Update() error {
	defer func() {
		g.frame++
		Toaster.Update(g.frame)
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
		return nil
	}
	state, ok := g.States[g.currentState]
	if ok {
		g.HasUpdatedLastFrame = true
		return state.Update()
	}
	return errors.New(fmt.Sprintf("Cannot update state %v, does not exist", g.currentState))
}
func (g *TerraNomina) Draw(screen *ebiten.Image) {
	defer func() {
		Toaster.Draw(screen)
	}()
	if g.initializing {
		g.Initializing(screen)
		msg := fmt.Sprintf("TPS: %0.1f, FPS: %0.1f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, msg)
		return
	}
	state, ok := g.States[g.currentState]
	if ok {
		state.Draw(screen)
		if g.HasUpdatedLastFrame {
			g.HasUpdatedLastFrame = false
			g.CheckRecorder(screen)
		}
		msg := fmt.Sprintf("TPS: %0.1f, FPS: %0.1f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, msg)
	}
}
func (g *TerraNomina) CheckRecorder(screen *ebiten.Image) {
	if RecordAll && screen != nil {
		RecorderLock.Lock()
		Recorder.NextFrame(screen)
		RecorderLock.Unlock()
		screenShotDown, screenShotChng := Keyli.GetMappedKeyState(record_key_1_id, screenshot_key_2_id)
		if screenShotDown && screenShotChng && !Recorder.IsSaving() {
			fileN := fmt.Sprintf("%s_%s", ScreenShotFile, GE.GetTime())
			go Recorder.SaveScreenShot(fileN)
			Toaster.New(fmt.Sprintf("Saved Image to %s", fileN), int(FPS*2.0), nil)
		}
		down, chng := Keyli.GetMappedKeyState(record_key_1_id, record_key_2_id)
		if down && chng && !Recorder.IsSaving() {
			g.currentSaveFileSC = fmt.Sprintf("%s_%s", ScreenCaptureFile, GE.GetTime())
			Toaster.New(fmt.Sprintf("Saving Video to %s...", g.currentSaveFileSC), 0, g.StopPrintSavingTo)
			go g.OnDoneSaving()
			RecorderLock.Lock()
			Recorder.Save(g.currentSaveFileSC, g.doneSavingSC)
			RecorderLock.Unlock()
		}
	}
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
	GE.StopProfiling(cpuprofile, memprofile)
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
		log.Println("User Interrupt")
		g.Close()
		log.Fatal("User Termination")
		return
	}()

	g.doneSavingSC = make(chan bool)
	g.StopPrintSavingTo = make(chan bool)

	Toaster.New("Loading soundtrack", int(FPS*2), nil)
	st, err := GE.LoadSoundTrack(F_SOUNDTRACK, 1)
	CheckErr(err)
	Soundtrack = st
	Soundtrack.OnFinished = func() {
		Soundtrack.NextTrack = Soundtrack.GetRandomTrack(time.Now().UnixNano())
	}
	go func() {
		Soundtrack.SetVolume(0)
		Soundtrack.Play(SOUNDTRACK_MAIN)
		time.Sleep(time.Duration(float64(time.Second) * (GE.STANDARD_FADE_TIME + 0.1)))
		Soundtrack.GetCurrent().Rewind()
		Soundtrack.SetVolume(StandardVolume)
	}()

	Toaster.New("Loading keyboard layout", int(FPS*2), nil)
	Keyli = &GE.KeyLi{}
	Keyli.Reset()

	left_key_id = Keyli.MappKey(ebiten.KeyA)
	right_key_id = Keyli.MappKey(ebiten.KeyD)
	up_key_id = Keyli.MappKey(ebiten.KeyW)
	down_key_id = Keyli.MappKey(ebiten.KeyS)
	ESC_KEY_ID = Keyli.MappKey(ebiten.KeyEscape)
	record_key_1_id = Keyli.MappKey(ebiten.KeyR)
	record_key_2_id = Keyli.MappKey(ebiten.KeyControl)
	screenshot_key_2_id = Keyli.MappKey(ebiten.KeyAlt)
	interaction_key = Keyli.MappKey(ebiten.KeyI)
	attack_key_1 = Keyli.MappKey(ebiten.Key1)
	attack_key_2 = Keyli.MappKey(ebiten.Key2)
	attack_key_3 = Keyli.MappKey(ebiten.Key3)

	Keyli.LoadConfig(F_KEYLI_MAPPER)
	//Keyli.RegisterKeyEventListener(ESC_KEY_ID, func(l *GE.KeyLi, state bool){fmt.Printf("Esc is %v\n", state)})

	Toaster.New("Creating client", int(FPS*2), nil)
	Client = GC.GetNewClient()
	ClientManager = GC.GetClientManager(Client)

	Toaster.New("Creating recorder", int(FPS*2), nil)
	ScreenCaptureFile = "./screencapture"
	ScreenShotFile = "./screenshot"
	ResetRecorder()

	Toaster.New("Initializing gamestates", int(FPS*2), nil)
	for _, state := range g.States {
		state.Init()
	}

	Toaster.New("Finished", int(FPS*1), nil)

	close(done)
}
func (g *TerraNomina) OnDoneSaving() {
	<-g.doneSavingSC
	g.StopPrintSavingTo <- true
	Toaster.New(fmt.Sprintf("Saved to %s", g.currentSaveFileSC), int(FPS*2), nil)
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
	Update() error
	Draw(screen *ebiten.Image)
}
