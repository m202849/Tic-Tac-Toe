package tic

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	tokens []int
}

func (b *Board) InitialBoard() {
	for i = 0; i < 9; i++ {
		b.tokens[i] = 0
	}
}

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = 2
	} else {
		fmt.Println("bad input")
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 2 {
		return "x"
	} else {
		return "・"
	}
}
