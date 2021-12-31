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

func (md *Metadesk) mdNode(n Node) C.MD_Node {
	return C.MD_Node{
		next:        md.mdNodeP(n.Next),
		prev:        md.mdNodeP(n.Prev),
		parent:      md.mdNodeP(n.Parent),
		first_child: md.mdNodeP(n.FirstChild),
		last_child:  md.mdNodeP(n.LastChild),

		first_tag: md.mdNodeP(n.FirstTag),
		last_tag:  md.mdNodeP(n.LastTag),

		kind:       C.MD_NodeKind(n.Kind),
		flags:      C.MD_NodeFlags(n.Flags),
		string:     mdStr(md.a, n.String),
		raw_string: mdStr(md.a, n.RawString),

		prev_comment: mdStr(md.a, n.PrevComment),
		next_comment: mdStr(md.a, n.NextComment),

		offset: C.MD_u64(n.Offset),

		ref_target: md.mdNodeP(n.RefTarget),
	}
}

func (md *Metadesk) mdNodeP(n *Node) *C.MD_Node {
	if n == nil {
		return nil
	}
	if existing, ok := md.go2c[n]; ok {
		return (*C.MD_Node)(existing)
	}

	np := (*C.MD_Node)(C.MD_ArenaPush(md.a, C.sizeof_MD_Node))
	md.track(n, unsafe.Pointer(np))
	*np = md.mdNode(*n)
	return np
}

func (md *Metadesk) goNode(n C.MD_Node) Node {
	return Node{
		Next:       md.goNodeP(n.next),
		Prev:       md.goNodeP(n.prev),
		Parent:     md.goNodeP(n.parent),
		FirstChild: md.goNodeP(n.first_child),
		LastChild:  md.goNodeP(n.last_child),

		FirstTag: md.goNodeP(n.first_tag),
		LastTag:  md.goNodeP(n.last_tag),

		Kind:      NodeKind(n.kind),
		Flags:     NodeFlags(n.flags),
		String:    goStr(n.string),
		RawString: goStr(n.raw_string),

		PrevComment: goStr(n.prev_comment),
		NextComment: goStr(n.next_comment),

		Offset: int(n.offset),

		RefTarget: md.goNodeP(n.ref_target),
	}
}

func (md *Metadesk) goNodeP(n *C.MD_Node) *Node {
	if n == nil {
		return nil
	}
	if existing, ok := md.c2go[unsafe.Pointer(n)]; ok {
		return existing.(*Node)
	}

	var res Node
	md.track(&res, unsafe.Pointer(n))
	res = md.goNode(*n)
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

func (md *Metadesk) mdParseResult(r ParseResult) C.MD_ParseResult {
	return C.MD_ParseResult{
		node:           md.mdNodeP(r.Node),
		string_advance: C.MD_u64(r.StringAdvance),
		errors:         md.mdMessageList(r.Errors),
	}
}

func (md *Metadesk) mdMessageListP(l *MessageList) *C.MD_MessageList {
	if l == nil {
		return nil
	}
	if existing, ok := md.go2c[l]; ok {
		return (*C.MD_MessageList)(existing)
	}

	lp := (*C.MD_MessageList)(C.MD_ArenaPush(md.a, C.sizeof_MD_MessageList))
	md.track(l, unsafe.Pointer(lp))
	*lp = md.mdMessageList(*l)
	return lp
}

func (md *Metadesk) mdMessageList(l MessageList) C.MD_MessageList {
	var ml C.MD_MessageList
	for i := range l.Messages {
		C.MD_MessageListPush(&ml, md.mdMessageP(&l.Messages[i]))
	}
	return ml
}

func (md *Metadesk) mdMessageP(m *Message) *C.MD_Message {
	if m == nil {
		return nil
	}
	if existing, ok := md.go2c[m]; ok {
		return (*C.MD_Message)(existing)
	}

	mp := (*C.MD_Message)(C.MD_ArenaPush(md.a, C.sizeof_MD_Message))
	md.track(m, unsafe.Pointer(mp))
	*mp = md.mdMessage(*m)
	return mp
}

func (md *Metadesk) mdMessage(m Message) C.MD_Message {
	return C.MD_Message{
		// don't set next; that happens on push
		node:   md.mdNodeP(m.Node),
		kind:   C.MD_MessageKind(m.Kind),
		string: mdStr(md.a, m.String),
	}
}

func (md *Metadesk) goParseResult(r C.MD_ParseResult) ParseResult {
	return ParseResult{
		Node:          md.goNodeP(r.node),
		StringAdvance: int(r.string_advance),
		Errors:        md.goMessageList(r.errors),
	}
}

func (md *Metadesk) goMessageListP(l *C.MD_MessageList) *MessageList {
	if l == nil {
		return nil
	}
	if existing, ok := md.c2go[unsafe.Pointer(l)]; ok {
		return existing.(*MessageList)
	}

	var res MessageList
	md.track(&res, unsafe.Pointer(l))
	res = md.goMessageList(*l)
	return &res
}

func (md *Metadesk) goMessageList(l C.MD_MessageList) MessageList {
	res := MessageList{
		MaxMessageKind: MessageKind(l.max_message_kind),
		Messages:       make([]Message, 0, l.node_count),
	}
	for msg := l.first; msg != nil; msg = msg.next {
		goMsg := md.goMessageP(msg)
		res.Messages = append(res.Messages, *goMsg)
	}
	return res
}

func (md *Metadesk) goMessageP(m *C.MD_Message) *Message {
	if m == nil {
		return nil
	}
	if existing, ok := md.c2go[unsafe.Pointer(m)]; ok {
		return existing.(*Message)
	}

	var res Message
	md.track(&res, unsafe.Pointer(m))
	res = md.goMessage(*m)
	return &res
}

func (md *Metadesk) goMessage(m C.MD_Message) Message {
	return Message{
		Node:   md.goNodeP(m.node),
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
