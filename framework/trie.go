package framework

import "strings"

type Tree struct {
	root *node
}

type node struct {
	path      string
	children  []*node
	handlers  ControllerHandler
	wildChild bool
}

type HandlersChain []ControllerHandler

func NewNode() *node {
	return &node{}
}

func NewTree() *Tree {
	return &Tree{NewNode()}
}

func isWildSegment(path string) bool {
	return strings.HasPrefix(path, ":")
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.children) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.children
	}

	nodes := make([]*node, 0, len(n.children))
	//过滤
	for _, cnode := range n.children {
		if isWildSegment(cnode.path) {
			nodes = append(nodes, cnode)
		} else if cnode.path == segment {
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

//判断路由是否在子节点上
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]

}
