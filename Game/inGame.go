package Game

import (
	"fmt"
	"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func GetInGame(g *TerraNomina) *InGame {
	return &InGame{parent:g}
}
type InGame struct {
	parent *TerraNomina
	wrld *GE.WorldStructure
}

func (i *InGame) Init(g *TerraNomina) {
	fmt.Println("Initializing InGame")
	
}
func (i *InGame) Start(g *TerraNomina, oldState int) {
	fmt.Print("--------> InGame     \n")
	i.wrld = WorldStructure
}
func (i *InGame) Stop(g *TerraNomina, newState int) {
	fmt.Print("InGame      -------->")
}
func (i *InGame) Update(screen *ebiten.Image) error {
	if i.parent.frame%2 == 0 {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			i.wrld.Move(-1,0)
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			i.wrld.Move(1,0)
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			i.wrld.Move(0,-1)
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			i.wrld.Move(0,1)
		}
	}
	i.wrld.UpdateLightLevel(1)
	i.wrld.DrawLights(false)
	
	i.wrld.DrawBack(screen)
	i.wrld.DrawFront(screen)
	
	msg := fmt.Sprintf(`TPS: %0.2f`, ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)
	
	return nil
}