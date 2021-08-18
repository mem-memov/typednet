package typednet

type Classes struct{
	storage storage
}

func newClasses(storage storage) Classes {
	return Classes{
		storage: storage,
	}
}

func (c Classes) add() (Class, error) {

}
