package Game

import (
	"fmt"
	"time"
	"marvin/GraphEng/GE"
	"marvin/GameConn/GC"
	"github.com/hajimehoshi/ebiten"
)

/**
Connects
Sends: MAP_REQUEST
Receives: [ MAP_REQUEST | (n)map bytes ]
**/

func GetConnecting(g *TerraNomina) *Connecting {
	return &Connecting{parent:g}
}
type Connecting struct {
	background *GE.ImageObj
	loadingAnim *GE.Animation
	parent *TerraNomina
	oldState int
	
	ipAddr string
	mapData []byte
}

func (t *Connecting) Init(g *TerraNomina) {
	fmt.Println("Initializing Connecting")
	lps := &GE.Params{}; err := lps.LoadFromFile(RES+CONNECTING_FILES+"/loading.txt"); CheckErr(err)
	limg, err := GE.LoadEbitenImg(RES+CONNECTING_FILES+"/loading.png")
	CheckErr(err)
	t.loadingAnim = GE.GetAnimationFromParams(0,0,XRES,YRES, lps, limg)
	t.loadingAnim.Init(nil,nil)
	
	t.background, err = GE.LoadImgObj(RES+CONNECTING_FILES+"/back.png", XRES, YRES, 0, 0, 0)
	CheckErr(err)
}
func (t *Connecting) Start(g *TerraNomina, oldState int) {
	fmt.Print("--------> Connecting \n")
	t.oldState = oldState
	t.ipAddr = USER_INPUT_IP_ADDR
	t.loadingAnim.Start(nil,nil)
	
	go func(){
		fmt.Printf("Connecting to '%s'\n", t.ipAddr)
		ClientManager.InputHandler = func(mt int, msg []byte, err error, c *GC.Client) (bool) {
			if msg[0] == MAP_REQUEST {
				t.mapData = msg[1:]
			}
			return true
		}
		err := Client.MakeConn(t.ipAddr)
		CheckErr(err)
		time.Sleep(time.Second)
		data := []byte{MAP_REQUEST}
		fmt.Println("sending: ", data)
		Client.Send(data)
	}()
}
func (t *Connecting) Stop(g *TerraNomina, newState int) {
	fmt.Print("Connecting  -------->")
	t.loadingAnim.Stop(nil,nil)
}
func (t *Connecting) Update(screen *ebiten.Image) error {
	if len(t.mapData) > 0 {
		go t.LoadWorld(t.mapData)
		t.mapData = nil
	}
	
	t.loadingAnim.Update(t.parent.frame)
	t.background.DrawImageObj(screen)
	t.loadingAnim.DrawImageObj(screen)
	return nil
}

func (t *Connecting) LoadWorld(data []byte) {
	wrld, err := GE.GetWorldStructureFromBytes(0,0,XRES,YRES, data, RES+TILE_FILES, RES+STRUCTURE_FILES)
	CheckErr(err)
	WorldStructure = wrld
	
	t.parent.ChangeState(INGAME_STATE)
}