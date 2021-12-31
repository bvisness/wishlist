// Generated from the official Metadesk reference. DO NOT EDIT!
package metadesk

// #include "md.h"
import "C"

var defaultInstance = NewMetadesk()

func StringFromNodeKind(kind NodeKind) string {
	return defaultInstance.StringFromNodeKind(kind)
}

func (md *Metadesk) StringFromNodeKind(kind NodeKind) string {
	if md == nil {
		md = &defaultInstance
	}

	_kind := C.MD_NodeKind(kind)
	_ret := C.MD_StringFromNodeKind(_kind)
	return goStr(_ret)
}

func StringListFromNodeFlags(flags NodeFlags) []string {
	return defaultInstance.StringListFromNodeFlags(flags)
}

func (md *Metadesk) StringListFromNodeFlags(flags NodeFlags) []string {
	if md == nil {
		md = &defaultInstance
	}

	_flags := C.MD_NodeFlags(flags)
	_ret := C.MD_StringListFromNodeFlags(md.a, _flags)
	return goStrList(_ret)
}

func ParseResultZero() ParseResult {
	return defaultInstance.ParseResultZero()
}

func (md *Metadesk) ParseResultZero() ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_ret := C.MD_ParseResultZero()
	return md.goParseResult(_ret)
}

func ParseNodeSet(_string string, offset int, parent *Node, rule ParseSetRule) ParseResult {
	return defaultInstance.ParseNodeSet(_string, offset, parent, rule)
}

func (md *Metadesk) ParseNodeSet(_string string, offset int, parent *Node, rule ParseSetRule) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	__string := mdStr(md.a, _string)
	_offset := C.MD_u64(offset)
	_parent := md.mdNodeP(parent)
	_rule := C.MD_ParseSetRule(rule)
	_ret := C.MD_ParseNodeSet(md.a, __string, _offset, _parent, _rule)
	return md.goParseResult(_ret)
}

func ParseOneNode(_string string, offset int) ParseResult {
	return defaultInstance.ParseOneNode(_string, offset)
}

func (md *Metadesk) ParseOneNode(_string string, offset int) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	__string := mdStr(md.a, _string)
	_offset := C.MD_u64(offset)
	_ret := C.MD_ParseOneNode(md.a, __string, _offset)
	return md.goParseResult(_ret)
}

func ParseWholeString(filename string, contents string) ParseResult {
	return defaultInstance.ParseWholeString(filename, contents)
}

func (md *Metadesk) ParseWholeString(filename string, contents string) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_filename := mdStr(md.a, filename)
	_contents := mdStr(md.a, contents)
	_ret := C.MD_ParseWholeString(md.a, _filename, _contents)
	return md.goParseResult(_ret)
}

func ParseWholeFile(filename string) ParseResult {
	return defaultInstance.ParseWholeFile(filename)
}

func (md *Metadesk) ParseWholeFile(filename string) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_filename := mdStr(md.a, filename)
	_ret := C.MD_ParseWholeFile(md.a, _filename)
	return md.goParseResult(_ret)
}

func MakeErrorMarkerNode(parse_contents string, offset int) *Node {
	return defaultInstance.MakeErrorMarkerNode(parse_contents, offset)
}

func (md *Metadesk) MakeErrorMarkerNode(parse_contents string, offset int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_parse_contents := mdStr(md.a, parse_contents)
	_offset := C.MD_u64(offset)
	_ret := C.MD_MakeErrorMarkerNode(md.a, _parse_contents, _offset)
	return md.goNodeP(_ret)
}

func MakeNodeError(node *Node, kind MessageKind, str string) *Message {
	return defaultInstance.MakeNodeError(node, kind, str)
}

func (md *Metadesk) MakeNodeError(node *Node, kind MessageKind, str string) *Message {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_kind := C.MD_MessageKind(kind)
	_str := mdStr(md.a, str)
	_ret := C.MD_MakeNodeError(md.a, _node, _kind, _str)
	return md.goMessageP(_ret)
}

func MakeTokenError(parse_contents string, token C.MD_Token, kind MessageKind, str string) *Message {
	return defaultInstance.MakeTokenError(parse_contents, token, kind, str)
}

func (md *Metadesk) MakeTokenError(parse_contents string, token C.MD_Token, kind MessageKind, str string) *Message {
	if md == nil {
		md = &defaultInstance
	}

	_parse_contents := mdStr(md.a, parse_contents)
	_kind := C.MD_MessageKind(kind)
	_str := mdStr(md.a, str)
	_ret := C.MD_MakeTokenError(md.a, _parse_contents, token, _kind, _str)
	return md.goMessageP(_ret)
}

func MessageListPush(list *MessageList, error *Message) {
	defaultInstance.MessageListPush(list, error)
}

func (md *Metadesk) MessageListPush(list *MessageList, error *Message) {
	if md == nil {
		md = &defaultInstance
	}

	_list := md.mdMessageListP(list)
	_error := md.mdMessageP(error)
	C.MD_MessageListPush(_list, _error)
}

func MessageListConcat(list *MessageList, to_push *MessageList) {
	defaultInstance.MessageListConcat(list, to_push)
}

func (md *Metadesk) MessageListConcat(list *MessageList, to_push *MessageList) {
	if md == nil {
		md = &defaultInstance
	}

	_list := md.mdMessageListP(list)
	_to_push := md.mdMessageListP(to_push)
	C.MD_MessageListConcat(_list, _to_push)
}

func NodeIsNil(node *Node) bool {
	return defaultInstance.NodeIsNil(node)
}

func (md *Metadesk) NodeIsNil(node *Node) bool {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_NodeIsNil(_node)
	return _ret != 0
}

func NilNode() *Node {
	return defaultInstance.NilNode()
}

func (md *Metadesk) NilNode() *Node {
	if md == nil {
		md = &defaultInstance
	}

	_ret := C.MD_NilNode()
	return md.goNodeP(_ret)
}

func PushChild(parent *Node, new_child *Node) {
	defaultInstance.PushChild(parent, new_child)
}

func (md *Metadesk) PushChild(parent *Node, new_child *Node) {
	if md == nil {
		md = &defaultInstance
	}

	_parent := md.mdNodeP(parent)
	_new_child := md.mdNodeP(new_child)
	C.MD_PushChild(_parent, _new_child)
}

func PushTag(node *Node, tag *Node) {
	defaultInstance.PushTag(node, tag)
}

func (md *Metadesk) PushTag(node *Node, tag *Node) {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_tag := md.mdNodeP(tag)
	C.MD_PushTag(_node, _tag)
}

func PushNewReference(list *Node, target *Node) *Node {
	return defaultInstance.PushNewReference(list, target)
}

func (md *Metadesk) PushNewReference(list *Node, target *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_list := md.mdNodeP(list)
	_target := md.mdNodeP(target)
	_ret := C.MD_PushNewReference(md.a, _list, _target)
	return md.goNodeP(_ret)
}

func FirstNodeWithString(first *Node, _string string, flags MatchFlags) *Node {
	return defaultInstance.FirstNodeWithString(first, _string, flags)
}

func (md *Metadesk) FirstNodeWithString(first *Node, _string string, flags MatchFlags) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_first := md.mdNodeP(first)
	__string := mdStr(md.a, _string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_FirstNodeWithString(_first, __string, _flags)
	return md.goNodeP(_ret)
}

func NodeAtIndex(first *Node, n int) *Node {
	return defaultInstance.NodeAtIndex(first, n)
}

func (md *Metadesk) NodeAtIndex(first *Node, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_first := md.mdNodeP(first)
	_n := C.int(n)
	_ret := C.MD_NodeAtIndex(_first, _n)
	return md.goNodeP(_ret)
}

func FirstNodeWithFlags(first *Node, flags NodeFlags) *Node {
	return defaultInstance.FirstNodeWithFlags(first, flags)
}

func (md *Metadesk) FirstNodeWithFlags(first *Node, flags NodeFlags) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_first := md.mdNodeP(first)
	_flags := C.MD_NodeFlags(flags)
	_ret := C.MD_FirstNodeWithFlags(_first, _flags)
	return md.goNodeP(_ret)
}

func IndexFromNode(node *Node) int {
	return defaultInstance.IndexFromNode(node)
}

func (md *Metadesk) IndexFromNode(node *Node) int {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_IndexFromNode(_node)
	return int(_ret)
}

func RootFromNode(node *Node) *Node {
	return defaultInstance.RootFromNode(node)
}

func (md *Metadesk) RootFromNode(node *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_RootFromNode(_node)
	return md.goNodeP(_ret)
}

func ChildFromString(node *Node, child_string string, flags MatchFlags) *Node {
	return defaultInstance.ChildFromString(node, child_string, flags)
}

func (md *Metadesk) ChildFromString(node *Node, child_string string, flags MatchFlags) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_child_string := mdStr(md.a, child_string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_ChildFromString(_node, _child_string, _flags)
	return md.goNodeP(_ret)
}

func TagFromString(node *Node, tag_string string, flags MatchFlags) *Node {
	return defaultInstance.TagFromString(node, tag_string, flags)
}

func (md *Metadesk) TagFromString(node *Node, tag_string string, flags MatchFlags) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_tag_string := mdStr(md.a, tag_string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_TagFromString(_node, _tag_string, _flags)
	return md.goNodeP(_ret)
}

func ChildFromIndex(node *Node, n int) *Node {
	return defaultInstance.ChildFromIndex(node, n)
}

func (md *Metadesk) ChildFromIndex(node *Node, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_n := C.int(n)
	_ret := C.MD_ChildFromIndex(_node, _n)
	return md.goNodeP(_ret)
}

func TagFromIndex(node *Node, n int) *Node {
	return defaultInstance.TagFromIndex(node, n)
}

func (md *Metadesk) TagFromIndex(node *Node, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_n := C.int(n)
	_ret := C.MD_TagFromIndex(_node, _n)
	return md.goNodeP(_ret)
}

func TagArgFromIndex(node *Node, tag_string string, flags MatchFlags, n int) *Node {
	return defaultInstance.TagArgFromIndex(node, tag_string, flags, n)
}

func (md *Metadesk) TagArgFromIndex(node *Node, tag_string string, flags MatchFlags, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_tag_string := mdStr(md.a, tag_string)
	_flags := C.MD_MatchFlags(flags)
	_n := C.int(n)
	_ret := C.MD_TagArgFromIndex(_node, _tag_string, _flags, _n)
	return md.goNodeP(_ret)
}

func TagArgFromString(node *Node, tag_string string, tag_str_flags MatchFlags, arg_string string, arg_str_flags MatchFlags) *Node {
	return defaultInstance.TagArgFromString(node, tag_string, tag_str_flags, arg_string, arg_str_flags)
}

func (md *Metadesk) TagArgFromString(node *Node, tag_string string, tag_str_flags MatchFlags, arg_string string, arg_str_flags MatchFlags) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_tag_string := mdStr(md.a, tag_string)
	_tag_str_flags := C.MD_MatchFlags(tag_str_flags)
	_arg_string := mdStr(md.a, arg_string)
	_arg_str_flags := C.MD_MatchFlags(arg_str_flags)
	_ret := C.MD_TagArgFromString(_node, _tag_string, _tag_str_flags, _arg_string, _arg_str_flags)
	return md.goNodeP(_ret)
}

func NodeHasChild(node *Node, _string string, flags MatchFlags) bool {
	return defaultInstance.NodeHasChild(node, _string, flags)
}

func (md *Metadesk) NodeHasChild(node *Node, _string string, flags MatchFlags) bool {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	__string := mdStr(md.a, _string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeHasChild(_node, __string, _flags)
	return _ret != 0
}

func NodeHasTag(node *Node, tag_string string, flags MatchFlags) bool {
	return defaultInstance.NodeHasTag(node, tag_string, flags)
}

func (md *Metadesk) NodeHasTag(node *Node, tag_string string, flags MatchFlags) bool {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_tag_string := mdStr(md.a, tag_string)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeHasTag(_node, _tag_string, _flags)
	return _ret != 0
}

func ChildCountFromNode(node *Node) int {
	return defaultInstance.ChildCountFromNode(node)
}

func (md *Metadesk) ChildCountFromNode(node *Node) int {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_ChildCountFromNode(_node)
	return int(_ret)
}

func TagCountFromNode(node *Node) int {
	return defaultInstance.TagCountFromNode(node)
}

func (md *Metadesk) TagCountFromNode(node *Node) int {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_TagCountFromNode(_node)
	return int(_ret)
}

func ResolveNodeFromReference(node *Node) *Node {
	return defaultInstance.ResolveNodeFromReference(node)
}

func (md *Metadesk) ResolveNodeFromReference(node *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_ResolveNodeFromReference(_node)
	return md.goNodeP(_ret)
}

func NodeNextWithLimit(node *Node, opl *Node) *Node {
	return defaultInstance.NodeNextWithLimit(node, opl)
}

func (md *Metadesk) NodeNextWithLimit(node *Node, opl *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_opl := md.mdNodeP(opl)
	_ret := C.MD_NodeNextWithLimit(_node, _opl)
	return md.goNodeP(_ret)
}

func PrevCommentFromNode(node *Node) string {
	return defaultInstance.PrevCommentFromNode(node)
}

func (md *Metadesk) PrevCommentFromNode(node *Node) string {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_PrevCommentFromNode(_node)
	return goStr(_ret)
}

func NextCommentFromNode(node *Node) string {
	return defaultInstance.NextCommentFromNode(node)
}

func (md *Metadesk) NextCommentFromNode(node *Node) string {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_NextCommentFromNode(_node)
	return goStr(_ret)
}

func StringFromMessageKind(kind MessageKind) string {
	return defaultInstance.StringFromMessageKind(kind)
}

func (md *Metadesk) StringFromMessageKind(kind MessageKind) string {
	if md == nil {
		md = &defaultInstance
	}

	_kind := C.MD_MessageKind(kind)
	_ret := C.MD_StringFromMessageKind(_kind)
	return goStr(_ret)
}

func FormatMessage(loc C.MD_CodeLoc, kind MessageKind, _string string) string {
	return defaultInstance.FormatMessage(loc, kind, _string)
}

func (md *Metadesk) FormatMessage(loc C.MD_CodeLoc, kind MessageKind, _string string) string {
	if md == nil {
		md = &defaultInstance
	}

	_kind := C.MD_MessageKind(kind)
	__string := mdStr(md.a, _string)
	_ret := C.MD_FormatMessage(md.a, loc, _kind, __string)
	return goStr(_ret)
}

func NodeMatch(a *Node, b *Node, flags MatchFlags) bool {
	return defaultInstance.NodeMatch(a, b, flags)
}

func (md *Metadesk) NodeMatch(a *Node, b *Node, flags MatchFlags) bool {
	if md == nil {
		md = &defaultInstance
	}

	_a := md.mdNodeP(a)
	_b := md.mdNodeP(b)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeMatch(_a, _b, _flags)
	return _ret != 0
}

func NodeDeepMatch(a *Node, b *Node, flags MatchFlags) bool {
	return defaultInstance.NodeDeepMatch(a, b, flags)
}

func (md *Metadesk) NodeDeepMatch(a *Node, b *Node, flags MatchFlags) bool {
	if md == nil {
		md = &defaultInstance
	}

	_a := md.mdNodeP(a)
	_b := md.mdNodeP(b)
	_flags := C.MD_MatchFlags(flags)
	_ret := C.MD_NodeDeepMatch(_a, _b, _flags)
	return _ret != 0
}

func ExprBakeOprTableFromList(list *C.MD_ExprOprList) C.MD_ExprOprTable {
	return defaultInstance.ExprBakeOprTableFromList(list)
}

func (md *Metadesk) ExprBakeOprTableFromList(list *C.MD_ExprOprList) C.MD_ExprOprTable {
	if md == nil {
		md = &defaultInstance
	}

	_ret := C.MD_ExprBakeOprTableFromList(md.a, list)
	return _ret
}

func ExprOprFromKindString(table *C.MD_ExprOprTable, kind C.MD_ExprOprKind, s string) *C.MD_ExprOpr {
	return defaultInstance.ExprOprFromKindString(table, kind, s)
}

func (md *Metadesk) ExprOprFromKindString(table *C.MD_ExprOprTable, kind C.MD_ExprOprKind, s string) *C.MD_ExprOpr {
	if md == nil {
		md = &defaultInstance
	}

	_s := mdStr(md.a, s)
	_ret := C.MD_ExprOprFromKindString(table, kind, _s)
	return _ret
}

func ExprParse(op_table *C.MD_ExprOprTable, first *Node, one_past_last *Node) C.MD_ExprParseResult {
	return defaultInstance.ExprParse(op_table, first, one_past_last)
}

func (md *Metadesk) ExprParse(op_table *C.MD_ExprOprTable, first *Node, one_past_last *Node) C.MD_ExprParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_first := md.mdNodeP(first)
	_one_past_last := md.mdNodeP(one_past_last)
	_ret := C.MD_ExprParse(md.a, op_table, _first, _one_past_last)
	return _ret
}

