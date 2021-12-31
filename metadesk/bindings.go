//go:generate rm -f md.go
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
	// value should be returned.

	switch b {
	case "int":
		return "int", "C.int(" + name + ")", "return int(_ret)"
	case "MD_b32":
		return "bool", "mdBool(" + name + ")", "return _ret == 0"
	case "MD_i64":
		return "int", "C.MD_i64(" + name + ")", "return int(_ret)"
	case "MD_u64":
		return "int", "C.MD_u64(" + name + ")", "return int(_ret)"
	case "MD_MatchFlags":
		return "MatchFlags", "C.MD_MatchFlags(" + name + ")", "return MatchFlags(_ret)"
	case "MD_MessageKind":
		return "MessageKind", "C.MD_MessageKind(" + name + ")", "return MessageKind(_ret)"
	case "MD_NodeFlags":
		return "NodeFlags", "C.MD_NodeFlags(" + name + ")", "return NodeFlags(_ret)"
	case "MD_NodeKind":
		return "NodeKind", "C.MD_NodeKind(" + name + ")", "return NodeKind(_ret)"
	case "MD_ParseResult":
		return "ParseResult", "mdParseResult(" + name + ")", "return md.goParseResult(_ret)"
	case "MD_ParseSetRule":
		return "ParseSetRule", "C.MD_ParseSetRule(" + name + ")", "return ParseSetRule(_ret)"
	case "MD_String8":
		return "string", "mdStr(md.a, " + name + ")", "return goStr(_ret)"
	case "MD_String8List":
		return "[]string", "md.mdStrList(" + name + ")", "return goStrList(_ret)"
	case "*MD_Message":
		return "*Message", "md.mdMessageP(" + name + ")", "return md.goMessageP(_ret)"
	case "*MD_MessageList":
		return "*MessageList", "md.mdMessageListP(" + name + ")", "return md.goMessageListP(_ret)"
	case "*MD_Node":
		return "*Node", "md.mdNodeP(" + name + ")", "return md.goNodeP(_ret)"
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
		res += bindingType(goStr(token.string))
	}
	return res
}

// Generates language bindings and returns the resulting Go source bytes so you can write them to a
// file or whatever.
//
// It doesn't use any of the nice bindings because, well, it generates those.
func GenBindings(reference string) []byte {
	a := C.MD_ArenaAlloc()

	ref := C.MD_ParseWholeString(a, mdStr(a, "reference"), mdStr(a, reference))

	funcStr := mdStr(a, "func")
	sendStr := mdStr(a, "send")
	returnStr := mdStr(a, "return")

	var out strings.Builder
	out.WriteString("// Generated from the official Metadesk reference. DO NOT EDIT!\n")
	out.WriteString("package metadesk\n")
	out.WriteString("\n")
	out.WriteString(`// #include "md.h"` + "\n")
	out.WriteString(`import "C"` + "\n")
	out.WriteString("\n")
	out.WriteString("var defaultInstance = NewMetadesk()\n")
	out.WriteString("\n")

nextfunc:
	for _, def := range AllNodes(ref.node.first_child) {
		isFunc := C.MD_NodeIsNil(C.MD_TagFromString(def, funcStr, 0)) == 0
		send := goStr(C.MD_TagFromString(def, sendStr, 0).first_child.string)

		if !isFunc {
			continue
		}

		switch send {
		case "ExpressionParser", "Parsing", "Nodes":
		default:
			continue
		}

		rawName := goStr(def.string)
		name := strings.TrimPrefix(rawName, "MD_")

		ret := C.MD_FirstNodeWithString(def.first_child, returnStr, 0)
		returnTypeStr := ""
		returnStr := ""
		if C.MD_NodeIsNil(ret) == 0 {
			returnType := parseType(ret.first_child)
			returnTypeStr, _, returnStr = returnType.MDShuffle("")
		}

		args := AllNodes(def.first_child)
		if len(args) == 0 {
			fmt.Printf("No documentation?? (%s)\n", rawName)
			continue
		}

		type inputArg struct {
			name, t string
		}
		var inputArgs []inputArg
		var conversionExprs []string
		var callArgs []string
		for _, argNode := range AllNodes(def.first_child) {
			originalName := goStr(argNode.string)
			if originalName == "return" {
				continue
			}
			if originalName == "..." {
				fmt.Printf("No bindings for varargs! (%s)\n", rawName)
				continue nextfunc
			}

			switch originalName {
			case "string":
				originalName = "_string"
			}

			t := parseType(argNode.first_child)
			inType, expr, _ := t.MDShuffle(originalName)

			if inType == "*void" {
				fmt.Printf("No bindings for void*! (%s)\n", rawName)
				continue nextfunc
			} else if inType == "*FILE" {
				fmt.Printf("No bindings for FILE*! (%s)\n", rawName)
				continue nextfunc
			} else if inType == "*C.MD_Arena" {
				// We always use the default arena
				callArgs = append(callArgs, "md.a")
			} else {
				inputArgs = append(inputArgs, inputArg{originalName, inType})
				if expr != "" {
					convertedName := "_" + originalName
					conversionExprs = append(conversionExprs, convertedName+" := "+expr)
					callArgs = append(callArgs, convertedName)
				} else {
					callArgs = append(callArgs, originalName)
				}
			}
		}

		// signature prep
		signature := ""
		signature += name + "("
		for i, arg := range inputArgs {
			if i > 0 {
				signature += ", "
			}
			signature += arg.name + " " + arg.t
		}
		signature += ")"
		if returnTypeStr != "" {
			signature += " " + returnTypeStr
		}

		//
		// Bare function (default instance)
		//
		{
			// signature
			out.WriteString("func " + signature + " {\n")

			// call to method on default instance
			out.WriteString(`	`)
			if returnTypeStr != "" {
				out.WriteString("return ")
			}
			out.WriteString("defaultInstance." + name + "(")
			for i, arg := range inputArgs {
				if i > 0 {
					out.WriteString(", ")
				}
				out.WriteString(arg.name)
			}
			out.WriteString(")\n")

			out.WriteString("}\n")
			out.WriteString("\n")
		}

		//
		// Method (any instance)
		//
		{
			// signature
			out.WriteString("func (md *Metadesk) " + signature + " {\n")

			// conversions
			for _, expr := range conversionExprs {
				out.WriteString(`	` + expr + "\n")
			}

			// call and return
			callArgsStr := ""
			for i, arg := range callArgs {
				if i > 0 {
					callArgsStr += ", "
				}
				callArgsStr += arg
			}
			if returnTypeStr == "" {
				out.WriteString(`	C.` + rawName + "(" + callArgsStr + ")" + "\n")
			} else {
				out.WriteString(`	_ret := C.` + rawName + "(" + callArgsStr + ")" + "\n")
				out.WriteString(`	` + returnStr + "\n")
			}

			out.WriteString("}\n")
			out.WriteString("\n")
		}
	}

	return []byte(out.String())
}
