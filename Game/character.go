package Game

import (
	"fmt"
	"io/ioutil"
	"os"
)

var Classes []*Class = []*Class{
	{"Fighter", []string{"Berserker", "Defender"}},
	{"Knight", []string{"Heavy Armor", "Combat Rager", "Guard", "Paladin"}},
}

type Class struct {
	Name     string
	Subclass []string
}

const ABIL_STRENGTH = 0
const ABIL_DEXTERITY = 1
const ABIL_INTELLIGENCE = 2
const ABIL_CHARISMA = 3

const SCORE_STRENGTH = 0
const SCORE_DEXTERITY = 1
const SCORE_INTELLIGENCE = 2
const SCORE_CHARISMA = 3
const SCORE_ENDURANCE = 4
const SCORE_PERSUASION = 5
const SCORE_DECEPTION = 6
const SCORE_PERFORMANCE = 7
const SCORE_INSIGHT = 8
const SCORE_THIEVERY = 9
const SCORE_STEALTH = 10
const SCORE_ACROBATICS = 11
const SCORE_NATURE = 12
const SCORE_ARCANA = 13
const SCORE_PERCEPTION = 14
const SCORE_CRAFTSMANSHIP = 15
const SCORE_DUNGEONEERING = 16

var Races []*Race = []*Race{
	{"Elv", []int8{0, 2, 2, 0}, []int{SCORE_NATURE, SCORE_ACROBATICS, SCORE_PERCEPTION, SCORE_INSIGHT, SCORE_INTELLIGENCE}, 1, []string{"Wood Elv", "High Elv"}},
	{"Human", []int8{1, 1, 1, 1}, []int{}, 4, []string{"Urban Human", "Country-Side Human", "Mountain Tribe"}},
	{"Half-Elv", []int8{0, 2, 0, 2}, []int{SCORE_ACROBATICS, SCORE_PERCEPTION, SCORE_INSIGHT, SCORE_DUNGEONEERING}, 1, []string{"Dark Elv", "City Elv"}},
	{"Ork", []int8{3, 0, 0, 0}, []int{SCORE_STRENGTH, SCORE_ENDURANCE}, 1, []string{"Mountain Ork", "Cave Ork"}},
	{"Goblin", []int8{-1, 2, 0, 2}, []int{SCORE_STEALTH, SCORE_THIEVERY, SCORE_ACROBATICS, SCORE_DECEPTION, SCORE_PERCEPTION}, 1, []string{"Ravin Goblin", "Sever Goblin"}},
	{"Dwarf", []int8{+2, 0, +2, 0}, []int{SCORE_STRENGTH, SCORE_CRAFTSMANSHIP, SCORE_DUNGEONEERING}, 1, []string{"Hill Dwarf", "Mountain Dwarf"}},
	{"Halfling", []int8{-1, 0, +1, +3}, []int{SCORE_PERSUASION, SCORE_DECEPTION, SCORE_DUNGEONEERING}, 1, []string{"Rock Halfling", "Forest Halfling"}},
}

type Race struct {
	Name       string
	Attributes []int8
	Profencies []int
	Extraprof  int
	Subraces   []string //will change later, placeholder
}

func LoadChar(name string) []byte {
	file, err := ioutil.ReadFile(F_CHARACTER + name + ".char")
	CheckErr(err)

	return file
}

func SaveChar(name string, race int8, class int8, attribute []int8, profencies []int8) {
	file, err := os.Create(F_CHARACTER + "/" + name + ".char")
	CheckErr(err)

	defer file.Close()

	file.Truncate(0)

	byteattrib := make([]byte, len(attribute))
	for i, attrib := range attribute {
		byteattrib[i] = byte(attrib)
	}

	byteprof := make([]byte, len(profencies))
	for i, prof := range profencies {
		byteprof[i] = byte(prof)
	}

	bytearray := make([]byte, 0)
	bytearray = append(bytearray, byte(race), byte(class))
	bytearray = append(bytearray, byteattrib...)
	bytearray = append(bytearray, byteprof...)

	fmt.Println(bytearray)
	file.Write(bytearray)
}
