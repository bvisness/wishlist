// Generated from the official Metadesk reference. DO NOT EDIT!
package metadesk

// #include "md.h"
import "C"

var defaultArena = C.MD_ArenaAlloc()

func StringFromNodeKind(kind NodeKind) string {
	_kind := C.MD_NodeKind(kind)
	_ret := C.MD_StringFromNodeKind(_kind)
	return goStr(_ret)
}

func StringListFromNodeFlags(flags C.MD_NodeFlags) C.MD_String8List {
	_ret := C.MD_StringListFromNodeFlags(defaultArena, flags)
	return _ret
}

func ParseResultZero() C.MD_ParseResult {
	_ret := C.MD_ParseResultZero()
	return _ret
}

func ParseNodeSet(_string string, offset int, parent *Node, rule C.MD_ParseSetRule) C.MD_ParseResult {
	__string := mdStr(defaultArena, _string)
	_offset := C.MD_u64(offset)
	_parent := mdNodeP(defaultArena, parent)
	_ret := C.MD_ParseNodeSet(defaultArena, __string, _offset, _parent, rule)
	return _ret
}

func ParseOneNode(_string string, offset int) C.MD_ParseResult {
	__string := mdStr(defaultArena, _string)
	_offset := C.MD_u64(offset)
	_ret := C.MD_ParseOneNode(defaultArena, __string, _offset)
	return _ret
}

func ParseWholeString(filename string, contents string) C.MD_ParseResult {
	_filename := mdStr(defaultArena, filename)
	_contents := mdStr(defaultArena, contents)
	_ret := C.MD_ParseWholeString(defaultArena, _filename, _contents)
	return _ret
}

func ParseWholeFile(filename string) C.MD_ParseResult {
	_filename := mdStr(defaultArena, filename)
	_ret := C.MD_ParseWholeFile(defaultArena, _filename)
	return _ret
}

func MakeErrorMarkerNode(parse_contents string, offset int) *Node {
	_parse_contents := mdStr(defaultArena, parse_contents)
	_offset := C.MD_u64(offset)
	_ret := C.MD_MakeErrorMarkerNode(defaultArena, _parse_contents, _offset)
	return goNodeP(_ret)
}

func MakeNodeError(node *Node, kind C.MD_MessageKind, str string) *C.MD_Message {
	_node := mdNodeP(defaultArena, node)
	_str := mdStr(defaultArena, str)
	_ret := C.MD_MakeNodeError(defaultArena, _node, kind, _str)
	return _ret
}

func MakeTokenError(parse_contents string, token C.MD_Token, kind C.MD_MessageKind, str string) *C.MD_Message {
	_parse_contents := mdStr(defaultArena, parse_contents)
	_str := mdStr(defaultArena, str)
	_ret := C.MD_MakeTokenError(defaultArena, _parse_contents, token, kind, _str)
	return _ret
}

func MessageListPush(list *C.MD_MessageList, error *C.MD_Message) {
	C.MD_MessageListPush(list, error)
}

func MessageListConcat(list *C.MD_MessageList, to_push *C.MD_MessageList) {
	C.MD_MessageListConcat(list, to_push)
}

func NodeIsNil(node *Node) bool {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_NodeIsNil(_node)
	return _ret == 0
}

func NilNode() *Node {
	_ret := C.MD_NilNode()
	return goNodeP(_ret)
}

func PushChild(parent *Node, new_child *Node) {
	_parent := mdNodeP(defaultArena, parent)
	_new_child := mdNodeP(defaultArena, new_child)
	C.MD_PushChild(_parent, _new_child)
}

func PushTag(node *Node, tag *Node) {
	_node := mdNodeP(defaultArena, node)
	_tag := mdNodeP(defaultArena, tag)
	C.MD_PushTag(_node, _tag)
}

func PushNewReference(list *Node, target *Node) *Node {
	_list := mdNodeP(defaultArena, list)
	_target := mdNodeP(defaultArena, target)
	_ret := C.MD_PushNewReference(defaultArena, _list, _target)
	return goNodeP(_ret)
}

func FirstNodeWithString(first *Node, _string string, flags MatchFlags) *Node {
	_first := mdNodeP(defaultArena, first)
	__string := mdStr(defaultArena, _string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_FirstNodeWithString(_first, __string, _flags)
	return goNodeP(_ret)
}

func NodeAtIndex(first *Node, n int) *Node {
	_first := mdNodeP(defaultArena, first)
	_n := C.int(n)
	_ret := C.MD_NodeAtIndex(_first, _n)
	return goNodeP(_ret)
}

func FirstNodeWithFlags(first *Node, flags C.MD_NodeFlags) *Node {
	_first := mdNodeP(defaultArena, first)
	_ret := C.MD_FirstNodeWithFlags(_first, flags)
	return goNodeP(_ret)
}

func IndexFromNode(node *Node) int {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_IndexFromNode(_node)
	return int(_ret)
}

func RootFromNode(node *Node) *Node {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_RootFromNode(_node)
	return goNodeP(_ret)
}

func ChildFromString(node *Node, child_string string, flags MatchFlags) *Node {
	_node := mdNodeP(defaultArena, node)
	_child_string := mdStr(defaultArena, child_string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_ChildFromString(_node, _child_string, _flags)
	return goNodeP(_ret)
}

func TagFromString(node *Node, tag_string string, flags MatchFlags) *Node {
	_node := mdNodeP(defaultArena, node)
	_tag_string := mdStr(defaultArena, tag_string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_TagFromString(_node, _tag_string, _flags)
	return goNodeP(_ret)
}

func ChildFromIndex(node *Node, n int) *Node {
	_node := mdNodeP(defaultArena, node)
	_n := C.int(n)
	_ret := C.MD_ChildFromIndex(_node, _n)
	return goNodeP(_ret)
}

func TagFromIndex(node *Node, n int) *Node {
	_node := mdNodeP(defaultArena, node)
	_n := C.int(n)
	_ret := C.MD_TagFromIndex(_node, _n)
	return goNodeP(_ret)
}

func TagArgFromIndex(node *Node, tag_string string, flags MatchFlags, n int) *Node {
	_node := mdNodeP(defaultArena, node)
	_tag_string := mdStr(defaultArena, tag_string)
	_flags := C.MD_MatchFlags(flags)
	_n := C.int(n)
	_ret := C.MD_TagArgFromIndex(_node, _tag_string, _flags, _n)
	return goNodeP(_ret)
}

func TagArgFromString(node *Node, tag_string string, tag_str_flags MatchFlags, arg_string string, arg_str_flags MatchFlags) *Node {
	_node := mdNodeP(defaultArena, node)
	_tag_string := mdStr(defaultArena, tag_string)
	_tag_str_flags := C.MD_MatchFlags(tag_str_flags)
	_arg_string := mdStr(defaultArena, arg_string)
	_arg_str_flags := C.MD_MatchFlags(arg_str_flags)
	_ret := C.MD_TagArgFromString(_node, _tag_string, _tag_str_flags, _arg_string, _arg_str_flags)
	return goNodeP(_ret)
}

func NodeHasChild(node *Node, _string string, flags MatchFlags) bool {
	_node := mdNodeP(defaultArena, node)
	__string := mdStr(defaultArena, _string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeHasChild(_node, __string, _flags)
	return _ret == 0
}

func NodeHasTag(node *Node, tag_string string, flags MatchFlags) bool {
	_node := mdNodeP(defaultArena, node)
	_tag_string := mdStr(defaultArena, tag_string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeHasTag(_node, _tag_string, _flags)
	return _ret == 0
}

func ChildCountFromNode(node *Node) int {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_ChildCountFromNode(_node)
	return int(_ret)
}

func TagCountFromNode(node *Node) int {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_TagCountFromNode(_node)
	return int(_ret)
}

func ResolveNodeFromReference(node *Node) *Node {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_ResolveNodeFromReference(_node)
	return goNodeP(_ret)
}

func NodeNextWithLimit(node *Node, opl *Node) *Node {
	_node := mdNodeP(defaultArena, node)
	_opl := mdNodeP(defaultArena, opl)
	_ret := C.MD_NodeNextWithLimit(_node, _opl)
	return goNodeP(_ret)
}

func PrevCommentFromNode(node *Node) string {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_PrevCommentFromNode(_node)
	return goStr(_ret)
}

func NextCommentFromNode(node *Node) string {
	_node := mdNodeP(defaultArena, node)
	_ret := C.MD_NextCommentFromNode(_node)
	return goStr(_ret)
}

func StringFromMessageKind(kind C.MD_MessageKind) string {
	_ret := C.MD_StringFromMessageKind(kind)
	return goStr(_ret)
}

func FormatMessage(loc C.MD_CodeLoc, kind C.MD_MessageKind, _string string) string {
	__string := mdStr(defaultArena, _string)
	_ret := C.MD_FormatMessage(defaultArena, loc, kind, __string)
	return goStr(_ret)
}

func NodeMatch(a *Node, b *Node, flags MatchFlags) bool {
	_a := mdNodeP(defaultArena, a)
	_b := mdNodeP(defaultArena, b)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeMatch(_a, _b, _flags)
	return _ret == 0
}

func NodeDeepMatch(a *Node, b *Node, flags MatchFlags) bool {
	_a := mdNodeP(defaultArena, a)
	_b := mdNodeP(defaultArena, b)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeDeepMatch(_a, _b, _flags)
	return _ret == 0
}

func ExprBakeOprTableFromList(list *C.MD_ExprOprList) C.MD_ExprOprTable {
	_ret := C.MD_ExprBakeOprTableFromList(defaultArena, list)
	return _ret
}

func ExprOprFromKindString(table *C.MD_ExprOprTable, kind C.MD_ExprOprKind, s string) *C.MD_ExprOpr {
	_s := mdStr(defaultArena, s)
	_ret := C.MD_ExprOprFromKindString(table, kind, _s)
	return _ret
}

func ExprParse(op_table *C.MD_ExprOprTable, first *Node, one_past_last *Node) C.MD_ExprParseResult {
	_first := mdNodeP(defaultArena, first)
	_one_past_last := mdNodeP(defaultArena, one_past_last)
	_ret := C.MD_ExprParse(defaultArena, op_table, _first, _one_past_last)
	return _ret
}

