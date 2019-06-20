package parser

import (
   "fmt"
   "reflect"
   . "github.com/zcp-lang/zcp/lexer"
   //. "github.com/zcp-lang/zcp/ast"
)

type Parser struct{
   p int
   Err string
   token map[int]Tkn
}


func typeof(v interface{}) string {

return reflect.TypeOf(v).String()

}


func New(token map[int]Tkn) *Parser {

return &Parser{p:0,token:token}

}


func (self *Parser) Run() {

exp := self.ArrayCall()

fmt.Println(exp)

}


func (self *Parser) GetToken() (Token,string,int) {

token := self.token[self.p].Token
value := self.token[self.p].Value
line := self.token[self.p].Line

self.p++

return token,value,line

}


func (self *Parser) GetTokenSkip() (Token,string,int) {

token := self.token[self.p].Token
value := self.token[self.p].Value
line := self.token[self.p].Line

if SPACE == token || LINE == token {

self.p++

token = self.token[self.p].Token
value = self.token[self.p].Value
line = self.token[self.p].Line

}


self.p++

return token,value,line

}