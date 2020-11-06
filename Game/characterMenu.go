package Game

import (
	"fmt"
	"image/color"
	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

func GetCharacterMenu(parent *TerraNomina) (cm *CharacterMenu) {
	cm = &CharacterMenu{parent: parent}
	return
}

type CharacterMenu struct {
	parent *TerraNomina

	racething *GE.Group
	races     []*GE.Group
	currRace  int
}

type RaceImg []*GE.WObj

func getRace(name string, x, y float64) (group *GE.Group) {
	title := &Image{GE.GetTextImage(name, x, y, 50, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})}
	group = GE.GetGroup(title)
	return
}

func (menu *CharacterMenu) Init(g *TerraNomina) {
	rx, ry := 100.0, 100.0

	img, _ := GE.LoadImgObj(F_CHARACTERMENU+"/racetemplate.png", 700, 500, rx, ry, 0)
	racebackground := &Image{img}

	menu.racething = GE.GetGroup(racebackground)
	menu.racething.Init(nil, nil)

	menu.races = []*GE.Group{
		getRace("Elv", rx+300, ry+50),
	}

	for _, race := range menu.races {
		race.Init(nil, nil)
	}

	menu.currRace = 0
}

func (menu *CharacterMenu) Start(g *TerraNomina, lastState int) {
	fmt.Print("--------> CharacterMenu   \n")
}

func (menu *CharacterMenu) Stop(g *TerraNomina, nextState int) {
	fmt.Print("CharacterMenu ------>")
}

func (menu *CharacterMenu) Update(screen *ebiten.Image) error {
	down, changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if changed && !down {
		menu.GetBack()
	}

	menu.racething.Update(0)
	menu.races[menu.currRace].Update(0)

	menu.racething.Draw(screen)
	menu.races[menu.currRace].DrawFuncs[0](screen)

	return nil
}

func (menu *CharacterMenu) GetBack() {
	menu.parent.ChangeState(TITLESCREEN_STATE)
}
