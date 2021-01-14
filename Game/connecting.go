package Game

import (
	"fmt"
	"time"

	//"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"

	"github.com/hajimehoshi/ebiten"
)

//Connects
//Sends: MAP_REQUEST
//Receives: [ MAP_REQUEST | (n)map bytes ]

func GetConnecting(g *TerraNomina) *Connecting {
	return &Connecting{parent: g}
}

type Connecting struct {
	background  *GE.ImageObj
	loadingAnim *GE.Animation
	parent      *TerraNomina
	oldState    int

	ipAddr  string
	mapData []byte
}

func (t *Connecting) Init() {
	fmt.Println("Initializing Connecting")
	lps := &GE.Params{}
	err := lps.LoadFromFile(F_CONNECTING + "/loading.txt")
	CheckErr(err)
	limg, err := GE.LoadEbitenImg(F_CONNECTING + "/loading.png")
	CheckErr(err)
	t.loadingAnim = GE.GetAnimationFromParams(0, 0, XRES, YRES, lps, limg)
	t.loadingAnim.Init(nil, nil)

	t.background, err = GE.LoadImgObj(F_CONNECTING+"/back.png", XRES, YRES, 0, 0, 0)
	CheckErr(err)
	
	sm,err := TNE.GetSmallWorld(0, 0, XRES, YRES, F_TILES, F_STRUCTURES, F_ENTITY)
	CheckErr(err)
	err = sm.ActivePlayer.SetPlayer(ActivePlayer)
	CheckErr(err)
	SmallWorld = sm
}
func (t *Connecting) Start(oldState int) {
	fmt.Print("--------> Connecting \n")
	t.oldState = oldState
	t.ipAddr = USER_INPUT_IP_ADDR
	t.loadingAnim.Start(nil, nil)

	go func() {
//		//fmt.Printf("Connecting to '%s'\n", t.ipAddr)
//		ClientManager.InputHandler = func(mt int, msg []byte, err error, c *GC.Client) bool {
//			//fmt.Printf("server send: msg: %v, err: %v\n", msg, err)
//			if msg[0] == MAP_REQUEST {
//				t.mapData = msg[1:]
//			}
//			return true
//		}
		err := Client.MakeConn(t.ipAddr)
		CheckErr(err)
		time.Sleep(time.Second)
		
		
		
//		data := []byte{MAP_REQUEST}
//		Client.Send(data)
//		Client.WaitForConfirmation()
//
//		data = append([]byte{CHAR_SEND}, LoadChar("char")...)
//		Client.Send(data)
	}()
}
func (t *Connecting) Stop(newState int) {
	fmt.Print("Connecting  -------->")
	t.loadingAnim.Stop(nil, nil)
}
func (t *Connecting) Update(screen *ebiten.Image) error {
	if SmallWorld.HasWorldStruct() {
		t.parent.ChangeState(INGAME_STATE)
	}

	t.loadingAnim.Update(t.parent.frame)
	t.background.DrawImageObj(screen)
	t.loadingAnim.DrawImageObj(screen)
	return nil
}
