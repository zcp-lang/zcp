package parser

import (
   "strconv"
   . "github.com/zcp-lang/zcp/ast"
   . "github.com/zcp-lang/zcp/lexer"
)

func (self *Parser) Null() Exp {

p := self.p

if t,_,l := self.GetTokenSkip();t==NULL {

return NullExp{l}

}

self.p = p

return nil

}

func (self *Parser) True() Exp {

p := self.p

if t,_,l := self.GetTokenSkip();t==TRUE {

return TrueExp{l}

}

self.p = p

return nil

}

func (self *Parser) False() Exp {

p := self.p

if t,_,l := self.GetTokenSkip();t==FALSE {

return FalseExp{l}

}

self.p = p

return nil

}

func (self *Parser) Name() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==IDENTIFIER {

return NameExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) String() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==STRING {

return StringExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Number() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==NUMBER || t==NNUMBER {

r,_:=strconv.Atoi(v)

return NumberExp{l,r}

}

self.p = p

return nil

}

func (self *Parser) Pointer() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==POINTER {

return PointerExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Quote() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==QUOTE {

return QuoteExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Comment() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==COMMENT {

return CommentExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Length() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==LENGTH {

return LengthExp{l,v}

}

self.p = p

return nil

}

func (self *Parser) Contrary() Exp {

p := self.p

if t,v,l := self.GetTokenSkip();t==CONTRARY {

return LogicalExp{l,"!","!",v}

}

self.p = p

return nil

}

func (self *Parser) Update() Exp {

p := self.p

left_t,left_v,left_l := self.GetTokenSkip();

right_t,right_v,right_l := self.GetTokenSkip();

if left_t == INCREMENT && right_t == IDENTIFIER {

return UpdateExp{right_l,"++",right_v,true}

}

if left_t == DECREMENT && right_t == IDENTIFIER {

return UpdateExp{right_l,"--",right_v,true}

}

if left_t == IDENTIFIER && right_t == INCREMENT {


return UpdateExp{left_l,"++",left_v,false}

}

if left_t == IDENTIFIER && right_t == DECREMENT {

return UpdateExp{left_l,"--",left_v,false}

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

func (self *Parser) ArrayItem() Exp {

p := self.p

exp1 := self.Exp()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,_,_ := self.GetTokenSkip()

if exp2 := self.Exp(); t == COLON && exp2 != nil {

return ArrayItem{exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1

}

func (self *Parser) ArrayExp() Exp {

p := self.p

_,_,line := self.GetTokenSkip()

self.p = p

result := make([]Exp,0)

exp1 := self.ArrayItem()

if exp1 == nil { self.p = p; return nil; }

result = append(result,exp1)

for {

p = self.p

t,_,l := self.GetTokenSkip()

line = l

if t != COMMA { self.p = p; return ArrayExp{l,result}; }

exp2 := self.ArrayItem()

if exp2 == nil { self.p = p; return ArrayExp{l,result}; }

result = append(result,exp2)

}

return ArrayExp{line,result}

}

func (self *Parser) Array() Exp {

p := self.p

t,_,l := self.GetTokenSkip();

if t != LEFT_BRACKET { self.p = p; return nil; }

p1 := self.p

exp := self.ArrayExp();

if exp == nil { self.p = p1; }

if t,_,_ := self.GetTokenSkip();t != RIGHT_BRACKET { self.p = p; return nil; }

if exp != nil { return exp; }else{ return ArrayExp{l,nil}; }

}

func (self *Parser) Exp0() Exp {

p := self.p

if exp:=self.Array();exp!=nil {

return exp

}

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

func (self *Parser) Exp1() Exp {

p := self.p

exp1 := self.Exp0()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,v,l := self.GetTokenSkip()

if exp2 := self.Exp1();(t == MULTIPLY || t == SLASH ||t == REMAINDER) && (exp2 != nil) {

return BinaryExp{l,v,exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1;

}

func (self *Parser) Exp2() Exp {

p := self.p

exp1 := self.Exp1()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,v,l := self.GetTokenSkip()

if exp2 := self.Exp2();(t == PLUS || t == MINUS) && (exp2 != nil) {

return BinaryExp{l,v,exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1;

}

func (self *Parser) Exp3() Exp {

p := self.p

exp1 := self.Exp2()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,v,l := self.GetTokenSkip()

if exp2 := self.Exp3(); t == CONCAT && exp2 != nil {

return BinaryExp{l,v,exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1;

}

func (self *Parser) Exp4() Exp {

p := self.p

exp1 := self.Exp3()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,v,l := self.GetTokenSkip()

if exp2 := self.Exp4();(t == LESS || t == GREATER || t == LESS_OR_EQUAL || t == GREATER_OR_EQUAL || t == NOT_EQUAL || t == EQUAL || t == STRICT_NOT_EQUAL || t == STRICT_EQUAL || t == STRING_START || t == STRING_END || t == QUESTION_MARK) && (exp2 != nil) {

return BinaryExp{l,v,exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1;

}

func (self *Parser) Exp5() Exp {

p := self.p

exp1 := self.Exp4()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,v,l := self.GetTokenSkip()

if exp2 := self.Exp5(); t == LOGICAL_AND && exp2 != nil {

return LogicalExp{l,v,exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1;

}

func (self *Parser) Exp6() Exp {

p := self.p

exp1 := self.Exp5()

if exp1 == nil { self.p = p; return nil; }

p = self.p

t,v,l := self.GetTokenSkip()

if exp2 := self.Exp6(); t == LOGICAL_OR && exp2 != nil {

return LogicalExp{l,v,exp1,exp2}

}else{ self.p = p; return exp1; }

self.p = p

return exp1;

}

func (self *Parser) Exp() Exp {

p := self.p

exp := self.Exp6()

if exp == nil { self.p = p; return nil; }

return exp

}