package main

import (

   . "fmt"
   . "github.com/zcp-lang/zcp/lexer"
   . "github.com/zcp-lang/zcp/parser"
   . "github.com/zcp-lang/zcp/compiler"
)

func main(){

str := `
a = 2 + 3 * 4;
`

l,_ := New_Lexer(str)

lex,_ := l.Run()

Println("词法分析结果:\n",lex)

Println("\n")

p := New_Parser(lex)

ast := p.Run()

Println("语法分析结果:\n",ast)

Println("中间指令集结果:\n")

compiler := New_Compiler(ast)

compiler.Run()

}