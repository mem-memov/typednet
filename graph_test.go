package typednet

import (
	"github.com/mem-memov/clew"
	"testing"
)

func TestNewGraph(t *testing.T) {
	NewGraph(clew.NewGraph(clew.NewSliceStorage()))
}

func TestGraph_AddClass(t *testing.T) {
	g := NewGraph(clew.NewGraph(clew.NewSliceStorage()))
}
