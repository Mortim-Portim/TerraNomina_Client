package Game

import (
	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	cmp "github.com/mortim-portim/GraphEng/compression"
	"github.com/mortim-portim/TN_Engine/TNE"

	"github.com/hajimehoshi/ebiten"
)

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
	Println("Initializing Connecting")
	lps := &GE.Params{}
	err := lps.LoadFromFile(F_CONNECTING + "/loading.txt")
	CheckErr(err)
	limg, err := GetEbitenImage(F_CONNECTING + "/loading.png")
	CheckErr(err)
	t.loadingAnim = GE.GetAnimationFromParams(0, 0, XRES, YRES, lps, limg)
	t.loadingAnim.Init(nil, nil)

	backImg, backImgE, err := GetImages(F_CONNECTING + "/back.png")
	CheckErr(err)
	t.background = GE.NewImageObj(backImg, backImgE, 0, 0, XRES, YRES, 0)

	ClientManager.OnCloseConnection = func() {
		close(ServerClosing)
	}
}
func (t *Connecting) Start(oldState int) {
	Print("--------> Connecting \n")
	t.loadingAnim.Start(nil, nil)

	ServerClosing = make(chan bool)
	GC.PRINT_LOG_PRIORITY = 3
	t.oldState = oldState
	t.ipAddr = USER_INPUT_IP_ADDR

	go func() {
		Printf("Connecting to '%s'\n", t.ipAddr)
		ClientManager.InputHandler = func(mt int, msg []byte, err error, c *GC.Client) bool {
			if msg[0] == GC.BINARYMSG && len(msg) == 10 {
				if string(msg[1:8]) == TNE.NumberOfSVACIDs_Msg {
					t.SVACIDs = int(cmp.BytesToInt16(msg[8:10]))
					Printf("Waiting for %v SyncVars to be registered\n", t.SVACIDs)
				}
			}
			return true
		}

		err := Client.MakeConn(t.ipAddr)
		//HANDLE ERR
		CheckErr(err)
	}()
}
func (t *Connecting) Stop(newState int) {
	Print("Connecting  -------->")
	t.loadingAnim.Stop(nil, nil)
}
func (t *Connecting) Update() error {
	if t.SVACIDs != 0 && len(ClientManager.SyncvarsByACID) == t.SVACIDs {
		SmallWorld.GetRegistered(ClientManager)
		Printf("%v SyncVars registered\n", t.SVACIDs)
		t.SVACIDs = 0
	}
	if SmallWorld.HasWorldStruct() {
		Println("WorldStructure received, setting player and reassigning all entities")
		SmallWorld.ActivePlayer.SetPlayer(OwnPlayer)
		SmallWorld.ReassignAllEntities()
		SmallWorld.ActivePlayer.UpdateSyncVars(ClientManager)
		t.parent.ChangeState(INGAME_STATE)
	}

	t.loadingAnim.Update(t.parent.frame)
	return nil
}
func (t *Connecting) Draw(screen *ebiten.Image) {
	t.background.DrawImageObj(screen)
	t.loadingAnim.DrawImageObj(screen)
}
