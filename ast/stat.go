package ast

type Stat interface{}

type BreakStat struct{ Line int }

type ReturnStat struct{
	Line    int
	Exp     Exp
}

type IfStat struct {
	Test       Exp
	Consequent Stat
	Alternate  Stat
}

type WhileStat struct {
	Test   Exp
	Block  Stat
}

type ForStat struct {
	Line int
	InitExp   Exp
	LimitExp  Exp
	StepExp   Exp
	Block     Stat
}

type AssignStat struct {
	Line   int
	Left   Stat
	Right  Exp
}

type FunctionStat struct {
	Id       Exp
	Params   Exp
	Body     Stat
}