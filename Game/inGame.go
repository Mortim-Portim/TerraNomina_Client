package Game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/mortim-portim/GraphEng/GE"
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

	SocialMenu        *SocialMenu
	ShowingSocialMenu bool

	AbilBars *GE.Group
}

func (i *InGame) Init() {
	Println("Initializing InGame")
	i.SocialMenu = GetNewSocialMenu(F_UI_ELEMENTS + "/dialog_pannel.png")

	healthBarEimg, err := GetEbitenImage(F_InGame + "/health_bar.png")
	CheckErr(err)
	staminaBarEimg, err := GetEbitenImage(F_InGame + "/stamina_bar.png")
	CheckErr(err)
	manaBarEimg, err := GetEbitenImage(F_InGame + "/mana_bar.png")
	CheckErr(err)
	Health := GE.GetAbilbar(healthBarEimg, 0, 0, XRES/3, 16, 5, 108, 6, color.RGBA{255, 0, 0, 255}, color.RGBA{15, 0, 0, 255})
	Stamina := GE.GetAbilbar(staminaBarEimg, XRES/3, 0, XRES/3, 16, 5, 108, 6, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 15, 0, 255})
	Mana := GE.GetAbilbar(manaBarEimg, XRES/3*2, 0, XRES/3, 16, 5, 108, 6, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 0, 15, 255})
	i.AbilBars = GE.GetGroup(Health, Stamina, Mana)

	for _, AttackParam := range TNE.Attacks {
		img, err := GetEbitenImage(fmt.Sprintf("%s/%s.png", F_SKILLS, AttackParam.GetName()))
		CheckErr(err)
		AttackParam.Init(img)
	}
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
	i.ShowingSocialMenu = false

	OwnPlayer.SetOnHealthChange(func(old, new float32) {
		i.AbilBars.Members[0].(*GE.Abilbar).Set(float64(OwnPlayer.HealthPercent()))
	})
	OwnPlayer.SetOnStaminaChange(func(old, new float32) {
		i.AbilBars.Members[1].(*GE.Abilbar).Set(float64(OwnPlayer.StaminaPercent()))
	})
	OwnPlayer.SetOnManaChange(func(old, new float32) {
		i.AbilBars.Members[2].(*GE.Abilbar).Set(float64(OwnPlayer.ManaPercent()))
	})
}
func (i *InGame) Stop(newState int) {
	Print("InGame      -------->")
}
func (i *InGame) Update() error {
	esc, esc_changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if esc_changed && !esc {
		i.OpenOptions()
	}

	i.UpdateOwnPlayerMovement()

	i.UpdateAttacking()

	//i.UpdateSocialMenu()

	i.sm.UpdateAll(false)
	i.sm.ActivePlayer.UpdateChanFromPlayer()
	i.sm.ActivePlayer.UpdateSyncVars(ClientManager)
	// smMsg, num := i.sm.Print(false)
	// if num > 0 {
	// 	Println(smMsg)
	// }
	i.sm.ResetActions()

	delay := time.Now().Sub(i.lastUpdate).Microseconds()
	i.lastUpdate = time.Now()
	i.isDelayed = float64(delay)/float64(i.meanDelay) > 1.2
	i.meanDelay = int(float64(i.meanDelay)*(9.0/10.0) + float64(delay)*(1.0/10.0))
	return nil
}
func (i *InGame) Draw(screen *ebiten.Image) {
	i.sm.Draw(screen)
	i.AbilBars.Draw(screen)

	if i.ShowingSocialMenu {
		i.SocialMenu.Draw(screen)
	}

	msg := fmt.Sprintf("Time: %v, TPS: %0.1f, Ping: %v", i.sm.Struct.CurrentTime, ebiten.CurrentTPS(), Client.Ping)
	if i.isDelayed {
		msg += fmt.Sprintf(", meanDelay: %v", i.meanDelay)
		//Toaster.New(fmt.Sprintf("%v/%v", i.meanDelay, delay), 6)
	}
	ebitenutil.DebugPrint(screen, msg)
}
func (i *InGame) UpdateAttacking() {
	prim, c1 := Keyli.GetMappedKeyState(attack_key_1)
	seco, c2 := Keyli.GetMappedKeyState(attack_key_2)
	tert, c3 := Keyli.GetMappedKeyState(attack_key_3)
	if prim && c1 {
		OwnPlayer.ChangeToAttack(0)
	}
	if seco && c2 {
		OwnPlayer.ChangeToAttack(1)
	}
	if tert && c3 {
		OwnPlayer.ChangeToAttack(2)
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		wxf, wyf := SmallWorld.Struct.GetTileOfCoordsFP(float64(x), float64(y))
		OwnPlayer.StartAttack(wxf, wyf, SmallWorld)
	}
}
func (i *InGame) UpdateSocialMenu() {
	if !i.ShowingSocialMenu {
		OwnPlayer.CheckNearbyDialogs(i.sm.Ents...)
	}
	interD, interC := Keyli.GetMappedKeyState(interaction_key)
	if interD && interC {
		if !i.ShowingSocialMenu && OwnPlayer.ShowsDialogSymbol {
			i.ShowingSocialMenu = true
			i.SocialMenu.Start(OwnPlayer.DialogEntity)
		} else if i.ShowingSocialMenu {
			i.ShowingSocialMenu = false
			i.SocialMenu.Stop()
		}
	}
	if i.ShowingSocialMenu {
		i.SocialMenu.Update()
	}
}
func (i *InGame) UpdateOwnPlayerMovement() {
	left, lC := Keyli.GetMappedKeyState(left_key_id)
	right, rC := Keyli.GetMappedKeyState(right_key_id)
	up, uC := Keyli.GetMappedKeyState(up_key_id)
	down, dC := Keyli.GetMappedKeyState(down_key_id)
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
		OwnPlayer.Move(0.02)
	}
	OwnPlayer.KeepMoving(moving)
}
func (i *InGame) Close() {
	i.parent.ChangeState(TITLESCREEN_STATE)
}
func (i *InGame) OpenOptions() {

}
