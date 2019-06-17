package compiler

import (
   //. "fmt"
   )

func (self *Compiler) StringDecode(str string) [][]string {

s := []rune(str);

var l []rune;var list,lists [][]rune;var result [][][]rune;

for _,ch := range s {

if ch == '\x0a' { list = append(list,l); l = make([]rune,0); }else{ l = append(l,ch); }

}

list = append(list,l)

for _,ch := range list {

if ch != nil { lists = append(lists,ch); }

}

//lists = append(lists[:len(lists)-1],lists[len(lists):]...)

for _,ch := range lists {

var ll []rune;var tmp [][]rune;

for _,chr := range ch {

if chr == '\x20' { tmp = append(tmp,ll); ll = make([]rune,0); }else{ ll = append(ll,chr); }

}

tmp = append(tmp,ll);result = append(result,tmp)

}

var tt [][]string;

for _,ch := range result {

var t []string;

for _,chr := range ch {

t = append(t,string(chr))

}

tt = append(tt,t)

}

return tt

}