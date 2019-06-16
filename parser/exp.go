package parser

import (
   "strconv"
   . "github.com/zcp-lang/zcp/ast"
   . "github.com/zcp-lang/zcp/lexer"
)

func (self *Parser) Null() Exp {

p := self.p

if t,_,l := self.GetTokenSkip();t==NULL {

return &NullExp{l}

}

self.p = p

return nil

}

func (self *Parser) True() Exp {

p := self.p

if t,_,l := self.GetTokenSkip();t==TRUE {

return &TrueExp{l}

}

self.p = p

return nil

}

func (self *Parser) False() Exp {

p := self.p

if t,_,l := self.GetTokenSkip();t==FALSE {

return &FalseExp{l}

}

self.p = p

return nil

}

func (self *Parser) Name() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==IDENTIFIER {

return &NameExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) String() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==STRING {

return &StringExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Number() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==NUMBER || t==NNUMBER {

r,_:=strconv.Atoi(v)

return &NumberExp{l,r}

}

self.p = p

return nil

}

func (self *Parser) Pointer() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==POINTER {

return &PointerExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Quote() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==QUOTE {

return &QuoteExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Comment() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==COMMENT {

return &CommentExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Length() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==LENGTH {

return &LengthExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Contrary() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==CONTRARY {

return &LogicalExp{l,"!","!",v}

}

self.p = p

return nil

}

func (self *Parser) Update() Exp {

p := self.p

left_t,left_v,left_l := self.GetTokenSkip();

right_t,right_v,right_l := self.GetTokenSkip();

if left_t == INCREMENT && right_t == IDENTIFIER {

return &UpdateExp{right_l,"++",right_v,true}

}

if left_t == DECREMENT && right_t == IDENTIFIER {

return &UpdateExp{right_l,"--",right_v,true}

}

if left_t == IDENTIFIER && right_t == INCREMENT {


return &UpdateExp{left_l,"++",left_v,false}

}

if left_t == IDENTIFIER && right_t == DECREMENT {

return &UpdateExp{left_l,"--",left_v,false}

}

self.p = p

return nil

}
/*
func (self *Parser) Staid() Exp {

p := self.p

left,_,_ := self.GetTokenSkip();

if left != LEFT_PARENTHESIS { self.p = p; return nil; }

exp := self.Exp();

if exp == nil { self.p = p; return nil; }

right,_,_ := self.GetTokenSkip();

if right != RIGHT_PARENTHESIS { self.p = p; return nil; }

return exp;

}
*/

func (self *Parser) Exp0() Exp {

p := self.p

if exp:=self.Update();exp!=nil {

return exp

}

if exp:=self.Null();exp!=nil {

return exp

}

if exp:=self.True();exp!=nil {

return exp

}

if exp:=self.False();exp!=nil {

return exp

}

if exp:=self.Name();exp!=nil {

return exp

}

if exp:=self.String();exp!=nil {

return exp

}

if exp:=self.Number();exp!=nil {

return exp

}

if exp:=self.Pointer();exp!=nil {

return exp

}

if exp:=self.Quote();exp!=nil {

return exp

}

if exp:=self.Comment();exp!=nil {

return exp

}

if exp:=self.Length();exp!=nil {

return exp

}

if exp:=self.Contrary();exp!=nil {

return exp

}

self.p = p

return nil

}