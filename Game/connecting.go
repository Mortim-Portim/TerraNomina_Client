package Game

import (
	"fmt"
	"time"

	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
	cmp "github.com/mortim-portim/GraphEng/Compression"

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
	SVACIDs int
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

	t.SVACIDs = 0
	go func() {
		err := Client.MakeConn(t.ipAddr)
		CheckErr(err)
		time.Sleep(time.Second)
		
		fmt.Printf("Connecting to '%s'\n", t.ipAddr)
		ClientManager.InputHandler = func(mt int, msg []byte, err error, c *GC.Client) bool {
			if msg[0] == GC.BINARYMSG && len(msg) == 10 {
				if string(msg[1:8]) == TNE.NumberOfSVACIDs_Msg {
					t.SVACIDs = int(cmp.BytesToInt16(msg[8:10]))
				}
			}
			return true
		}
	}()
}
func (t *Connecting) Stop(newState int) {
	fmt.Print("Connecting  -------->")
	t.loadingAnim.Stop(nil, nil)
}
func (t *Connecting) Update(screen *ebiten.Image) error {
	if t.SVACIDs != 0 && len(ClientManager.SyncvarsByACID) == t.SVACIDs {
		
		
		
		t.parent.ChangeState(INGAME_STATE)
	}

	t.loadingAnim.Update(t.parent.frame)
	t.background.DrawImageObj(screen)
	t.loadingAnim.DrawImageObj(screen)
	return nil
}
