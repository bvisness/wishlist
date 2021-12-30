package metadesk

// #include <string.h>
// #include "md.h"
import "C"

import "unsafe"

// Converts a Go string to a Metadesk MD_String8. Allocates on the default arena.
func Str(a *C.MD_Arena, s string) C.MD_String8 {
	sp := C.MD_ArenaPush(a, C.MD_u64(len(s)))
	C.memcpy(sp, unsafe.Pointer(C._GoStringPtr(s)), C.size_t(len(s)))

	return C.MD_String8{
		str:  (*C.MD_u8)(sp),
		size: C.MD_u64(len(s)),
	}
}

func GoStr(s C.MD_String8) string {
	return string(C.GoBytes(unsafe.Pointer(s.str), (C.int)(s.size)))
}

func AllNodes(first *C.MD_Node) []*C.MD_Node {
	var res []*C.MD_Node
	for it := first; C.MD_NodeIsNil(it) == 0; it = it.next {
		res = append(res, it)
	}
	return res
}
