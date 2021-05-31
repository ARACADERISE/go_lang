package main

import (
	"go_lang/src"
	"go_lang/go_lang_packager/packager"
)

func main() {
	packager.Package("main.tf")
	var lexer src.LexerI

	lexer = src.Init_lexer("main.tf")
	lexer.Lex()
	parser := src.Init_parser(lexer)
	parser.Parse()
}
