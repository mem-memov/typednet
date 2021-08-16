package typednet

type storage interface {
	Has(uint) (bool, error)
	Create() (uint, error)
	ReadSources(target uint) ([]uint, error)
	ReadTargets(source uint) ([]uint, error)
	Connect(source uint, target uint) error
	Disconnect(source uint, target uint) error
	Delete(source uint) error
}
