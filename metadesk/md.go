package metadesk

/*
#cgo linux CFLAGS: -DMD_OS_LINUX=1
#cgo darwin CFLAGS: -DMD_OS_MAC=1
#cgo windows CFLAGS: -DMD_OS_WINDOWS=1

#cgo amd64 CFLAGS: -DMD_ARCH_X64=1
#cgo 386 CFLAGS: -DMD_ARCH_X86=1
#cgo arm64 CFLAGS: -DMD_ARCH_ARM64=1
#cgo arm CFLAGS: -DMD_ARCH_ARM32=1

#cgo CFLAGS: -DMD_THREAD_LOCAL=_Thread_local

#include "md.h"

MD_String8 S8Lit(const char *s) {
	return MD_S8Lit(s);
}

#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

var a = C.MD_ArenaAlloc()

// Converts a Go string to a Metadesk MD_String8. Allocates on the default arena.
func Str(s string) C.MD_String8 {
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

func blep() {
	res := C.MD_ParseWholeString(a, Str("my test file"), Str("1 2 3 4"))

	var out C.MD_String8List
	C.MD_DebugDumpFromNode(a, &out, res.node, 1, Str("  "), C.MD_GenerateFlags_Tree)

	dump := C.MD_S8ListJoin(a, out, &C.MD_StringJoin{
		mid: Str(""),
	})

	fmt.Println(GoStr(dump))
}
