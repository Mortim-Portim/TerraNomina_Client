package Game

import (
	"github.com/hajimehoshi/ebiten"
)
func GetTitleScreen(g *TerraNomina) *TitleScreen {
	return &TitleScreen{parent:g}
}
type TitleScreen struct {
	
	parent *TerraNomina
}

func (t *TitleScreen) Start(g *TerraNomina) {
	
}
func (t *TitleScreen) Stop(g *TerraNomina) {
	
}
func (t *TitleScreen) Update(screen *ebiten.Image) error {
	TITLE_BackImg.Update(t.parent.frame)
	TITLE_BackImg.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)
	return nil
}