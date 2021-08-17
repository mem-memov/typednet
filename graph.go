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
	hops := len(classes) + 1

	for i := hops; i > 0; i-- {
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
		if classMark != uint(1) {
			return []uint{}, fmt.Errorf("Class mark not 1")
		}
	}

	getLastInChain := func(neighbour uint) (uint, error) {
		for {
			neighbours, err := g.storage.ReadTargets(neighbour)
			if err != nil {
				return 0, err
			}
			if len(neighbours) > 1 {
				return 0, fmt.Errorf("too many targets at Class chaine source %d", neighbour)
			}

			if len(neighbours) == 0 {
				return neighbour, nil
			}

			neighbour = neighbours[0]
		}
	}

	chains, err := g.storage.ReadTargets(classMark)
	if err != nil {
		return []uint{}, err
	}

	classes := make([]uint, 0)

	for _, chain := range chains {

		last, err := getLastInChain(chain)
		if err != nil {
			return []uint{}, err
		}

		classes = append(classes, last)
	}

	return classes, nil
}

func (g *Graph) CreateInstance(class Class) (Instance, error) {
	root, err := g.storage.Create()
	if err != nil {
		return Instance{}, err
	}

	incoming, err := g.storage.Create()
	if err != nil {
		return Instance{}, err
	}

	outgoing, err := g.storage.Create()
	if err != nil {
		return Instance{}, err
	}

	err = g.storage.Connect(root, class.toInteger())
	if err != nil {
		return Instance{}, err
	}

	err = g.storage.Connect(incoming, root)
	if err != nil {
		return Instance{}, err
	}

	err = g.storage.Connect(root, outgoing)
	if err != nil {
		return Instance{}, err
	}

	return newInstance(root, class, incoming, outgoing), nil
}

func (g *Graph) ReadIncoming(instance uint) (map[uint][]uint, error) {
	byClass := make(map[uint][]uint)

	incoming, err := g.storage.ReadSources(instance)
	if err != nil {
		return byClass, err
	}

	if len(incoming) != 1 {
		return byClass, fmt.Errorf("incoming collection invalid at target Instance %d", instance)
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
			return byClass, fmt.Errorf("incoming collection invalid at source Instance %d", instance)
		}

		nodeInstance := nodeInstances[0]

		nodeInstanceTargets, err := g.storage.ReadTargets(nodeInstance)
		if err != nil {
			return byClass, err
		}

		if len(nodeInstanceTargets) != 2 {
			return byClass, fmt.Errorf("invalid Instance target number at source Instance %d", nodeInstance)
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

//func (g *Graph) ReadOutgoing(Instance uint) (map[uint][]uint, error) {
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
