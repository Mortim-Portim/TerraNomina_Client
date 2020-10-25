package Game

import "github.com/hajimehoshi/ebiten"

func getCharacterMenu() (cm *CharacterMenu) {
	cm = &CharacterMenu{}
	return
}

type CharacterMenu struct {
}

func (menu *CharacterMenu) Init(g *TerraNomina) {

}

func (menu *CharacterMenu) Start(g *TerraNomina, lastState int) {

}

func (menu *CharacterMenu) Stop(g *TerraNomina, nextState int) {

}

func (menu *CharacterMenu) Update(screen *ebiten.Image) error {
	return nil
}
