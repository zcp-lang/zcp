package ast

type Exp interface{}

type NullExp struct{ Line int }
type TrueExp struct{ Line int }
type FalseExp struct{ Line int }
type VarargExp struct{ Line int }

type IntegerExp struct {
	Line int
	Val  int64
}

type FloatExp struct {
	Line int
	Val  float64
}

type NumberExp struct {
	Line int
	Val  int
}

type StringExp struct {
	Line int
	Str  string
}

type UnopExp struct {
	Line int
	Op   string
	Exp  Exp
}

type BinopExp struct {
	Line int
	Op   string
	Exp1 Exp
	Exp2 Exp
}

type LogicalExp struct {
	Line int
	Op   string
	Exp1 Exp
	Exp2 Exp
}

type UpdateExp struct {
	Line int
	Op   string
	Argument string
	Prefix bool
}

type ConcatExp struct {
	Line int
	Exps []Exp
}

type ArrayConstructorExp struct {
	Line     int
	LastLine int
	KeyExps  []Exp
	ValExps  []Exp
}

type FuncDefExp struct {
	Line     int
	LastLine int
	ParList  []string
	IsVararg bool
	Block    *Block
}

type NameExp struct {
	Line int
	Name string
}

type QuoteExp struct {
	Line int
	Name string
}

type PointerExp struct {
	Line int
	Name string
}

type LengthExp struct {
	Line int
	Name string
}

type CommentExp struct {
	Line int
	Val  string
}

type ParensExp struct {
	Exp Exp
}

type ArrayAccessExp struct {
	LastLine  int
	PrefixExp Exp
	KeyExp    Exp
}

type CallExp struct {
	Line      int
	Prefix    Exp
	Name      StringExp
	Args      []Exp
}

type FuncCallExp struct {
	Line      int
	LastLine  int
	PrefixExp Exp
	NameExp   *StringExp
	Args      []Exp
}
