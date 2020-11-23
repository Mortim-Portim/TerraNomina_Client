package Game

import (
	"fmt"

	"github.com/mortim-portim/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func GetInGame(g *TerraNomina) *InGame {
	return &InGame{parent: g}
}

type InGame struct {
	parent *TerraNomina

	left, right, up, down, changer bool
	wrld                           *GE.WorldStructure

	playerAnim *GE.DayNightAnim
}

func (i *InGame) Init() {
	fmt.Println("Initializing InGame")
	Keyli.MappIDToKey(left_key_id, ebiten.KeyLeft)
	Keyli.MappIDToKey(right_key_id, ebiten.KeyRight)
	Keyli.MappIDToKey(up_key_id, ebiten.KeyUp)
	Keyli.MappIDToKey(down_key_id, ebiten.KeyDown)

	w := 16.0
	h := PLAYER_MODELL_HEIGHT * w
	//TODO hight
	anim, err := GE.GetDayNightAnimFromParams(XRES/2-w/2, YRES/2-h*0.75, w, h, F_DAYNIGHT+"/jump.txt", F_DAYNIGHT+"/jump.png")
	CheckErr(err)
	i.playerAnim = anim
}
func (i *InGame) Start(oldState int) {
	fmt.Print("--------> InGame     \n")
	i.wrld = WorldStructure
	w := i.wrld.GetTileS()
	h := PLAYER_MODELL_HEIGHT * w
	i.playerAnim.SetParams(XRES/2-w/2, YRES/2-h*0.75, w, h)
	i.wrld.SetLightLevel(30)
	
	Soundtrack.Play(SOUNDTRACK_MAIN)
}
func (i *InGame) Stop(newState int) {
	fmt.Print("InGame      -------->")
}
func (i *InGame) Update(screen *ebiten.Image) error {
	i.left, _ = Keyli.GetMappedKeyState(left_key_id)
	i.right, _ = Keyli.GetMappedKeyState(right_key_id)
	i.up, _ = Keyli.GetMappedKeyState(up_key_id)
	i.down, _ = Keyli.GetMappedKeyState(down_key_id)

	if i.parent.frame%MOVEMENT_UPDATE_PERIOD == 0 {
		vert := 0
		hori := 0
		if i.left || i.right {
			if i.left {
				hori = -1
			}
			if i.right {
				hori = 1
			}
		}
		if i.up || i.down {
			if i.up {
				vert = -1
			}
			if i.down {
				vert = 1
			}
		}

		if vert != 0 && hori != 0 {
			if i.changer {
				vert = 0
			} else {
				hori = 0
			}
		}
		x, y := i.wrld.Middle()
		if !i.wrld.Collides(x+hori, y+vert) {
			i.wrld.Move(hori, vert, true, false)
		}
		i.changer = !i.changer
	}
	i.wrld.UpdateLightLevel(1)

	i.playerAnim.Update(i.parent.frame)
	i.playerAnim.LightLevel = int16(i.wrld.GetLightLevel())

	i.wrld.Draw(screen)

	msg := fmt.Sprintf(`TPS: %0.2f`, ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)

	return nil
}
