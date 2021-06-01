package src

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"fmt"
)

// Tokens
const (
	Default = 1
	T_LSB	= 2
	T_RSB	= 3
	T_SEMI	= 4
	T_STR	= 5
	T_DOT	= 6
	T_WRAP	= 7
	K_VAR_N = 8
	K_LET	= 9
	K_REQ	= 10
	K_PRINT = 11
	T_EOF	= 12
	T_LB	= 13
	T_RB	= 14
	T_LP	= 15
	T_RP	= 16
)

// Errors
const (
	NoDir		= 128
	ReadErr		= 134
	NoFile		= 138
	InvalidToken	= 140
	UnknownKeyword 	= 144
)

type Lexer struct {
	File_content	[]byte
	File_size	int
	Errors		[]error
	Current_Token	int
	Token_value	string
	index		int
}

type LexerI interface {
	Lex() *Lexer
	Current_token() int
}

func Init_lexer(filename string) *Lexer {
	dir, err := filepath.Abs(filename)

	if err != nil {
		log.Fatal(fmt.Sprintf("[ERROR %d] -> Error find main path to file", NoDir))
	}

	file, Err := os.Stat(dir)

	if Err != nil {
		log.Fatal(fmt.Sprintf("[ERROR %d] -> Error openeing %s", ReadErr, filename))
	}

	info := Lexer{ File_size: int(file.Size()), Current_Token: Default, index: 0 }

	data, E := ioutil.ReadFile(dir)

	if E != nil {
		log.Fatal(fmt.Sprintf("[ERROR %d] -> error reading file %s", ReadErr, filename))
	}

	info.File_content = data

	return &info
}

func (lexer *Lexer) advance_with_token(token_type int, token_value string) *Lexer {
	lexer.index += 1
	lexer.Current_Token = token_type
	lexer.Token_value = token_value

	return lexer
}

func is_alpha(val byte) bool {
	if(val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') {
		return true
	}
	return false
}

func (lexer *Lexer) pickup_keyword() (string, *Lexer) {
	keyword := ""

	for {
		if lexer.File_content[lexer.index] == lexer.File_content[len(lexer.File_content) - 1] {
			keyword += string(lexer.File_content[lexer.index])
			break
		}
		if lexer.File_content[lexer.index] == '|' {
			//lexer.index += 1
			break
		}
		if lexer.File_content[lexer.index] == '.' {
			break
		}
		if lexer.File_content[lexer.index] == ' ' {
			break
		}
		if lexer.File_content[lexer.index] == '(' || lexer.File_content[lexer.index] == ')' {
			break
		}

		keyword += string(lexer.File_content[lexer.index])

		lexer.index += 1
	}

	return keyword, lexer
}

func (lexer *Lexer) pickup_str() (string, *Lexer) {
	str_value := ""

	lexer.index += 1
	for {
		if lexer.File_content[lexer.index] == '"' {
			break
		}

		str_value += string(lexer.File_content[lexer.index])
		lexer.index += 1
	}

	return str_value, lexer
}

func (lexer *Lexer) Current_token() int {
	return lexer.Current_Token
}

func (lexer *Lexer) Lex() *Lexer {
	for {
		if lexer.index >= len(lexer.File_content) - 1 {
			break
		}
		if is_alpha(lexer.File_content[lexer.index]) {
			// Do Something to pickup keyword
			keyword, new_lex := lexer.pickup_keyword()
			lexer = new_lex

			switch keyword {
				case "print": return lexer.advance_with_token(K_PRINT, "print")
				case "let": return lexer.advance_with_token(K_LET, "let")
				case "require": {
					return lexer.advance_with_token(K_REQ, "require")
				}
				default: return lexer.advance_with_token(K_VAR_N, keyword)
			}

			return lexer.advance_with_token(K_VAR_N, keyword)
		}
		switch lexer.File_content[lexer.index] {
			case ' ': {
				for {
					if lexer.File_content[lexer.index] != ' ' {
						break
					}
					lexer.index += 1
				}
				continue
			}
			case '\n': {
				for {
					lexer.index += 1
					if lexer.File_content[lexer.index] != '\n' {
						break
					}
				}
				continue
			}
			case '"': {
				str_value, _ := lexer.pickup_str()
				//lexer = lex

				//lexer.index += 1

				return lexer.advance_with_token(T_STR, str_value)
			}
			case '.': return lexer.advance_with_token(T_DOT, ".")
			case ';': return lexer.advance_with_token(T_SEMI, ";")
			case '#': return lexer.advance_with_token(T_WRAP, "#")
			case '*': {
				for {
					lexer.index += 1
					if lexer.File_content[lexer.index] == '*' {
						lexer.index += 1
						break
					}
				}
			}
			case '{': return lexer.advance_with_token(T_LSB, "{")
			case '}': return lexer.advance_with_token(T_RSB, "}")
			case '[': return lexer.advance_with_token(T_LB, "[")
			case ']': return lexer.advance_with_token(T_RB, "]")
			case ')': return lexer.advance_with_token(T_LP, "(")
			case '(': return lexer.advance_with_token(T_RP, ")")
			default: {
				log.Fatal(fmt.Sprintf("[ERROR %d] -> Invalid Character: %c", InvalidToken, lexer.File_content[lexer.index]))
			}
		}
		lexer.index += 1

		if lexer.index == len(lexer.File_content) - 1 {
			break
		}

	}

	lexer.Current_Token = T_EOF

	return lexer
}
