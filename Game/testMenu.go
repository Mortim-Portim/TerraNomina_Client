package Game

//import (
//	"image/color"
//
//	"github.com/hajimehoshi/ebiten"
//	"github.com/mortim-portim/GraphEng/GE"
//)
//
//func getTestMenu(parent *TerraNomina) *TestMenu {
//	return &TestMenu{parent: parent}
//}
//
//type TestMenu struct {
//	parent *TerraNomina
//
//	skills *GE.Group
//}
//
//var keys []ebiten.Key = []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9, ebiten.Key0}
//
//func (menu *TestMenu) Init() {
//	xmid := XRES * 0.5
//	ymid := YRES * 0.9
//	width := YRES * 0.07
//
//	buttons := make([]GE.UpdateAble, 0)
//	for y := -1.0; y < 1; y++ {
//		for x := -5.0; x < 5; x++ {
//			img, err := GE.LoadEbitenImg(F_IMAGES + "/Skills/Void.png")
//			CheckErr(err)
//
//			button := GE.GetImageButton(img, xmid+(x*(width+3)), ymid+(y*(width+3)), width, width)
//
//			Keyli.RegisterKeyEventListener(Keyli.MappKey(keys[int(x)+5]), func(*GE.KeyLi, bool) {
//				button.LPressed = true
//			})
//
//			buttons = append(buttons, button)
//		}
//	}
//
//	menu.skills = GE.GetGroup(buttons...)
//	menu.skills.Init(nil, nil)
//}
//
//func (menu *TestMenu) Start(lastState int) {
//
//}
//
//func (menu *TestMenu) Stop(nextState int) {
//
//}
//
//func (menu *TestMenu) Update(screen *ebiten.Image) error {
//	screen.Fill(color.RGBA{168, 255, 68, 255})
//
//	menu.skills.Update(menu.parent.frame)
//	menu.skills.Draw(screen)
//
//	return nil
//}
