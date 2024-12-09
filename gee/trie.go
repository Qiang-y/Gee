package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 匹配的路由，如/p/:lang
	part     string  // 路由中的一部分，该node真正的值，如 p 或 :lang
	children []*node // 子节点
	isWild   bool    // 精准匹配标志，当part以 : 或 * 开头时为true
}

func newNode(pattern string, part string) *node {
	n := &node{
		pattern:  pattern,
		part:     part,
		children: make([]*node, 0),
		isWild:   false,
	}
	if part[0] == ':' || part[0] == '*' {
		n.isWild = true
	}
	return n
}

// 工具，获取第一个匹配的子节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 工具，获取所有匹配的子节点，用于查找
func (n *node) matchChildren(part string) (children []*node) {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		if n.pattern != "" {
			panic("Route Conflict!")
		}
		n.pattern = pattern
		return
	}

	// 寻找下一个节点
	child := n.matchChild(parts[height])
	if child == nil {
		child = newNode(pattern, parts[height])
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

	children := n.matchChildren(parts[height])
	for _, child := range children {
		if result := child.search(parts, height+1); result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel((list))
	}
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}
