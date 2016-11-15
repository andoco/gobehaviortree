package gobehaviortree

import "fmt"

func Example() {
	b := NewBuilder()
	t := b.Node("seq", b.Node("act")).Build()

	walkdf(t.Root, func(n *Node) { fmt.Printf("%s:", n.Name) })
	// Output: root:seq:act:
}
