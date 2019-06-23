package compiler

import (
   . "fmt"
   //"reflect"
   "strconv"
   //. "github.com/zcp-lang/zcp/lexer"
   . "github.com/zcp-lang/zcp/ast"
)

type Compiler struct{
   p   int
   ast Stat
   bytecode [][]string
   Err string
}

func New_Compiler(block Stat) *Compiler {

return &Compiler{0,block,make([][]string,0),""}

}

func (self *Compiler) Run() {

self.Init(self.ast)

Println(self.bytecode)

}

func (self *Compiler) Init(op interface{}) bool {

//t := self.Type(op)

if result,ok:=op.([]Stat);ok == true {

return self.Stats(result)

}else if result,ok:=op.(AssignStat);ok == true {

return self.AssignStat(result)

}else if result,ok:=op.(BinaryExp);ok == true {

return self.BinaryExp(result)

}else if result,ok:=op.(NumberExp);ok == true {

return self.NumberExp(result)

}

return false

}

func (self *Compiler) Stats(op []Stat) bool {

for _,op := range op {

self.Init(op)

}

return true

}

func (self *Compiler) AssignStat(op AssignStat) bool {

l,_ := op.Left.(Exps)

r,_ := op.Right.(Exps)

var i int

for i=0;i<len(l.Exp);i++ {

la,_ := l.Exp[i].(NameExp);v := la.Name
exp,_ := r.Exp[i].(Exp);

self.Init(exp)

self.Set([]string{"STORE",v})

}

return true

}

func (self *Compiler) BinaryExp(op BinaryExp) bool {

if left := self.Init(op.Left);left == false { return false; }

if right := self.Init(op.Right);right == false { return false; }

if op.Op == "+" {

if result := self.Set([]string{"ADD"});result == false { return false; }

}else if op.Op == "*" {

if result := self.Set([]string{"MUL"});result == false { return false; }

}

return true

}

func (self *Compiler) NumberExp(op NumberExp) bool {

number := strconv.Itoa(op.Val)

return self.Set([]string{"PUSH",number})

}

func (self *Compiler) Set(op []string) bool {

self.bytecode = append(self.bytecode,op)

return true

}