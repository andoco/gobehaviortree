package gobehaviortree

import "fmt"

func Example() {
	b := DefaultBuilder{}
	b.Node("seq", []*nodeBuilder{b.Node("act", nil)})
	t := b.Build()

	walkdf(t.Root, func(n *Node) { fmt.Printf("%s:", n.Name) })
	// Output: root:seq:act:
}
