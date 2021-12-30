// Generated from the official Metadesk reference. DO NOT EDIT!
package metadesk

// #include "md.h"
import "C"

var defaultArena = C.MD_ArenaAlloc()

func StringFromNodeKind(kind C.MD_NodeKind) string {
	_ret := C.MD_StringFromNodeKind(kind)
	return GoStr(_ret)
}

func StringListFromNodeFlags(arena *C.MD_Arena, flags C.MD_NodeFlags) C.MD_String8List {
	_ret := C.MD_StringListFromNodeFlags(arena, flags)
	return _ret
}

func NodeIsNil(node *C.MD_Node) bool {
	_ret := C.MD_NodeIsNil(node)
	return _ret == 0
}

func NilNode() *C.MD_Node {
	_ret := C.MD_NilNode()
	return _ret
}

func PushChild(parent *C.MD_Node, new_child *C.MD_Node) {
	C.MD_PushChild(parent, new_child)
}

func PushTag(node *C.MD_Node, tag *C.MD_Node) {
	C.MD_PushTag(node, tag)
}

func PushNewReference(list *C.MD_Node, target *C.MD_Node) *C.MD_Node {
	_ret := C.MD_PushNewReference(list, target)
	return _ret
}

func FirstNodeWithString(first *C.MD_Node, _string string, flags C.MD_MatchFlags) *C.MD_Node {
	__string := Str(defaultArena, _string)
	_ret := C.MD_FirstNodeWithString(first, __string, flags)
	return _ret
}

func NodeAtIndex(first *C.MD_Node, n int) *C.MD_Node {
	_n := C.int(n)
	_ret := C.MD_NodeAtIndex(first, _n)
	return _ret
}

func FirstNodeWithFlags(first *C.MD_Node, flags C.MD_NodeFlags) *C.MD_Node {
	_ret := C.MD_FirstNodeWithFlags(first, flags)
	return _ret
}

func IndexFromNode(node *C.MD_Node) int {
	_ret := C.MD_IndexFromNode(node)
	return int(_ret)
}

func RootFromNode(node *C.MD_Node) *C.MD_Node {
	_ret := C.MD_RootFromNode(node)
	return _ret
}

func ChildFromString(node *C.MD_Node, child_string string, flags C.MD_MatchFlags) *C.MD_Node {
	_child_string := Str(defaultArena, child_string)
	_ret := C.MD_ChildFromString(node, _child_string, flags)
	return _ret
}

func TagFromString(node *C.MD_Node, tag_string string, flags C.MD_MatchFlags) *C.MD_Node {
	_tag_string := Str(defaultArena, tag_string)
	_ret := C.MD_TagFromString(node, _tag_string, flags)
	return _ret
}

func ChildFromIndex(node *C.MD_Node, n int) *C.MD_Node {
	_n := C.int(n)
	_ret := C.MD_ChildFromIndex(node, _n)
	return _ret
}

func TagFromIndex(node *C.MD_Node, n int) *C.MD_Node {
	_n := C.int(n)
	_ret := C.MD_TagFromIndex(node, _n)
	return _ret
}

func TagArgFromIndex(node *C.MD_Node, tag_string string, flags C.MD_MatchFlags, n int) *C.MD_Node {
	_tag_string := Str(defaultArena, tag_string)
	_n := C.int(n)
	_ret := C.MD_TagArgFromIndex(node, _tag_string, flags, _n)
	return _ret
}

func TagArgFromString(node *C.MD_Node, tag_string string, tag_str_flags C.MD_MatchFlags, arg_string int, arg_str_flags C.MD_MatchFlags) *C.MD_Node {
	_tag_string := Str(defaultArena, tag_string)
	_arg_string := C.int(arg_string)
	_ret := C.MD_TagArgFromString(node, _tag_string, tag_str_flags, _arg_string, arg_str_flags)
	return _ret
}

func NodeHasChild(node *C.MD_Node, _string string, flags C.MD_MatchFlags) bool {
	__string := Str(defaultArena, _string)
	_ret := C.MD_NodeHasChild(node, __string, flags)
	return _ret == 0
}

func NodeHasTag(node *C.MD_Node, tag_string string, flags C.MD_MatchFlags) bool {
	_tag_string := Str(defaultArena, tag_string)
	_ret := C.MD_NodeHasTag(node, _tag_string, flags)
	return _ret == 0
}

func ChildCountFromNode(node *C.MD_Node) int {
	_ret := C.MD_ChildCountFromNode(node)
	return int(_ret)
}

func TagCountFromNode(node *C.MD_Node) int {
	_ret := C.MD_TagCountFromNode(node)
	return int(_ret)
}

func ResolveNodeFromReference(node *C.MD_Node) *C.MD_Node {
	_ret := C.MD_ResolveNodeFromReference(node)
	return _ret
}

func NodeNextWithLimit(node *C.MD_Node, opl *C.MD_Node) *C.MD_Node {
	_ret := C.MD_NodeNextWithLimit(node, opl)
	return _ret
}

func PrevCommentFromNode(node *C.MD_Node) string {
	_ret := C.MD_PrevCommentFromNode(node)
	return GoStr(_ret)
}

func NextCommentFromNode(node *C.MD_Node) string {
	_ret := C.MD_NextCommentFromNode(node)
	return GoStr(_ret)
}

func StringFromMessageKind(kind C.MD_MessageKind) string {
	_ret := C.MD_StringFromMessageKind(kind)
	return GoStr(_ret)
}

func FormatMessage(arena *C.MD_Arena, loc C.MD_CodeLoc, kind C.MD_MessageKind, _string string) string {
	__string := Str(defaultArena, _string)
	_ret := C.MD_FormatMessage(arena, loc, kind, __string)
	return GoStr(_ret)
}

func PrintMessage(file *FILE, loc C.MD_CodeLoc, kind C.MD_MessageKind, _string string) {
	__string := Str(defaultArena, _string)
	C.MD_PrintMessage(file, loc, kind, __string)
}

func NodeMatch(a *C.MD_Node, b *C.MD_Node, flags C.MD_MatchFlags) bool {
	_ret := C.MD_NodeMatch(a, b, flags)
	return _ret == 0
}

func NodeDeepMatch(a *C.MD_Node, b *C.MD_Node, flags C.MD_MatchFlags) bool {
	_ret := C.MD_NodeDeepMatch(a, b, flags)
	return _ret == 0
}

