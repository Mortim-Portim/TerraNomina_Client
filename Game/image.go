package Game

import (
	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

type Image struct {
	image *GE.ImageObj
}

func (img *Image) Init(screen *ebiten.Image, data interface{}) (GE.UpdateFunc, GE.DrawFunc) {
	return img.Update, img.Draw
}

func (img *Image) Start(screen *ebiten.Image, data interface{}) {}

func (img *Image) Stop(screen *ebiten.Image, data interface{}) {}

func (img *Image) Update(frame int) {}

func (img *Image) Draw(screen *ebiten.Image) {
	img.image.DrawImageObj(screen)
}
