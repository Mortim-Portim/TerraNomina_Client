package Game

import (
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

//var number []*ebiten.Image

type SelectStatsMenu struct {
	parent *TerraNomina

	statsthing *GE.Group
	name       *GE.EditText

	attpicture []*GE.ImageObj
	attributes []int8
	sum        *GE.ImageObj

	profselect    []*GE.Button
	profpicture   []*GE.ImageObj
	proficiencies []int8
	profcount     int
}

//var stats []string = []string{"Strength", "Dexterity", "Intelligence", "Charisma"}
//var proficiencies []string = []string{"Strength (STR)", "Dexterity (DEX)", "Intelligence (INT)", "Charisma (CHA)", "Endurance (STR)", "Persuasion (CHA)", "Deception (CHA)", "Performance (CHA)", "Insight (INT)", "Thievery (DEX)", "Stealth (DEX)", "Acrobatics (DEX)", "Nature (INT)", "Arcana (INT)", "Perception (INT)", "Craftsmanship (INT)", "Dungeoneering"}
//var abtoprof []int = []int{ABIL_STRENGTH, ABIL_DEXTERITY, ABIL_INTELLIGENCE, ABIL_CHARISMA, ABIL_STRENGTH, ABIL_CHARISMA, ABIL_CHARISMA, ABIL_CHARISMA, ABIL_INTELLIGENCE, ABIL_DEXTERITY, ABIL_DEXTERITY, ABIL_DEXTERITY, ABIL_INTELLIGENCE, ABIL_INTELLIGENCE, ABIL_INTELLIGENCE, ABIL_INTELLIGENCE, ABIL_INTELLIGENCE}

func (menu *SelectStatsMenu) Init() {
	number = make([]*ebiten.Image, 16)
	for i := -2; i <= 13; i++ {
		number[i+2] = GE.MakePopUpOld(strconv.Itoa(i), 50, GE.StandardFont, color.Black, color.Transparent)
	}

	abiliscore := make([]GE.UpdateAble, len(stats)*3)
	nums := make([]*GE.ImageObj, len(stats))

	for i, stat := range stats {
		abiliscore[i*3] = GE.GetTextImage(stat, XRES*0.05, YRES*(0.17+float64(i)*0.07), YRES*0.05, GE.StandardFont, color.Black, color.Transparent)

		lbuttonimg := &GE.ImageObj{arrow, nil, XRES * 0.05, YRES * 0.05, XRES * 0.25, YRES * (0.17 + float64(i)*0.07), 0}
		lbutton := GE.GetButton(lbuttonimg, nil)
		lbutton.Data = i
		lbutton.RegisterOnLeftEvent(func(button *GE.Button) {
			if button.LPressed {
				index := button.Data.(int)
				if menu.attributes[index] > 0 {
					menu.changeAttribute(index, -1)
				}
			}
		})
		abiliscore[i*3+1] = lbutton

		rbuttonimg := &GE.ImageObj{arrow, nil, XRES * 0.05, YRES * 0.05, XRES * 0.35, YRES * (0.17 + float64(i)*0.07), 180}
		rbutton := GE.GetButton(rbuttonimg, nil)
		rbutton.Data = i
		rbutton.RegisterOnLeftEvent(func(button *GE.Button) {
			if button.LPressed {
				index := button.Data.(int)
				if menu.attributes[index] < 8 {
					menu.changeAttribute(index, +1)
				}
			}
		})
		abiliscore[i*3+2] = rbutton

		numimg := &GE.ImageObj{Y: YRES * (0.17 + float64(i)*0.07), H: YRES * 0.05, X: XRES * 0.325}
		nums[i] = numimg
	}

	profscore := make([]GE.UpdateAble, len(proficiencies))
	menu.profselect = make([]*GE.Button, len(proficiencies))
	menu.profpicture = make([]*GE.ImageObj, len(proficiencies))
	menu.proficiencies = make([]int8, len(proficiencies))

	clear, _ := GE.LoadEbitenImg(F_CHARACTERMENU + "/CheckboxClear.png")
	checked, _ := GE.LoadEbitenImg(F_CHARACTERMENU + "/CheckboxChecked.png")
	for i, prof := range proficiencies {
		profscore[i] = GE.GetTextImage(prof, XRES*0.6, YRES*(0.05+float64(i)*0.05), YRES*0.03, GE.StandardFont, color.Black, color.Transparent)

		profselect := GE.GetButton(&GE.ImageObj{clear, nil, XRES * 0.03, XRES * 0.03, XRES * 0.77, YRES * (0.04 + float64(i)*0.05), 0}, checked)
		profselect.Data = i
		profselect.RegisterOnLeftEvent(func(btn *GE.Button) {
			if btn.LPressed {
				if btn.DrawDark {
					menu.profcount--
					menu.changeProfencies(btn.Data.(int), -5)
				} else {
					if menu.profcount < charinmaking.Race.Extraprof {
						menu.profcount++
						menu.changeProfencies(btn.Data.(int), 5)
					} else {
						return
					}
				}
				btn.DrawDark = !btn.DrawDark
			}
		})
		menu.profselect[i] = profselect

		menu.profpicture[i] = &GE.ImageObj{Y: YRES * (0.05 + float64(i)*0.05), H: YRES * 0.03, X: XRES * 0.82}
	}

	sumlabel := GE.GetTextImage("10 / 10", 0, YRES*0.47, YRES*0.05, GE.StandardFont, color.Black, color.Transparent)
	sumlabel.X = XRES*0.4 - sumlabel.W
	menu.sum = sumlabel

	savebutton := GE.GetTextButton("Save", "", GE.StandardFont, XRES*0.1, YRES*0.83, YRES*0.12, color.Black, &color.RGBA{255, 0, 0, 255})
	savebutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			charinmaking.Attributes = menu.attributes
			charinmaking.Proficiencies = menu.proficiencies
			charinmaking.Name = menu.name.GetText()
			SaveChar(charinmaking)

			menu.parent.ChangeState(TITLESCREEN_STATE)
		}
	})

	menu.name = GE.GetEditText("Name", XRES*0.05, YRES*0.03, YRES*0.08, 15, GE.StandardFont, color.Black, color.RGBA{255, 0, 0, 255})

	backbutton := GE.GetTextButton("Back", "", GE.StandardFont, XRES*0.25, YRES*0.83, YRES*0.12, color.Black, color.Transparent)
	backbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			menu.parent.ChangeState(SELCLASS_STATE)
		}
	})
	abiliscore = append(abiliscore, backbutton)

	statsthing := []GE.UpdateAble{}
	statsthing = append(statsthing, savebutton, backbutton)
	statsthing = append(statsthing, abiliscore...)
	statsthing = append(statsthing, profscore...)

	menu.statsthing = GE.GetGroup(statsthing...)
	menu.statsthing.Init(nil, nil)

	menu.attpicture = nums

	menu.attributes = make([]int8, 4)
}

func SaveChar(char *TNE.Character) {
	file, _ := os.Create(F_CHARACTER + "/" + char.Name + ".char")
	defer file.Close()
	file.Truncate(0)
	file.Write(char.ToByte())
}

func (menu *SelectStatsMenu) changeAttribute(index int, deltavalue int8) {
	menu.attributes[index] += deltavalue
	menu.attpicture[index].Img = number[menu.attributes[index]+2]
	curnumber, oldimg := number[menu.attributes[index]+2], menu.attpicture[index]
	w, h := curnumber.Size()
	numimg := &GE.ImageObj{curnumber, nil, float64(w) * ((oldimg.H) / float64(h)), oldimg.H, 0, 0, 0}
	numimg.SetMiddle(oldimg.GetMiddle())
	menu.attpicture[index] = numimg

	score := 10
	for i := range menu.attributes {
		for l := charinmaking.Race.Attributes[i]; l < menu.attributes[i]; l++ {
			score -= pointmap[l+1]
		}
	}

	for i, value := range abtoprof {
		if value == index {
			menu.changeProfencies(i, deltavalue)
		}
	}

	menu.sum = GE.GetTextImage(strconv.Itoa(score)+" / 10", 0, menu.sum.Y, menu.sum.H, GE.StandardFont, color.Black, color.Transparent)
	menu.sum.X = XRES*0.4 - menu.sum.W
}

func (menu *SelectStatsMenu) changeProfencies(index int, deltavalue int8) {
	menu.proficiencies[index] += deltavalue

	curnumber, oldimg := number[menu.proficiencies[index]+2], menu.profpicture[index]
	w, h := curnumber.Size()
	numimg := &GE.ImageObj{curnumber, nil, float64(w) * ((oldimg.H) / float64(h)), oldimg.H, 0, 0, 0}
	numimg.SetMiddle(oldimg.GetMiddle())
	menu.profpicture[index] = numimg
}

func (menu *SelectStatsMenu) resetStats() {
	for i := range menu.attpicture {
		menu.attributes[i] = 0
		menu.changeAttribute(i, 0)
	}

	for i, but := range menu.profselect {
		menu.proficiencies[i] = 0
		menu.changeProfencies(i, 0)

		but.DrawDark = false
		but.Active = true
	}
}

func (menu *SelectStatsMenu) Start(laststate int) {
	menu.resetStats()

	for i, stat := range charinmaking.Race.Attributes {
		menu.changeAttribute(i, stat)
	}

	for _, prof := range charinmaking.Race.Profencies {
		menu.changeProfencies(prof, 5)
		menu.profselect[prof].Active = false
	}
}
func (menu *SelectStatsMenu) Stop(nextstate int) {}

func (menu *SelectStatsMenu) Update(screen *ebiten.Image) error {
	menu.statsthing.Update(menu.parent.frame)
	menu.name.Update(menu.parent.frame)

	screen.Fill(color.RGBA{168, 255, 68, 255})
	menu.statsthing.Draw(screen)

	for _, img := range menu.attpicture {
		img.Draw(screen)
	}

	for _, img := range menu.profselect {
		img.Update(menu.parent.frame)
		img.Draw(screen)
	}

	for _, img := range menu.profpicture {
		img.Update(menu.parent.frame)
		img.Draw(screen)
	}

	menu.sum.Draw(screen)
	menu.name.Draw(screen)

	return nil
}
