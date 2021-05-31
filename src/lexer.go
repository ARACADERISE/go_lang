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
	K_PRINT = 6
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

func advance_token(lexer *Lexer, token_type int) *Lexer {
	lexer.index += 1
	lexer.Current_Token = token_type

	return lexer
}

func is_alpha(val byte) bool {
	if(val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') {
		return true
	}
	return false
}

func (lexer *Lexer) Lex() *Lexer {
	for {
		if is_alpha(lexer.File_content[lexer.index]) {
			// Do Something to pickup keyword
			fmt.Println("HERE")
		}
		switch lexer.File_content[lexer.index] {
			case '"': {
				// Do Something To Get String
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

	}

	return lexer
}
