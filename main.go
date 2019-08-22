package main

import (

   . "fmt"
     "os"
     //"strconv"
     "io/ioutil"
   . "github.com/zcp-lang/zcp/lexer"
   . "github.com/zcp-lang/zcp/parser"
   . "github.com/zcp-lang/zcp/compiler"
   . "github.com/zcp-lang/zcp/vm"
)

func Get(file string) (string,bool) {

f, err := os.OpenFile(file,os.O_RDONLY,0600)

defer f.Close()

if err != nil {

return "",false

}else{

data,errs := ioutil.ReadAll(f)

if errs != nil { return "",false; }

return string(data),true

}

return "",false

}

func Put(file,data string) bool {

f, err := os.OpenFile(file,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0600)

defer f.Close()

if err != nil {

return false

}else{

_,errs := f.Write([]byte(data))

if errs != nil { return false; }

return true

}

}

func main(){

if len(os.Args) > 1{

if os.Args[1] == "bug" {

str,_ := Get(os.Args[2])

l,_ := New_Lexer(str)
lex,_ := l.Run()
p := New_Parser(lex)
ast := p.Run()
compiler := New_Compiler(ast)
asm := compiler.Run()

Put("/sdcard/go/app/zcp/main.asm",compiler.StringEncode(asm))

Println(compiler.StringEncode(asm),"\n")

Println(ast,"\n")

vm := New_Vm()

vm.Run(asm)

}else{

str,_ := Get(os.Args[1])

l,_ := New_Lexer(str)
lex,_ := l.Run()
p := New_Parser(lex)
ast := p.Run()
compiler := New_Compiler(ast)
asm := compiler.Run()

if len(os.Args)>2 && os.Args[2] != "" {

Put(os.Args[2],compiler.StringEncode(asm))

}

vm := New_Vm()

vm.Run(asm)

}

}else{

l,_ := New_Lexer(os.Args[0])
lex,_ := l.Run()
p := New_Parser(lex)
ast := p.Run()
compiler := New_Compiler(ast)
asm := compiler.Run()

Println(asm)

}

}