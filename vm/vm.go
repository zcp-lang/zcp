package vm

import (
   . "fmt"
     "strconv"
   )

type Vm struct{
   P int
   F string
   Asm [][]string
   Stack [][]string
   Msm map[string]map[string][]string
}

func New_Vm(asm [][]string) *Vm {

return &Vm{0,"main",asm,make([][]string,6),nil}

}

func (self *Vm) Run() {

for _,v := range self.Asm {

self.Init(v)

}

Println(self.Stack)

}

func (self *Vm) Init(op []string) bool {

if op[0] == "PUSH" {

return self.Push(op)

}else if op[0] == "ADD" {

return self.Add(op)

}else if op[0] == "SUB" {

return self.Sub(op)

}else if op[0] == "MUL" {

return self.Mul(op)

}else if op[0] == "DIV" {

return self.Div(op)

}else if op[0] == "STORE" {

return self.Store(op)

}

return false

}

func (self *Vm) Push(op []string) bool {

self.Set(self.P,[]string{op[1],"int"})

self.P++

return true

}

func (self *Vm) Add(op []string) bool {

var a,b int

if as:=self.Get(self.P-2);as!=nil { a,_ = strconv.Atoi(as[0]); }

if bs:=self.Get(self.P-1);bs!=nil { b,_ = strconv.Atoi(bs[0]); }

self.P -= 2

self.Set(self.P,[]string{strconv.Itoa(a+b),"int"})

return true

}

func (self *Vm) Sub(op []string) bool {

var a,b int

if as:=self.Get(self.P-2);as!=nil { a,_ = strconv.Atoi(as[0]); }

if bs:=self.Get(self.P-1);bs!=nil { b,_ = strconv.Atoi(bs[0]); }

self.P -= 2

self.Set(self.P,[]string{strconv.Itoa(a-b),"int"})

return true

}

func (self *Vm) Mul(op []string) bool {

var a,b int

if as:=self.Get(self.P-2);as!=nil { a,_ = strconv.Atoi(as[0]); }

if bs:=self.Get(self.P-1);bs!=nil { b,_ = strconv.Atoi(bs[0]); }

self.P -= 2

self.Set(self.P,[]string{strconv.Itoa(a*b),"int"})

return true

}

func (self *Vm) Div(op []string) bool {

var a,b int

if as:=self.Get(self.P-2);as!=nil { a,_ = strconv.Atoi(as[0]); }

if bs:=self.Get(self.P-1);bs!=nil { b,_ = strconv.Atoi(bs[0]); }

self.P -= 2

self.Set(self.P,[]string{strconv.Itoa(a/b),"int"})

return true

}

func (self *Vm) Store(op []string) bool {

Println(self.Stack[self.P])

return true

}

func (self *Vm) Set(p int,op []string) bool {

if len(self.Stack) <= p {

size := p - len(self.Stack)

var i int

for i=0;i<size+1;i++ {

self.App()

}

}

self.Stack[p] = op

return true

}

func (self *Vm) App() {

self.Stack = append(self.Stack,[]string{})

}

func (self *Vm) Pa(p int) {

var i int

for i=p;i<len(self.Stack);i++ {

self.Stack[i] = []string{}

}

}

func (self *Vm) Get(p int) []string {

if len(self.Stack) > p {

return self.Stack[p]

}

return nil

}