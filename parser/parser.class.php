<?php
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

}

?>