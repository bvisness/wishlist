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
*/
import "C"

import (
	"fmt"
	"unsafe"
)

var a = C.MD_ArenaAlloc()

func Str(s string) (str C.MD_String8, free func()) {
	cs := C.CString(s)
	return C.S8Lit(cs), func() {
		C.free(unsafe.Pointer(cs))
	}
}

func FromStr(s C.MD_String8) string {
	return string(C.GoBytes(unsafe.Pointer(s.str), (C.int)(s.size)))
}

func blep() {
	filename, _ := Str("mytestfile")
	source, _ := Str("1 2 3 4")
	res := C.MD_ParseWholeString(a, filename, source)

	indent, _ := Str("  ")
	var out C.MD_String8List
	C.MD_DebugDumpFromNode(a, &out, res.node, 1, indent, C.MD_GenerateFlags_Tree)

	mid, _ := Str("")
	dump := C.MD_S8ListJoin(a, out, &C.MD_StringJoin{
		mid: mid,
	})

	fmt.Println(FromStr(dump))
}
