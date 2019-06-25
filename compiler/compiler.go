package compiler

import (
   //. "fmt"
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

func (self *Compiler) Run() [][]string {

self.Init(self.ast)

return self.bytecode

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

}else if result,ok:=op.(NullExp);ok == true {

return self.NullExp(result)

}else if result,ok:=op.(TrueExp);ok == true {

return self.TrueExp(result)

}else if result,ok:=op.(FalseExp);ok == true {

return self.FalseExp(result)

}else if result,ok:=op.(StringExp);ok == true {

return self.StringExp(result)

}else if result,ok:=op.(NameExp);ok == true {

return self.NameExp(result)

}else if result,ok:=op.(LengthExp);ok == true {

return self.LengthExp(result)

}else if result,ok:=op.(ReturnStat);ok == true {

return self.ReturnStat(result)

}else if result,ok:=op.(LogicalExp);ok == true {

return self.LogicalExp(result)

}else if result,ok:=op.(IfStat);ok == true {

return self.IfStat(result)

}else if result,ok:=op.(WhileStat);ok == true {

return self.WhileStat(result)

}else if result,ok:=op.(ForStat);ok == true {

return self.ForStat(result)

}else if result,ok:=op.(FunctionStat);ok == true {

return self.FunctionStat(result)

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

}else if op.Op == "-" {

if result := self.Set([]string{"SUB"});result == false { return false; }

}else if op.Op == "*" {

if result := self.Set([]string{"MUL"});result == false { return false; }

}else if op.Op == "/" {

if result := self.Set([]string{"DIV"});result == false { return false; }

}else if op.Op == "%" {

if result := self.Set([]string{"YUC"});result == false { return false; }

}else if op.Op == "<" {

if result := self.Set([]string{"JB"});result == false { return false; }

}else if op.Op == ">" {

if result := self.Set([]string{"JG"});result == false { return false; }

}else if op.Op == "<=" {

if result := self.Set([]string{"JBE"});result == false { return false; }

}else if op.Op == ">=" {

if result := self.Set([]string{"JGE"});result == false { return false; }

}else if op.Op == "==" {

if result := self.Set([]string{"JE"});result == false { return false; }

}else if op.Op == "!=" {

if result := self.Set([]string{"JNE"});result == false { return false; }

}else if op.Op == "===" {

if result := self.Set([]string{"JEO"});result == false { return false; }

}else if op.Op == "!==" {

if result := self.Set([]string{"JNEO"});result == false { return false; }

}else if op.Op == "?*" {

if result := self.Set([]string{"STRING_START"});result == false { return false; }

}else if op.Op == "*?" {

if result := self.Set([]string{"STRING_END"});result == false { return false; }

}else if op.Op == "?" {

if result := self.Set([]string{"QUESTION_MARK"});result == false { return false; }

}

return true

}

func (self *Compiler) LogicalExp(op LogicalExp) bool {

if op.Op == "!" {

if result := self.Set([]string{"LOAD",op.Right.(string)});result == false { return false; }

if result := self.Set([]string{"NOT"});result == false { return false; }

}else{

if left := self.Init(op.Left);left == false { return false; }

if right := self.Init(op.Right);right == false { return false; }

if op.Op == "&&" {

if result := self.Set([]string{"AND"});result == false { return false; }

}else if op.Op == "||" {

if result := self.Set([]string{"OR"});result == false { return false; }

}

}

return true

}

func (self *Compiler) IfStat(op IfStat) bool {

if op.Test == nil { return true; }

if test := self.Init(op.Test);test == false { return false; }

if result := self.Set([]string{"JIF"});result == false { return false; }

jp := len(self.bytecode)

if result := self.Set([]string{"JMP","0"});result == false { return false; }

self.Init(op.Consequent);

cendp := len(self.bytecode)

if result := self.Set([]string{"JMP","0"});result == false { return false; }

jpp := len(self.bytecode)

self.Init(op.Alternate);

aendp := len(self.bytecode)


jpps := strconv.Itoa(jpp)

self.bytecode[jp] = []string{"JMP",jpps}

aendps := strconv.Itoa(aendp)

self.bytecode[cendp] = []string{"JMP",aendps}

return true

}

func (self *Compiler) WhileStat(op WhileStat) bool {

p := len(self.bytecode)

if op.Test == nil { return true; }

if test := self.Init(op.Test);test == false { return false; }

if result := self.Set([]string{"JIF"});result == false { return false; }

jp := len(self.bytecode)

if result := self.Set([]string{"JMP","0"});result == false { return false; }

self.Init(op.Block)

self.Set([]string{"JMP",strconv.Itoa(p)})

self.bytecode[jp] = []string{"JMP",strconv.Itoa(len(self.bytecode))}

return true

}

func (self *Compiler) ForStat(op ForStat) bool {

self.Init(op.InitExp)

p := len(self.bytecode)

self.Init(op.LimitExp)

if result := self.Set([]string{"JIF"});result == false { return false; }

jp := len(self.bytecode)

if result := self.Set([]string{"JMP","0"});result == false { return false; }

self.Init(op.Block)

self.Init(op.StepExp)

if result := self.Set([]string{"JMP",strconv.Itoa(p)});result == false { return false; }

self.bytecode[jp] = []string{"JMP",strconv.Itoa(len(self.bytecode))}

return true

}

func (self *Compiler) FunctionStat(op FunctionStat) bool {

pa,_ := op.Id.(NameExp)

params,_ := op.Params.(Exps)

self.Set([]string{"FUNCTION",pa.Name,strconv.Itoa(len(params.Exp)),strconv.Itoa(len(self.bytecode))})

for _,v := range params.Exp {

t,_ := v.(NameExp)

self.Set([]string{"PARAMS",t.Name})

}

self.Init(op.Body)

return true

}

func (self *Compiler) NullExp(op NullExp) bool {

return self.Set([]string{"PUSHN","null"})

}

func (self *Compiler) TrueExp(op TrueExp) bool {

return self.Set([]string{"PUSHT","true"})

}

func (self *Compiler) FalseExp(op FalseExp) bool {

return self.Set([]string{"PUSHF","false"})

}

func (self *Compiler) NumberExp(op NumberExp) bool {

number := strconv.Itoa(op.Val)

return self.Set([]string{"PUSH",number})

}

func (self *Compiler) StringExp(op StringExp) bool {

return self.Set([]string{"PUSHS",op.Str})

}

func (self *Compiler) NameExp(op NameExp) bool {

return self.Set([]string{"LOAD",op.Name})

}

func (self *Compiler) LengthExp(op LengthExp) bool {

if result := self.Set([]string{"LOAD",op.Name});result == false { return false; }

if result := self.Set([]string{"LENGTH"});result == false { return false; }

return true

}

func (self *Compiler) ReturnStat(op ReturnStat) bool {

r,_ := op.Exp.(Exps)

var i int

for i=0;i<len(r.Exp);i++ {

exp,_ := r.Exp[i].(Exp);

self.Init(exp)

}

return self.Set([]string{"RET",strconv.Itoa(len(r.Exp))})

}

func (self *Compiler) Set(op []string) bool {

self.bytecode = append(self.bytecode,op)

return true

}