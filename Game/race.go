package Game

var Races []*Race

func InitRaces() {
	elv := &Race{"Elv", 0, []int8{0, 2, 2, 0}, []string{"Wood Elv", "High Elv"}}

	Races = []*Race{elv}
}

type Race struct {
	name       string
	imgID      int
	attributes []int8
	subraces   []string //will change later, placeholder
}
