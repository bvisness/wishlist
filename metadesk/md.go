// Generated from the official Metadesk reference. DO NOT EDIT!
package metadesk

// #include "md.h"
import "C"

var defaultArena = C.MD_ArenaAlloc()

func StringFromNodeKind(kind C.MD_NodeKind) string {
	_ret := C.MD_StringFromNodeKind(kind)
	return GoStr(_ret)
}

func StringListFromNodeFlags(flags C.MD_NodeFlags) C.MD_String8List {
	_ret := C.MD_StringListFromNodeFlags(defaultArena, flags)
	return _ret
}

func ParseResultZero() C.MD_ParseResult {
	_ret := C.MD_ParseResultZero()
	return _ret
}

func ParseNodeSet(_string string, offset C.MD_u64, parent *C.MD_Node, rule C.MD_ParseSetRule) C.MD_ParseResult {
	__string := Str(defaultArena, _string)
	_ret := C.MD_ParseNodeSet(defaultArena, __string, offset, parent, rule)
	return _ret
}

func ParseOneNode(_string string, offset C.MD_u64) C.MD_ParseResult {
	__string := Str(defaultArena, _string)
	_ret := C.MD_ParseOneNode(defaultArena, __string, offset)
	return _ret
}

func ParseWholeString(filename string, contents string) C.MD_ParseResult {
	_filename := Str(defaultArena, filename)
	_contents := Str(defaultArena, contents)
	_ret := C.MD_ParseWholeString(defaultArena, _filename, _contents)
	return _ret
}

func ParseWholeFile(filename string) C.MD_ParseResult {
	_filename := Str(defaultArena, filename)
	_ret := C.MD_ParseWholeFile(defaultArena, _filename)
	return _ret
}

func MakeErrorMarkerNode(parse_contents string, offset C.MD_u64) *C.MD_Node {
	_parse_contents := Str(defaultArena, parse_contents)
	_ret := C.MD_MakeErrorMarkerNode(defaultArena, _parse_contents, offset)
	return _ret
}

func MakeNodeError(node *C.MD_Node, kind C.MD_MessageKind, str string) *C.MD_Message {
	_str := Str(defaultArena, str)
	_ret := C.MD_MakeNodeError(defaultArena, node, kind, _str)
	return _ret
}

func MakeTokenError(parse_contents string, token C.MD_Token, kind C.MD_MessageKind, str string) *C.MD_Message {
	_parse_contents := Str(defaultArena, parse_contents)
	_str := Str(defaultArena, str)
	_ret := C.MD_MakeTokenError(defaultArena, _parse_contents, token, kind, _str)
	return _ret
}

func MessageListPush(list *C.MD_MessageList, error *C.MD_Message) {
	C.MD_MessageListPush(list, error)
}

func MessageListConcat(list *C.MD_MessageList, to_push *C.MD_MessageList) {
	C.MD_MessageListConcat(list, to_push)
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
	_ret := C.MD_PushNewReference(defaultArena, list, target)
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

func TagArgFromString(node *C.MD_Node, tag_string string, tag_str_flags C.MD_MatchFlags, arg_string string, arg_str_flags C.MD_MatchFlags) *C.MD_Node {
	_tag_string := Str(defaultArena, tag_string)
	_arg_string := Str(defaultArena, arg_string)
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

func FormatMessage(loc C.MD_CodeLoc, kind C.MD_MessageKind, _string string) string {
	__string := Str(defaultArena, _string)
	_ret := C.MD_FormatMessage(defaultArena, loc, kind, __string)
	return GoStr(_ret)
}

func NodeMatch(a *C.MD_Node, b *C.MD_Node, flags C.MD_MatchFlags) bool {
	_ret := C.MD_NodeMatch(a, b, flags)
	return _ret == 0
}

func NodeDeepMatch(a *C.MD_Node, b *C.MD_Node, flags C.MD_MatchFlags) bool {
	_ret := C.MD_NodeDeepMatch(a, b, flags)
	return _ret == 0
}

func ExprBakeOprTableFromList(list *C.MD_ExprOprList) C.MD_ExprOprTable {
	_ret := C.MD_ExprBakeOprTableFromList(defaultArena, list)
	return _ret
}

func ExprOprFromKindString(table *C.MD_ExprOprTable, kind C.MD_ExprOprKind, s string) *C.MD_ExprOpr {
	_s := Str(defaultArena, s)
	_ret := C.MD_ExprOprFromKindString(table, kind, _s)
	return _ret
}

func ExprParse(op_table *C.MD_ExprOprTable, first *C.MD_Node, one_past_last *C.MD_Node) C.MD_ExprParseResult {
	_ret := C.MD_ExprParse(defaultArena, op_table, first, one_past_last)
	return _ret
}

