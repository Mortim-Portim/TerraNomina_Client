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
	{"Elv", []int8{0, 2, 2, 0}, []string{"Wood Elv", "High Elv"}},
	{"Human", []int8{1, 1, 1, 1}, []string{"Urban Human", "Country-Side Human", "Mountain Tribe"}},
	{"Half-Elv", []int8{0, 2, 0, 2}, []string{"Dark Elv", "City Elv"}},
	{"Ork", []int8{3, 0, 0, 0}, []string{"Mountain Ork", "Cave Ork"}},
	{"Goblin", []int8{-1, 2, 0, 2}, []string{"Ravin Goblin", "Sever Goblin"}},
	{"Dwarf", []int8{+2, 0, +2, 0}, []string{"Hill Dwarf", "Mountain Dwarf"}},
	{"Halfling", []int8{0, -1, +1, +3}, []string{"Rock Halfling", "Forest Halfling"}},
}

type Race struct {
	name       string
	attributes []int8
	subraces   []string //will change later, placeholder
}

func LoadChar(name string) []byte {
	file, err := ioutil.ReadFile(F_CHARACTER + name + ".char")
	CheckErr(err)

	return file
}

func SaveChar(name string, attribute []int8) {
	file, err := os.Create(F_CHARACTER + name + ".char")
	CheckErr(err)

	defer file.Close()

	bytearray := make([]byte, len(attribute))

	for i, attrib := range attribute {
		bytearray[i] = byte(attrib)
	}

	file.Write(bytearray)
}
