package Game

import (
	"log"
	"time"
	"github.com/hajimehoshi/ebiten"
)
func GetTitleScreen(g *TerraNomina) *TitleScreen {
	return &TitleScreen{parent:g}
}
type TitleScreen struct {
	
	parent *TerraNomina
}

func (t *TitleScreen) Start(g *TerraNomina) {
	log.Println("Starting TitleScreen")
	MainTheme.FadeToR(time.Now().UnixNano(), 3)
	MainTheme.PlayInfinite()
}
func (t *TitleScreen) Stop(g *TerraNomina) {
	log.Println("Stopping TitleScreen")
	MainTheme.FadeOut(1)
	time.Sleep(time.Second)
}
func (t *TitleScreen) Update(screen *ebiten.Image) error {
	TITLE_BackImg.Update(t.parent.frame)
	TITLE_BackImg.DrawImageObj(screen)
	TITLE_Name.DrawImageObj(screen)
	return nil
}