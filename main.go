package main

import (

   . "fmt"
   "github.com/zcp-lang/zcp/lexer"
   "github.com/zcp-lang/zcp/parser"
   //"github.com/zcp-lang/zcp/compiler"
)

func main(){

str := `(2+3)*4+(5-2)*9`

l,_ := lexer.New(str)

result,err := l.Run()

Println("词法分析结果:\n",result)

Println(err)

Println("\n")

Println("语法分析结果:\n")

p := parser.New(result)

p.Run()
/*
Println("\n")

asm := `
PUSH 2
PUSH 3
PLUS
CALL print`
compiler := compiler.New()

Println(compiler.StringDecode(asm))
*/
}