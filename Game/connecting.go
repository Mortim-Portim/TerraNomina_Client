package Game

import (
	"fmt"
	//"marvin/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
)

func GetConnecting(g *TerraNomina) *Connecting {
	return &Connecting{parent:g}
}
type Connecting struct {
	parent *TerraNomina
	oldState int
	
	ipAddr string
}

func (t *Connecting) Init(g *TerraNomina) {
	fmt.Println("Initializing Connecting")
	
}
func (t *Connecting) Start(g *TerraNomina, oldState int) {
	fmt.Print("--------> Connecting \n")
	t.oldState = oldState
	t.ipAddr = USER_INPUT_IP_ADDR
}
func (t *Connecting) Stop(g *TerraNomina, newState int) {
	fmt.Print("Connecting  -------->")
}
func (t *Connecting) Update(screen *ebiten.Image) error {
	
	return nil
}