package main

import (
	"fmt"
	"go_lang/src"
)

func main() {
	var info src.LexerI

	info = src.Init_lexer("dude.t")
	info.Lex()
	fmt.Println(info)
}
