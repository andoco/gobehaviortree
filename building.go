package gobehaviortree

type Tree struct {
	Root *Node
}

func NewBuilder() *DefaultBuilder {
	return &DefaultBuilder{}
}

type DefaultBuilder struct {
	name     string
	children []DefaultBuilder
	args     map[string]interface{}
}

func (b DefaultBuilder) Node(name string, children ...DefaultBuilder) *DefaultBuilder {
	return &DefaultBuilder{name: name, children: children}
}

func (b DefaultBuilder) Build() Tree {
	root := NewNode("root")
	buildWalker(b, root)

	return Tree{Root: root}
}

func buildWalker(b DefaultBuilder, parent *Node) {
	n := NewNode(b.name)
	parent.Children = append(parent.Children, n)

	for _, cb := range b.children {
		buildWalker(cb, n)
	}
}

func walkdf(current *Node, f func(*Node)) {
	f(current)

	for _, n := range current.Children {
		walkdf(n, f)
	}
}
