// Generated from the official Metadesk reference. DO NOT EDIT!
package metadesk

// #include "md.h"
import "C"

var defaultInstance = NewMetadesk()

// Returns a string that contains a name matching 'kind'.
func StringFromNodeKind(kind NodeKind) string {
	return defaultInstance.StringFromNodeKind(kind)
}

// Returns a string that contains a name matching 'kind'.
func (md *Metadesk) StringFromNodeKind(kind NodeKind) string {
	if md == nil {
		md = &defaultInstance
	}

	_kind := C.MD_NodeKind(kind)
	_ret := C.MD_StringFromNodeKind(_kind)
	return goStr(_ret)
}

// Builds a string list for all bits set in 'flags', with each string being the name of one of the flags that is set.
func StringListFromNodeFlags(flags NodeFlags) []string {
	return defaultInstance.StringListFromNodeFlags(flags)
}

// Builds a string list for all bits set in 'flags', with each string being the name of one of the flags that is set.
func (md *Metadesk) StringListFromNodeFlags(flags NodeFlags) []string {
	if md == nil {
		md = &defaultInstance
	}

	_flags := C.MD_NodeFlags(flags)
	_ret := C.MD_StringListFromNodeFlags(md.a, _flags)
	return goStrList(_ret)
}

// Constructs a default MD_ParseResult, which indicates that nothing was parsed.
func ParseResultZero() ParseResult {
	return defaultInstance.ParseResultZero()
}

// Constructs a default MD_ParseResult, which indicates that nothing was parsed.
func (md *Metadesk) ParseResultZero() ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_ret := C.MD_ParseResultZero()
	return md.goParseResult(_ret)
}

// Parses a single Metadesk node set, starting at 'offset' bytes into 'string'. Parses the associated set delimiters in accordance with 'rule'.
func ParseNodeSet(_string string, offset int, parent *Node, rule ParseSetRule) ParseResult {
	return defaultInstance.ParseNodeSet(_string, offset, parent, rule)
}

// Parses a single Metadesk node set, starting at 'offset' bytes into 'string'. Parses the associated set delimiters in accordance with 'rule'.
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

// Parses a single Metadesk subtree, starting at 'offset' bytes into 'string'.
func ParseOneNode(_string string, offset int) ParseResult {
	return defaultInstance.ParseOneNode(_string, offset)
}

// Parses a single Metadesk subtree, starting at 'offset' bytes into 'string'.
func (md *Metadesk) ParseOneNode(_string string, offset int) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	__string := mdStr(md.a, _string)
	_offset := C.MD_u64(offset)
	_ret := C.MD_ParseOneNode(md.a, __string, _offset)
	return md.goParseResult(_ret)
}

// Parses an entire string encoding Metadesk. Parents all parsed nodes with a node with 'MD_NodeKind_File' set as its kind.
func ParseWholeString(filename string, contents string) ParseResult {
	return defaultInstance.ParseWholeString(filename, contents)
}

// Parses an entire string encoding Metadesk. Parents all parsed nodes with a node with 'MD_NodeKind_File' set as its kind.
func (md *Metadesk) ParseWholeString(filename string, contents string) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_filename := mdStr(md.a, filename)
	_contents := mdStr(md.a, contents)
	_ret := C.MD_ParseWholeString(md.a, _filename, _contents)
	return md.goParseResult(_ret)
}

// Uses the C standard library to load the file associated with 'filename', and parses all of it to return a single tree for the whole file.
func ParseWholeFile(filename string) ParseResult {
	return defaultInstance.ParseWholeFile(filename)
}

// Uses the C standard library to load the file associated with 'filename', and parses all of it to return a single tree for the whole file.
func (md *Metadesk) ParseWholeFile(filename string) ParseResult {
	if md == nil {
		md = &defaultInstance
	}

	_filename := mdStr(md.a, filename)
	_ret := C.MD_ParseWholeFile(md.a, _filename)
	return md.goParseResult(_ret)
}

// Constructs an MD_Node on 'arena' for the purpose of marking the location at which an error in source text input occurred.
func MakeErrorMarkerNode(parse_contents string, offset int) *Node {
	return defaultInstance.MakeErrorMarkerNode(parse_contents, offset)
}

// Constructs an MD_Node on 'arena' for the purpose of marking the location at which an error in source text input occurred.
func (md *Metadesk) MakeErrorMarkerNode(parse_contents string, offset int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_parse_contents := mdStr(md.a, parse_contents)
	_offset := C.MD_u64(offset)
	_ret := C.MD_MakeErrorMarkerNode(md.a, _parse_contents, _offset)
	return md.goNodeP(_ret)
}

// Allocates and initializes an MD_Message associated with a particular MD_Node.
func MakeNodeError(node *Node, kind MessageKind, str string) *Message {
	return defaultInstance.MakeNodeError(node, kind, str)
}

// Allocates and initializes an MD_Message associated with a particular MD_Node.
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

// Allocates and initializes an MD_Message associated with a particular MD_Token.
func MakeTokenError(parse_contents string, token Token, kind MessageKind, str string) *Message {
	return defaultInstance.MakeTokenError(parse_contents, token, kind, str)
}

// Allocates and initializes an MD_Message associated with a particular MD_Token.
func (md *Metadesk) MakeTokenError(parse_contents string, token Token, kind MessageKind, str string) *Message {
	if md == nil {
		md = &defaultInstance
	}

	_parse_contents := mdStr(md.a, parse_contents)
	_token := md.mdToken(token)
	_kind := C.MD_MessageKind(kind)
	_str := mdStr(md.a, str)
	_ret := C.MD_MakeTokenError(md.a, _parse_contents, _token, _kind, _str)
	return md.goMessageP(_ret)
}

// Pushes a constructed MD_Message into an MD_MessageList.
func MessageListPush(list *MessageList, error *Message) {
	defaultInstance.MessageListPush(list, error)
}

// Pushes a constructed MD_Message into an MD_MessageList.
func (md *Metadesk) MessageListPush(list *MessageList, error *Message) {
	if md == nil {
		md = &defaultInstance
	}

	_list := md.mdMessageListP(list)
	_error := md.mdMessageP(error)
	C.MD_MessageListPush(_list, _error)
}

// Pushes the contents of 'to_push' into 'list'. Zeroes 'to_push'; the memory used in forming 'to_push' will be used in 'list', and nothing will be copied.
func MessageListConcat(list *MessageList, to_push *MessageList) {
	defaultInstance.MessageListConcat(list, to_push)
}

// Pushes the contents of 'to_push' into 'list'. Zeroes 'to_push'; the memory used in forming 'to_push' will be used in 'list', and nothing will be copied.
func (md *Metadesk) MessageListConcat(list *MessageList, to_push *MessageList) {
	if md == nil {
		md = &defaultInstance
	}

	_list := md.mdMessageListP(list)
	_to_push := md.mdMessageListP(to_push)
	C.MD_MessageListConcat(_list, _to_push)
}

// Returns '1' if the 'node' is nil, or '0' otherwise. A nil node pointer is not equivalent to a null pointer. It can still be dereferenced, and is treated as a dummy placeholder node value.
func NodeIsNil(node *Node) bool {
	return defaultInstance.NodeIsNil(node)
}

// Returns '1' if the 'node' is nil, or '0' otherwise. A nil node pointer is not equivalent to a null pointer. It can still be dereferenced, and is treated as a dummy placeholder node value.
func (md *Metadesk) NodeIsNil(node *Node) bool {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_NodeIsNil(_node)
	return _ret != 0
}

// Returns a nil node pointer.
func NilNode() *Node {
	return defaultInstance.NilNode()
}

// Returns a nil node pointer.
func (md *Metadesk) NilNode() *Node {
	if md == nil {
		md = &defaultInstance
	}

	_ret := C.MD_NilNode()
	return md.goNodeP(_ret)
}

// Links 'new_child' up as being a child of 'parent', inserting it into the end of `parent`'s children list.
func PushChild(parent *Node, new_child *Node) {
	defaultInstance.PushChild(parent, new_child)
}

// Links 'new_child' up as being a child of 'parent', inserting it into the end of `parent`'s children list.
func (md *Metadesk) PushChild(parent *Node, new_child *Node) {
	if md == nil {
		md = &defaultInstance
	}

	_parent := md.mdNodeP(parent)
	_new_child := md.mdNodeP(new_child)
	C.MD_PushChild(_parent, _new_child)
}

// Links 'tag' up as being a tag of 'node', inserting it into the end of `node`'s tag list.
func PushTag(node *Node, tag *Node) {
	defaultInstance.PushTag(node, tag)
}

// Links 'tag' up as being a tag of 'node', inserting it into the end of `node`'s tag list.
func (md *Metadesk) PushTag(node *Node, tag *Node) {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_tag := md.mdNodeP(tag)
	C.MD_PushTag(_node, _tag)
}

// Creates a new reference node, pointing at 'target', and links it up as a child of 'list'.
func PushNewReference(list *Node, target *Node) *Node {
	return defaultInstance.PushNewReference(list, target)
}

// Creates a new reference node, pointing at 'target', and links it up as a child of 'list'.
func (md *Metadesk) PushNewReference(list *Node, target *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_list := md.mdNodeP(list)
	_target := md.mdNodeP(target)
	_ret := C.MD_PushNewReference(md.a, _list, _target)
	return md.goNodeP(_ret)
}

// Finds a node in the range defined by 'first' and 'one_past_last', with the string matching 'string' in accordance with 'flags', or returns a nil node pointer if it is not found.
func FirstNodeWithString(first *Node, _string string, flags MatchFlags) *Node {
	return defaultInstance.FirstNodeWithString(first, _string, flags)
}

// Finds a node in the range defined by 'first' and 'one_past_last', with the string matching 'string' in accordance with 'flags', or returns a nil node pointer if it is not found.
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

// Finds the 'n'th node in the range defined by 'first' and 'one_past_last', or returns a nil node pointer if it is not found. '0' would match 'first', '1' would match 'first->next', and so on.
func NodeAtIndex(first *Node, n int) *Node {
	return defaultInstance.NodeAtIndex(first, n)
}

// Finds the 'n'th node in the range defined by 'first' and 'one_past_last', or returns a nil node pointer if it is not found. '0' would match 'first', '1' would match 'first->next', and so on.
func (md *Metadesk) NodeAtIndex(first *Node, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_first := md.mdNodeP(first)
	_n := C.int(n)
	_ret := C.MD_NodeAtIndex(_first, _n)
	return md.goNodeP(_ret)
}

// Given a starting node 'first', will scan across the node's siblings in-order to find a node that has flags that overlap the passed 'flags'. Useful when, for example, finding the set of node ranges delimited by commas or semicolons inside of a single MD_Node children list.
func FirstNodeWithFlags(first *Node, flags NodeFlags) *Node {
	return defaultInstance.FirstNodeWithFlags(first, flags)
}

// Given a starting node 'first', will scan across the node's siblings in-order to find a node that has flags that overlap the passed 'flags'. Useful when, for example, finding the set of node ranges delimited by commas or semicolons inside of a single MD_Node children list.
func (md *Metadesk) FirstNodeWithFlags(first *Node, flags NodeFlags) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_first := md.mdNodeP(first)
	_flags := C.MD_NodeFlags(flags)
	_ret := C.MD_FirstNodeWithFlags(_first, _flags)
	return md.goNodeP(_ret)
}

// Finds the child index of 'node', with '0' being the first child, '1' being the second, and so on.
func IndexFromNode(node *Node) int {
	return defaultInstance.IndexFromNode(node)
}

// Finds the child index of 'node', with '0' being the first child, '1' being the second, and so on.
func (md *Metadesk) IndexFromNode(node *Node) int {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_IndexFromNode(_node)
	return int(_ret)
}

// Finds the highest-most node in the parent chain of 'node', starting with 'node'. If 'node' has no parent, then 'node' is returned.
func RootFromNode(node *Node) *Node {
	return defaultInstance.RootFromNode(node)
}

// Finds the highest-most node in the parent chain of 'node', starting with 'node'. If 'node' has no parent, then 'node' is returned.
func (md *Metadesk) RootFromNode(node *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_RootFromNode(_node)
	return md.goNodeP(_ret)
}

// Finds a child of 'node' with a string matching 'child_string', where the rules of matching are determined by 'flags'.
func ChildFromString(node *Node, child_string string, flags MatchFlags) *Node {
	return defaultInstance.ChildFromString(node, child_string, flags)
}

// Finds a child of 'node' with a string matching 'child_string', where the rules of matching are determined by 'flags'.
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

// Finds a tag on 'node' with a string matching 'tag_string', where the rules of matching are determined by 'flags'.
func TagFromString(node *Node, tag_string string, flags MatchFlags) *Node {
	return defaultInstance.TagFromString(node, tag_string, flags)
}

// Finds a tag on 'node' with a string matching 'tag_string', where the rules of matching are determined by 'flags'.
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

// Finds a child of 'node' with an index matching 'n'. Returns a nil node pointer if no such child is found.
func ChildFromIndex(node *Node, n int) *Node {
	return defaultInstance.ChildFromIndex(node, n)
}

// Finds a child of 'node' with an index matching 'n'. Returns a nil node pointer if no such child is found.
func (md *Metadesk) ChildFromIndex(node *Node, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_n := C.int(n)
	_ret := C.MD_ChildFromIndex(_node, _n)
	return md.goNodeP(_ret)
}

// Finds a tag on 'node' with an index matching 'n'. Returns a nil node pointer if no such tag is found.
func TagFromIndex(node *Node, n int) *Node {
	return defaultInstance.TagFromIndex(node, n)
}

// Finds a tag on 'node' with an index matching 'n'. Returns a nil node pointer if no such tag is found.
func (md *Metadesk) TagFromIndex(node *Node, n int) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_n := C.int(n)
	_ret := C.MD_TagFromIndex(_node, _n)
	return md.goNodeP(_ret)
}

// Finds the 'n'th tag argument of the tag matching 'tag_string' on 'node', with the matching on 'tag_string' being controlled by 'flags'. Returns a nil node pointer if no such node was found.
func TagArgFromIndex(node *Node, tag_string string, flags MatchFlags, n int) *Node {
	return defaultInstance.TagArgFromIndex(node, tag_string, flags, n)
}

// Finds the 'n'th tag argument of the tag matching 'tag_string' on 'node', with the matching on 'tag_string' being controlled by 'flags'. Returns a nil node pointer if no such node was found.
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

// Finds the tag argument with a string matching 'arg_string', of the tag matching 'tag_string', on 'node'. Matching 'tag_string' is controlled by 'tag_str_flags'. Matching 'arg_string' is controlled by 'arg_str_flags'. Returns a nil node pointer if no such node was found.
func TagArgFromString(node *Node, tag_string string, tag_str_flags MatchFlags, arg_string string, arg_str_flags MatchFlags) *Node {
	return defaultInstance.TagArgFromString(node, tag_string, tag_str_flags, arg_string, arg_str_flags)
}

// Finds the tag argument with a string matching 'arg_string', of the tag matching 'tag_string', on 'node'. Matching 'tag_string' is controlled by 'tag_str_flags'. Matching 'arg_string' is controlled by 'arg_str_flags'. Returns a nil node pointer if no such node was found.
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

// Returns '1' if 'node' has a child with a string matching 'string', with the matching rules being controlled by 'flags', or '0' otherwise.
func NodeHasChild(node *Node, _string string, flags MatchFlags) bool {
	return defaultInstance.NodeHasChild(node, _string, flags)
}

// Returns '1' if 'node' has a child with a string matching 'string', with the matching rules being controlled by 'flags', or '0' otherwise.
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

// Returns '1' if 'node' has a tag with a string matching 'tag_string', with the matching rules being controlled by 'flags', or '0' otherwise.
func NodeHasTag(node *Node, tag_string string, flags MatchFlags) bool {
	return defaultInstance.NodeHasTag(node, tag_string, flags)
}

// Returns '1' if 'node' has a tag with a string matching 'tag_string', with the matching rules being controlled by 'flags', or '0' otherwise.
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

// Returns the number of children of 'node'.
func ChildCountFromNode(node *Node) int {
	return defaultInstance.ChildCountFromNode(node)
}

// Returns the number of children of 'node'.
func (md *Metadesk) ChildCountFromNode(node *Node) int {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_ChildCountFromNode(_node)
	return int(_ret)
}

// Returns the number of tags on 'node'.
func TagCountFromNode(node *Node) int {
	return defaultInstance.TagCountFromNode(node)
}

// Returns the number of tags on 'node'.
func (md *Metadesk) TagCountFromNode(node *Node) int {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_TagCountFromNode(_node)
	return int(_ret)
}

// If 'node' is of kind 'MD_NodeKind_Reference', will follow the chain of 'ref_target's until the final referenced node.
func ResolveNodeFromReference(node *Node) *Node {
	return defaultInstance.ResolveNodeFromReference(node)
}

// If 'node' is of kind 'MD_NodeKind_Reference', will follow the chain of 'ref_target's until the final referenced node.
func (md *Metadesk) ResolveNodeFromReference(node *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_ResolveNodeFromReference(_node)
	return md.goNodeP(_ret)
}

// Moves to the next sibling of 'node', unless it is 'opl', in which case it returns a nil node.
func NodeNextWithLimit(node *Node, opl *Node) *Node {
	return defaultInstance.NodeNextWithLimit(node, opl)
}

// Moves to the next sibling of 'node', unless it is 'opl', in which case it returns a nil node.
func (md *Metadesk) NodeNextWithLimit(node *Node, opl *Node) *Node {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_opl := md.mdNodeP(opl)
	_ret := C.MD_NodeNextWithLimit(_node, _opl)
	return md.goNodeP(_ret)
}

// Gets the string of the comment that immediately preceded 'node', if any.
func PrevCommentFromNode(node *Node) string {
	return defaultInstance.PrevCommentFromNode(node)
}

// Gets the string of the comment that immediately preceded 'node', if any.
func (md *Metadesk) PrevCommentFromNode(node *Node) string {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_PrevCommentFromNode(_node)
	return goStr(_ret)
}

// Gets the string of the comment that immediately followed 'node', if any.
func NextCommentFromNode(node *Node) string {
	return defaultInstance.NextCommentFromNode(node)
}

// Gets the string of the comment that immediately followed 'node', if any.
func (md *Metadesk) NextCommentFromNode(node *Node) string {
	if md == nil {
		md = &defaultInstance
	}

	_node := md.mdNodeP(node)
	_ret := C.MD_NextCommentFromNode(_node)
	return goStr(_ret)
}

// Returns a string encoding the name of the passed 'kind' value.
func StringFromMessageKind(kind MessageKind) string {
	return defaultInstance.StringFromMessageKind(kind)
}

// Returns a string encoding the name of the passed 'kind' value.
func (md *Metadesk) StringFromMessageKind(kind MessageKind) string {
	if md == nil {
		md = &defaultInstance
	}

	_kind := C.MD_MessageKind(kind)
	_ret := C.MD_StringFromMessageKind(_kind)
	return goStr(_ret)
}

// Provides a standard way to format a message string that is associated with an MD_CodeLoc and an MD_MessageKind.
func FormatMessage(loc CodeLoc, kind MessageKind, _string string) string {
	return defaultInstance.FormatMessage(loc, kind, _string)
}

// Provides a standard way to format a message string that is associated with an MD_CodeLoc and an MD_MessageKind.
func (md *Metadesk) FormatMessage(loc CodeLoc, kind MessageKind, _string string) string {
	if md == nil {
		md = &defaultInstance
	}

	_loc := md.mdCodeLoc(loc)
	_kind := C.MD_MessageKind(kind)
	__string := mdStr(md.a, _string)
	_ret := C.MD_FormatMessage(md.a, _loc, _kind, __string)
	return goStr(_ret)
}

// Compares the passed MD_Node nodes 'a' and 'b' non-recursively, and determines whether or not they match. 'flags' determines the rules used in the matching algorithm, including tag-sensitivity and case-sensitivity.
func NodeMatch(a *Node, b *Node, flags MatchFlags) bool {
	return defaultInstance.NodeMatch(a, b, flags)
}

// Compares the passed MD_Node nodes 'a' and 'b' non-recursively, and determines whether or not they match. 'flags' determines the rules used in the matching algorithm, including tag-sensitivity and case-sensitivity.
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

// Compares the passed MD_Node trees 'a' and 'b' recursively, and determines whether or not they and their children match. 'flags' determines the rules used in the matching algorithm, including tag-sensitivity and case-sensitivity.
func NodeDeepMatch(a *Node, b *Node, flags MatchFlags) bool {
	return defaultInstance.NodeDeepMatch(a, b, flags)
}

// Compares the passed MD_Node trees 'a' and 'b' recursively, and determines whether or not they and their children match. 'flags' determines the rules used in the matching algorithm, including tag-sensitivity and case-sensitivity.
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

