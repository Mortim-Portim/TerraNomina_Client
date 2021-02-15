package Game

import (
	"github.com/mortim-portim/GameConn/GC"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
	cmp "github.com/mortim-portim/GraphEng/Compression"

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
	limg, err := GE.LoadEbitenImg(F_CONNECTING + "/loading.png")
	CheckErr(err)
	t.loadingAnim = GE.GetAnimationFromParams(0, 0, XRES, YRES, lps, limg)
	t.loadingAnim.Init(nil, nil)

	t.background, err = GE.LoadImgObj(F_CONNECTING+"/back.png", XRES, YRES, 0, 0, 0)
	CheckErr(err)
	
	ClientManager.OnCloseConnection = func(){
		close(ServerClosing)
	}
}
func (t *Connecting) Start(oldState int) {
	Print("--------> Connecting \n")
	t.loadingAnim.Start(nil, nil)
	
	ServerClosing = make(chan bool)
	GC.PRINT_LOG = false
	t.oldState = oldState
	t.ipAddr = USER_INPUT_IP_ADDR
	
	sm,err := TNE.GetSmallWorld(0, 0, XRES, YRES, F_TILES, F_STRUCTURES, F_ENTITY, nil)
	CheckErr(err)
	sm.RegisterOnEntityChangeListeners()
	
	ple, err := sm.Ef.GetByName("Dwarf")
	CheckErr(err)
	OwnPlayer = &TNE.Player{Race:&TNE.Race{Entity:ple}}
	SmallWorld = sm

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
		CheckErr(err)
	}()
}
func (t *Connecting) Stop(newState int) {
	Print("Connecting  -------->")
	t.loadingAnim.Stop(nil, nil)
}
func (t *Connecting) Update(screen *ebiten.Image) error {
	if t.SVACIDs != 0 && len(ClientManager.SyncvarsByACID) == t.SVACIDs {
		SmallWorld.GetRegistered(ClientManager)
		Printf("%v SyncVars registered\n", t.SVACIDs)
		t.SVACIDs = 0
	}
	if SmallWorld.HasWorldStruct() {
		Println("WorldStructure received, setting player and reassigning all entities")
		SmallWorld.ActivePlayer.SetPlayer(OwnPlayer)
		SmallWorld.ActivePlayer.UpdateSyncVars(ClientManager)
		Client.WaitForConfirmation()
		t.parent.ChangeState(INGAME_STATE)
	}

	t.loadingAnim.Update(t.parent.frame)
	t.background.DrawImageObj(screen)
	t.loadingAnim.DrawImageObj(screen)
	return nil
}
