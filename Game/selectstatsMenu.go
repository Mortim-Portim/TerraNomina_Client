package Game

import (
	"fmt"
	"image/color"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
)

func GetSelectStatsMenu(parent *TerraNomina) *SelectStatsMenu {
	return &SelectStatsMenu{parent: parent}
}

var number []*ebiten.Image

type SelectStatsMenu struct {
	parent *TerraNomina

	char *TNE.Character

	statsthing *GE.Group
	name       *GE.EditText
	attedit    []*GE.EditText
}

var stats []string = []string{"Speed", "Max-Health", "Max-Mana", "Max-Stamina", "Stamina-Regen", "Mana-Regen"}
var defstats []int8 = []int8{100, 100, 100, 100, 5, 5}

func (menu *SelectStatsMenu) Init() {
	fmt.Println("Initializing InGame")
	number = make([]*ebiten.Image, 16)
	for i := -2; i <= 13; i++ {
		number[i+2] = GE.MakePopUp(strconv.Itoa(i), GE.StandardFont, color.Black, color.Transparent)
	}

	abilcomps := make([]GE.UpdateAble, len(stats))
	menu.attedit = make([]*GE.EditText, len(stats))
	for i, stat := range stats {
		abilcomps[i] = GE.GetTextImage(stat, XRES*0.3, YRES*(0.17+float64(i)*0.06), YRES*0.04, GE.StandardFont, color.Black, color.Transparent)
		menu.attedit[i] = GE.GetEditText(strconv.Itoa(int(defstats[i])), XRES*0.45, YRES*(0.17+float64(i)*0.06), YRES*0.04, 6, GE.StandardFont, color.Black, color.Black)
	}

	savebutton := GE.GetTextButton("Save", "", GE.StandardFont, XRES*0.1, YRES*0.83, YRES*0.12, color.Black, &color.RGBA{255, 0, 0, 255})
	savebutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			//menu.char.Attributes = menu.attributes
			//menu.char.Name = menu.name.GetText()
			//SaveChar(menu.char)

			menu.parent.ChangeState(TITLESCREEN_STATE)
		}
	})

	menu.name = GE.GetEditText("Name", XRES*0.05, YRES*0.03, YRES*0.08, 15, GE.StandardFont, color.Black, color.RGBA{255, 0, 0, 255})

	backbutton := GE.GetTextButton("Back", "", GE.StandardFont, XRES*0.25, YRES*0.83, YRES*0.12, color.Black, color.Transparent)
	backbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			menu.parent.ChangeState(TITLESCREEN_STATE)
		}
	})

	statsthing := []GE.UpdateAble{}
	statsthing = append(statsthing, savebutton, backbutton)
	statsthing = append(statsthing, abilcomps...)

	menu.statsthing = GE.GetGroup(statsthing...)
	menu.statsthing.Init(nil, nil)
}

func SaveChar(char *TNE.Character) {
	file, _ := os.Create(F_CHARACTER + "/" + char.Name + ".char")
	defer file.Close()
	file.Truncate(0)
	file.Write(char.ToByte())
}

func (menu *SelectStatsMenu) Start(laststate int) {
	menu.char = TNE.GetBlankCharacter()
}

func (menu *SelectStatsMenu) Stop(nextstate int) {}

func (menu *SelectStatsMenu) Update() error {
	menu.statsthing.Update(menu.parent.frame)
	menu.name.Update(menu.parent.frame)

	for _, edit := range menu.attedit {
		edit.Update(menu.parent.frame)
	}

	return nil
}

func (menu *SelectStatsMenu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{168, 255, 68, 255})
	menu.statsthing.Draw(screen)

	for _, edit := range menu.attedit {
		edit.Draw(screen)
	}
	menu.name.Draw(screen)
}
