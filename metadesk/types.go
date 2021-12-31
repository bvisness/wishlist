package metadesk

type Node struct {
	Next        *Node
	Prev        *Node
	Parent      *Node
	FirstChild  *Node
	LastChild   *Node
	FirstTag    *Node
	LastTag     *Node
	Kind        NodeKind
	Flags       NodeFlags
	String      string
	RawString   string
	PrevComment string
	NextComment string
	Offset      int
	RefTarget   *Node
}

// These constants distinguish major roles of MD_Node in the Metadesk abstract syntax tree data structure.
type NodeKind int

const (
	// The Nil node is a unique node representing the lack of information, for example iterating off the end of a list, or up to the parent of a root node results in Nil.
	NodeKind_Nil NodeKind = iota
	// A File node represents parsed Metadesk source text.
	NodeKind_File
	// An ErrorMarker node is generated when reporting errors. It is used to record the location of an error that occurred in the lexing phase of a parse.
	NodeKind_ErrorMarker
	// A Main node represents the main structure of the metadesk abstract syntax tree. Some of these nodes have children which will also be Main nodes. These nodes can be given their text by identifiers, numerics, string and character literals, and operator symbols.
	NodeKind_Main
	// A Tag node represents a tag attached to a label node with the @identifer syntax. The children of a tag node represent the arguments placed in the tag.
	NodeKind_Tag
	// A List node serves as the root of an externally chained list of nodes. Its children are nodes with the NodeKind_Reference kind.
	NodeKind_List
	// A Reference node is an indirection to another node. The node field RefTarget contains a pointer to the referenced node. These nodes are typically used for creating externally chained linked lists that gather nodes from a parse tree.
	NodeKind_Reference
)

// These flags are set on Node to indicate particular details about the strings that were parsed to create the node.
type NodeFlags int

const (
	// This node's children open with (
	NodeFlag_HasParenLeft NodeFlags = 1 << 0
	// This node's children close with )
	NodeFlag_HasParenRight NodeFlags = 1 << 1
	// This node's children open with [
	NodeFlag_HasBracketLeft NodeFlags = 1 << 2
	// This node's children close with ]
	NodeFlag_HasBracketRight NodeFlags = 1 << 3
	// This node's children open with {
	NodeFlag_HasBraceLeft NodeFlags = 1 << 4
	// This node's children close with }
	NodeFlag_HasBraceRight NodeFlags = 1 << 5
	// The delimiter between this node and its next sibling is a ;
	NodeFlag_IsBeforeSemicolon NodeFlags = 1 << 6
	// The delimiter between this node and its previous sibling is a ;
	NodeFlag_IsAfterSemicolon NodeFlags = 1 << 7
	// The delimiter between this node and its next sibling is a ,
	NodeFlag_IsBeforeComma NodeFlags = 1 << 8
	// The delimiter between this node and its previous sibling is a ,
	NodeFlag_IsAfterComma NodeFlags = 1 << 9
	// This is a string literal, with ' character(s) marking the boundaries.
	NodeFlag_StringSingleQuote NodeFlags = 1 << 10
	// This is a string literal, with " character(s) marking the boundaries.
	NodeFlag_StringDoubleQuote NodeFlags = 1 << 11
	// This is a string literal, with ` character(s) marking the boundaries.
	NodeFlag_StringTick NodeFlags = 1 << 12
	// This is a string literal that used triplets (three of its boundary characters in a row, on either side) to mark its boundaries, making it multiline.
	NodeFlag_StringTriplet NodeFlags = 1 << 13
	// The label on this node comes from a token with the MD_TokenKind_Numeric kind.
	NodeFlag_Numeric NodeFlags = 1 << 14
	// The label on this node comes from a token with the MD_TokenKind_Identifier kind.
	NodeFlag_Identifier NodeFlags = 1 << 15
	// The label on this node comes from a token with the MD_TokenKind_StringLiteral kind.
	NodeFlag_StringLiteral NodeFlags = 1 << 16
)

// These flags control matching rules in routines that perform matching on strings and Node trees. Not all flags are within this enum. These flags must not be overlapping with those in the StringMatchFlags enum, nor those in the NodeMatchFlags enum. This allows all flags to be associated with their respective behaviors, but also be combined when appropriate.
//
// NOTE FOR BINDINGS:
// There are no other types such as StringMatchFlags. They never overlap anyway, and in Go it's too annoying to manually cast to MatchFlags all the time.
type MatchFlags int

const (
	// For routines returning the location of a substring, alters the behavior to return the last match instead of the first match.
	MatchFlag_FindLast MatchFlags = 1 << 0

	// When comparing strings, consider lower case letters equivalent to upper case equivalents in the ASCII range.
	StringMatchFlag_CaseInsensitive MatchFlags = 1 << 4
	// When comparing strings, do not require the strings to be the same length. If one of the strings is a prefix of another, the two strings will count as a match.
	StringMatchFlag_RightSideSloppy MatchFlags = 1 << 5
	// When comparing strings, consider forward slash and backward slash to be equivalents.
	StringMatchFlag_SlashInsensitive MatchFlags = 1 << 6

	// When comparing nodes with this flag set, differences in the order and names of tags on a node count as differences in the input nodes. Without this flag tags are ignored in tree comparisons.
	NodeMatchFlag_Tags MatchFlags = 1 << 16
	// When comparing nodes with this flag set in addition to MD_NodeMatchFlag_Tags, the differences in the arguments of each tag (the tag's children in the tree) count as differences in the input nodes. Tag arguments are compared with fully recursive compares, whether or not the compare routine would be recursive or not.
	NodeMatchFlag_TagArguments MatchFlags = 1 << 17
)
