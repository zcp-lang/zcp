package vm

import (
   //. "fmt"
   )

type _nativeFunction func(NativeCall) Value

type NativeCall struct {
      Vm       *Vm
      Arg      string
}

func (call *NativeCall) Argument(i int) string {

return call.Arg

}