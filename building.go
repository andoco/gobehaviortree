package gobehaviortree

type Tree struct {
	Root *Node
}

type Builder interface {
	Node(name string, children []nodeBuilder) *nodeBuilder
	Build() Tree
}

type DefaultBuilder struct {
	rootBuilder *nodeBuilder
}

func (b *DefaultBuilder) Node(name string, children []*nodeBuilder) *nodeBuilder {
	nb := &nodeBuilder{name: name, children: children}
	b.rootBuilder = nb
	return nb
}

func (b DefaultBuilder) Build() Tree {
	root := NewNode("root")
	buildWalker(b.rootBuilder, root)

	return Tree{Root: root}
}

type nodeBuilder struct {
	name     string
	children []*nodeBuilder
	args     map[string]interface{}
}

func buildWalker(nb *nodeBuilder, parent *Node) {
	n := NewNode(nb.name)
	parent.Children = append(parent.Children, n)

	for _, cnb := range nb.children {
		buildWalker(cnb, n)
	}
}

func walkdf(current *Node, f func(*Node)) {
	f(current)

	for _, n := range current.Children {
		walkdf(n, f)
	}
}
