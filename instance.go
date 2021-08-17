package typednet

type Instance struct {
	root     uint
	class    Class
	incoming uint
	outgoing uint
}

func newInstance(root uint, class Class, incoming uint, outgoing uint) Instance {
	return Instance{
		root:     root,
		class:    class,
		incoming: incoming,
		outgoing: outgoing,
	}
}

//func (i Instance) getIncoming() ([]Instance, error) {
//
//}
//
//func (i Instance) getOutgoing() ([]Instance, error) {
//
//}
