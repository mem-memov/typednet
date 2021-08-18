package typednet

import "fmt"

type Graph struct {
	storage storage
	classMark uint
}

func NewGraph(storage storage) *Graph {
	return &Graph{
		storage: storage,
		classMark: uint(1),
	}
}

func (g *Graph) AddClass() (Class, error) {
	classMarkExists, err := g.storage.Has(g.classMark)
	if err != nil {
		return 0, err
	}

	if !classMarkExists {
		node, err := g.storage.Create()
		if err != nil {
			return 0, err
		}
		if node != uint(g.classMark) {
			return 0, fmt.Errorf("Class mark not 1")
		}
	}

	node, err := g.storage.Create()
	if err != nil {
		return 0, err
	}

	err = g.storage.Connect(node, g.classMark)
	if err != nil {
		return 0, err
	}

	return newClass(node), nil
}

func (g *Graph) GetClasses() ([]Class, error) {

	classMarkExists, err := g.storage.Has(g.classMark)
	if err != nil {
		return []Class{}, err
	}

	if !classMarkExists {
		return []Class{}, nil
	}

	sources, err := g.storage.ReadSources(g.classMark)
	if err != nil {
		return []Class{}, err
	}

	classes := make([]Class, 0)

	for _, source := range sources {
		classes = append(classes, newClass(source))
	}

	return classes, nil
}

func (g *Graph) GetClassInstances(class Class) ([]Instance, error) {


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

func (g *Graph) ReadIncoming(instance Instance) (map[Class][]Instance, error) {
	byClass := make(map[Class][]Instance)

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
			nodeClass := newClass(nodeInstanceTarget)
			byClass[nodeClass] = append(byClass[nodeClass], newInstance(nodeInstance, nodeClass, ))
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
