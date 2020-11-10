package Game

import (
	"fmt"
	"image/color"
	"io"
	"marvin/GraphEng/GE"
	"os"

	"github.com/hajimehoshi/ebiten"
)

func GetCharacterMenu(parent *TerraNomina) (cm *CharacterMenu) {
	cm = &CharacterMenu{parent: parent}
	return
}

type CharacterMenu struct {
	parent *TerraNomina
	state  int

	buttons *GE.Group

	//Races
	rbackground []*GE.ImageObj
	races       []*GE.Group
	currRace    int

	//Classes
	classthing *GE.Group
	class      []*GE.Group
	currClass  int
}

func (menu *CharacterMenu) Init(g *TerraNomina) {
	lbuttonimg, _ := GE.LoadImgObj(F_CHARACTERMENU+"/arrow.png", XRES*0.04, YRES*0.04, XRES*0.01, YRES*0.46, 0)
	leftbutton := GE.GetButton(lbuttonimg, lbuttonimg.Img)

	rbuttonimg, _ := GE.LoadImgObj(F_CHARACTERMENU+"/arrow.png", XRES*0.04, YRES*0.04, XRES*0.95, YRES*0.46, 180)
	rightbutton := GE.GetButton(rbuttonimg, rbuttonimg.Img)

	nextbutton := GE.GetTextButton("Next", "", GE.StandardFont, XRES*0.05, YRES*0.82, YRES*0.1, color.Black, &color.RGBA{168, 255, 68, 255})
	nextbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.save()
			menu.GetBack()
		}
	})

	menu.buttons = GE.GetGroup(leftbutton, rightbutton, nextbutton)
	menu.buttons.Init(nil, nil)

	menu.initRace()
	//menu.initClass()
}

func (menu *CharacterMenu) initRace() {
	menu.buttons.Member[0].(*GE.Button).RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeRace(-1)
		}
	})

	menu.buttons.Member[1].(*GE.Button).RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeRace(1)
		}
	})

	menu.races = make([]*GE.Group, len(Races))
	menu.rbackground = make([]*GE.ImageObj, len(Races))

	for i, race := range Races {
		menu.races[i], menu.rbackground[i] = getRace(race)
	}
}

func getRace(race *Race) (group *GE.Group, background *GE.ImageObj) {
	var err error
	background, err = GE.LoadImgObj(F_CHARACTERMENU+"/background"+race.name+".png", XRES-20, YRES-20, 10, 10, 0)
	CheckErr(err)

	title := GE.GetTextImage(race.name, XRES*0.06, YRES*0.1, YRES*0.09, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})
	stats := GE.GetTextImage(fmt.Sprintf("STR: %v DEX: %v INT: %v CHA: %v", race.attributes[0], race.attributes[1], race.attributes[2], race.attributes[3]), XRES*0.5, YRES*0.33, YRES*0.05, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})

	subraces := make([]GE.UpdateAble, len(race.subraces))
	for i, subrace := range race.subraces {
		subraces[i] = GE.GetTextImage(subrace, XRES*0.5, YRES*(0.48+float64(i)*0.05), YRES*0.04, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})
	}

	group = GE.GetGroup(append(subraces, title, stats)...)
	group.Init(nil, nil)
	return
}

func (menu *CharacterMenu) initClass() {
	x, y := 100.0, 550.0

	lbuttonimg, _ := GE.LoadImgObj(F_CHARACTERMENU+"/arrow.png", 0, 0, x-80, y+200, 0)
	lbuttonimg.ScaleToOriginalSize()
	leftbutton := GE.GetButton(lbuttonimg, lbuttonimg.Img)
	leftbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeClass(-1)
		}
	})

	rbuttonimg, _ := GE.LoadImgObj(F_CHARACTERMENU+"/arrow.png", 0, 0, x+910, y+200, 180)
	rbuttonimg.ScaleToOriginalSize()
	rightbutton := GE.GetButton(rbuttonimg, rbuttonimg.Img)
	rightbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeClass(1)
		}
	})

	menu.classthing = GE.GetGroup(leftbutton, rightbutton)
	menu.classthing.Init(nil, nil)

	for _, class := range Classes {
		group := getClass(class, x, y)
		menu.class = append(menu.class, group)
	}
}

func getClass(class *Class, x, y float64) (group *GE.Group) {
	title := GE.GetTextImage(class.name, XRES*0.06, YRES*0.1, YRES*0.9, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})

	subclass := make([]GE.UpdateAble, len(class.subclass))
	for i, subclas := range class.subclass {
		subclass[i] = GE.GetTextImage(subclas, x+220, y+160+float64(i*40), 30, GE.StandardFont, color.Black, &color.RGBA{168, 255, 68, 255})
	}

	group = GE.GetGroup(title)
	group.Init(nil, nil)
	return
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

	screen.Fill(color.RGBA{168, 255, 68, 255})

	switch menu.state {
	case 0:
		menu.races[menu.currRace].Update(0)
		menu.buttons.Update(0)

		menu.rbackground[menu.currRace].Draw(screen)
		menu.buttons.Draw(screen)
		menu.races[menu.currRace].Draw(screen)
	case 1:
		menu.classthing.Update(0)
		menu.class[menu.currClass].Update(0)

		menu.classthing.Draw(screen)
		menu.class[menu.currClass].Draw(screen)
	}

	return nil
}

func (menu *CharacterMenu) GetBack() {
	menu.parent.ChangeState(TITLESCREEN_STATE)
}

//Change which race is displayed
func (menu *CharacterMenu) changeRace(delta int) {
	menu.currRace += delta

	if menu.currRace < 0 {
		menu.currRace = len(menu.races) - 1
	}

	if menu.currRace >= len(menu.races) {
		menu.currRace = 0
	}
}

func (menu *CharacterMenu) changeClass(delta int) {
	menu.currClass += delta

	if menu.currClass < 0 {
		menu.currClass = len(menu.class) - 1
	}

	if menu.currClass >= len(menu.class) {
		menu.currClass = 0
	}
}

func (menu *CharacterMenu) save() {
	file, err := os.Create(F_CHARACTER + "/char.char")
	CheckErr(err)

	defer file.Close()

	io.WriteString(file, Races[menu.currRace].name)
}
