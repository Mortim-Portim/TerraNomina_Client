package Game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetSelectRaceMenu(parent *TerraNomina) *SelectRaceMenu {
	return &SelectRaceMenu{parent: parent}
}

type SelectRaceMenu struct {
	parent *TerraNomina

	racething   *GE.Group
	rbackground []*GE.ImageObj
	races       []*GE.Group
	currRace    int
}

var charinmaking *Character
var arrow *ebiten.Image

func SetupCharacterMenu() {
	var err error
	arrow, err = GE.LoadEbitenImg(F_CHARACTERMENU + "/arrow.png")
	CheckErr(err)
}

func (menu *SelectRaceMenu) Init() {
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

	nextbutton := GE.GetTextButton("Next", "", GE.StandardFont, XRES*0.1, YRES*0.83, YRES*0.12, color.Black, color.Transparent)
	nextbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			charinmaking.race = Races[menu.currRace]
			menu.parent.ChangeState(SELCLASS_STATE)
		}
	})

	backbutton := GE.GetTextButton("Back", "", GE.StandardFont, XRES*0.25, YRES*0.83, YRES*0.12, color.Black, color.Transparent)
	backbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			menu.parent.ChangeState(TITLESCREEN_STATE)
		}
	})

	menu.racething = GE.GetGroup(leftbutton, rightbutton, nextbutton, backbutton)
	menu.racething.Init(nil, nil)

	menu.races = make([]*GE.Group, len(Races))
	menu.rbackground = make([]*GE.ImageObj, len(Races))

	for i, race := range Races {
		menu.races[i] = getRace(race)
	}
}

func getRace(race *Race) (group *GE.Group) {
	title := GE.GetTextImage(race.Name, 0, 0, YRES*0.15, GE.StandardFont, color.Black, color.Transparent)
	title.SetMiddle(XRES*0.25, YRES*0.12)
	stats := GE.GetTextImage(fmt.Sprintf("STR:%v | DEX:%v | INT:%v | CHA:%v", race.Attributes[0], race.Attributes[1], race.Attributes[2], race.Attributes[3]), XRES*0.5, YRES*0.32, YRES*0.06, GE.StandardFont, color.Black, color.Transparent)
	anim, err := GE.GetDayNightAnimFromParams(0, 0, 0, 0, F_CREATURE+"/"+race.Name+"/idle_R.txt", F_CREATURE+"/"+race.Name+"/idle_R.png")
	GE.ShitImDying(err)
	anim.ScaleToOriginalSize()
	anim.ScaleDim(YRES*0.48, 1)
	anim.SetMiddle(XRES*0.15, YRES*0.53)

	subraces := make([]GE.UpdateAble, len(race.Subraces))
	for i, subrace := range race.Subraces {
		subraces[i] = GE.GetTextImage(subrace, XRES*0.5, YRES*(0.48+float64(i)*0.05), YRES*0.04, GE.StandardFont, color.Black, color.Transparent)
	}

	group = GE.GetGroup(append(subraces, title, stats, anim)...)
	group.Init(nil, nil)
	return
}

//Change which race is displayed
func (menu *SelectRaceMenu) changeRace(delta int) {
	menu.currRace += delta

	if menu.currRace < 0 {
		menu.currRace = len(menu.races) - 1
	}

	if menu.currRace >= len(menu.races) {
		menu.currRace = 0
	}

	Soundtrack.Play(Races[menu.currRace].Name)
}

func (menu *SelectRaceMenu) Start(laststate int) {
	var err error
	for i, race := range Races {
		menu.rbackground[i], err = GE.LoadImgObj(F_CHARACTERMENU+"/race/background"+race.Name+".png", XRES, YRES, 0, 0, 0)
		CheckErr(err)
	}

	charinmaking = &Character{}
}

func (menu *SelectRaceMenu) Stop(nextstate int) {
	for i := range menu.rbackground {
		menu.rbackground[i] = nil
	}
}

func (menu *SelectRaceMenu) Update(screen *ebiten.Image) error {
	menu.races[menu.currRace].Update(menu.parent.frame)
	menu.racething.Update(menu.parent.frame)

	if menu.rbackground[menu.currRace] != nil {
		menu.rbackground[menu.currRace].Draw(screen)
	}
	menu.races[menu.currRace].Draw(screen)
	menu.racething.Draw(screen)

	return nil
}
