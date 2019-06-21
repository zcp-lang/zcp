package compiler

import (
   //. "fmt"
   )

type Compiler struct{
   p int
   Err string
}

func New_Compiler() *Compiler {

return &Compiler{0,""}

}