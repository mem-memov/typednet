package typednet

import "fmt"

type Graph struct {
	storage storage
}

func NewGraph(storage storage) *Graph {
	return &Graph{
		storage: storage,
	}
}

func (g *Graph) AddClass() (uint, error) {
	classes, err := g.GetClasses()
	if err != nil {
		return 0, err
	}

	classMark := uint(1)

	current := classMark

	for i := len(classes); i > 0; i-- {
		next, err := g.storage.Create()
		if err != nil {
			return 0, err
		}

		err = g.storage.Connect(current, next)
		if err != nil {
			return 0, err
		}

		current = next
	}

	return current, nil
}

func (g *Graph) GetClasses() ([]uint, error) {
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

	neighbour := classMark

	for {
		neighbours, err := g.storage.ReadTargets(neighbour)
		if err != nil {
			return []uint{}, err
		}

		if len(neighbours) == 0 {
			break
		}

		for _, neighbour := range neighbours {
			neighbours, err := g.storage.ReadTargets(neighbour)
			if err != nil {
				return []uint{}, err
			}

			if len(neighbours) == 0 {
				classes = append(classes, neighbour)
			}
		}
	}

	return classes, nil
}

func (g *Graph) CreateInstance(class uint) (uint, error) {
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

		nodeInstance := nodeInstances[0]

		nodeInstanceTargets, err := g.storage.ReadTargets(nodeInstance)
		if err != nil {
			return byClass, err
		}

		if len(nodeInstanceTargets) != 2 {
			return byClass, fmt.Errorf("invalid instance target number at source instance %d", nodeInstance)
		}

		for _, nodeInstanceTarget := range nodeInstanceTargets {
			if nodeInstanceTarget == node {
				continue
			}
			nodeClass := nodeInstanceTarget
			byClass[nodeClass] = append(byClass[nodeClass], nodeInstance)
		}
	}

	return byClass, nil
}

//func (g *Graph) ReadOutgoing(instance uint) (map[uint][]uint, error) {
//
//}
//
//func (g *Graph) Connect(source uint, target uint) error {
//
//}
//
//func (g *Graph) Disconnect(source uint, target uint) error {
//
//}
//
//func (g *Graph) Delete(source uint) error {
//
//}
