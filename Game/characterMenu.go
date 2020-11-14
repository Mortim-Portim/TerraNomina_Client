package Game

import (
	"fmt"
	"image/color"
	"marvin/GraphEng/GE"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

//In die GetGroup Funktion Init reinpacken
//eventuell imageobject nicht * bitte bitte bitte

func GetCharacterMenu(parent *TerraNomina) (cm *CharacterMenu) {
	cm = &CharacterMenu{parent: parent}
	return
}

var arrow *ebiten.Image

type CharacterMenu struct {
	parent *TerraNomina
	state  int

	//Races
	racething   *GE.Group
	rbackground []*GE.ImageObj
	races       []*GE.Group
	currRace    int

	//Classes
	classthing  *GE.Group
	cbackground []*GE.ImageObj
	classes     []*GE.Group
	currClass   int

	//Stats
	statsthing *GE.Group
	attpicture []*GE.ImageObj
	sum        *GE.ImageObj
	attributes []int8
}

func (menu *CharacterMenu) Init(g *TerraNomina) {
	arrow, _ = GE.LoadEbitenImg(F_CHARACTERMENU + "/arrow.png")

	menu.initRace()
	menu.initClass()
	menu.initStats()
}

func (menu *CharacterMenu) initRace() {
	larrowimg := &GE.ImageObj{arrow, nil, XRES * 0.04, YRES * 0.06, XRES * 0.01, YRES * 0.46, 0}
	larrowimg.SetMiddleY(YRES / 2)
	leftbutton := GE.GetButton(larrowimg, nil)
	leftbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			menu.changeRace(-1)
		}
	})

	rarrowimg := &GE.ImageObj{arrow, nil, XRES * 0.04, YRES * 0.06, XRES * 0.95, 0, 180}
	rarrowimg.SetMiddleY(YRES / 2)
	rightbutton := GE.GetButton(rarrowimg, nil)
	rightbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeRace(1)
		}
	})

	nextbutton := GE.GetTextButton("Next", "", GE.StandardFont, XRES*0.1, YRES*0.83, YRES*0.12, color.Black, &color.RGBA{255, 0, 0, 255})
	nextbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			menu.state = 1

			for i, num := range Races[menu.currRace].attributes {
				menu.attributes[i] = num
				menu.changeAttribute(i, int(num))
			}
		}
	})

	menu.racething = GE.GetGroup(leftbutton, rightbutton, nextbutton)
	menu.racething.Init(nil, nil)

	menu.races = make([]*GE.Group, len(Races))
	menu.rbackground = make([]*GE.ImageObj, len(Races))

	for i, race := range Races {
		menu.races[i] = getRace(race)
	}
}

func getRace(race *Race) (group *GE.Group) {
	title := GE.GetTextImage(race.name, 0, 0, YRES*0.15, GE.StandardFont, color.Black, color.Transparent)
	title.SetMiddle(XRES*0.25, YRES*0.14)
	stats := GE.GetTextImage(fmt.Sprintf("STR: %v DEX: %v INT: %v CHA: %v", race.attributes[0], race.attributes[1], race.attributes[2], race.attributes[3]), XRES*0.52, YRES*0.32, YRES*0.06, GE.StandardFont, color.Black, color.Transparent)
	anim, err := GE.GetDayNightAnimFromParams(0, 0, 0, 0, F_CREATURE+"/"+race.name+"/idle_R.txt", F_CREATURE+"/"+race.name+"/idle_R.png")
	CheckErr(err)
	anim.ScaleToOriginalSize()
	anim.ScaleDim(YRES*0.48, 1)
	anim.SetMiddle(XRES*0.15, YRES*0.53)

	subraces := make([]GE.UpdateAble, len(race.subraces))
	for i, subrace := range race.subraces {
		subraces[i] = GE.GetTextImage(subrace, XRES*0.5, YRES*(0.48+float64(i)*0.05), YRES*0.04, GE.StandardFont, color.Black, color.Transparent)
	}

	group = GE.GetGroup(append(subraces, title, stats, anim)...)
	group.Init(nil, nil)
	return
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

//Change which class is displayed
func (menu *CharacterMenu) initClass() {
	larrowimg := &GE.ImageObj{arrow, nil, XRES * 0.04, YRES * 0.04, XRES * 0.01, YRES * 0.46, 0}
	leftbutton := GE.GetButton(larrowimg, nil)
	leftbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeClass(-1)
		}
	})

	rarrowimg := &GE.ImageObj{arrow, nil, XRES * 0.04, YRES * 0.04, XRES * 0.95, YRES * 0.46, 180}
	rightbutton := GE.GetButton(rarrowimg, nil)
	rightbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.changeClass(1)
		}
	})

	nextbutton := GE.GetTextButton("Next", "", GE.StandardFont, XRES*0.05, YRES*0.82, YRES*0.1, color.Black, &color.RGBA{255, 0, 0, 255})
	nextbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed {
			menu.state = 2
		}
	})

	menu.classthing = GE.GetGroup(leftbutton, rightbutton, nextbutton)
	menu.classthing.Init(nil, nil)

	menu.classes = make([]*GE.Group, len(Classes))
	menu.cbackground = make([]*GE.ImageObj, len(Classes))

	for i, class := range Classes {
		group := getClass(class)
		menu.classes[i] = group
	}
}

func getClass(class *Class) (group *GE.Group) {
	title := GE.GetTextImage(class.name, XRES*0.06, YRES*0.1, YRES*0.09, GE.StandardFont, color.Black, color.Transparent)

	subclass := make([]GE.UpdateAble, len(class.subclass))
	for i, subclas := range class.subclass {
		subclass[i] = GE.GetTextImage(subclas, XRES*0.1, YRES*(0.4+float64(i)*0.05), YRES*0.04, GE.StandardFont, color.Black, color.Transparent)
	}

	group = GE.GetGroup(append(subclass, title)...)
	group.Init(nil, nil)
	return
}

func (menu *CharacterMenu) changeClass(delta int) {
	menu.currClass += delta

	if menu.currClass < 0 {
		menu.currClass = len(menu.classes) - 1
	}

	if menu.currClass >= len(menu.classes) {
		menu.currClass = 0
	}
}

func (menu *CharacterMenu) initStats() {
	stats := []string{"Strength", "Dexterity", "Intelligence", "Charisma"}

	buttons := make([]GE.UpdateAble, len(stats)*3)
	nums := make([]*GE.ImageObj, len(stats))
	for i, stat := range stats {
		buttons[i*3] = GE.GetTextImage(stat, XRES*0.05, YRES*(0.1+float64(i)*0.07), YRES*0.05, GE.StandardFont, color.Black, color.Transparent)

		lbuttonimg := &GE.ImageObj{arrow, nil, XRES * 0.05, YRES * 0.05, XRES * 0.25, YRES * (0.1 + float64(i)*0.07), 0}
		lbutton := GE.GetButton(lbuttonimg, nil)
		lbutton.Data = i
		lbutton.RegisterOnLeftEvent(func(button *GE.Button) {
			if button.LPressed {
				index := button.Data.(int)
				if menu.attributes[index] > 0 {
					menu.attributes[index]--
					menu.changeAttribute(index, int(menu.attributes[index]))
				}
			}
		})
		buttons[i*3+1] = lbutton

		rbuttonimg := &GE.ImageObj{arrow, nil, XRES * 0.05, YRES * 0.05, XRES * 0.35, YRES * (0.1 + float64(i)*0.07), 180}
		rbutton := GE.GetButton(rbuttonimg, nil)
		rbutton.Data = i
		rbutton.RegisterOnLeftEvent(func(button *GE.Button) {
			if button.LPressed {
				index := button.Data.(int)
				if menu.attributes[index] < 8 {
					menu.attributes[index]++
					menu.changeAttribute(index, int(menu.attributes[index]))
				}
			}
		})
		buttons[i*3+2] = rbutton

		numimg := GE.GetTextImage("0", 0, YRES*(0.1+float64(i)*0.07), YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
		numimg.SetMiddleX(XRES * 0.325)
		nums[i] = numimg
	}

	sumlabel := GE.GetTextImage("10 / 10", 0, YRES*0.4, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	sumlabel.X = XRES*0.4 - sumlabel.W
	menu.sum = sumlabel

	menu.statsthing = GE.GetGroup(buttons...)
	menu.statsthing.Init(nil, nil)

	menu.attpicture = nums

	menu.attributes = make([]int8, 4)
}

var pointmap []int = []int{1, 1, 2, 2, 3, 3, 4, 4}

func (menu *CharacterMenu) changeAttribute(index, newvalue int) {
	oldimg := menu.attpicture[index]
	menu.attpicture[index] = GE.GetTextImage(strconv.Itoa(newvalue), 0, oldimg.Y, oldimg.H, GE.StandardFont, color.Black, color.Transparent)
	menu.attpicture[index].SetMiddleX(XRES * 0.325)

	menu.attributes[index] = int8(newvalue)

	score := 10

	for i := range menu.attributes {
		for l := Races[menu.currRace].attributes[i]; l < menu.attributes[i]; l++ {
			score -= pointmap[l]
		}
		fmt.Print("\n")
	}

	menu.sum = GE.GetTextImage(strconv.Itoa(score)+" / 10", 0, YRES*0.4, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	menu.sum.X = XRES*0.4 - menu.sum.W
}

func (menu *CharacterMenu) Start(g *TerraNomina, lastState int) {
	fmt.Print("--------> CharacterMenu   \n")

	for i, race := range Races {
		menu.rbackground[i], _ = GE.LoadImgObj(F_CHARACTERMENU+"/background"+race.name+".png", XRES, YRES, 0, 0, 0)
	}

	for i, class := range Classes {
		menu.cbackground[i], _ = GE.LoadImgObj(F_CHARACTERMENU+"/background"+class.name+".png", XRES, YRES, 0, 0, 0)
	}

	menu.state = 0
}

func (menu *CharacterMenu) Stop(g *TerraNomina, nextState int) {
	fmt.Print("CharacterMenu ------>")

	for i := range menu.rbackground {
		menu.rbackground[i] = nil
	}

	for i := range menu.cbackground {
		menu.cbackground[i] = nil
	}
}

func (menu *CharacterMenu) Update(screen *ebiten.Image) error {
	down, changed := Keyli.GetMappedKeyState(ESC_KEY_ID)
	if changed && !down {
		menu.GetBack()
		return nil
	}

	screen.Fill(color.RGBA{168, 255, 68, 255})

	switch menu.state {
	case 0:
		menu.races[menu.currRace].Update(menu.parent.frame)
		menu.racething.Update(menu.parent.frame)

		menu.rbackground[menu.currRace].Draw(screen)
		menu.races[menu.currRace].Draw(screen)
		menu.racething.Draw(screen)
	case 1:
		menu.classes[menu.currClass].Update(menu.parent.frame)
		menu.classthing.Update(menu.parent.frame)

		menu.cbackground[menu.currClass].Draw(screen)
		menu.classes[menu.currClass].Draw(screen)
		menu.classthing.Draw(screen)
	case 2:
		menu.statsthing.Update(menu.parent.frame)
		menu.statsthing.Draw(screen)

		for _, img := range menu.attpicture {
			img.Draw(screen)
		}

		menu.sum.Draw(screen)
	}

	return nil
}

func (menu *CharacterMenu) GetBack() {
	menu.parent.ChangeState(TITLESCREEN_STATE)
}
