//go:generate rm md.go
//go:generate go run ./gen/gen.go
package metadesk

/*
#include <stdlib.h>
#include <string.h>

#include "md.h"
*/
import "C"

import (
	"fmt"
	"strings"
)

type bindingType string

func (b bindingType) MDShuffle(name string) (string, string, string) {
	// The returned value will be named "_ret", and the result
	// value should be returned
	switch b {
	case "MD_String8":
		return "string", "Str(defaultArena, " + name + ")", "return GoStr(_ret)"
	case "MD_b32":
		return "bool", name + " == 0", "return _ret == 0"
	case "MD_i64":
		return "int", "C.MD_i64(" + name + ")", "return int(_ret)"
	case "int":
		return "int", "C.int(" + name + ")", "return int(_ret)"
	default:
		goType := string(b)
		if strings.HasPrefix(goType, "MD_") {
			goType = "C." + goType
		} else if strings.HasPrefix(goType, "*MD_") {
			goType = "*C." + goType[1:]
		}
		return goType, "", "return _ret"
	}
}

func parseType(first *C.MD_Node) bindingType {
	children := AllNodes(first)
	var res bindingType
	for _, token := range children {
		res += bindingType(GoStr(token.string))
	}
	return res
}

// Generates language bindings and returns the resulting Go source bytes so you can write them to a
// file or whatever.
//
// It doesn't use any of the nice bindings because, well, it generates those.
func GenBindings(reference string) []byte {
	a := C.MD_ArenaAlloc()

	ref := C.MD_ParseWholeString(a, Str(a, "reference"), Str(a, reference))

	funcStr := Str(a, "func")
	sendStr := Str(a, "send")
	returnStr := Str(a, "return")

	var out strings.Builder
	out.WriteString("// Generated from the official Metadesk reference. DO NOT EDIT!\n")
	out.WriteString("package metadesk\n")
	out.WriteString("\n")
	out.WriteString(`// #include "md.h"` + "\n")
	out.WriteString(`import "C"` + "\n")
	out.WriteString("\n")
	out.WriteString("var defaultArena = C.MD_ArenaAlloc()\n")
	out.WriteString("\n")

nextfunc:
	for _, def := range AllNodes(ref.node.first_child) {
		isFunc := C.MD_NodeIsNil(C.MD_TagFromString(def, funcStr, 0)) == 0
		send := GoStr(C.MD_TagFromString(def, sendStr, 0).first_child.string)

		if !isFunc {
			continue
		}

		if send != "Nodes" {
			continue
		}

		rawName := GoStr(def.string)
		name := strings.TrimPrefix(rawName, "MD_")

		if rawName == "MD_PrintMessage" {
			continue
		}

		ret := C.MD_FirstNodeWithString(def.first_child, returnStr, 0)
		returnTypeStr := ""
		returnStr := ""
		if C.MD_NodeIsNil(ret) == 0 {
			returnType := parseType(ret.first_child)
			returnTypeStr, _, returnStr = returnType.MDShuffle("")
		}

		type argStuff struct {
			t            string
			originalName string
			newName      string
			expr         string
		}
		var args []argStuff
		for _, argNode := range AllNodes(def.first_child) {
			originalName := GoStr(argNode.string)
			if originalName == "return" {
				continue
			}
			if originalName == "..." {
				fmt.Println("No bindings for varargs!")
				continue nextfunc
			}

			switch originalName {
			case "string":
				originalName = "_string"
			}

			t := parseType(argNode.first_child)
			inType, expr, _ := t.MDShuffle(originalName)

			var arg argStuff
			if expr == "" {
				// no conversion necessary
				arg = argStuff{
					t:            inType,
					originalName: originalName,
					newName:      originalName,
					expr:         expr,
				}
			} else {
				arg = argStuff{
					t:            inType,
					originalName: originalName,
					newName:      "_" + originalName,
					expr:         expr,
				}
			}
			args = append(args, arg)
		}

		// signature
		out.WriteString("func " + name + "(")
		for i, arg := range args {
			if i > 0 {
				out.WriteString(", ")
			}
			out.WriteString(arg.originalName + " " + arg.t)
		}
		out.WriteString(")")
		if returnTypeStr != "" {
			out.WriteString(" " + returnTypeStr)
		}
		out.WriteString(" {\n")

		// args
		var extraExprs []string
		var callArgs []string
		for _, arg := range args {
			callArgs = append(callArgs, arg.newName)
			if arg.expr != "" {
				extraExprs = append(extraExprs, arg.newName+" := "+arg.expr)
			}
		}
		for _, expr := range extraExprs {
			out.WriteString(`	` + expr + "\n")
		}
		callArgsStr := ""
		for i, arg := range callArgs {
			if i > 0 {
				callArgsStr += ", "
			}
			callArgsStr += arg
		}

		// call and return
		if returnTypeStr == "" {
			out.WriteString(`	C.` + rawName + "(" + callArgsStr + ")" + "\n")
		} else {
			out.WriteString(`	_ret := C.` + rawName + "(" + callArgsStr + ")" + "\n")
			out.WriteString(`	` + returnStr + "\n")
		}

		out.WriteString("}\n")
		out.WriteString("\n")
	}

	return []byte(out.String())
}
