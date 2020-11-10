package Game

type classList []*Class

var Classes classList

func InitClasses() {
	fighter := &Class{"Fighter", []string{"Berserker", "Defender"}}
	knight := &Class{"Knight", []string{"Heavy Armor", "Combat Rager", "Guard", "Paladin"}}

	Classes = []*Class{fighter, knight}
}

func (classes classList) search(name string) *Class {
	for _, class := range classes {
		if class.name == name {
			return class
		}
	}

	return nil
}

type Class struct {
	name     string
	subclass []string
}
