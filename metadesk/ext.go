package metadesk

// #include "md.h"
import "C"

import "strings"

func (n *Node) Children() []*Node {
	var res []*Node
	for node := n.FirstChild; !node.IsNil(); node = node.Next {
		res = append(res, node)
	}
	return res
}

func (n *Node) ChildFromString(child_string string, flags MatchFlags) *Node {
	return n.md.ChildFromString(n, child_string, flags)
}

func (n *Node) IsNil() bool {
	return n.md.NodeIsNil(n)
}

// The string generation functions are weird and mutate a string list...
// it's much easier to get the desired experience by manually binding them.

func DebugDumpFromNode(node *Node, indent bool, indent_string string, flags GenerateFlags) string {
	return defaultInstance.DebugDumpFromNode(node, indent, indent_string, flags)
}

func (md *Metadesk) DebugDumpFromNode(node *Node, indent bool, indent_string string, flags GenerateFlags) string {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_indent := C.int(mdBool(indent)) // weird, but whatever
	_indent_string := mdStr(md.a, indent_string)
	_flags := C.MD_GenerateFlags(flags)

	var out C.MD_String8List
	C.MD_DebugDumpFromNode(md.a, &out, _node, _indent, _indent_string, _flags)

	return strings.Join(goStrList(out), "")
}

func (n *Node) DebugDump(indent bool, indent_string string, flags GenerateFlags) string {
	return n.md.DebugDumpFromNode(n, indent, indent_string, flags)
}
