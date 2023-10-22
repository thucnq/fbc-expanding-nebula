package main

import (
	"fmt"
	"time"
)

type Params struct {
	// OriginState, LastState   [][]int
	// RowCount, ColCount, I, J int

	I, J int
}

var count = 0

var cFromP = [2][][]int{
	{{0, 0, 0, 0}, {0, 0, 1, 1}, {0, 1, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 1, 0}, {1, 0, 1, 1}, {1, 1, 0, 0}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}},
	{{0, 0, 0, 1}, {0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}},
}

var c2x2FromP = map[string][][][]int{
	"0100": {{{0, 0, 0}, {0, 0, 1}, {0, 0, 1}}, {{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}, {{0, 0, 0}, {0, 0, 1}, {1, 1, 1}}, {{0, 0, 0}, {1, 1, 0}, {0, 0, 1}}, {{0, 0, 0}, {1, 1, 0}, {0, 1, 0}}, {{0, 0, 0}, {1, 1, 0}, {0, 1, 1}}, {{0, 0, 0}, {1, 1, 0}, {1, 0, 1}}, {{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, {{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}, {{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}, {{0, 0, 1}, {0, 0, 0}, {1, 1, 1}}, {{0, 1, 0}, {1, 0, 0}, {0, 1, 1}}, {{0, 1, 0}, {1, 0, 0}, {1, 0, 0}}, {{0, 1, 0}, {1, 0, 0}, {1, 1, 1}}, {{1, 0, 0}, {0, 1, 0}, {0, 1, 0}}, {{1, 0, 0}, {0, 1, 0}, {0, 1, 1}}, {{1, 0, 0}, {0, 1, 0}, {1, 0, 1}}, {{1, 0, 0}, {0, 1, 0}, {1, 1, 0}}, {{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}, {{1, 0, 0}, {1, 0, 1}, {0, 1, 0}}, {{1, 0, 0}, {1, 0, 1}, {0, 1, 1}}, {{1, 0, 0}, {1, 0, 1}, {1, 0, 1}}, {{1, 0, 0}, {1, 0, 1}, {1, 1, 0}}, {{1, 0, 0}, {1, 0, 1}, {1, 1, 1}}, {{1, 0, 0}, {1, 1, 0}, {0, 0, 1}}, {{1, 0, 0}, {1, 1, 0}, {0, 1, 0}}, {{1, 0, 0}, {1, 1, 0}, {0, 1, 1}}, {{1, 0, 0}, {1, 1, 0}, {1, 0, 1}}, {{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, {{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}, {{1, 0, 1}, {1, 0, 0}, {0, 1, 1}}, {{1, 0, 1}, {1, 0, 0}, {1, 0, 0}}, {{1, 0, 1}, {1, 0, 0}, {1, 1, 1}}, {{1, 1, 0}, {0, 0, 0}, {0, 0, 0}}, {{1, 1, 0}, {0, 0, 0}, {1, 1, 1}}, {{1, 1, 0}, {1, 0, 0}, {0, 1, 1}}, {{1, 1, 0}, {1, 0, 0}, {1, 0, 0}}, {{1, 1, 0}, {1, 0, 0}, {1, 1, 1}}},
	"1111": {{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, {{0, 0, 0}, {1, 0, 1}, {0, 0, 0}}, {{0, 0, 1}, {1, 0, 0}, {0, 0, 1}}, {{0, 1, 0}, {0, 0, 0}, {0, 1, 0}}, {{0, 1, 0}, {0, 0, 0}, {1, 0, 1}}, {{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}, {{1, 0, 1}, {0, 0, 0}, {0, 1, 0}}, {{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
	"1101": {{{0, 0, 0}, {0, 1, 0}, {1, 0, 0}}, {{0, 0, 0}, {1, 0, 1}, {1, 0, 0}}, {{0, 0, 1}, {1, 0, 0}, {0, 1, 0}}, {{0, 0, 1}, {1, 0, 0}, {1, 0, 1}}, {{0, 0, 1}, {1, 0, 0}, {1, 1, 0}}, {{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}, {{0, 1, 0}, {0, 0, 0}, {1, 1, 0}}, {{1, 0, 0}, {0, 0, 1}, {0, 0, 0}}, {{1, 0, 1}, {0, 0, 0}, {0, 0, 1}}, {{1, 0, 1}, {0, 0, 0}, {1, 1, 0}}},
	"1010": {{{0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, {{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}, {{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}, {{0, 0, 1}, {0, 1, 0}, {0, 0, 1}}, {{0, 0, 1}, {0, 1, 1}, {0, 0, 0}}, {{0, 0, 1}, {0, 1, 1}, {0, 0, 1}}, {{0, 0, 1}, {1, 0, 1}, {0, 0, 1}}, {{0, 1, 0}, {0, 0, 1}, {0, 1, 0}}, {{0, 1, 0}, {0, 0, 1}, {0, 1, 1}}, {{0, 1, 0}, {0, 0, 1}, {1, 0, 1}}, {{0, 1, 1}, {0, 0, 0}, {0, 1, 1}}, {{0, 1, 1}, {0, 0, 0}, {1, 0, 0}}, {{0, 1, 1}, {0, 0, 1}, {0, 1, 0}}, {{0, 1, 1}, {0, 0, 1}, {0, 1, 1}}, {{0, 1, 1}, {0, 0, 1}, {1, 0, 1}}, {{1, 0, 0}, {0, 0, 0}, {0, 1, 1}}, {{1, 0, 0}, {0, 0, 0}, {1, 0, 0}}, {{1, 0, 1}, {0, 0, 1}, {0, 1, 0}}, {{1, 0, 1}, {0, 0, 1}, {0, 1, 1}}, {{1, 0, 1}, {0, 0, 1}, {1, 0, 1}}},
	"1000": {{{0, 0, 0}, {0, 1, 1}, {0, 1, 0}}, {{0, 0, 0}, {0, 1, 1}, {0, 1, 1}}, {{0, 0, 0}, {0, 1, 1}, {1, 0, 0}}, {{0, 0, 0}, {0, 1, 1}, {1, 0, 1}}, {{0, 0, 0}, {0, 1, 1}, {1, 1, 0}}, {{0, 0, 0}, {0, 1, 1}, {1, 1, 1}}, {{0, 0, 0}, {1, 0, 0}, {0, 1, 1}}, {{0, 0, 0}, {1, 0, 0}, {1, 0, 0}}, {{0, 0, 0}, {1, 0, 0}, {1, 1, 1}}, {{0, 0, 1}, {0, 1, 0}, {0, 1, 0}}, {{0, 0, 1}, {0, 1, 0}, {0, 1, 1}}, {{0, 0, 1}, {0, 1, 0}, {1, 0, 1}}, {{0, 0, 1}, {0, 1, 0}, {1, 1, 0}}, {{0, 0, 1}, {0, 1, 0}, {1, 1, 1}}, {{0, 0, 1}, {0, 1, 1}, {0, 1, 0}}, {{0, 0, 1}, {0, 1, 1}, {0, 1, 1}}, {{0, 0, 1}, {0, 1, 1}, {1, 0, 0}}, {{0, 0, 1}, {0, 1, 1}, {1, 0, 1}}, {{0, 0, 1}, {0, 1, 1}, {1, 1, 0}}, {{0, 0, 1}, {0, 1, 1}, {1, 1, 1}}, {{0, 0, 1}, {1, 0, 1}, {0, 1, 0}}, {{0, 0, 1}, {1, 0, 1}, {0, 1, 1}}, {{0, 0, 1}, {1, 0, 1}, {1, 0, 1}}, {{0, 0, 1}, {1, 0, 1}, {1, 1, 0}}, {{0, 0, 1}, {1, 0, 1}, {1, 1, 1}}, {{0, 1, 0}, {0, 0, 1}, {0, 0, 1}}, {{0, 1, 0}, {0, 0, 1}, {1, 1, 0}}, {{0, 1, 0}, {0, 0, 1}, {1, 1, 1}}, {{0, 1, 1}, {0, 0, 0}, {0, 0, 0}}, {{0, 1, 1}, {0, 0, 0}, {1, 1, 1}}, {{0, 1, 1}, {0, 0, 1}, {0, 0, 1}}, {{0, 1, 1}, {0, 0, 1}, {1, 1, 0}}, {{0, 1, 1}, {0, 0, 1}, {1, 1, 1}}, {{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}, {{1, 0, 0}, {0, 0, 0}, {1, 1, 1}}, {{1, 0, 1}, {0, 0, 1}, {0, 0, 1}}, {{1, 0, 1}, {0, 0, 1}, {1, 1, 0}}, {{1, 0, 1}, {0, 0, 1}, {1, 1, 1}}},
	"0001": {{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}}, {{0, 0, 0}, {0, 0, 0}, {1, 1, 0}}, {{0, 0, 1}, {0, 0, 1}, {0, 0, 0}}, {{0, 0, 1}, {1, 1, 0}, {0, 0, 0}}, {{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}, {{0, 1, 0}, {0, 1, 0}, {1, 0, 0}}, {{0, 1, 0}, {1, 0, 1}, {1, 0, 0}}, {{0, 1, 0}, {1, 1, 0}, {0, 0, 0}}, {{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}, {{0, 1, 1}, {0, 1, 0}, {1, 0, 0}}, {{0, 1, 1}, {1, 0, 0}, {0, 1, 0}}, {{0, 1, 1}, {1, 0, 0}, {1, 0, 1}}, {{0, 1, 1}, {1, 0, 0}, {1, 1, 0}}, {{0, 1, 1}, {1, 0, 1}, {1, 0, 0}}, {{0, 1, 1}, {1, 1, 0}, {0, 0, 0}}, {{0, 1, 1}, {1, 1, 0}, {1, 0, 0}}, {{1, 0, 0}, {1, 0, 0}, {0, 1, 0}}, {{1, 0, 0}, {1, 0, 0}, {1, 0, 1}}, {{1, 0, 0}, {1, 0, 0}, {1, 1, 0}}, {{1, 0, 1}, {0, 1, 0}, {1, 0, 0}}, {{1, 0, 1}, {1, 0, 1}, {1, 0, 0}}, {{1, 0, 1}, {1, 1, 0}, {0, 0, 0}}, {{1, 0, 1}, {1, 1, 0}, {1, 0, 0}}, {{1, 1, 0}, {0, 0, 1}, {0, 0, 0}}, {{1, 1, 0}, {0, 1, 0}, {1, 0, 0}}, {{1, 1, 0}, {1, 0, 1}, {1, 0, 0}}, {{1, 1, 0}, {1, 1, 0}, {0, 0, 0}}, {{1, 1, 0}, {1, 1, 0}, {1, 0, 0}}, {{1, 1, 1}, {0, 0, 0}, {0, 0, 1}}, {{1, 1, 1}, {0, 0, 0}, {1, 1, 0}}, {{1, 1, 1}, {0, 0, 1}, {0, 0, 0}}, {{1, 1, 1}, {0, 1, 0}, {1, 0, 0}}, {{1, 1, 1}, {1, 0, 0}, {0, 1, 0}}, {{1, 1, 1}, {1, 0, 0}, {1, 0, 1}}, {{1, 1, 1}, {1, 0, 0}, {1, 1, 0}}, {{1, 1, 1}, {1, 0, 1}, {1, 0, 0}}, {{1, 1, 1}, {1, 1, 0}, {0, 0, 0}}, {{1, 1, 1}, {1, 1, 0}, {1, 0, 0}}},
	"0011": {{{0, 0, 0}, {0, 0, 0}, {0, 1, 0}}, {{0, 0, 0}, {0, 0, 0}, {1, 0, 1}}, {{0, 0, 1}, {0, 0, 1}, {1, 0, 0}}, {{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}, {{0, 1, 0}, {1, 0, 1}, {0, 0, 0}}, {{0, 1, 1}, {0, 1, 0}, {0, 0, 0}}, {{0, 1, 1}, {1, 0, 0}, {0, 0, 1}}, {{0, 1, 1}, {1, 0, 1}, {0, 0, 0}}, {{1, 0, 0}, {1, 0, 0}, {0, 0, 1}}, {{1, 0, 1}, {0, 1, 0}, {0, 0, 0}}, {{1, 0, 1}, {1, 0, 1}, {0, 0, 0}}, {{1, 1, 0}, {0, 0, 1}, {1, 0, 0}}, {{1, 1, 0}, {0, 1, 0}, {0, 0, 0}}, {{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, {{1, 1, 1}, {0, 0, 0}, {0, 1, 0}}, {{1, 1, 1}, {0, 0, 0}, {1, 0, 1}}, {{1, 1, 1}, {0, 0, 1}, {1, 0, 0}}, {{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}, {{1, 1, 1}, {1, 0, 0}, {0, 0, 1}}, {{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}},
	"0101": {{{0, 0, 0}, {0, 0, 1}, {0, 0, 0}}, {{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}, {{0, 0, 0}, {1, 1, 0}, {1, 0, 0}}, {{0, 0, 1}, {0, 0, 0}, {0, 0, 1}}, {{0, 0, 1}, {0, 0, 0}, {1, 1, 0}}, {{0, 1, 0}, {1, 0, 0}, {0, 1, 0}}, {{0, 1, 0}, {1, 0, 0}, {1, 0, 1}}, {{0, 1, 0}, {1, 0, 0}, {1, 1, 0}}, {{1, 0, 0}, {0, 1, 0}, {1, 0, 0}}, {{1, 0, 0}, {1, 0, 1}, {1, 0, 0}}, {{1, 0, 0}, {1, 1, 0}, {0, 0, 0}}, {{1, 0, 0}, {1, 1, 0}, {1, 0, 0}}, {{1, 0, 1}, {1, 0, 0}, {0, 1, 0}}, {{1, 0, 1}, {1, 0, 0}, {1, 0, 1}}, {{1, 0, 1}, {1, 0, 0}, {1, 1, 0}}, {{1, 1, 0}, {0, 0, 0}, {0, 0, 1}}, {{1, 1, 0}, {0, 0, 0}, {1, 1, 0}}, {{1, 1, 0}, {1, 0, 0}, {0, 1, 0}}, {{1, 1, 0}, {1, 0, 0}, {1, 0, 1}}, {{1, 1, 0}, {1, 0, 0}, {1, 1, 0}}},
	"1110": {{{0, 0, 0}, {0, 1, 0}, {0, 0, 1}}, {{0, 0, 0}, {1, 0, 1}, {0, 0, 1}}, {{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, {{0, 1, 0}, {0, 0, 0}, {0, 1, 1}}, {{0, 1, 0}, {0, 0, 0}, {1, 0, 0}}, {{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}, {{1, 0, 0}, {0, 0, 1}, {0, 1, 1}}, {{1, 0, 0}, {0, 0, 1}, {1, 0, 1}}, {{1, 0, 1}, {0, 0, 0}, {0, 1, 1}}, {{1, 0, 1}, {0, 0, 0}, {1, 0, 0}}},
	"1100": {{{0, 0, 0}, {0, 1, 0}, {0, 1, 0}}, {{0, 0, 0}, {0, 1, 0}, {0, 1, 1}}, {{0, 0, 0}, {0, 1, 0}, {1, 0, 1}}, {{0, 0, 0}, {0, 1, 0}, {1, 1, 0}}, {{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, {{0, 0, 0}, {1, 0, 1}, {0, 1, 0}}, {{0, 0, 0}, {1, 0, 1}, {0, 1, 1}}, {{0, 0, 0}, {1, 0, 1}, {1, 0, 1}}, {{0, 0, 0}, {1, 0, 1}, {1, 1, 0}}, {{0, 0, 0}, {1, 0, 1}, {1, 1, 1}}, {{0, 0, 1}, {1, 0, 0}, {0, 1, 1}}, {{0, 0, 1}, {1, 0, 0}, {1, 0, 0}}, {{0, 0, 1}, {1, 0, 0}, {1, 1, 1}}, {{0, 1, 0}, {0, 0, 0}, {0, 0, 0}}, {{0, 1, 0}, {0, 0, 0}, {1, 1, 1}}, {{1, 0, 0}, {0, 0, 1}, {0, 0, 1}}, {{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}, {{1, 0, 0}, {0, 0, 1}, {1, 1, 1}}, {{1, 0, 1}, {0, 0, 0}, {0, 0, 0}}, {{1, 0, 1}, {0, 0, 0}, {1, 1, 1}}},
	"1011": {{{0, 0, 0}, {1, 0, 0}, {0, 0, 1}}, {{0, 0, 1}, {0, 1, 0}, {0, 0, 0}}, {{0, 0, 1}, {1, 0, 1}, {0, 0, 0}}, {{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}, {{0, 1, 1}, {0, 0, 0}, {0, 1, 0}}, {{0, 1, 1}, {0, 0, 0}, {1, 0, 1}}, {{0, 1, 1}, {0, 0, 1}, {1, 0, 0}}, {{1, 0, 0}, {0, 0, 0}, {0, 1, 0}}, {{1, 0, 0}, {0, 0, 0}, {1, 0, 1}}, {{1, 0, 1}, {0, 0, 1}, {1, 0, 0}}},
	"1001": {{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}}, {{0, 0, 0}, {1, 0, 0}, {1, 0, 1}}, {{0, 0, 0}, {1, 0, 0}, {1, 1, 0}}, {{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}, {{0, 0, 1}, {1, 0, 1}, {1, 0, 0}}, {{0, 1, 0}, {0, 0, 1}, {0, 0, 0}}, {{0, 1, 1}, {0, 0, 0}, {0, 0, 1}}, {{0, 1, 1}, {0, 0, 0}, {1, 1, 0}}, {{0, 1, 1}, {0, 0, 1}, {0, 0, 0}}, {{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}, {{1, 0, 0}, {0, 0, 0}, {1, 1, 0}}, {{1, 0, 1}, {0, 0, 1}, {0, 0, 0}}},
	"0000": {{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, {{0, 0, 0}, {0, 0, 0}, {1, 1, 1}}, {{0, 0, 0}, {1, 1, 1}, {0, 0, 0}}, {{0, 0, 0}, {1, 1, 1}, {0, 0, 1}}, {{0, 0, 0}, {1, 1, 1}, {0, 1, 0}}, {{0, 0, 0}, {1, 1, 1}, {0, 1, 1}}, {{0, 0, 0}, {1, 1, 1}, {1, 0, 0}}, {{0, 0, 0}, {1, 1, 1}, {1, 0, 1}}, {{0, 0, 0}, {1, 1, 1}, {1, 1, 0}}, {{0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, {{0, 0, 1}, {0, 0, 1}, {0, 0, 1}}, {{0, 0, 1}, {0, 0, 1}, {1, 1, 0}}, {{0, 0, 1}, {0, 0, 1}, {1, 1, 1}}, {{0, 0, 1}, {1, 1, 0}, {0, 0, 1}}, {{0, 0, 1}, {1, 1, 0}, {0, 1, 0}}, {{0, 0, 1}, {1, 1, 0}, {0, 1, 1}}, {{0, 0, 1}, {1, 1, 0}, {1, 0, 1}}, {{0, 0, 1}, {1, 1, 0}, {1, 1, 0}}, {{0, 0, 1}, {1, 1, 0}, {1, 1, 1}}, {{0, 0, 1}, {1, 1, 1}, {0, 0, 0}}, {{0, 0, 1}, {1, 1, 1}, {0, 0, 1}}, {{0, 0, 1}, {1, 1, 1}, {0, 1, 0}}, {{0, 0, 1}, {1, 1, 1}, {0, 1, 1}}, {{0, 0, 1}, {1, 1, 1}, {1, 0, 0}}, {{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}, {{0, 0, 1}, {1, 1, 1}, {1, 1, 0}}, {{0, 0, 1}, {1, 1, 1}, {1, 1, 1}}, {{0, 1, 0}, {0, 1, 0}, {0, 1, 0}}, {{0, 1, 0}, {0, 1, 0}, {0, 1, 1}}, {{0, 1, 0}, {0, 1, 0}, {1, 0, 1}}, {{0, 1, 0}, {0, 1, 0}, {1, 1, 0}}, {{0, 1, 0}, {0, 1, 0}, {1, 1, 1}}, {{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}, {{0, 1, 0}, {0, 1, 1}, {0, 1, 1}}, {{0, 1, 0}, {0, 1, 1}, {1, 0, 0}}, {{0, 1, 0}, {0, 1, 1}, {1, 0, 1}}, {{0, 1, 0}, {0, 1, 1}, {1, 1, 0}}, {{0, 1, 0}, {0, 1, 1}, {1, 1, 1}}, {{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, {{0, 1, 0}, {1, 0, 1}, {0, 1, 1}}, {{0, 1, 0}, {1, 0, 1}, {1, 0, 1}}, {{0, 1, 0}, {1, 0, 1}, {1, 1, 0}}, {{0, 1, 0}, {1, 0, 1}, {1, 1, 1}}, {{0, 1, 0}, {1, 1, 0}, {0, 0, 1}}, {{0, 1, 0}, {1, 1, 0}, {0, 1, 0}}, {{0, 1, 0}, {1, 1, 0}, {0, 1, 1}}, {{0, 1, 0}, {1, 1, 0}, {1, 0, 1}}, {{0, 1, 0}, {1, 1, 0}, {1, 1, 0}}, {{0, 1, 0}, {1, 1, 0}, {1, 1, 1}}, {{0, 1, 0}, {1, 1, 1}, {0, 0, 0}}, {{0, 1, 0}, {1, 1, 1}, {0, 0, 1}}, {{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, {{0, 1, 0}, {1, 1, 1}, {0, 1, 1}}, {{0, 1, 0}, {1, 1, 1}, {1, 0, 0}}, {{0, 1, 0}, {1, 1, 1}, {1, 0, 1}}, {{0, 1, 0}, {1, 1, 1}, {1, 1, 0}}, {{0, 1, 0}, {1, 1, 1}, {1, 1, 1}}, {{0, 1, 1}, {0, 1, 0}, {0, 1, 0}}, {{0, 1, 1}, {0, 1, 0}, {0, 1, 1}}, {{0, 1, 1}, {0, 1, 0}, {1, 0, 1}}, {{0, 1, 1}, {0, 1, 0}, {1, 1, 0}}, {{0, 1, 1}, {0, 1, 0}, {1, 1, 1}}, {{0, 1, 1}, {0, 1, 1}, {0, 1, 0}}, {{0, 1, 1}, {0, 1, 1}, {0, 1, 1}}, {{0, 1, 1}, {0, 1, 1}, {1, 0, 0}}, {{0, 1, 1}, {0, 1, 1}, {1, 0, 1}}, {{0, 1, 1}, {0, 1, 1}, {1, 1, 0}}, {{0, 1, 1}, {0, 1, 1}, {1, 1, 1}}, {{0, 1, 1}, {1, 0, 0}, {0, 1, 1}}, {{0, 1, 1}, {1, 0, 0}, {1, 0, 0}}, {{0, 1, 1}, {1, 0, 0}, {1, 1, 1}}, {{0, 1, 1}, {1, 0, 1}, {0, 1, 0}}, {{0, 1, 1}, {1, 0, 1}, {0, 1, 1}}, {{0, 1, 1}, {1, 0, 1}, {1, 0, 1}}, {{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, {{0, 1, 1}, {1, 0, 1}, {1, 1, 1}}, {{0, 1, 1}, {1, 1, 0}, {0, 0, 1}}, {{0, 1, 1}, {1, 1, 0}, {0, 1, 0}}, {{0, 1, 1}, {1, 1, 0}, {0, 1, 1}}, {{0, 1, 1}, {1, 1, 0}, {1, 0, 1}}, {{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}, {{0, 1, 1}, {1, 1, 0}, {1, 1, 1}}, {{0, 1, 1}, {1, 1, 1}, {0, 0, 0}}, {{0, 1, 1}, {1, 1, 1}, {0, 0, 1}}, {{0, 1, 1}, {1, 1, 1}, {0, 1, 0}}, {{0, 1, 1}, {1, 1, 1}, {0, 1, 1}}, {{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}, {{0, 1, 1}, {1, 1, 1}, {1, 0, 1}}, {{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}, {{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 0, 0}, {0, 1, 1}, {0, 1, 0}}, {{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}, {{1, 0, 0}, {0, 1, 1}, {1, 0, 0}}, {{1, 0, 0}, {0, 1, 1}, {1, 0, 1}}, {{1, 0, 0}, {0, 1, 1}, {1, 1, 0}}, {{1, 0, 0}, {0, 1, 1}, {1, 1, 1}}, {{1, 0, 0}, {1, 0, 0}, {0, 1, 1}}, {{1, 0, 0}, {1, 0, 0}, {1, 0, 0}}, {{1, 0, 0}, {1, 0, 0}, {1, 1, 1}}, {{1, 0, 0}, {1, 1, 1}, {0, 0, 0}}, {{1, 0, 0}, {1, 1, 1}, {0, 0, 1}}, {{1, 0, 0}, {1, 1, 1}, {0, 1, 0}}, {{1, 0, 0}, {1, 1, 1}, {0, 1, 1}}, {{1, 0, 0}, {1, 1, 1}, {1, 0, 0}}, {{1, 0, 0}, {1, 1, 1}, {1, 0, 1}}, {{1, 0, 0}, {1, 1, 1}, {1, 1, 0}}, {{1, 0, 0}, {1, 1, 1}, {1, 1, 1}}, {{1, 0, 1}, {0, 1, 0}, {0, 1, 0}}, {{1, 0, 1}, {0, 1, 0}, {0, 1, 1}}, {{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, {{1, 0, 1}, {0, 1, 0}, {1, 1, 0}}, {{1, 0, 1}, {0, 1, 0}, {1, 1, 1}}, {{1, 0, 1}, {0, 1, 1}, {0, 1, 0}}, {{1, 0, 1}, {0, 1, 1}, {0, 1, 1}}, {{1, 0, 1}, {0, 1, 1}, {1, 0, 0}}, {{1, 0, 1}, {0, 1, 1}, {1, 0, 1}}, {{1, 0, 1}, {0, 1, 1}, {1, 1, 0}}, {{1, 0, 1}, {0, 1, 1}, {1, 1, 1}}, {{1, 0, 1}, {1, 0, 1}, {0, 1, 0}}, {{1, 0, 1}, {1, 0, 1}, {0, 1, 1}}, {{1, 0, 1}, {1, 0, 1}, {1, 0, 1}}, {{1, 0, 1}, {1, 0, 1}, {1, 1, 0}}, {{1, 0, 1}, {1, 0, 1}, {1, 1, 1}}, {{1, 0, 1}, {1, 1, 0}, {0, 0, 1}}, {{1, 0, 1}, {1, 1, 0}, {0, 1, 0}}, {{1, 0, 1}, {1, 1, 0}, {0, 1, 1}}, {{1, 0, 1}, {1, 1, 0}, {1, 0, 1}}, {{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, {{1, 0, 1}, {1, 1, 0}, {1, 1, 1}}, {{1, 0, 1}, {1, 1, 1}, {0, 0, 0}}, {{1, 0, 1}, {1, 1, 1}, {0, 0, 1}}, {{1, 0, 1}, {1, 1, 1}, {0, 1, 0}}, {{1, 0, 1}, {1, 1, 1}, {0, 1, 1}}, {{1, 0, 1}, {1, 1, 1}, {1, 0, 0}}, {{1, 0, 1}, {1, 1, 1}, {1, 0, 1}}, {{1, 0, 1}, {1, 1, 1}, {1, 1, 0}}, {{1, 0, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 0}, {0, 0, 1}, {0, 0, 1}}, {{1, 1, 0}, {0, 0, 1}, {1, 1, 0}}, {{1, 1, 0}, {0, 0, 1}, {1, 1, 1}}, {{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, {{1, 1, 0}, {0, 1, 0}, {0, 1, 1}}, {{1, 1, 0}, {0, 1, 0}, {1, 0, 1}}, {{1, 1, 0}, {0, 1, 0}, {1, 1, 0}}, {{1, 1, 0}, {0, 1, 0}, {1, 1, 1}}, {{1, 1, 0}, {0, 1, 1}, {0, 1, 0}}, {{1, 1, 0}, {0, 1, 1}, {0, 1, 1}}, {{1, 1, 0}, {0, 1, 1}, {1, 0, 0}}, {{1, 1, 0}, {0, 1, 1}, {1, 0, 1}}, {{1, 1, 0}, {0, 1, 1}, {1, 1, 0}}, {{1, 1, 0}, {0, 1, 1}, {1, 1, 1}}, {{1, 1, 0}, {1, 0, 1}, {0, 1, 0}}, {{1, 1, 0}, {1, 0, 1}, {0, 1, 1}}, {{1, 1, 0}, {1, 0, 1}, {1, 0, 1}}, {{1, 1, 0}, {1, 0, 1}, {1, 1, 0}}, {{1, 1, 0}, {1, 0, 1}, {1, 1, 1}}, {{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, {{1, 1, 0}, {1, 1, 0}, {0, 1, 0}}, {{1, 1, 0}, {1, 1, 0}, {0, 1, 1}}, {{1, 1, 0}, {1, 1, 0}, {1, 0, 1}}, {{1, 1, 0}, {1, 1, 0}, {1, 1, 0}}, {{1, 1, 0}, {1, 1, 0}, {1, 1, 1}}, {{1, 1, 0}, {1, 1, 1}, {0, 0, 0}}, {{1, 1, 0}, {1, 1, 1}, {0, 0, 1}}, {{1, 1, 0}, {1, 1, 1}, {0, 1, 0}}, {{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}, {{1, 1, 0}, {1, 1, 1}, {1, 0, 0}}, {{1, 1, 0}, {1, 1, 1}, {1, 0, 1}}, {{1, 1, 0}, {1, 1, 1}, {1, 1, 0}}, {{1, 1, 0}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {0, 0, 0}, {0, 0, 0}}, {{1, 1, 1}, {0, 0, 0}, {1, 1, 1}}, {{1, 1, 1}, {0, 0, 1}, {0, 0, 1}}, {{1, 1, 1}, {0, 0, 1}, {1, 1, 0}}, {{1, 1, 1}, {0, 0, 1}, {1, 1, 1}}, {{1, 1, 1}, {0, 1, 0}, {0, 1, 0}}, {{1, 1, 1}, {0, 1, 0}, {0, 1, 1}}, {{1, 1, 1}, {0, 1, 0}, {1, 0, 1}}, {{1, 1, 1}, {0, 1, 0}, {1, 1, 0}}, {{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}, {{1, 1, 1}, {0, 1, 1}, {0, 1, 0}}, {{1, 1, 1}, {0, 1, 1}, {0, 1, 1}}, {{1, 1, 1}, {0, 1, 1}, {1, 0, 0}}, {{1, 1, 1}, {0, 1, 1}, {1, 0, 1}}, {{1, 1, 1}, {0, 1, 1}, {1, 1, 0}}, {{1, 1, 1}, {0, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 0, 0}, {0, 1, 1}}, {{1, 1, 1}, {1, 0, 0}, {1, 0, 0}}, {{1, 1, 1}, {1, 0, 0}, {1, 1, 1}}, {{1, 1, 1}, {1, 0, 1}, {0, 1, 0}}, {{1, 1, 1}, {1, 0, 1}, {0, 1, 1}}, {{1, 1, 1}, {1, 0, 1}, {1, 0, 1}}, {{1, 1, 1}, {1, 0, 1}, {1, 1, 0}}, {{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 0}, {0, 0, 1}}, {{1, 1, 1}, {1, 1, 0}, {0, 1, 0}}, {{1, 1, 1}, {1, 1, 0}, {0, 1, 1}}, {{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, {{1, 1, 1}, {1, 1, 0}, {1, 1, 0}}, {{1, 1, 1}, {1, 1, 0}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {0, 0, 0}}, {{1, 1, 1}, {1, 1, 1}, {0, 0, 1}}, {{1, 1, 1}, {1, 1, 1}, {0, 1, 0}}, {{1, 1, 1}, {1, 1, 1}, {0, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 0, 0}}, {{1, 1, 1}, {1, 1, 1}, {1, 0, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}},
	"0010": {{{0, 0, 0}, {0, 0, 0}, {0, 1, 1}}, {{0, 0, 0}, {0, 0, 0}, {1, 0, 0}}, {{0, 0, 1}, {0, 0, 1}, {0, 1, 0}}, {{0, 0, 1}, {0, 0, 1}, {0, 1, 1}}, {{0, 0, 1}, {0, 0, 1}, {1, 0, 1}}, {{0, 1, 0}, {0, 1, 0}, {0, 0, 1}}, {{0, 1, 0}, {0, 1, 1}, {0, 0, 0}}, {{0, 1, 0}, {0, 1, 1}, {0, 0, 1}}, {{0, 1, 0}, {1, 0, 1}, {0, 0, 1}}, {{0, 1, 1}, {0, 1, 0}, {0, 0, 1}}, {{0, 1, 1}, {0, 1, 1}, {0, 0, 0}}, {{0, 1, 1}, {0, 1, 1}, {0, 0, 1}}, {{0, 1, 1}, {1, 0, 0}, {0, 0, 0}}, {{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}, {{1, 0, 0}, {0, 1, 1}, {0, 0, 0}}, {{1, 0, 0}, {0, 1, 1}, {0, 0, 1}}, {{1, 0, 0}, {1, 0, 0}, {0, 0, 0}}, {{1, 0, 1}, {0, 1, 0}, {0, 0, 1}}, {{1, 0, 1}, {0, 1, 1}, {0, 0, 0}}, {{1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, {{1, 0, 1}, {1, 0, 1}, {0, 0, 1}}, {{1, 1, 0}, {0, 0, 1}, {0, 1, 0}}, {{1, 1, 0}, {0, 0, 1}, {0, 1, 1}}, {{1, 1, 0}, {0, 0, 1}, {1, 0, 1}}, {{1, 1, 0}, {0, 1, 0}, {0, 0, 1}}, {{1, 1, 0}, {0, 1, 1}, {0, 0, 0}}, {{1, 1, 0}, {0, 1, 1}, {0, 0, 1}}, {{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, {{1, 1, 1}, {0, 0, 0}, {0, 1, 1}}, {{1, 1, 1}, {0, 0, 0}, {1, 0, 0}}, {{1, 1, 1}, {0, 0, 1}, {0, 1, 0}}, {{1, 1, 1}, {0, 0, 1}, {0, 1, 1}}, {{1, 1, 1}, {0, 0, 1}, {1, 0, 1}}, {{1, 1, 1}, {0, 1, 0}, {0, 0, 1}}, {{1, 1, 1}, {0, 1, 1}, {0, 0, 0}}, {{1, 1, 1}, {0, 1, 1}, {0, 0, 1}}, {{1, 1, 1}, {1, 0, 0}, {0, 0, 0}}, {{1, 1, 1}, {1, 0, 1}, {0, 0, 1}}},
	"0110": {{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}}, {{0, 0, 0}, {0, 0, 1}, {0, 1, 1}}, {{0, 0, 0}, {0, 0, 1}, {1, 0, 1}}, {{0, 0, 1}, {0, 0, 0}, {0, 1, 1}}, {{0, 0, 1}, {0, 0, 0}, {1, 0, 0}}, {{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}, {{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, {{1, 0, 0}, {1, 0, 1}, {0, 0, 1}}, {{1, 0, 1}, {1, 0, 0}, {0, 0, 0}}, {{1, 1, 0}, {0, 0, 0}, {0, 1, 1}}, {{1, 1, 0}, {0, 0, 0}, {1, 0, 0}}, {{1, 1, 0}, {1, 0, 0}, {0, 0, 0}}},
	"0111": {{{0, 0, 0}, {0, 0, 1}, {1, 0, 0}}, {{0, 0, 1}, {0, 0, 0}, {0, 1, 0}}, {{0, 0, 1}, {0, 0, 0}, {1, 0, 1}}, {{0, 1, 0}, {1, 0, 0}, {0, 0, 1}}, {{1, 0, 0}, {0, 1, 0}, {0, 0, 0}}, {{1, 0, 0}, {1, 0, 1}, {0, 0, 0}}, {{1, 0, 1}, {1, 0, 0}, {0, 0, 1}}, {{1, 1, 0}, {0, 0, 0}, {0, 1, 0}}, {{1, 1, 0}, {0, 0, 0}, {1, 0, 1}}, {{1, 1, 0}, {1, 0, 0}, {0, 0, 1}}},
}

func main() {
	var lastStates = [][][]int{
		{
			{0, 1},
		},
		{
			{0},
			{1},
		},
		{
			{1, 1},
			{1, 1},
		},
		{
			{1, 0},
			{0, 0},
		},
		{
			{0, 1},
			{1, 0},
		},
		{
			{0, 1},
			{1, 1},
		},
		{
			{1, 1},
			{1, 0},
		},
		//{
		//	{1, 0},
		//	{0, 1},
		//},
	}
	for _, lastState := range lastStates {
		start := time.Now()
		count = 0
		var rowCount = len(lastState)
		var colCount = len(lastState[0])
		var originState = genOriginState(rowCount, colCount)
		exec(originState, lastState, rowCount, colCount, 0, 0)
		var caseCount = genCaseCount(rowCount, colCount)
		coolExec(caseCount, lastState, rowCount, colCount)

		fmt.Printf("cc %+v\n", cc)
		cc = CaseCount{}
		fmt.Printf("result: %d - process in %v\n", count, time.Since(start))
	}
}

// cộng 1 vào 0 tức là số case chuỗi gốc trừ chuỗi công 1 với đuối thay bằng 1
// cộng số case chuỗi gốc với đuối thay bằng 1

// cộng 1 vào 1 tức là nhân đôi case trước của chuỗi gốc
// trừ 2 lần case trước nữa

// cộng 0 vào 1 tức là thê đc 4 case ở đuôi (x4)
// trừ số case nếu cộng 1

// cộng 0 vào 0 tức là têm đc 4 case ở đuôi (x4)
// trừ đi số case nếu cộng 1

var res []int

type CaseCount struct {
	OneTopRight       int64
	OneBottomLeft     int64
	OneBottomRight    int64
	ZeroRight         int64
	ZeroBottom        int64
	OneRight          int64
	OneBottom         int64
	HasZeroBottomLeft int64
	HasOneBottomLeft  int64
}

var counted = []CaseCount{{
	OneTopRight:       3,
	OneBottomLeft:     3,
	OneBottomRight:    3,
	ZeroRight:         2,
	ZeroBottom:        2,
	OneRight:          4,
	OneBottom:         4,
	HasZeroBottomLeft: 5,
	HasOneBottomLeft:  7,
}, {
	OneTopRight:       1,
	OneBottomLeft:     1,
	OneBottomRight:    1,
	ZeroRight:         2,
	ZeroBottom:        2,
	HasZeroBottomLeft: 3,
	HasOneBottomLeft:  1,
}}

func (c CaseCount) Count() int64 {
	return c.OneTopRight + c.OneBottomRight + c.ZeroRight + c.OneRight
}

// total 1 = top right + bottom right +  zero right
// total 0 = top right + bottom right + zero right + one right
func coolExec(caseCount [][]CaseCount, lastState [][]int, rowCount, colCount int) {

	caseCount[0][0] = counted[lastState[0][0]]
	for i := 1; i < rowCount; i++ {
		if lastState[i][0] == 0 {
			caseCount[i][0] = CaseCount{
				OneTopRight:    2*caseCount[i-1][0].OneBottom + caseCount[i-1][0].OneBottomRight,
				OneBottomLeft:  caseCount[i-1][0].OneBottom + caseCount[i-1][0].OneBottomLeft + caseCount[i-1][0].OneTopRight,
				OneBottomRight: caseCount[i-1][0].ZeroBottom + 2*caseCount[i-1][0].OneBottomLeft,
				ZeroRight:      caseCount[i-1][0].ZeroBottom + caseCount[i-1][0].OneBottomLeft,
				ZeroBottom:     caseCount[i-1][0].OneBottom + caseCount[i-1][0].ZeroBottom,
				OneRight:       2*caseCount[i-1][0].OneBottom + 2*caseCount[i-1][0].OneBottomRight,
				OneBottom:      caseCount[i-1][0].OneBottom + caseCount[i-1][0].ZeroBottom + caseCount[i-1][0].OneBottomLeft + caseCount[i-1][0].OneBottomRight,
			}
			caseCount[i][0].HasOneBottomLeft = caseCount[i][0].OneRight + caseCount[i][0].OneBottomRight
			caseCount[i][0].HasZeroBottomLeft = caseCount[i][0].OneTopRight + caseCount[i][0].ZeroRight
		} else {
			caseCount[i][0] = CaseCount{
				OneTopRight:    caseCount[i-1][0].OneBottomRight,
				OneBottomLeft:  caseCount[i-1][0].ZeroBottom,
				OneBottomRight: caseCount[i-1][0].ZeroBottom,
				ZeroRight:      caseCount[i-1][0].OneBottomLeft + caseCount[i-1][0].ZeroBottom,
				ZeroBottom:     caseCount[i-1][0].OneBottomRight + caseCount[i-1][0].OneBottomLeft,
			}
			caseCount[i][0].HasOneBottomLeft = caseCount[i][0].OneRight + caseCount[i][0].OneBottomRight
			caseCount[i][0].HasZeroBottomLeft = caseCount[i][0].OneTopRight + caseCount[i][0].ZeroRight
		}
	}

	for j := 1; j < colCount; j++ {
		if lastState[0][j] == 0 {
			caseCount[0][j] = CaseCount{
				OneTopRight:    caseCount[0][j-1].OneTopRight + caseCount[0][j-1].OneBottomRight + caseCount[0][j-1].OneRight,
				OneBottomLeft:  caseCount[0][j-1].OneBottomRight + 2*caseCount[0][j-1].OneRight,
				OneBottomRight: caseCount[0][j-1].OneTopRight + caseCount[0][j-1].OneBottomRight + caseCount[0][j-1].OneRight,
				ZeroRight:      caseCount[0][j-1].OneRight + caseCount[0][j-1].ZeroRight,
				ZeroBottom:     caseCount[0][j-1].OneTopRight + caseCount[0][j-1].ZeroRight,
				OneRight:       caseCount[0][j-1].ZeroRight + caseCount[0][j-1].OneBottomRight + caseCount[0][j-1].OneTopRight + caseCount[0][j-1].OneRight,
				OneBottom:      caseCount[0][j-1].OneTopRight + caseCount[0][j-1].OneBottomRight + caseCount[0][j-1].OneRight,
			}
			caseCount[0][j].HasOneBottomLeft = caseCount[0][j].OneRight + caseCount[0][j].OneBottomRight
			caseCount[0][j].HasZeroBottomLeft = caseCount[0][j].OneTopRight + caseCount[0][j].ZeroRight
		} else {
			caseCount[0][j] = CaseCount{
				OneTopRight:    caseCount[0][j-1].ZeroRight,
				OneBottomLeft:  caseCount[0][j-1].OneBottomRight,
				OneBottomRight: caseCount[0][j-1].ZeroRight,
				ZeroRight:      caseCount[0][j-1].OneTopRight + caseCount[0][j-1].OneBottomRight,
				ZeroBottom:     caseCount[0][j-1].ZeroRight + caseCount[0][j-1].OneTopRight,
			}
			caseCount[0][j].HasOneBottomLeft = caseCount[0][j].OneRight + caseCount[0][j].OneBottomRight
			caseCount[0][j].HasZeroBottomLeft = caseCount[0][j].OneTopRight + caseCount[0][j].ZeroRight
		}
	}
	for i := 1; i < rowCount; i++ {
		for j := 1; j < colCount; j++ {
			fmt.Println(caseCount[i-1][j-1].HasZeroBottomLeft, caseCount[i-1][j-1].HasOneBottomLeft)
			if lastState[i][j] == 0 {
				caseCount[i][j] = CaseCount{
					OneTopRight:    caseCount[i][j-1].OneTopRight + caseCount[i][j-1].OneBottomRight + caseCount[i][j-1].OneRight + 2*caseCount[i-1][j].OneBottom + caseCount[i-1][j].OneBottomRight - caseCount[i-1][j-1].HasZeroBottomLeft - caseCount[i-1][j-1].HasOneBottomLeft,
					OneBottomLeft:  caseCount[i][j-1].OneBottomRight + 2*caseCount[i][j-1].OneRight + caseCount[i-1][j].OneBottom + caseCount[i-1][j].OneBottomLeft + caseCount[i-1][j].OneTopRight - caseCount[i-1][j-1].HasZeroBottomLeft - caseCount[i-1][j-1].HasOneBottomLeft,
					OneBottomRight: caseCount[i][j-1].OneTopRight + caseCount[i][j-1].OneBottomRight + caseCount[i][j-1].OneRight + caseCount[i-1][j].ZeroBottom + 2*caseCount[i-1][j].OneBottomLeft - 1 - caseCount[i-1][j-1].HasOneBottomLeft - caseCount[i-1][j-1].HasZeroBottomLeft,
					ZeroRight:      caseCount[i][j-1].OneRight + caseCount[i][j-1].ZeroRight + caseCount[i-1][j].ZeroBottom + caseCount[i-1][j].OneBottomLeft - 1 - caseCount[i-1][j-1].HasZeroBottomLeft,
					ZeroBottom:     caseCount[i][j-1].OneTopRight + caseCount[i][j-1].ZeroRight + caseCount[i-1][j].OneBottom + caseCount[i-1][j].ZeroBottom - 1 - caseCount[i-1][j-1].HasZeroBottomLeft,
					OneRight:       caseCount[i][j-1].ZeroRight + caseCount[i][j-1].OneBottomRight + caseCount[i][j-1].OneTopRight + caseCount[i][j-1].OneRight + 2*caseCount[i-1][j].OneBottom + 2*caseCount[i-1][j].OneBottomRight - 1 - 2*caseCount[i-1][j-1].HasZeroBottomLeft,
					OneBottom:      caseCount[i][j-1].OneTopRight + caseCount[i][j-1].OneBottomRight + caseCount[i][j-1].OneRight + caseCount[i-1][j].OneBottom + caseCount[i-1][j].ZeroBottom + caseCount[i-1][j].OneBottomLeft + caseCount[i-1][j].OneBottomRight - 2*caseCount[i-1][j-1].HasZeroBottomLeft,
				}
				caseCount[i][j].HasOneBottomLeft = caseCount[i][j].OneRight + caseCount[i][j].OneBottomRight
				caseCount[i][j].HasZeroBottomLeft = caseCount[i][j].OneTopRight + caseCount[i][j].ZeroRight
			} else {
				caseCount[i][j] = CaseCount{
					OneTopRight:    caseCount[i][j-1].ZeroRight + caseCount[i-1][j].OneBottomRight - caseCount[i-1][j-1].HasZeroBottomLeft,
					OneBottomLeft:  caseCount[i][j-1].OneBottomRight + caseCount[i-1][j].ZeroBottom - caseCount[i-1][j-1].HasZeroBottomLeft,
					OneBottomRight: caseCount[i][j-1].ZeroRight + caseCount[i-1][j].ZeroBottom - caseCount[i-1][j-1].HasZeroBottomLeft,
					ZeroRight:      caseCount[i][j-1].OneTopRight + caseCount[i][j-1].OneBottomRight + caseCount[i-1][j].OneBottomLeft + caseCount[i-1][j].ZeroBottom - 2*caseCount[i-1][j-1].HasZeroBottomLeft,
					ZeroBottom:     caseCount[i][j-1].ZeroRight + caseCount[i][j-1].OneTopRight + caseCount[i-1][j].OneBottomRight + caseCount[i-1][j].OneBottomLeft - 2*caseCount[i-1][j-1].HasZeroBottomLeft,
				}
				caseCount[i][j].HasOneBottomLeft = caseCount[i][j].OneRight + caseCount[i][j].OneBottomRight
				caseCount[i][j].HasZeroBottomLeft = caseCount[i][j].OneTopRight + caseCount[i][j].ZeroRight
			}

		}
	}

	fmt.Printf("count %+v %v\n", caseCount[rowCount-1][colCount-1], caseCount[rowCount-1][colCount-1].Count())
}

var cc = CaseCount{}

func exec(originState, lastState [][]int, rowCount, colCount, i, j int) {
	for _, c := range cFromP[lastState[i][j]] {
		if check(originState, i, j, c[0], c[1], c[2], c[3]) {
			originState[i][j] = c[0]
			originState[i][j+1] = c[1]
			originState[i+1][j] = c[2]
			originState[i+1][j+1] = c[3]
			if j < colCount-1 {
				exec(originState, lastState, rowCount, colCount, i, j+1)
			} else if i < rowCount-1 {
				exec(originState, lastState, rowCount, colCount, i+1, 0)
			} else {
				if originState[i][j+1] == 1 && originState[i+1][j+1] == 0 {
					cc.OneTopRight++
				}
				if originState[i+1][j] == 1 && originState[i+1][j+1] == 0 {
					cc.OneBottomLeft++
				}
				if originState[i][j+1] == 0 && originState[i+1][j+1] == 1 {
					cc.OneBottomRight++
				}
				if originState[i][j+1] == 0 && originState[i+1][j+1] == 0 {
					cc.ZeroRight++
				}
				if originState[i+1][j] == 0 && originState[i+1][j+1] == 0 {
					cc.ZeroBottom++
				}
				if originState[i][j+1] == 1 && originState[i+1][j+1] == 1 {
					cc.OneRight++
				}
				if originState[i+1][j] == 1 && originState[i+1][j+1] == 1 {
					cc.OneBottom++
				}
				if originState[i+1][j+1] == 1 {
					cc.HasOneBottomLeft++
				}
				if originState[i+1][j+1] == 0 {
					cc.HasZeroBottomLeft++
				}
				//for _, v := range originState {
				//	fmt.Printf("%v %v %+v\n", i, j, v)
				//}
				//fmt.Println()
				count++
			}
		}
	}
}

// func exec2x2(originState, lastState [][]int, rowCount, colCount, i, j int) {
// 	key := fmt.Sprintf("%d%d%d%d", lastState[i][j], lastState[i][j+1], lastState[i+1][j], lastState[i+1][j+1])
// 	for _, c := range c2x2FromP[key] {
// 		if check2x2(originState, i, j, c) {
// 			originState[i][j] = c[0]
// 			originState[i][j+1] = c[1]
// 			originState[i+1][j] = c[2]
// 			originState[i+1][j+1] = c[3]
// 			if j < colCount-1 {
// 				exec(originState, lastState, rowCount, colCount, i, j+1)
// 			} else if i < rowCount-1 {
// 				exec(originState, lastState, rowCount, colCount, i+1, 0)
// 			} else {
// 				// for _, v := range originState {
// 				// 	fmt.Printf("%v %v %+v\n", i, j, v)
// 				// }
// 				count++
// 			}
// 		}
// 	}
// }

func check2x2(originState [][]int, i, j int, p [][]int) bool {

	return false
}

func check(originState [][]int, i, j, p00, p01, p10, p11 int) bool {
	if i == 0 && j == 0 {
		return true
	}

	if i == 0 {
		if originState[i][j] == p00 && originState[i+1][j] == p10 {
			return true
		}
		return false
	}

	if j == 0 {
		if originState[i][j] == p00 && originState[i][j+1] == p01 {
			return true
		}
		return false
	}

	if originState[i][j] == p00 && originState[i+1][j] == p10 && originState[i][j+1] == p01 {
		return true
	}

	return false
}

func genOriginState(rowCount, colCount int) [][]int {
	originState := [][]int{}
	for i := 0; i <= rowCount; i++ {
		arr := []int{}
		for j := 0; j <= colCount; j++ {
			arr = append(arr, 0)
		}
		originState = append(originState, arr)
	}
	return originState
}

func genCaseCount(rowCount, colCount int) [][]CaseCount {
	caseCount := [][]CaseCount{}
	for i := 0; i <= rowCount; i++ {
		arr := []CaseCount{}
		for j := 0; j <= colCount; j++ {
			arr = append(arr, CaseCount{})
		}
		caseCount = append(caseCount, arr)
	}
	return caseCount
}

func min(as ...int64) int64 {
	m := as[0]
	for i := 1; i < len(as); i++ {
		if as[i] < m {
			m = as[i]
		}
	}
	return m
}
