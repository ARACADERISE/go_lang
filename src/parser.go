package src

type Parser struct {
	lexer	interface{}
}

func Init_parser(lexer interface{}) *Parser {
	return &Parser{ lexer: lexer }
}
