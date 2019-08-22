package vm

import (
   . "fmt"
     "strconv"
)

type Call struct{
   Vm *Vm
   Id int
   Pos int
   Name string
   A int
   P int
   Len int
   Asm [][]string
   Stack [][]string
   Msm map[string][]string
 }

func (self *Call) Next() bool {

if self.A >= len(self.Asm) { return false; }

return self.Analysis(self.Asm[self.A])

}

//分析指令
func (self *Call) Analysis(op []string) bool {

if self.P < 0 { self.P = 1; }

if op[0] == "PUSH" {

return self.Push("INT",op)

}else if op[0] == "PUSHS" {

return self.Push("STRING",op)

}else if op[0] == "PUSHF" {

return self.Push("BOOL",[]string{"","false"})

}else if op[0] == "PUSHT" {

return self.Push("BOOL",[]string{"","true"})

}else if op[0] == "PUSHN" {

return self.Push("NULL",op)

}else if op[0] == "ADD" {

return self.Add(op)

}else if op[0] == "SUB" {

return self.Sub(op)

}else if op[0] == "MUL" {

return self.Mul(op)

}else if op[0] == "DIV" {

return self.Div(op)

}else if op[0] == "YUC" {

return self.Yuc(op)

}else if op[0] == "STORE" {

return self.Store(op)

}else if op[0] == "LOAD" {

return self.Load(op)

}else if op[0] == "JB" {

return self.Jb(op)

}else if op[0] == "JG" {

return self.Jg(op)

}else if op[0] == "JBE" {

return self.Jbe(op)

}else if op[0] == "JGE" {

return self.Jge(op)

}else if op[0] == "JE" {

return self.Je(op)

}else if op[0] == "JNE" {

return self.Jne(op)

}else if op[0] == "JEO" {

return self.Jeo(op)

}else if op[0] == "JNEO" {

return self.Jneo(op)

}else if op[0] == "AND" {

return self.And(op)

}else if op[0] == "OR" {

return self.Or(op)

}else if op[0] == "JIF" {

return self.Jif(op)

}else if op[0] == "JMP" {

return self.Jmp(op)

}else if op[0] == "PRS" {

return self.Prs(op)

}else if op[0] == "RET" {

return self.Ret(op)

}else if op[0] == "END" {

return self.End(op)

}else if op[0] == "CALL" {

return self.Call(op)

}

self.Pa()

return false

}

//压Stack指令
func (self *Call) Push(types string,op []string) bool {

self.Set([]string{types,op[1]})

return true

}

//加法指令
func (self *Call) Add(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0] == "INT" {

a1,_ := strconv.Atoi(a[1])
b1,err := strconv.Atoi(b[1])

if err!=nil {

self.Set([]string{"STRING",a[1]+b[1]})

}else{

self.Set([]string{"INT",strconv.Itoa(a1+b1)})

}

}

if a[0] == "STRING" {

self.Set([]string{"STRING",a[1]+b[1]})

}

return true

}

//减法指令
func (self *Call) Sub(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0] == "INT" {

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

return self.Set([]string{"INT",strconv.Itoa(a1-b1)})

}

return self.Set([]string{"NULL","null"})

}

//乘法指令
func (self *Call) Mul(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0] == "INT" {

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

return self.Set([]string{"INT",strconv.Itoa(a1*b1)})

}

return self.Set([]string{"NULL","null"})

}

//除法指令
func (self *Call) Div(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0] == "INT" {

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

return self.Set([]string{"INT",strconv.Itoa(a1/b1)})

}

return self.Set([]string{"NULL","null"})

}

//求余指令
func (self *Call) Yuc(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0] == "INT" {

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

return self.Set([]string{"INT",strconv.Itoa(a1%b1)})

}

return self.Set([]string{"NULL","null"})

}

//向内存写入数据
func (self *Call) Store(op []string) bool {

var v []string

if self.P > 0 { v = self.Get(self.P-1); }else{ v = self.Stack[self.P]; }

self.Stack_set(self.P-2)

self.Msm[op[1]] = v

self.Pa()

return true

}

//从内存加载数据到Stack
func (self *Call) Load(op []string) bool {

result := self.Set(self.Msm[op[1]])

return result

}

// <
func (self *Call) Jb(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

if a1 < b1 { return self.Set([]string{"BOOL","true"}); }

return self.Set([]string{"BOOL","false"})

}

// >
func (self *Call) Jg(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

if a1 > b1 { return self.Set([]string{"BOOL","true"}); }

return self.Set([]string{"BOOL","false"})

}

// <=
func (self *Call) Jbe(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

if a1 <= b1 { return self.Set([]string{"BOOL","true"}); }

return self.Set([]string{"BOOL","false"})

}

// >=
func (self *Call) Jge(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

if a1 >= b1 { return self.Set([]string{"BOOL","true"}); }

return self.Set([]string{"BOOL","false"})

}

// ==
func (self *Call) Je(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

if a1 == b1 { return self.Set([]string{"BOOL","true"}); }

return self.Set([]string{"BOOL","false"})

}

// !=
func (self *Call) Jne(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

a1,_ := strconv.Atoi(a[1])
b1,_ := strconv.Atoi(b[1])

if a1 != b1 { return self.Set([]string{"BOOL","true"}); }

return self.Set([]string{"BOOL","false"})

}

// ===
func (self *Call) Jeo(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0]!=b[0] || a[1]!=b[1] { return self.Set([]string{"BOOL","false"}); }

return self.Set([]string{"BOOL","true"})

}

// !==
func (self *Call) Jneo(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if a[0]==b[0] { return self.Set([]string{"BOOL","false"}); }

return self.Set([]string{"BOOL","true"})

}

// and(&&)
func (self *Call) And(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if (a[1]!="null"&&a[1]!="false") && (b[1]!="null"&&b[1]!="false") {

return self.Set([]string{"BOOL","true"});

}

return self.Set([]string{"BOOL","false"})

}

// or(||)
func (self *Call) Or(op []string) bool {

a := self.Get(self.P-2)
b := self.Get(self.P-1)

if a==nil || b==nil { return false; }

self.Stack_set(self.P-2)

if (a[1]!="null"&&a[1]!="false") || (b[1]!="null"&&b[1]!="false") {

return self.Set([]string{"BOOL","true"});

}

return self.Set([]string{"BOOL","false"})

}

// if
func (self *Call) Jif(op []string) bool {

exp := self.Get(self.P-1)

self.Stack_set(self.P-1)

if exp[1] == "true" {

self.A++

}

return true

}

//跳转指令
func (self *Call) Jmp(op []string) bool {

i,err := strconv.Atoi(op[1])

if err!=nil { return false; }

i = i - self.Pos

self.A = i - 1

return true

}

//形参加载到内存
func (self *Call) Prs(op []string) bool {

var v []string

if self.P > 0 { v = self.Get(self.P-1); }else{ v = self.Stack[self.P]; }

self.Stack_set(self.P-2)

if v==nil { v = []string{"NULL","null"}; }

self.Msm[op[1]] = v

self.Pa()

return true

}

//函数返回
func (self *Call) Ret(op []string) bool {

self.A = len(self.Asm)

return true

}

//函数结束
func (self *Call) End(op []string) bool {

self.A = len(self.Asm)

return true

}

//函数调用
func (self *Call) Call(op []string) bool {

fn := self.Vm.Fn[op[1]]

call := &Call{self.Vm,self.Id,fn.Pos,fn.Name,1,0,6,fn.Asm,make([][]string,6),make(map[string][]string)}

Println(call)

/*
for j:=0;j<len(parame[i]);j++ {

call.Set(parame[i][j])

}
*/

self.Vm.Call[self.Id][len(self.Vm.Call[self.Id])] = call

return true

}

//向Stack写入数据
func (self *Call) Set(op []string) bool {

if len(self.Stack) <= self.Len { self.Stack = append(self.Stack,[]string{}); }

if len(self.Stack) <= self.P {

for i:=0;i<(self.P+3)-len(self.Stack);i++ {

self.Stack = append(self.Stack,[]string{})

}

}

if self.P < 0 { self.P = 0; }

self.Stack[self.P] = op

self.P++

self.Pa()

return true

}

//获取指定Stack位置数据
func (self *Call) Get(p int) []string {

if p >= 0 && len(self.Stack) >= p { return self.Stack[p]; }

self.P = 1

return nil

}

//设置Stack指针位置
func (self *Call) Stack_set(p int) bool {

if p < 0 { self.P = 0; }else{ self.P = p; }

if len(self.Stack) <= self.P {

for i:=0;i<(self.P+3)-len(self.Stack);i++ {

self.Stack = append(self.Stack,[]string{})

}

}

self.Pa()

return true

}

//清空Stack
func (self *Call) Pa() {

if self.P < 0 { self.P = 0; }

for i:=self.P;i<len(self.Stack);i++ {

self.Stack[i] = []string{}

}

}