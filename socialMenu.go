package main

import (
	"fmt"
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/TN_Engine/TNE"
)

func GetNewSocialMenu(back_path string) *SocialMenu {
	sm := &SocialMenu{}
	bEimg, err := GetEbitenImage(back_path)
	CheckErr(err)
	sm.Back = GE.NewImageObj(nil, bEimg, 0, 0, 1, 1, 0)
	return sm
}

type SocialMenu struct {
	Back, NPC_Img, NPC_Name_Img *GE.ImageObj
	NPC_Name                    string

	NPC *TNE.SyncEntity
}

func (sm *SocialMenu) Start(e *TNE.SyncEntity) {
	sm.NPC = e
	sm.NPC_Name = e.Entity.Char.Name
	if e.Entity.ID >= 0 {
		OwnPlayer.StartInteraction(int16(e.Entity.ID))
	}

	var (
		SOCIALMENU_H = YRES * 0.6
		SOCIALMENU_W = (100 / 50) * SOCIALMENU_H
		SOCIALMENU_X = (XRES - SOCIALMENU_W) / 2
		SOCIALMENU_Y = (YRES - SOCIALMENU_H) / 2

		SOCIALMENU_NPC_IMG_X    = SOCIALMENU_X + SOCIALMENU_W*(75.0/100.0)
		SOCIALMENU_NPC_IMG_Y    = SOCIALMENU_Y + SOCIALMENU_H*(17.0/50.0)
		SOCIALMENU_NPC_IMG_BNDS = SOCIALMENU_W * (16.0 / 100.0)

		SOCIALMENU_NPC_NAME_X = SOCIALMENU_X + SOCIALMENU_W*(74.0/100.0)
		SOCIALMENU_NPC_NAME_Y = SOCIALMENU_Y + SOCIALMENU_H*(39.0/50.0)
		//SOCIALMENU_NPC_NAME_W = SOCIALMENU_W * (18.0 / 100.0)
		SOCIALMENU_NPC_NAME_H = SOCIALMENU_H * (6.0 / 50.0)
	)

	sm.Back.SetXYWH(SOCIALMENU_X, SOCIALMENU_Y, SOCIALMENU_W, SOCIALMENU_H)
	sm.NPC_Name_Img = GE.GetTextImage(sm.NPC_Name, SOCIALMENU_NPC_NAME_X, SOCIALMENU_NPC_NAME_Y, SOCIALMENU_NPC_NAME_H, GE.StandardFont, color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 0})
	npcEImgIcon, err := GetEbitenImage(F_NPCICON + fmt.Sprintf("/%v.png", e.Entity.ID))
	if err != nil {
		standardIcon, err := GetEbitenImage(F_NPCICON + "/standard.png")
		CheckErr(err)
		npcEImgIcon = standardIcon
	}
	sm.NPC_Img = GE.NewImageObj(nil, npcEImgIcon, SOCIALMENU_NPC_IMG_X, SOCIALMENU_NPC_IMG_Y, SOCIALMENU_NPC_IMG_BNDS, SOCIALMENU_NPC_IMG_BNDS, 0)
}
func (sm *SocialMenu) Stop() {
	if sm.NPC.Entity.ID >= 0 {
		OwnPlayer.StopInteraction(int16(sm.NPC.Entity.ID))
	}
}
func (sm *SocialMenu) Update() {

}
func (sm *SocialMenu) Draw(screen *ebiten.Image) {
	// if sm.Back != nil {

	// }
	// if sm.NPC_Img != nil {

	// }
	// if sm.NPC_Name_Img != nil {

	// }
	sm.Back.Draw(screen)
	sm.NPC_Img.Draw(screen)
	sm.NPC_Name_Img.Draw(screen)
}
