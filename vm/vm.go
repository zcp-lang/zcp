package vm

import (
   . "fmt"
    // "time"
     "strconv"
   )

type Vm struct{
   P string
   Fn map[string]*Fn
   Co map[int]*Call
   Call map[int]map[int]*Call
   Msm map[string]map[string][]string
}

func New_Vm() *Vm {

return &Vm{"",make(map[string]*Fn),make(map[int]*Call),make(map[int]map[int]*Call),make(map[string]map[string][]string)}

}

func (self *Vm) Run(asm [][]string) bool {

self.Init(asm)

return self.Start()

}

func (self *Vm) Init(asm [][]string) bool {

var ret,pos,pos2 int
var main,fn [][]string

for i:=0;i<len(asm);i++ {

if asm[i][0]=="FN" {

pos = i

for i=i;asm[i][0]!="END";i++ {

if asm[i][0]=="RET" {

ret,_ = strconv.Atoi(asm[i][1])

}

fn = append(fn,asm[i])

}

fn = append(fn,asm[i])

params,_ := strconv.Atoi(fn[0][2])

start,_ := strconv.Atoi(fn[0][3])

end,_ := strconv.Atoi(fn[len(fn)-1][1])

self.Fn[fn[0][1]] = &Fn{fn[0][1],pos,params,ret,start,end,fn}

fn = [][]string{}

}else{

if main==nil { pos2 = i; }

main = append(main,asm[i])

}

}

self.Fn["main"] = &Fn{"main",pos2,0,1,0,0,main}

return true

}

func (self *Vm) Start() bool {

fn := self.Fn["main"]

self.Co[len(self.Co)] = &Call{self,len(self.Co),fn.Pos,fn.Name,0,0,6,fn.Asm,make([][]string,6),make(map[string][]string)}

parame := [][][]string{[][]string{},[][]string{[]string{"INT","3"}},[][]string{[]string{"INT","6"}}}

self.Coroutine([]string{"aaa","bbb","ccc"},parame)

self.Corun()

return true

}

func (self *Vm) Coroutine(name []string,parame [][][]string) bool {

for i:=0;i<len(name);i++ {

fn := self.Fn[name[i]]

call := &Call{self,len(self.Co),fn.Pos,fn.Name,1,0,6,fn.Asm,make([][]string,6),make(map[string][]string)}

for j:=0;j<len(parame[i]);j++ {

call.Set(parame[i][j])

}

self.Co[len(self.Co)] = call

}

return true

}

func (self *Vm) Corun() bool {

var i int

for {

if self.Code()==false { break;}

if i >= len(self.Co) { i = 0; }

if self.Co[i].A < len(self.Co[i].Asm){

self.Co[i].Next()

self.Co[i].A++

}

i++

}

Println(self.Co[0].Stack)
Println(self.Co[0].Msm)

return true

}

func (self *Vm) Code() bool {

for i:=0;i<len(self.Co);i++ {

if self.Co[i].A < len(self.Co[i].Asm) { return true; }

}

return false

}

/*
func (self *Vm) Call() ([][]string,bool) {


}
*/