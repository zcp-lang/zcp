package main

import (

   . "fmt"
   . "github.com/zcp-lang/zcp/lexer"
   . "github.com/zcp-lang/zcp/parser"
   //"github.com/zcp-lang/zcp/compiler"
)

func main(){

str := `
function add(a,b){

return a + b;

}

a,b,c = 2333,"moid",true;

result = add(a,666);

print(result);
`

l,_ := New_Lexer(str)

result,err := l.Run()

Println("词法分析结果:\n",result)

Println(err)

Println("\n")

Println("语法分析结果:\n")

p := New_Parser(result)

p.Run()
/*
Println("\n")

asm := `
PUSH 2
PUSH 3
PLUS
CALL print`
compiler := compiler.New_Compiler()

Println(compiler.StringDecode(asm))
*/
}