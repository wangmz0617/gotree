package tree

import (
	"github.com/fredericlemoine/bitset"
)

type Edge struct {
	left, right *Node   // Left and right nodes
	length      float64 // length of branch
	support     float64 // -1 if no support
	// a Bit at index i in the bitset corresponds to the position of the tip i
	//left:0/right:1 .
	// i is the index of the tip in the sorted tip name array
	bitset *bitset.BitSet // Bitset of length Number of taxa each
}

/* Edge functions */
/******************/

func (e *Edge) setLeft(left *Node) {
	e.left = left
}
func (e *Edge) setRight(right *Node) {
	e.right = right
}
func (e *Edge) SetLength(length float64) {
	e.length = length
}

func (e *Edge) SetSupport(support float64) {
	e.support = support
}

func (e *Edge) Length() float64 {
	return e.length
}

func (e *Edge) Support() float64 {
	return e.support
}

func (e *Edge) DumpBitSet() string {
	if e.bitset == nil {
		return "nil"
	}
	return e.bitset.DumpAsBits()
}