package parser

import (
   //"strconv"
   . "github.com/zcp-lang/zcp/ast"
   . "github.com/zcp-lang/zcp/lexer"
)

func (self *Parser) Stat() Stat {

p := self.p

var result Stat

if ifs:=self.If();ifs != nil {

result = ifs

} else if assign:=self.Assign();assign != nil {

result = assign

} else if fors:=self.For();fors != nil {

result = fors

} else if while:=self.While();while != nil {

result = while

} else if function:=self.Function();function != nil {

result = function

} else if returns:=self.Return();returns != nil {

result = returns

} else if breaks:=self.Break();breaks != nil {

result = breaks

} else if exp:=self.Exp();exp != nil {

result = exp

}else{

self.p = p; return nil;

}

p1 := self.p

if t,_,_ := self.GetTokens(); t == LINE || t == SEMICOLON { self.p = p1; self.p++; }else{ self.p = p1; }

return result

}

func (self *Parser) Block() Stat {

p := self.p

stat := self.Stat()

if stat == nil { self.p = p; return nil; }

result := make([]Stat,0)

result = append(result,stat)

for {

p = self.p

stat2 := self.Stat()

if stat2 == nil { self.p = p; return result; }

result = append(result,stat2)

}

return Blocks{result}

}