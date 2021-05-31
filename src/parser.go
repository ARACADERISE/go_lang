package src

import (
	"fmt"
	"log"
	//"encoding/json"
	"go_lang/go_lang_packager/packager"
)

type Parser struct {
	lexer		interface{}
	lang_info	*packager.LangInfo
}

func Init_parser(lexer interface{}) *Parser {
	return &Parser{ lexer: lexer }
}

func (parser *Parser) get_next_token(lexer *Lexer) *Lexer {
	return lexer.Lex()
}

func (parser *Parser) parse_require(lexer *Lexer) *Parser {
	parser.get_next_token(lexer)

	switch lexer.Current_token() {
		case T_STR: {
			parser.get_next_token(lexer)

			info := packager.Read_info_package(lexer.Token_value)

			parser.lang_info = info
		}
		default: log.Fatal(fmt.Sprintf("[PARSING ERROR] -> Expected string, got %s", lexer.Token_value))
	}

	parser.get_next_token(lexer)

	if lexer.Current_token() == T_SEMI {
		parser.get_next_token(lexer)
	}
	return parser
}

func (parser *Parser) Parse() *Parser {

	lexer := parser.lexer.(*Lexer)

	for {
		switch lexer.Current_token() {
			case K_REQ: parser.parse_require(lexer)
			case T_EOF: return parser
			default: log.Fatal("Unexpected token")
		}
	}
}
