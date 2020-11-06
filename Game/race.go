package Game

var Races []*Race

func InitRaces() {
	human := &Race{"Human", 0, []int8{1, 1, 1, 1}, []string{"Urban Human", "Country-Side Human", "Mountain Tribe"}}
	elv := &Race{"Elv", 1, []int8{0, 2, 2, 0}, []string{"Wood Elv", "High Elv"}}
	halfelv := &Race{"Half-Elv", 2, []int8{0, 2, 0, 2}, []string{"Dark Elv", "City Elv"}}
	ork := &Race{"Ork", 3, []int8{3, 0, 0, 0}, []string{"Mountain Ork", "Cave Ork"}}

	Races = []*Race{human, elv, halfelv, ork}
}

type Race struct {
	name       string
	imgID      int
	attributes []int8
	subraces   []string //will change later, placeholder
}
