package main

import (
	"fmt"
	"io"
)

// SelectStatement represents a SQL SELECT statement.
type Statement struct {
	Left  string
	Right string
	Op    string
	Dest  string
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (*Statement, error) {
	stmt := &Statement{}

	// First token should be an IDENT or NOT.
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case IDENT:
		stmt.Left = lit
	case NOT:
		fallthrough
	case ECHO:
		p.unscan()
	default:
		return nil, fmt.Errorf("found %q, expected IDENT or NOT", lit)
	}

	tok, lit = p.scanIgnoreWhitespace()
	switch tok {
	case ARROW:
		//continue
	case NOT:
		fallthrough
	case ECHO:
		fallthrough
	case AND:
		fallthrough
	case OR:
		fallthrough
	case LSHIFT:
		fallthrough
	case RSHIFT:
		stmt.Op = lit

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT {
			return nil, fmt.Errorf("found %q, expected IDENT", lit)
		}
		stmt.Right = lit

		tok, lit = p.scanIgnoreWhitespace()
		if tok != ARROW {
			return nil, fmt.Errorf("found %q, expected ARROW", lit)
		}
	default:
		return nil, fmt.Errorf("found %q, expected ARROW or OPERATOR", lit)
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected IDENT", lit)
	}
	stmt.Dest = lit

	// Return the successfully parsed statement.
	return stmt, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
