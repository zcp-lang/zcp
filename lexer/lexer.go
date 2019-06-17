package lexer

import (
   "strconv"
)

type Tkn struct{
   Token Token
   Value string
   Line int
}

type Lexer struct{
   p int
   line int
   src []rune
   list map[int]string
   err bool
   Errors string
   result map[int]Tkn
}

//初始化
func New(str string) (*Lexer,bool) {

var l Lexer

e := l.New(str)

return &l,e

}

//内部初始化
func (self *Lexer) New(str string) bool {

self.line = 1

self.list = make(map[int]string)

self.result = make(map[int]Tkn)

if self.src = []rune(str);self.src != nil {

return true

}

return false

}

//开始解析全部token
func (self *Lexer) Run() (map[int]Tkn,string) {

for {

if len(self.src) > self.p && self.Errors == "" {

chr := self.src[self.p]

if chr != -1 {

self.result[len(self.result)] = self.Next(); 

}else{ return self.result,self.Errors; }

}else{

return self.result,self.Errors

}

}

}

//获取下个token
func (self *Lexer) Next() Tkn {

var t Tkn
var tmp string

ch := self.Get()

if len(self.src) <= self.p { return t; }

//空格
if ch == "\x0a" {

t.Token = LINE;
t.Value = ch;
t.Line = self.line

self.line++

//换行
}else if ch == "\x20" {

t.Token = SPACE
t.Value = ch
t.Line = self.line

//标识符
}else if "a" <= ch && ch <= "z" || "A" <= ch && ch <= "Z" {

str := self.Identifier()

token := IDENTIFIER

if toke,_:=self.Keyword(str);toke!=ERROR {

/*
if toke == TRUE || toke == FALSE {

token = BOOLEAN

}else{ token = toke; }

*/

token = toke

}

t.Token = token
t.Value = str
t.Line = self.line

//数字
}else if "0" <= ch && ch <= "9" {

t.Token = NUMBER
t.Value = self.Number()
t.Line = self.line

//字符串
}else if ch == "\"" || ch == "'" {

str := self.String()

t.Token = STRING
t.Value = str
t.Line = self.line

//+
}else if ch == "+" {

self.p++

chr := self.Get()

//+=
if chr == "=" {

t.Token = ADD_ASSIGN
t.Value = ch+chr
t.Line = self.line

//++
}else if chr == "+" {

t.Token = INCREMENT
t.Value = ch+chr
t.Line = self.line

//+
}else{

self.p--

t.Token = PLUS
t.Value = ch
t.Line = self.line

}

//-
}else if ch == "-" {

self.p++

chr := self.Get()

//-=
if chr == "=" {

t.Token = SUBTRACT_ASSIGN
t.Value = ch+chr
t.Line = self.line

//--
}else if chr == "-" {

t.Token = DECREMENT
t.Value = ch+chr
t.Line = self.line

//-
}else{

self.p--

t.Token = MINUS
t.Value = ch
t.Line = self.line

}

//*
}else if ch == "*" {

self.p++

chr := self.Get()

//*指针
if "a" <= chr && chr <= "z" || "A" <= chr && chr <= "Z" {

t.Token = POINTER
t.Value = self.Identifier()
t.Line = self.line

//*=
}else if chr == "=" {

t.Token = MULTIPLY_ASSIGN
t.Value = ch+chr
t.Line = self.line

//*?
}else if chr == "?" {

t.Token = STRING_END
t.Value = ch+chr
t.Line = self.line

//*
}else{

self.p--

t.Token = MULTIPLY
t.Value = ch
t.Line = self.line

}

// /
}else if ch == "/" {

self.p++

chr := self.Get()

// /=
if chr == "=" {

t.Token = QUOTIENT_ASSIGN
t.Value = ch+chr
t.Line = self.line

// //
}else if chr == "/" {

var value string

line := self.line

for {

self.p++

if len(self.src) <= self.p { break; }

chrs := self.src[self.p]

if chrs == '\x0a' || chrs == -0 { self.p--; break; }

value += string(chrs)

}

t.Token = COMMENT
t.Value = value
t.Line = line

// /* .... */
}else if chr == "*" {

var value string

line := self.line

for {

self.p++;

if len(self.src) <= self.p { break; }

left := self.src[self.p]

if len(self.src) <= self.p { break; }

right := self.src[self.p+1]

if left == -0 { self.p++; break; }

if(left == '*' && right == '/'){

self.p++; break;

}

if left == '\x0a' { self.line++; }

value += string(left)

}

t.Token = COMMENT
t.Value = value
t.Line = line

// /
}else{

self.p--

t.Token = SLASH
t.Value = ch
t.Line = self.line

}

//%
}else if ch == "%" {

self.p++

chr := self.Get()

//%=
if chr == "=" {

t.Token = REMAINDER_ASSIGN
t.Value = ch+chr
t.Line = self.line

//%
}else{

self.p--

t.Token = REMAINDER
t.Value = ch
t.Line = self.line

}

//&
}else if ch == "&" {

self.p++

chr := self.Get()

//&地址引用
if "a" <= chr && chr <= "z" || "A" <= chr && chr <= "Z" {

t.Token = QUOTE
t.Value = self.Identifier()
t.Line = self.line

//&^
}else if chr == "^" {

t.Token = AND_NOT
t.Value = ch+chr
t.Line = self.line

//&=
}else if chr == "=" {

t.Token = AND_ASSIGN
t.Value = ch+chr
t.Line = self.line

//&&
}else if chr == "&" {

t.Token = LOGICAL_AND
t.Value = ch+chr
t.Line = self.line

//&
}else{

self.p--

t.Token = AND
t.Value = ch
t.Line = self.line

}

//|
}else if ch == "|" {

self.p++

chr := self.Get()

//|=
if chr == "=" {

t.Token = OR_ASSIGN
t.Value = ch+chr
t.Line = self.line

//||
}else if chr == "|" {

t.Token = LOGICAL_OR
t.Value = ch+chr
t.Line = self.line

//|
}else{

self.p--

t.Token = OR
t.Value = ch
t.Line = self.line

}

//^
}else if ch == "^" {

self.p++

chr := self.Get()

//^=
if chr == "=" {

t.Token = EXCLUSIVE_OR_ASSIGN
t.Value = ch+chr
t.Line = self.line

//^
}else{

self.p--

t.Token = EXCLUSIVE_OR
t.Value = ch
t.Line = self.line

}

//=
}else if ch == "=" {

self.p++

chr := self.Get()

//==
if chr == "=" {

self.p++

chrs := self.Get()

//===
if chrs == "=" {

t.Token = STRICT_EQUAL
t.Value = ch+chr+chrs
t.Line = self.line

//==
}else{

self.p--

t.Token = EQUAL
t.Value = ch+chr
t.Line = self.line

}

//=
}else{

self.p--

t.Token = ASSIGN
t.Value = ch
t.Line = self.line

}

//<
}else if ch == "<" {

self.p++

chr := self.Get()

//<=
if chr == "=" {

t.Token = LESS_OR_EQUAL
t.Value = ch+chr
t.Line = self.line

//<<
}else if chr == "<" {

t.Token = SHIFT_LEFT
t.Value = ch+chr
t.Line = self.line

//<
}else{

self.p--

t.Token = LESS
t.Value = ch
t.Line = self.line

}

//>
}else if ch == ">" {

self.p++

chr := self.Get()

//>=
if chr == "=" {

t.Token = GREATER_OR_EQUAL
t.Value = ch+chr
t.Line = self.line

//>>
}else if chr == ">" {

t.Token = SHIFT_RIGHT
t.Value = ch+chr
t.Line = self.line

//>
}else{

self.p--

t.Token = GREATER
t.Value = ch
t.Line = self.line

}

//!
}else if ch == "!" {

self.p++

chr := self.Get()

//!反意
if "a" <= chr && chr <= "z" || "A" <= chr && chr <= "Z" {

t.Token = CONTRARY
t.Value = self.Identifier()
t.Line = self.line

//!=
}else if chr == "=" {

self.p++

chrs := self.Get()

//!==
if chrs == "=" {

t.Token = STRICT_NOT_EQUAL
t.Value = ch+chr+chrs
t.Line = self.line

//!=
}else{

self.p--

t.Token = NOT_EQUAL
t.Value = ch+chr
t.Line = self.line

}

//!
}else{

self.p--

t.Token = NOT
t.Value = ch
t.Line = self.line

}

//~
}else if ch == "~" {

t.Token = BITWISE_NOT
t.Value = ch
t.Line = self.line

//(
}else if ch == "(" {

t.Token = LEFT_PARENTHESIS
t.Value = ch
t.Line = self.line

//[
}else if ch == "[" {

t.Token = LEFT_BRACKET
t.Value = ch
t.Line = self.line

//{
}else if ch == "{" {

t.Token = LEFT_BRACE
t.Value = ch
t.Line = self.line

//)
}else if ch == ")" {

t.Token = RIGHT_PARENTHESIS
t.Value = ch
t.Line = self.line

//]
}else if ch == "]" {

t.Token = RIGHT_BRACKET
t.Value = ch
t.Line = self.line

//}
}else if ch == "}" {

t.Token = RIGHT_BRACE
t.Value = ch
t.Line = self.line

//,
}else if ch == "," {

t.Token = COMMA
t.Value = ch
t.Line = self.line

//.
}else if ch == "." {

self.p++

chr := self.Get()

//..
if chr == "." {

t.Token = CONCAT
t.Value = ch+chr
t.Line = self.line

//.
}else{

self.p--

t.Token = PERIOD
t.Value = ch
t.Line = self.line

}

//;
}else if ch == ";" {

t.Token = SEMICOLON
t.Value = ch
t.Line = self.line

//:
}else if ch == ":" {

t.Token = COLON
t.Value = ch
t.Line = self.line

//?
}else if ch == "?" {

self.p++

chr := self.Get()

//?*
if chr == "*" {

t.Token = STRING_START
t.Value = ch+chr
t.Line = self.line

//?
}else{

self.p--

t.Token = QUESTION_MARK
t.Value = ch
t.Line = self.line

}

//#
}else if ch == "#" {

self.p++

chr := self.Get()

//#获取length长度
if "a" <= chr && chr <= "z" || "A" <= chr && chr <= "Z" {

t.Token = LENGTH
t.Value = self.Identifier()
t.Line = self.line

//#
}else{

self.p--

t.Token = ERROR
t.Value = ch
t.Line = self.line

}

}else{

if ch != "" {

self.error(self.line,ch)

}

}

self.p++
self.err = true

if tmp == "" {

tmp = t.Value

}

if tmp != "\x0a" {

self.list[self.line] = self.list[self.line]+tmp;

}

return t

}

func (self *Lexer) error(line int,str string) {

self.Errors = "词法分析错误:在第"+strconv.Itoa(line)+"行,error数据["+str+"],具体代码--->"+self.list[line]+self.Line();

}

func (self *Lexer) Line() string {

var result string

for {

if len(self.src) <= self.p { break; }

ch := self.src[self.p]

if ch == '\x0a' || ch == -0 { return result; }

self.p++; result += string(ch);

}

return result

}

//获取单个字符
func (self *Lexer) Get() string {

if len(self.src) <= self.p { return ""; }

return string(self.src[self.p])

}

//标识符解析
func (self *Lexer) Identifier() string {

var result string

for {

if len(self.src) <= self.p { break; }

ch := self.src[self.p]

if ch == -0 { self.p--; break; }

if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' || ch == '_' {

result += string(ch); self.p++;

}else{ self.p--; break; }

}

return result;

}

//字符串解析
func (self *Lexer) String() string {

var result string

p := self.src[self.p]

self.p++

for {

if len(self.src) <= self.p { break; }

ch := self.src[self.p]

if(ch == '\x0a'){ self.line++; }

if(ch == -0){ return ""; }

if(ch != p){ result += string(ch); self.p++; }else{ break; }

}

return result;

}

//数字截取
func (self *Lexer) Number() string {

var code bool
var next rune
var result string

for {

if len(self.src) <= self.p { break; }

ch := self.src[self.p]

if len(self.src) > self.p+1 { next = self.src[self.p+1] }

if ch == -1 { self.p--; break; }

if ('0' <= ch && ch <= '9') || (ch == '.' && '0' <= next && next <= '9') { 

if ch == '.' && code { return result; }else if ch == '.' { code = true; }

result += string(ch); self.p++;

}else{ self.p--; break; }

}

return string(result)

}

//关键字读取
func (self *Lexer) Keyword(str string) (Token,string) {

for k,v := range keyword {

if v == str { return k,v }

}

return ERROR,""

}