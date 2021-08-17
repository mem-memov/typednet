package typednet

import (
	"github.com/mem-memov/clew"
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	NewGraph(clew.NewGraph(clew.NewSliceStorage()))
}

func TestGraph_AddClass(t *testing.T) {
	s := clew.NewGraph(clew.NewSliceStorage())
	g := NewGraph(s)

	x, err := g.AddClass()
	if err != nil {
		t.Fail()
	}

	if x != 2 {
		t.Fail()
	}

	classes, err := g.GetClasses()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(classes, []uint{2}) {
		t.Errorf("%v", classes)
	}
}

func TestGraph_AddClass_MultipleTimes(t *testing.T) {
	s := clew.NewGraph(clew.NewSliceStorage())
	g := NewGraph(s)

	x, err := g.AddClass()
	if err != nil {
		t.Fail()
	}

	y, err := g.AddClass()
	if err != nil {
		t.Fail()
	}

	z, err := g.AddClass()
	if err != nil {
		t.Fail()
	}

	classes, err := g.GetClasses()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(classes, []uint{x, y, z}) {
		t.Errorf("%v", classes)
	}
}

func TestGraph_GetClasses_WhenNoAdded(t *testing.T) {
	s := clew.NewGraph(clew.NewSliceStorage())
	g := NewGraph(s)

	classes, err := g.GetClasses()
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(classes, []uint{}) {
		t.Errorf("%v", classes)
	}
}

func TestGraph_CreateInstance(t *testing.T) {
	s := clew.NewGraph(clew.NewSliceStorage())
	g := NewGraph(s)

	class, err := g.AddClass()
	if err != nil {
		t.Fail()
	}

	instance, err := g.CreateInstance(class)
	if err != nil {
		t.Fail()
	}

	sources, err := s.ReadSources(instance)
	if err != nil {
		t.Fail()
	}

	if len(sources) != 1 {
		t.Fail()
	}

	targets, err := s.ReadTargets(instance)
	if err != nil {
		t.Fail()
	}

	if len(targets) != 2 {
		t.Fail()
	}

	if targets[0] != class && targets[1] != class {
		t.Fail()
	}
}
