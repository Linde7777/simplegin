package simplegin

import (
	"fmt"
	"strings"
)

/*
given request:
GET - /
GET - /search
GET - /support
GET - /support/JetBrains

create such a tree:
`GET`->`/`->`search`

	|
	-> `support`->`JetBrains`

for node `/`, its pattern is "/", its part is ""
for node `search`, its pattern is "/search", its part is "search"
for node `JetBrains`, its pattern is "/support/JetBrains", its part is "JetBrains"
*/
type node struct {
	pattern    string
	part       string
	children   []*node
	fuzzyMatch bool // if the part contain `:` or `*`, it will be set to true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, fuzzyMatch=%t}", n.pattern, n.part, n.fuzzyMatch)
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// strictly match or fuzzy match
		if child.part == part || child.fuzzyMatch {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.fuzzyMatch {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// the param `height` will be used in recursion
func (n *node) insert(pattern string, parts []string, height int) {
	// all the parts has been checked, quit the recursion
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, fuzzyMatch: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
