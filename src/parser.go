package src

import (
	"fmt"
	"log"
	//"encoding/json"
	"go_lang/go_lang_packager/packager"
)

type Settings struct {
	inter_access	bool
}

type Parser struct {
	lexer		interface{}
	lang_info	interface{}
	settings	Settings
}

func Init_parser(lexer interface{}) *Parser {
	return &Parser{ lexer: lexer, settings: Settings{inter_access: false} }
}

func (parser *Parser) get_next_token(lexer *Lexer) *Lexer {
	return lexer.Lex()
}

func (parser *Parser) parse_require(lexer *Lexer) *Parser {
	parser.get_next_token(lexer)


	if lexer.Current_Token == T_STR {

		info := packager.Read_info_package(lexer.Token_value)

		switch lexer.Token_value {
			case "lang_info.json": parser.lang_info = info.(*packager.LangInfo)
		}

		parser.get_next_token(lexer)

		if lexer.Current_Token == T_SEMI {
			parser.get_next_token(lexer)
		}

		parser.lexer = lexer

		return parser
	}

	log.Fatal("Parsing Error")
	return parser
}

// Function prone to change
func (parser *Parser) parse_print(lexer *Lexer) *Parser {
	parser.get_next_token(lexer)

	switch lexer.Current_Token {
		case K_VAR_N: {
			switch lexer.Token_value {
				case "inter": {
					if parser.settings.inter_access == false {
						log.Fatal("[Inter Access Error] -> You do not have permission to use inter")
					}
					// Strict Syntax(For now)
					parser.get_next_token(lexer)

					if lexer.Current_Token == T_DOT {
						parser.get_next_token(lexer)
					} else {
						log.Fatal("[PARSING ERROR] -> Expected `.`")
					}

					switch lexer.Token_value {
						case "LangInfo": {
							parser.get_next_token(lexer)

							if lexer.Current_Token == T_DOT {
								parser.get_next_token(lexer)
							} else {
								log.Fatal("[PARSING ERROR] -> Expected `.`")
							}
							//parser.get_next_token(lexer)
							LI := parser.lang_info.(*packager.LangInfo)

							switch lexer.Token_value {
								case "lang_name":fmt.Println(LI.LangName)
								case "lang_version":fmt.Println(LI.LangVersion)
							}
						}
					}
				}
			}
		}
		case T_STR: {
			fmt.Println(lexer.Token_value)
		}
	}

	parser.get_next_token(lexer)

	return parser
}

func (parser *Parser) parse_wrap(lexer *Lexer) *Parser {
	parser.get_next_token(lexer)

	if lexer.Current_Token == T_LSB {
		parser.get_next_token(lexer)

		switch lexer.Token_value {
			case "allow": {
				parser.get_next_token(lexer)

				if lexer.Current_Token == T_LP {
					parser.get_next_token(lexer)

					switch lexer.Token_value {
						case "inter_access": {
							parser.settings.inter_access = true
							parser.get_next_token(lexer)
						}
					}

					if lexer.Current_Token == T_RP {
						parser.get_next_token(lexer)

						if lexer.Current_Token == T_RSB {
							parser.get_next_token(lexer)

							return parser
						}
						log.Fatal(fmt.Sprintf("[PARSING ERROR] -> Expected ']', got `%s` instead", lexer.Token_value))
					}
					log.Fatal(fmt.Sprintf("[PARSING ERROR] -> Expected ')', got `%s` instead", lexer.Token_value))
				}
				log.Fatal(fmt.Sprintf("[PARSING ERROR] -> Expected '(' got `%s` instead", lexer.Token_value))
			}
			log.Fatal(fmt.Sprintf("[PARSING ERROR] -> Expected ')', got `%s` instead", lexer.Token_value))
		}
		log.Fatal("[PARSING ERROR] -> Error parsing wrapper")
	}

	log.Fatal("[PARSING ERROR] -> Invalid syntax to wrapper")
	return parser;
}

func (parser *Parser) parse_func(lexer *Lexer) *Parser {
	parser.get_next_token(lexer)

	if lexer.Current_Token == K_VAR_N {
		parser.get_next_token(lexer)

		if lexer.Current_Token == T_LP {
			parser.get_next_token(lexer)

			if !(lexer.Current_Token == T_RP) {
				// do something
			} else {
				parser.get_next_token(lexer)

				if !(lexer.Current_Token == T_DA) {
					log.Fatal("Expected -> ReturnType for function")
				}

				parser.get_next_token(lexer)

				if !(lexer.Current_Token == K_VAR_N) {
					log.Fatal("ReturnType required after `->`")
				}
				parser.get_next_token(lexer)

				if !(lexer.Current_Token == T_LB) {
					log.Fatal("Expected beginning of function body, `{`")
				}
				parser.get_next_token(lexer)

				parser.Parse_Function_Code(lexer)

				if !(lexer.Current_Token == T_RB) {
					log.Fatal("Expected ending of function body, `}`")
				}

				parser.get_next_token(lexer)
			}

			return parser
		}
		log.Fatal(fmt.Sprintf("[PARSING ERROR] Expected '(', got `%s`", lexer.Token_value))
	}
	return parser
}

func (parser *Parser) Parse_Function_Code(lexer *Lexer) *Parser {

	esc := 1
	for {
		switch lexer.Current_Token {
			case K_PRINT: parser.parse_print(lexer)
			default: esc = 0
		}
		if esc == 0 {
			break
		}
	}

	return parser
}

func (parser *Parser) Parse() *Parser {

	lexer := parser.lexer.(*Lexer)
	for {
		switch lexer.Current_Token {
			case K_REQ: parser.parse_require(lexer)
			case K_PRINT: parser.parse_print(lexer)
			case K_FUNC: parser.parse_func(lexer)
			case T_WRAP: parser.parse_wrap(lexer)
			case T_EOF: return parser
			default: log.Fatal(fmt.Sprintf("[PARSING ERROR] -> Error whilst parsing(%s)", lexer.Token_value))
		}
	}

	return parser
}
