package compiler

import (
   //. "fmt"
   )

type Compiler struct{
   p int
   Err string
}

func New() *Compiler {

return &Compiler{0,""}

}