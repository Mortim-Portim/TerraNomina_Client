package Game

import (
	"io/ioutil"
	"os"
)

var Classes []*Class = []*Class{
	{"Fighter", []string{"Berserker", "Defender"}},
	{"Knight", []string{"Heavy Armor", "Combat Rager", "Guard", "Paladin"}},
}

type Class struct {
	name     string
	subclass []string
}

var Races []*Race = []*Race{
	{"Elv", 1, []uint8{0, 2, 2, 0}, []string{"Wood Elv", "High Elv"}},
	{"Human", 0, []uint8{1, 1, 1, 1}, []string{"Urban Human", "Country-Side Human", "Mountain Tribe"}},
	{"Half-Elv", 2, []uint8{0, 2, 0, 2}, []string{"Dark Elv", "City Elv"}},
	{"Ork", 3, []uint8{3, 0, 0, 0}, []string{"Mountain Ork", "Cave Ork"}},
}

type Race struct {
	name       string
	imgID      int
	attributes []uint8
	subraces   []string //will change later, placeholder
}

func LoadChar(name string) []byte {
	file, err := ioutil.ReadFile(F_CHARACTER + name + ".char")
	CheckErr(err)

	return file
}

func SaveChar(name string, attribute []uint8) {
	file, err := os.Create(F_CHARACTER + name + ".char")
	CheckErr(err)

	defer file.Close()

	bytearray := make([]byte, len(attribute))

	for i, attrib := range attribute {
		bytearray[i] = byte(attrib)
	}

	file.Write(bytearray)
}
