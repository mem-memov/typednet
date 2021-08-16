package typednet

import "fmt"

type Graph struct{
	storage storage
}

func NewGraph(storage storage) *Graph {
	return &Graph{
		storage: storage,
	}
}

func (g *Graph) addClass() (uint, error) {
	classes, err := g.getClasses()
	if err != nil {
		return 0, err
	}

	last := classes[len(classes) - 1]

	next, err := g.storage.Create()
	if err != nil {
		return 0, err
	}

	err = g.storage.Connect(last, next)

	return next, nil
}

func (g *Graph) getClasses() ([]uint, error) {
	classMark := uint(1)

	classMarkExists, err := g.storage.Has(classMark)
	if err != nil {
		return []uint{}, err
	}

	if !classMarkExists {
		_, err = g.storage.Create()
		if err != nil {
			return []uint{}, err
		}
	}

	classes := make([]uint, 1)

	next := classMark

	for {
		targets, err := g.storage.ReadTargets(next)
		if err != nil {
			return []uint{}, err
		}

		if len(targets) == 0 {
			break
		}

		class := targets[0]

		classes = append(classes, class)

		next = class
	}

	return classes, nil
}

func (g *Graph) Create(class uint) (uint, error) {
	instance, err := g.storage.Create()
	if err != nil {
		return 0, err
	}

	incoming, err := g.storage.Create()
	if err != nil {
		return 0, err
	}

	outgoing, err := g.storage.Create()
	if err != nil {
		return 0, err
	}

	err = g.storage.Connect(instance, class)
	if err != nil {
		return 0, err
	}

	err = g.storage.Connect(incoming, instance)
	if err != nil {
		return 0, err
	}

	err = g.storage.Connect(instance, outgoing)
	if err != nil {
		return 0, err
	}

	return instance, nil
}

func (g *Graph) ReadIncoming(instance uint) (map[uint][]uint, error) {
	byClass := make(map[uint][]uint)

	incoming, err := g.storage.ReadSources(instance)
	if err != nil {
		return byClass, err
	}

	if len(incoming) != 1 {
		return byClass, fmt.Errorf("incoming collection invalid at target instance %d", instance)
	}

	nodes, err := g.storage.ReadSources(incoming[0])
	if err != nil {
		return byClass, err
	}

	for _, node := range nodes {
		nodeInstances, err := g.storage.ReadSources(node)
		if err != nil {
			return byClass, err
		}

		if len(nodeInstances) != 1 {
			return byClass, fmt.Errorf("incoming collection invalid at source instance %d", instance)
		}

		source, err := g.storage.ReadSources(node)
		if err != nil {
			return byClass, err
		}
	}

	return byClass, nil
}

func (g *Graph) ReadOutgoing(instance uint) (map[uint][]uint, error) {

}

func (g *Graph) Connect(source uint, target uint) error {

}

func (g *Graph) Disconnect(source uint, target uint) error {

}

func (g *Graph) Delete(source uint) error {

}
