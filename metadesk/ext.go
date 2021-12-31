package metadesk

// #include "md.h"
import "C"

import (
	"strings"
)

func (n *Node) Children() []*Node {
	var res []*Node
	for node := n.FirstChild; !node.IsNil(); node = node.Next {
		res = append(res, node)
	}
	return res
}

// Finds the child index of 'node', with '0' being the first child, '1' being the second, and so on.
func (n *Node) Index() int {
	return n.md.IndexFromNode(n)
}

// Finds the highest-most node in the parent chain of 'node', starting with 'node'. If 'node' has no parent, then 'node' is returned.
func (n *Node) Root() *Node {
	return n.md.RootFromNode(n)
}

// Finds a child of 'node' with a string matching 'child_string', where the rules of matching are determined by 'flags'.
func (n *Node) ChildFromString(child_string string, flags MatchFlags) *Node {
	return n.md.ChildFromString(n, child_string, flags)
}

// Finds a tag on 'node' with a string matching 'tag_string', where the rules of matching are determined by 'flags'.
func (n *Node) TagFromString(tag_string string, flags MatchFlags) *Node {
	return n.md.TagFromString(n, tag_string, flags)
}

// Finds a child of 'node' with an index matching 'i'. Returns a nil node pointer if no such child is found.
func (n *Node) ChildFromIndex(i int) *Node {
	return n.md.ChildFromIndex(n, i)
}

// Finds a tag on 'node' with an index matching 'i'. Returns a nil node pointer if no such tag is found.
func (n *Node) TagFromIndex(i int) *Node {
	return n.md.TagFromIndex(n, i)
}

// Finds the 'i'th tag argument of the tag matching 'tag_string' on 'node', with the matching on 'tag_string' being controlled by 'flags'. Returns a nil node pointer if no such node was found.
func (n *Node) TagArgFromIndex(tag_string string, flags MatchFlags, i int) *Node {
	return n.md.TagArgFromIndex(n, tag_string, flags, i)
}

// Finds the tag argument with a string matching 'arg_string', of the tag matching 'tag_string', on 'node'. Matching 'tag_string' is controlled by 'tag_str_flags'. Matching 'arg_string' is controlled by 'arg_str_flags'. Returns a nil node pointer if no such node was found.
func (n *Node) TagArgFromString(tag_string string, tag_str_flags MatchFlags, arg_string string, arg_str_flags MatchFlags) *Node {
	return n.md.TagArgFromString(n, tag_string, tag_str_flags, arg_string, arg_str_flags)
}

// Returns '1' if 'node' has a child with a string matching 'string', with the matching rules being controlled by 'flags', or '0' otherwise.
func (n *Node) HasChild(_string string, flags MatchFlags) bool {
	return n.md.NodeHasChild(n, _string, flags)
}

// Returns '1' if 'node' has a tag with a string matching 'tag_string', with the matching rules being controlled by 'flags', or '0' otherwise.
func (n *Node) HasTag(tag_string string, flags MatchFlags) bool {
	return n.md.NodeHasTag(n, tag_string, flags)
}

// Returns the number of children of 'node'.
func (n *Node) ChildCount() int {
	return n.md.ChildCountFromNode(n)
}

// Returns the number of tags on 'node'.
func (n *Node) TagCountFromNode() int {
	return n.md.TagCountFromNode(n)
}

// Returns '1' if the 'node' is nil, or '0' otherwise. A nil node pointer is not equivalent to a null pointer. It can still be dereferenced, and is treated as a dummy placeholder node value.
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
