package main

import (
	"fmt"
	"go_lang/src"
	"go_lang/go_lang_packager/packager"
)

func main() {
	packager.Package("dude.t")
	var info src.LexerI

	info = src.Init_lexer("dude.t")
	info.Lex()
	fmt.Println(info)
}
