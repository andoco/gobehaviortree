package gobehaviortree

import "fmt"

func Example() {
	b := DefaultBuilder{}
	t := b.Node("seq", []*DefaultBuilder{b.Node("act", nil)}).Build()

	walkdf(t.Root, func(n *Node) { fmt.Printf("%s:", n.Name) })
	// Output: root:seq:act:
}
