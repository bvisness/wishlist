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

func mdNodeP(a *C.MD_Arena, n *Node) *C.MD_Node {
	var s shufflenator
	return s.mdNodeP(a, n)
}

func goNodeP(n *C.MD_Node) *Node {
	var s shufflenator
	return s.goNodeP(n)
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

func AllNodes(first *C.MD_Node) []*C.MD_Node {
	var res []*C.MD_Node
	for it := first; C.MD_NodeIsNil(it) == 0; it = it.next {
		res = append(res, it)
	}
	return res
}
