package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> " // prompt for the REPL

func Start(in io.Reader, out io.Writer) {
	// 开始 REPL
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT) // 打印 >>
		scanned := scanner.Scan()
		if !scanned {
			return // 退出 REPL
		}

		line := scanner.Text()
		l := lexer.New(line)                                                   // 创建lexer
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // 循环解析所有的token
			fmt.Printf("%+v\n", tok) // print the token
		}
	}
}
