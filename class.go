package typednet

type Class struct{
	node uint
	storage storage
}

func newClass(node uint, storage storage) Class {
	return Class{
		node: node,
		storage: storage,
	}
}

func (c Class) add() (Instance, error) {

}

func (c Class) toInteger() uint {
	return uint(c)
}
