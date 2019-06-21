package parser

import (
   //"strconv"
   . "github.com/zcp-lang/zcp/ast"
   . "github.com/zcp-lang/zcp/lexer"
)

func (self *Parser) Id() Stat {

p := self.p

name := self.Name()

if name == nil { self.p = p; return nil; }

result := make([]Exp,0)

result = append(result,name)

for {

p = self.p

if t,_,_ := self.GetTokenSkip(); t != COMMA { self.p = p; return Exps{result}; }

name2 := self.Name()

if name2 == nil { self.p = p; return Exps{result}; }

result = append(result,name2)

}

return Exps{result}

}

func (self *Parser) Assign() Stat {

p := self.p

id := self.Id()

if id == nil { self.p = p; return nil; }

t,_,l := self.GetTokenSkip();

if t != ASSIGN { self.p = p; return nil; }

exp := self.Exp()

if exp == nil { self.p = p; return nil; }

return AssignStat{l,id,exp}

}

func (self *Parser) If() Stat {

p := self.p

if t,_,_ := self.GetTokenSkip(); t != IF { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_PARENTHESIS { self.p = p; return nil; }

pa := self.p

exp := self.Exp6()

if exp == nil { self.p = pa; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_PARENTHESIS { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_BRACE { self.p = p; return nil; }

p1 := self.p

consequent := self.Block()

if consequent == nil { self.p = p1; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_BRACE { self.p = p; return nil; }

p2 := self.p

if t,_,_ := self.GetTokenSkip(); t != ELSE { self.p = p2; return IfStat{exp,consequent,nil}; }

alternate := self.If()

if alternate != nil { return IfStat{exp,consequent,alternate}; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_BRACE { self.p = p2; return IfStat{exp,consequent,nil}; }

p3 := self.p

alternate = self.Block()

if alternate == nil { self.p = p3; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_BRACE { self.p = p2; return IfStat{exp,consequent,nil}; }

return IfStat{exp,consequent,alternate};

}

func (self *Parser) While() Stat {

p := self.p

if t,_,_ := self.GetTokenSkip(); t != WHILE { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_PARENTHESIS { self.p = p; return nil; }

p1 := self.p

exp := self.Exp6()

if exp == nil { self.p = p1; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_PARENTHESIS { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_BRACE { self.p = p; return nil; }

p2 := self.p

body := self.Block()

if body == nil { self.p = p2; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_BRACE { self.p = p; return nil; }

return WhileStat{exp,body};

}

func (self *Parser) For() Stat {

p := self.p

t,_,l := self.GetTokenSkip();

if t != FOR { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_PARENTHESIS { self.p = p; return nil; }

p1 := self.p

exp1 := self.Assign()

if exp1 == nil { self.p = p1; }

if t,_,_ := self.GetTokenSkip(); t != SEMICOLON { self.p = p; return nil; }

p2 := self.p

exp2 := self.Exp6()

if exp2 == nil { self.p = p2; }

if t,_,_ := self.GetTokenSkip(); t != SEMICOLON { self.p = p; return nil; }

p3 := self.p

exp3 := self.Assign()

if exp3 == nil {

self.p = p3;

exp3 = self.Exp6()

if exp3 == nil { self.p = p3; }

}

if t,_,_ := self.GetTokenSkip(); t != RIGHT_PARENTHESIS { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_BRACE { self.p = p; return nil; }

p4 := self.p

body := self.Block()

if body == nil { self.p = p4; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_BRACE { self.p = p; return nil; }

return ForStat{l,exp1,exp2,exp3,body};

}

func (self *Parser) Function() Stat {

p := self.p

if t,_,_ := self.GetTokenSkip(); t != FUNCTION { self.p = p; return nil; }

id := self.Name()

if id == nil { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_PARENTHESIS { self.p = p; return nil; }

p1 := self.p

params := self.Id()

if params == nil { self.p = p1; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_PARENTHESIS { self.p = p; return nil; }

if t,_,_ := self.GetTokenSkip(); t != LEFT_BRACE { self.p = p; return nil; }

p2 := self.p

body := self.Block()

if body == nil { self.p = p2; }

if t,_,_ := self.GetTokenSkip(); t != RIGHT_BRACE { self.p = p; return nil; }

return FunctionStat{id,params,body};

}

func (self *Parser) Return() Stat {

p := self.p

t,_,l := self.GetTokenSkip();

if t != RETURN { self.p = p; return nil; }

p1 := self.p

if t,_,_:=self.GetToken(); t != SPACE { self.p = p1; return ReturnStat{l,nil}; }

exp := self.Exp()

if exp == nil { self.p = p1; return ReturnStat{l,nil}; }

return ReturnStat{l,exp}

}

func (self *Parser) Break() Stat {

p := self.p

t,_,l := self.GetTokenSkip();

if t != BREAK { self.p = p; return nil; }

return BreakStat{l};

}