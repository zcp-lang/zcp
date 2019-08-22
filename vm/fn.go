package vm

type Fn struct{
   Name string
   Pos int
   Params int
   Ret int
   Start int
   End int
   Asm [][]string
 }