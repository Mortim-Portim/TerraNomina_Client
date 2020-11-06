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

func getRace(race *Race, x, y float64) (group *GE.Group) {
	title := GE.GetTextImage(race.name, x+400, y+50, 50, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})
	stats := GE.GetTextImage(fmt.Sprintf("STR: %v DEX: %v INT: %v CHA: %v", race.attributes[0], race.attributes[1], race.attributes[2], race.attributes[3]), x+480, y+160, 30, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})

	subraces := make([]GE.UpdateAble, len(race.subraces))
	for i, subrace := range race.subraces {
		subraces[i] = GE.GetTextImage(subrace, x+220, y+160+float64(i*40), 30, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})
	}

	group = GE.GetGroup(append(subraces, title, stats)...)
	return
}

func (menu *CharacterMenu) Init(g *TerraNomina) {
	rx, ry := 100.0, 100.0

	racebackground, _ := GE.LoadImgObj(F_CHARACTERMENU+"/racetemplate.png", 0, 0, rx, ry, 0)
	racebackground.ScaleToOriginalSize()

	lbuttonimg, _ := GE.LoadImgObj(F_CHARACTERMENU+"/arrow.png", 0, 0, 20, 300, 0)
	lbuttonimg.ScaleToOriginalSize()
	leftbuton := GE.GetButton(lbuttonimg, lbuttonimg.Img)
	leftbuton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.ChangeRace(-1)
		}
	})

	rbuttonimg, _ := GE.LoadImgObj(F_CHARACTERMENU+"/arrow.png", 0, 0, 1010, 300, 180)
	rbuttonimg.ScaleToOriginalSize()
	rightbuton := GE.GetButton(rbuttonimg, rbuttonimg.Img)
	rightbuton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.ChangeRace(1)
		}
	})

	menu.racething = GE.GetGroup(racebackground, leftbuton, rightbuton)
	menu.racething.Init(nil, nil)

	for _, race := range Races {
		group := getRace(race, rx, ry)
		group.Init(nil, nil)
		menu.races = append(menu.races, group)
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

	screen.Fill(color.RGBA{128, 128, 128, 255})

	menu.racething.Draw(screen)
	menu.races[menu.currRace].Draw(screen)

	return nil
}

func (menu *CharacterMenu) GetBack() {
	menu.parent.ChangeState(TITLESCREEN_STATE)
}

func (menu *CharacterMenu) ChangeRace(delta int) {
	menu.currRace += delta

	if menu.currRace < 0 {
		menu.currRace = len(menu.races) - 1
	}

	if menu.currRace >= len(menu.races) {
		menu.currRace = 0
	}

	fmt.Printf("Changed to %v \n", menu.currRace)
}
