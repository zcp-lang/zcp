<?php

public function Exp1():array {

$p = $this->p;

$exp1 = $this->Exp0();

if(!$exp1){ $this->p = $p; return []; }

$p = $this->p;

$op = $this->token_get();

if(!$op){ $this->p = $p; return $exp1; }

$exp2 = $this->Exp1();

if(($op['token']=='MULTIPLY'||$op['token']=='SLASH'||$op['token']=='REMAINDER')&&($exp2)){

return ['type'=>'BinaryExpression','op'=>$op['value'],'left'=>$exp1,'right'=>$exp2];

}

$this->p = $p;

return $exp1;

}

public function Exp2():array {

$p = $this->p;

$exp1 = $this->Exp1();

if(!$exp1){ $this->p = $p; return []; }

$p = $this->p;

$op = $this->token_get();

if(!$op){ $this->p = $p; return $exp1; }

$exp2 = $this->Exp2();

if(($op['token']=='PLUS'||$op['token']=='MINUS')&&($exp2)){

return ['type'=>'BinaryExpression','op'=>$op['value'],'left'=>$exp1,'right'=>$exp2];

}

$this->p = $p;

return $exp1;

}

public function Exp3():array {

$p = $this->p;

$exp1 = $this->Exp2();

if(!$exp1){ $this->p = $p; return []; }

$p = $this->p;

$op = $this->token_get();

if(!$op){ $this->p = $p; return $exp1; }

$exp2 = $this->Exp3();

if(($op['token']=='CONCAT')&&($exp2)){

return ['type'=>'BinaryExpression','op'=>$op['value'],'left'=>$exp1,'right'=>$exp2];

}

$this->p = $p;

return $exp1;

}

public function Exp4():array {

$p = $this->p;

$exp1 = $this->Exp3();

if(!$exp1){ $this->p = $p; return []; }

$p = $this->p;

$op = $this->token_get();

if(!$op){ $this->p = $p; return $exp1; }

$exp2 = $this->Exp4();

if(($op['token']=='LESS'||$op['token']=='GREATER'||$op['token']=='LESS_OR_EQUAL'||$op['token']=='GREATER_OR_EQUAL'||$op['token']=='NOT_EQUAL'||$op['token']=='EQUAL'||$op['token']=='STRICT_NOT_EQUAL'||$op['token']=='STRICT_EQUAL'||$op['token']=='STRING_START'||$op['token']=='STRING_END'||$op['token']=='QUESTION_MARK')&&($exp2)){

return ['type'=>'BinaryExpression','op'=>$op['value'],'left'=>$exp1,'right'=>$exp2];

}

$this->p = $p;

return $exp1;

}

public function Exp5():array {

$p = $this->p;

$exp1 = $this->Exp4();

if(!$exp1){ $this->p = $p; return []; }

$p = $this->p;

$op = $this->token_get();

if(!$op){ $this->p = $p; return $exp1; }

$exp2 = $this->Exp5();

if(($op['token']=='LOGICAL_AND')&&($exp2)){

return ['type'=>'LogicalExpression','op'=>$op['value'],'left'=>$exp1,'right'=>$exp2];

}

$this->p = $p;

return $exp1;

}

public function Exp6():array {

$p = $this->p;

$exp1 = $this->Exp5();

if(!$exp1){ $this->p = $p; return []; }

$p = $this->p;

$op = $this->token_get();

if(!$op){ $this->p = $p; return $exp1; }

$exp2 = $this->Exp6();

if(($op['token']=='LOGICAL_OR')&&($exp2)){

return ['type'=>'LogicalExpression','op'=>$op['value'],'left'=>$exp1,'right'=>$exp2];

}

$this->p = $p;

return $exp1;

}

public function Exp():array {

$p = $this->p;

$exp = $this->Exp6();

if($exp){ return $exp; }

$this->p = $p;

return [];

}

public function Var():array {

$p = $this->p;

$var = $this->token_get();

if($var['token']!='S'&&$var['token']!='SS'&&$var['token']!='SSS'){ $this->p = $p; return []; }

$id = $this->token_get();

if($id['token']!='POINTER'&&$id['token']!='IDENTIFIER'){ $this->p = $p; return []; }

if($this->token_get()['token']!='ASSIGN'){ $this->p = $p; return []; }

$exp = $this->Exp();

if(!$exp){ $this->p = $p; return [];}

return ['type'=>'VariableDeclarator','var'=>$var['token'],'id'=>$id,'init'=>$exp];

}

public function Assign():array {

$p = $this->p;

$id = $this->token_get();

$id['type'] = $id['token'];

if($id['token']!='POINTER'&&$id['token']!='IDENTIFIER'){ $this->p = $p; return []; }

if($this->token_get()['token']!='ASSIGN'){ $this->p = $p; return []; }

$exp = $this->Exp();

if(!$exp){ $this->p = $p; return [];}

return ['type'=>'AssignmentExpression','op'=>'=','left'=>$id,'right'=>$exp];

}

public function M():array {

$result = [];

$p = $this->p;

$exp1 = $this->Exp();

if(!$exp1){ $this->p = $p; return []; }

$result[] = $exp1;

while(true){

$p = $this->p;

$op = $this->token_get();

if(!$op||$op['token']!='COMMA'){ $this->p = $p; return $result; }

$exp2 = $this->Exp();

if(!$exp2){ $this->p = $p; return $result; }

$result[] = $exp2;

}

return $result;

}

public function Mm():array {

$result = [];

$p = $this->p;

$id1 = $this->token_get();

if($id1['token']!='POINTER'&&$id1['token']!='IDENTIFIER'){ $this->p = $p; return []; }

$result[] = $id1;

while(true){

$p = $this->p;

$op = $this->token_get();

if(!$op||$op['token']!='COMMA'){ $this->p = $p; return $result; }

$id2 = $this->token_get();

if($id2['token']!='POINTER'&&$id2['token']!='IDENTIFIER'){ $this->p = $p; return $result; }

$result[] = $id2;

}

return $result;

}

public function Call():array {

$p = $this->p;

$id = $this->token_get();

$id['type'] = $id['token'];

if($id['token']!='IDENTIFIER'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_PARENTHESIS'){ $this->p = $p; return []; }

$exp = $this->M();

$p2 = $this->p;

if(!$exp){ $this->p = $p2; }

if($this->token_get()['token']!='RIGHT_PARENTHESIS'){ $this->p = $p; return []; }

return ['type'=>'CallExpression','callee'=>$id,'arguments'=>$exp];

}

public function F():array {

$p = $this->p;

if($this->token_get()['token']!='F'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_PARENTHESIS'){ $this->p = $p; return []; }

$p2 = $this->p;

$exp = $this->Exp();

if(!$exp){ $this->p = $p2; }

if($this->token_get()['token']!='RIGHT_PARENTHESIS'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_BRACE'){ $this->p = $p; return []; }

$p3 = $this->p;

$stat = $this->Ss(false);

if(!$stat){ $this->p = $p3; }

if($this->token_get()['token']!='RIGHT_BRACE'){ $this->p = $p; return []; }

$p4 = $this->p;

if($this->token_get()['token']!='ELSE'){ $this->p = $p4; return ['type'=>'IfStatement','test'=>$exp,'consequent'=>$stat]; }

$p5 = $this->p;

$fs = $this->F();

if($fs){ return ['type'=>'IfStatement','test'=>$exp,'consequent'=>$stat,'alternate'=>[$fs]]; }

$this->p = $p5;

if($this->token_get()['token']!='LEFT_BRACE'){ $this->p = $p4; return ['type'=>'IfStatement','test'=>$exp,'consequent'=>$stat]; }

$p6 = $this->p;

$stat2 = $this->Ss(false);

if(!$stat2){ $this->p = $p6; }

if($this->token_get()['token']!='RIGHT_BRACE'){ $this->p = $p4; return ['type'=>'IfStatement','test'=>$exp,'consequent'=>$stat]; }

return ['type'=>'IfStatement','test'=>$exp,'consequent'=>$stat,'alternate'=>$stat2];

}

public function For():array {

$p = $this->p;

if($this->token_get()['token']!='FOR'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_PARENTHESIS'){ $this->p = $p; return []; }

$p1 = $this->p;

$left = $this->Exp();

if(!$left){ $this->p = $p1; }

if($this->token_get()['token']!='SEMICOLON'){ $this->p = $p; return []; }

$p2 = $this->p;

$right = $this->Exp();

if(!$right){ $this->p = $p2; }

if($this->token_get()['token']!='RIGHT_PARENTHESIS'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_BRACE'){ $this->p = $p; return []; }

$p3 = $this->p;

$stat = $this->Ss(false);

if(!$stat){ $this->p = $p3; }

if($this->token_get()['token']!='RIGHT_BRACE'){ $this->p = $p; return []; }

return ['type'=>'ForStatement','left'=>$left,'right'=>$right,'body'=>$stat];

}

public function W():array {

$p = $this->p;

if($this->token_get()['token']!='W'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_PARENTHESIS'){ $this->p = $p; return []; }

$p2 = $this->p;

$exp = $this->Exp();

if(!$exp){ $this->p = $p2; }

if($this->token_get()['token']!='RIGHT_PARENTHESIS'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_BRACE'){ $this->p = $p; return []; }

$p3 = $this->p;

$stat = $this->Ss(false);

if(!$stat){ $this->p = $p3; }

if($this->token_get()['token']!='RIGHT_BRACE'){ $this->p = $p; return []; }

return ['type'=>'WhileStatement','test'=>$exp,'body'=>$stat];

}

public function T():array {

$p = $this->p;

if($this->token_get()['token']!='T'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_PARENTHESIS'){ $this->p = $p; return []; }

$p2 = $this->p;

$exp = $this->Exp();

if(!$exp){ $this->p = $p2; }

if($this->token_get()['token']!='RIGHT_PARENTHESIS'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_BRACE'){ $this->p = $p; return []; }

$p3 = $this->p;

$stat = $this->Ss(false);

if(!$stat){ $this->p = $p3; }

if($this->token_get()['token']!='RIGHT_BRACE'){ $this->p = $p; return []; }

return ['type'=>'ThreadStatement','test'=>$exp,'body'=>$stat];

}

public function Fn():array {

$p = $this->p;

if($this->token_get()['token']!='FN'){ $this->p = $p; return []; }

$id = $this->token_get();

$id['type'] = $id['token'];

if($id['token']!='IDENTIFIER'){ $this->p = $p; return []; }

if($this->token_get()['token']!='LEFT_PARENTHESIS'){ $this->p = $p; return []; }

$p1 = $this->p;

$exp = $this->Mm();

if(!$exp){ $this->p = $p1; }

if($this->token_get()['token']!='RIGHT_PARENTHESIS'){ $this->p = $p; return []; }

$p2 = $this->p;

$stat = $this->Ss(false);

if(!$stat){ $this->p = $p2; }

if($this->token_get()['token']!='END'){ $this->p = $p; return []; }

if($this->token_get()['token']!='FN'){ $this->p = $p; return []; }

return ['type'=>'FunctionDeclaration','id'=>$id,'params'=>$exp,'body'=>$stat];

}

public function Return():array {

$p = $this->p;

$t = $this->token_get();

if($t['token']!='ENDCODE'&&$t['token']!='RETURN'){ $this->p = $p; return []; }

$p1 = $this->p;

$exp = $this->Exp();

if(!$exp){ $this->p = $p1; return ['type'=>'ReturnStatement']; }

return ['type'=>'ReturnStatement','argument'=>$exp];

}

public function S():array {

$p = $this->p;

if($fn=$this->FN()){ return $fn; }

$this->p = $p;

if($t=$this->T()){ return $t; }

$this->p = $p;

if($f=$this->F()){ return $f; }

$this->p = $p;

if($for=$this->For()){ return $for; }

$this->p = $p;

if($w=$this->W()){ return $w; }

$this->p = $p;

if($var=$this->Var()){ return $var; }

$this->p = $p;

if($assign=$this->Assign()){ return $assign; }

$this->p = $p;

if($return=$this->Return()){ return $return; }

$this->p = $p;

if($exp=$this->Exp()){ return $exp; }

$this->p = $p;

if($comment=$this->Comment()){ return $comment; }

$this->p = $p;

return [];

}

public function Ss($code):array {

$result = [];

while(true){

$p = $this->p;

$ast = $this->S();

if(!$ast && $code && $this->err && ($t=$this->token_get())){

$this->error = "语法分析错误:在第".$t['line']."行,类型为(".$t['token']."),具体代码--->".$this->list[$t['line']];

return $result;

}

if(!$ast){

return $result;

}

$result[] = $ast;

$this->err = true;

}

}

public function token():array {

$result = $this->token[$this->p];

if(!$result){ return []; }

$this->p++;

return $result;

}

public function token_get():array {

while(true){

$result = $this->token[$this->p];

if(!$result){ return []; }

if($result['token'] != 'LINE' && $result['token'] != 'SPACE'){

$this->p++;

return $result;

}

$this->p++;

}

}

}

?>