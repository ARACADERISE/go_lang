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
	T_LB	= 2
	T_RB	= 3
	T_SEMI	= 4
	K_LET	= 5
	K_REQ	= 6
	K_PRINT = 7
)

// Errors
const (
	NoDir		= 128
	ReadErr		= 134
	NoFile		= 138
	InvalidToken	= 140
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
		if lexer.File_content[lexer.index] == ' ' {
			break
		}

		keyword += string(lexer.File_content[lexer.index])

		lexer.index += 1
	}

	return keyword, lexer
}

func (lexer *Lexer) Lex() *Lexer {
	for {
		if lexer.index == len(lexer.File_content) - 1 {
			break
		}
		if is_alpha(lexer.File_content[lexer.index]) {
			// Do Something to pickup keyword
			keyword, new_lex := lexer.pickup_keyword()
			lexer = new_lex

			switch keyword {
				case "let": {
					fmt.Println("LET KEYWORD")
					return lexer.advance_with_token(K_LET, "let")
				}
				case "require": {
					fmt.Println("REQUIRE KEYWORD")
					return lexer.advance_with_token(K_REQ, "require")
				}
			}
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
			case '"': {
				return lexer
			}
			case '*': {
				for {
					lexer.index += 1

					if lexer.File_content[lexer.index] == '*' {
						break
					}
				}
			}
			case '{': {
				// Do Something
			}
			default: {
				log.Fatal(fmt.Sprintf("[ERROR %d] -> Invalid Character: %c", InvalidToken, lexer.File_content[lexer.index]))
			}
		}
		lexer.index += 1

		if lexer.index == len(lexer.File_content) - 1 {
			break
		}

	}

	return lexer
}
