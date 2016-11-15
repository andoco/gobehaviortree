package gobehaviortree

type Tree struct {
	Root *Node
}

type Builder interface {
	Node(name string, children []Builder) *Builder
	Build() Tree
}

type DefaultBuilder struct {
	name     string
	children []*DefaultBuilder
	args     map[string]interface{}
}

func (b *DefaultBuilder) Node(name string, children []*DefaultBuilder) *DefaultBuilder {
	nb := &DefaultBuilder{name: name, children: children}
	return nb
}

func (b DefaultBuilder) Build() Tree {
	root := NewNode("root")
	buildWalker(&b, root)

	return Tree{Root: root}
}

func buildWalker(nb *DefaultBuilder, parent *Node) {
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
