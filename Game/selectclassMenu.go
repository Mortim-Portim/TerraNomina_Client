package Game

/*import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
)

type SelectClassMenu struct {
	parent      *TerraNomina
	classthing  *GE.Group
	cbackground []*GE.ImageObj
	classes     []*GE.Group
	currClass   int
}

func GetSelectClassMenu(parent *TerraNomina) *SelectClassMenu {
	return &SelectClassMenu{parent: parent}
}

func (menu *SelectClassMenu) Init() {
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

	nextbutton := GE.GetTextButton("Next", "", GE.StandardFont, XRES*0.09, YRES*0.83, YRES*0.12, color.Black, &color.RGBA{255, 0, 0, 255})
	nextbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			charinmaking.Class = TNE.Classes[menu.currClass]
			menu.parent.ChangeState(SELSTATS_STATE)
		}
	})

	backbutton := GE.GetTextButton("Back", "", GE.StandardFont, XRES*0.28, YRES*0.83, YRES*0.12, color.Black, color.Transparent)
	backbutton.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			menu.parent.ChangeState(SELRACE_STATE)
		}
	})

	menu.classthing = GE.GetGroup(leftbutton, rightbutton, nextbutton, backbutton)
	menu.classthing.Init(nil, nil)

	menu.classes = make([]*GE.Group, len(TNE.Classes))
	menu.cbackground = make([]*GE.ImageObj, len(TNE.Classes))

	for i, class := range TNE.Classes {
		group := getClass(class)
		menu.classes[i] = group
	}
}

/*
func (class *Class) getClass(group *GE.Group) {
	title := GE.GetTextImage(class.Name, 0, 0, YRES*0.15, GE.StandardFont, color.Black, color.Transparent)
	title.SetMiddle(XRES*0.25, YRES*0.12)

	subclass := make([]GE.UpdateAble, len(class.Subclass))
	for i, subclas := range class.Subclass {
		subclass[i] = GE.GetTextImage(subclas, XRES*0.5, YRES*(0.48+float64(i)*0.05), YRES*0.04, GE.StandardFont, color.Black, color.Transparent)
	}

	group = GE.GetGroup(append(subclass, title)...)
	group.Init(nil, nil)
	return
}

func (menu *SelectClassMenu) changeClass(delta int) {
	menu.currClass += delta

	if menu.currClass < 0 {
		menu.currClass = len(menu.classes) - 1
	}

	if menu.currClass >= len(menu.classes) {
		menu.currClass = 0
	}
}

func (menu *SelectClassMenu) Start(laststate int) {
	var err error
	for i, class := range TNE.Classes {
		menu.cbackground[i], err = GE.LoadImgObj(F_CHARACTERMENU+"/class/background"+class.Name+".png", XRES, YRES, 0, 0, 0)
		CheckErr(err)
	}
}

func (menu *SelectClassMenu) Stop(nextstate int) {
	for i := range menu.cbackground {
		menu.cbackground[i] = nil
	}
}

func (menu *SelectClassMenu) Update() error {
	menu.classes[menu.currClass].Update(menu.parent.frame)
	menu.classthing.Update(menu.parent.frame)
	return nil
}
func (menu *SelectClassMenu) Draw(screen *ebiten.Image) {
	if menu.cbackground[menu.currClass] != nil {
		menu.cbackground[menu.currClass].Draw(screen)
	}
	menu.classes[menu.currClass].Draw(screen)
	menu.classthing.Draw(screen)
}
*/
