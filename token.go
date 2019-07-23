package main

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // main

	// Misc characters
	ARROW // ->

	// Keywords
	NOT
	AND
	OR
	LSHIFT
	RSHIFT
	ECHO
)
