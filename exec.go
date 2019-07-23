package main

import (
	"log"
	"strconv"
	"strings"
)

type Runtime map[string]uint16

func (rt Runtime) Run(program []string) {
	for _, line := range program {
		parser := NewParser(strings.NewReader(line))
		stmt, err := parser.Parse()
		if err != nil {
			log.Fatal(err)
		}
		rt.Exec(stmt)
	}
}

func (rt Runtime) Exec(s *Statement) {
	rt[s.Dest] = rt.evaluate(s.Left, s.Op, s.Right)
}

func (rt Runtime) evaluate(left, op, right string) uint16 {
	var l, r uint16
	if lInt, err := strconv.Atoi(left); err == nil {
		l = uint16(lInt)
	} else {
		l = rt[left]
	}

	if rInt, err := strconv.Atoi(right); err == nil {
		r = uint16(rInt)
	} else {
		r = rt[right]
	}

	switch op {
	case "":
		return l
	case "AND":
		return l & r
	case "OR":
		return l | r
	case "LSHIFT":
		return l << r
	case "RSHIFT":
		return l >> r
	case "NOT":
		return ^r
	case "ECHO":
		log.Printf("ECHO %q: %d", right, r)
		return r
	default:
		log.Fatal("Bad operation: %q", op)
	}
	// never happens
	return ^uint16(0)
}
