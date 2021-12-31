package metadesk

// #include <string.h>
// #include "md.h"
import "C"

import "unsafe"

// Converts a Go string to a Metadesk MD_String8. Allocates on the default arena.
func mdStr(a *C.MD_Arena, s string) C.MD_String8 {
	sp := C.MD_ArenaPush(a, C.MD_u64(len(s)))
	C.memcpy(sp, unsafe.Pointer(C._GoStringPtr(s)), C.size_t(len(s)))

	return C.MD_String8{
		str:  (*C.MD_u8)(sp),
		size: C.MD_u64(len(s)),
	}
}

func goStr(s C.MD_String8) string {
	return string(C.GoBytes(unsafe.Pointer(s.str), (C.int)(s.size)))
}

func mdBool(b bool) C.MD_b32 {
	if b {
		return C.MD_b32(0)
	} else {
		return C.MD_b32(1)
	}
}

// this guy shufflenates stuff from Go to C and the other way.
// if you just follow a bunch of pointers then things go very
// sad because there are cycles everywhere.
type shufflenator struct {
	go2c map[*Node]*C.MD_Node
	c2go map[*C.MD_Node]*Node
}

func (s *shufflenator) mdNode(a *C.MD_Arena, n Node) C.MD_Node {
	return C.MD_Node{
		next:        s.mdNodeP(a, n.Next),
		prev:        s.mdNodeP(a, n.Prev),
		parent:      s.mdNodeP(a, n.Parent),
		first_child: s.mdNodeP(a, n.FirstChild),
		last_child:  s.mdNodeP(a, n.LastChild),

		first_tag: s.mdNodeP(a, n.FirstTag),
		last_tag:  s.mdNodeP(a, n.LastTag),

		kind:       C.MD_NodeKind(n.Kind),
		flags:      C.MD_NodeFlags(n.Flags),
		string:     mdStr(a, n.String),
		raw_string: mdStr(a, n.RawString),

		prev_comment: mdStr(a, n.PrevComment),
		next_comment: mdStr(a, n.NextComment),

		offset: C.MD_u64(n.Offset),

		ref_target: s.mdNodeP(a, n.RefTarget),
	}
}

func mdNodeP(a *C.MD_Arena, n *Node) *C.MD_Node {
	var s shufflenator
	return s.mdNodeP(a, n)
}

func (s *shufflenator) mdNodeP(a *C.MD_Arena, n *Node) *C.MD_Node {
	if n == nil {
		return nil
	}
	if existing, ok := s.go2c[n]; ok {
		return existing
	}

	np := (*C.MD_Node)(C.MD_ArenaPush(a, C.sizeof_MD_Node))
	*np = s.mdNode(a, *n)
	if s.go2c == nil {
		s.go2c = make(map[*Node]*C.MD_Node)
	}
	s.go2c[n] = np
	return np
}

func (s *shufflenator) goNode(n C.MD_Node) Node {
	return Node{
		Next:       s.goNodeP(n.next),
		Prev:       s.goNodeP(n.prev),
		Parent:     s.goNodeP(n.parent),
		FirstChild: s.goNodeP(n.first_child),
		LastChild:  s.goNodeP(n.last_child),

		FirstTag: s.goNodeP(n.first_tag),
		LastTag:  s.goNodeP(n.last_tag),

		Kind:      NodeKind(n.kind),
		Flags:     NodeFlags(n.flags),
		String:    goStr(n.string),
		RawString: goStr(n.raw_string),

		PrevComment: goStr(n.prev_comment),
		NextComment: goStr(n.next_comment),

		Offset: int(n.offset),

		RefTarget: s.goNodeP(n.ref_target),
	}
}

func goNodeP(n *C.MD_Node) *Node {
	var s shufflenator
	return s.goNodeP(n)
}

func (s *shufflenator) goNodeP(n *C.MD_Node) *Node {
	if n == nil {
		return nil
	}
	if existing, ok := s.c2go[n]; ok {
		return existing
	}

	res := s.goNode(*n)
	if s.c2go == nil {
		s.c2go = make(map[*C.MD_Node]*Node)
	}
	s.c2go[n] = &res
	return &res
}

func mdStrList(a *C.MD_Arena, l []string) C.MD_String8List {
	ml := C.MD_String8List{}
	for _, s := range l {
		C.MD_S8ListPush(a, &ml, mdStr(a, s))
	}
	return ml
}

func goStrList(l C.MD_String8List) []string {
	res := make([]string, 0, l.node_count)
	for node := l.first; node != nil; node = node.next {
		res = append(res, goStr(node.string))
	}
	return res
}

func mdParseResult(a *C.MD_Arena, r ParseResult) C.MD_ParseResult {
	var s shufflenator
	return s.mdParseResult(a, r)
}

func (s *shufflenator) mdParseResult(a *C.MD_Arena, r ParseResult) C.MD_ParseResult {
	return C.MD_ParseResult{
		node:           s.mdNodeP(a, r.Node),
		string_advance: C.MD_u64(r.StringAdvance),
		errors:         s.mdMessageList(a, r.Errors),
	}
}

func mdMessageListP(a *C.MD_Arena, l *MessageList) *C.MD_MessageList {
	var s shufflenator
	return s.mdMessageListP(a, l)
}

func (s *shufflenator) mdMessageListP(a *C.MD_Arena, l *MessageList) *C.MD_MessageList {
	if l == nil {
		return nil
	}

	lp := (*C.MD_MessageList)(C.MD_ArenaPush(a, C.sizeof_MD_MessageList))
	*lp = s.mdMessageList(a, *l)
	return lp
}

func (s *shufflenator) mdMessageList(a *C.MD_Arena, l MessageList) C.MD_MessageList {
	var ml C.MD_MessageList
	for i := range l.Messages {
		C.MD_MessageListPush(&ml, s.mdMessageP(a, &l.Messages[i]))
	}
	return ml
}

func mdMessageP(a *C.MD_Arena, m *Message) *C.MD_Message {
	var s shufflenator
	return s.mdMessageP(a, m)
}

func (s *shufflenator) mdMessageP(a *C.MD_Arena, m *Message) *C.MD_Message {
	if m == nil {
		return nil
	}
	// no existence check; messages don't cyclically refer to each other

	mp := (*C.MD_Message)(C.MD_ArenaPush(a, C.sizeof_MD_Message))
	*mp = s.mdMessage(a, *m)
	return mp
}

func (s *shufflenator) mdMessage(a *C.MD_Arena, m Message) C.MD_Message {
	return C.MD_Message{
		// don't set next; that happens on push
		node:   s.mdNodeP(a, m.Node),
		kind:   C.MD_MessageKind(m.Kind),
		string: mdStr(a, m.String),
	}
}

func goParseResult(r C.MD_ParseResult) ParseResult {
	var s shufflenator
	return s.goParseResult(r)
}

func (s *shufflenator) goParseResult(r C.MD_ParseResult) ParseResult {
	return ParseResult{
		Node:          s.goNodeP(r.node),
		StringAdvance: int(r.string_advance),
		Errors:        s.goMessageList(r.errors),
	}
}

func (s *shufflenator) goMessageListP(l *C.MD_MessageList) *MessageList {
	if l == nil {
		return nil
	}

	res := s.goMessageList(*l)
	return &res
}

func (s *shufflenator) goMessageList(l C.MD_MessageList) MessageList {
	res := MessageList{
		MaxMessageKind: MessageKind(l.max_message_kind),
		Messages:       make([]Message, 0, l.node_count),
	}
	for msg := l.first; msg != nil; msg = msg.next {
		goMsg := s.goMessageP(msg)
		res.Messages = append(res.Messages, *goMsg)
	}
	return res
}

func goMessageP(m *C.MD_Message) *Message {
	var s shufflenator
	return s.goMessageP(m)
}

func (s *shufflenator) goMessageP(m *C.MD_Message) *Message {
	if m == nil {
		return nil
	}
	// no existence check; messages don't cyclically refer to each other

	res := s.goMessage(*m)
	return &res
}

func (s *shufflenator) goMessage(m C.MD_Message) Message {
	return Message{
		Node:   s.goNodeP(m.node),
		Kind:   MessageKind(m.kind),
		String: goStr(m.string),
	}
}

func AllNodes(first *C.MD_Node) []*C.MD_Node {
	var res []*C.MD_Node
	for it := first; C.MD_NodeIsNil(it) == 0; it = it.next {
		res = append(res, it)
	}
	return res
}
