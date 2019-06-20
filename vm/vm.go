package vm

import (
   . "github.com/zcp-lang/zcp/compiler"
   )

type ZcpCall func(Vm.FunctionCall) Vm.Value

type Vm struct{
   p int
   Func 
}

func New() *Vm {

return &Vm{0,""}

}