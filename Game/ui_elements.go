package Game

import (
	"github.com/mortim-portim/GraphEng/GE"
	"github.com/hajimehoshi/ebiten"
	"image"
	"strings"
	"fmt"
)

var ERROR_UNKNOWN_IMAGE = "Unknown Image %s"
var ALL_IMAGES map[string]*image.Image
var ALL_EBITEN_IMAGES map[string]*ebiten.Image
func InitImages() {
	ALL_IMAGES = make(map[string]*image.Image)
	ALL_EBITEN_IMAGES = make(map[string]*ebiten.Image)
	img_files, err := GE.GetAllFiles(F_IMAGES)
	CheckErr(err)
	for _,f := range img_files {
		err, img := GE.LoadImg(f)
		if err == nil {
			ALL_IMAGES[f] = img
			eimg, err := GE.ImgToEbitenImg(img)
			if err == nil {
				ALL_EBITEN_IMAGES[f] = eimg
			}
		}
	}
}
func GetImages(path string) (*image.Image, *ebiten.Image, error) {
	img, err1 := GetImage(path)
	eimg, err2 := GetEbitenImage(path)
	if err1 != nil {return img, eimg, err1}
	return img, eimg, err2
}
func GetImage(path string) (*image.Image, error) {
	img, ok := ALL_IMAGES[path]
	if !ok {return nil, fmt.Errorf(ERROR_UNKNOWN_IMAGE, path)}
	return img, nil
}
func GetEbitenImage(path string) (*ebiten.Image, error) {
	img, ok := ALL_EBITEN_IMAGES[path]
	if !ok {return nil, fmt.Errorf(ERROR_UNKNOWN_IMAGE, path)}
	return img, nil
}
var ERROR_UNKNOWN_BUTTON = "Unknown Button %s"
var ALL_BUTTONS map[string]*GE.Button
func InitButtons() {
	ALL_BUTTONS = make(map[string]*GE.Button)
	path := F_BUTTONS+"/"
	files,err := GE.OSReadDir(path)
	CheckErr(err)
	names := make([]string, 0)
	for _,f := range(files) {
		name := strings.Split(f, "_")[0]
		if !containsS(names, name) {
			names = append(names, name)
			up, err1 := GetEbitenImage(path+name+"_u.png")
			dw, err2 := GetEbitenImage(path+name+"_d.png")
			if err1 == nil && err2 == nil {
				btn := GE.GetUpDownImageButton(up, dw, 0,0, 100, 100)
				btn.ChangeDrawDarkOnLeft = true
				ALL_BUTTONS[name] = btn
			}
		}
	}
}
func GetButtonImg(name string, up bool) (*ebiten.Image, error) {
	btn, ok := ALL_BUTTONS[name]
	if !ok {return nil, fmt.Errorf(ERROR_UNKNOWN_BUTTON, name)}
	if up {
		return btn.UpImg(), nil
	}
	return btn.DownImg(), nil
}
func GetButton(name string, X,Y,W,H float64, changeToDarkOnLeft bool) (*GE.Button, error) {
	obtn, ok := ALL_BUTTONS[name]
	if !ok {return nil, fmt.Errorf(ERROR_UNKNOWN_BUTTON, name)}
	btn := obtn.Copy()
	btn.Img.X = X
	btn.Img.Y = Y
	btn.Img.W = W
	btn.Img.H = H
	btn.ChangeDrawDarkOnLeft = changeToDarkOnLeft
	return btn, nil
}
//Returns true if e is in s
func containsS(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}