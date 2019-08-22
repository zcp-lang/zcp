<?php

class Vm{

public $p = 0;

public $f;

public $m = [];

public $fn = [];

public $fsp = [];

public $func = [];

public $stack = [];

public $opcode = [];

public $current = "main";

public function __construct(string $str){

$this->opcode = self::string_decode($str);

}

public function run():bool {

for($this->f=0;$this->f<@count($this->opcode);$this->f++){

echo "-----------<p>";

var_dump(json_encode($this->fn));

var_dump(json_encode($this->opcode[$this->f]));

$this->getclosure($this->opcode[$this->f]);

$this->stack_delete();

var_dump(json_encode($this->stack));

echo $this->p;

echo "<p>";

}

return true;

}

public function getclosure(array $opcode):bool {

if($opcode){

return @call_user_func([$this,$opcode[0]],$opcode);

}

return false;

}

public function PUSH(array $opcode):bool {

$this->stack[$this->p] = (int)$opcode[1];

if($this->stack[$this->p]!=$opcode[1]) return false;

$this->p++;

return true;

}

public function PUSHN(array $opcode):bool {

$this->stack[$this->p] = null;

if($this->stack[$this->p]!=null) return false;

$this->p++;

return true;

}

public function PUSHT(array $opcode):bool {

$this->stack[$this->p] = true;

if($this->stack[$this->p]!=true) return false;

$this->p++;

return true;

}

public function PUSHF(array $opcode):bool {

$this->stack[$this->p] = false;

if($this->stack[$this->p]!=false) return false;

$this->p++;

return true;

}

public function PUSHS(array $opcode):bool {

$this->stack[$this->p] = (string)$opcode[1];

if($this->stack[$this->p]!=$opcode[1]) return false;

$this->p++;

return true;

}

public function ADD(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = $v1 + $v2;

if($this->stack[$this->p]!=$v1+$v2) return false;

$this->p++;

return true;

}

public function SUB(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = $v2 - $v1;

if($this->stack[$this->p]!=$v2-$v1) return false;

$this->p++;

return true;

}

public function MUL(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = $v1 * $v2;

if($this->stack[$this->p]!=$v1*$v2) return false;

$this->p++;

return true;

}

public function DIV(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = $v2 / $v1;

if($this->stack[$this->p]!=$v2/$v1) return false;

$this->p++;

return true;

}

public function YUC(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = $v2 % $v1;

if($this->stack[$this->p]!=$v2%$v1) return false;

$this->p++;

return true;

}

public function STORE(array $opcode):bool {

$v = $this->stack[$this->p-1];

$this->m[$this->current][$opcode[1]] = $v;

if($this->m[$this->current][$opcode[1]]!=$v) return false;

$this->p--;

return true;

}

public function LOAD(array $opcode):bool {

$v = $this->m[$this->current][$opcode[1]];

$this->stack[$this->p] = $v;

if($this->stack[$this->p]!=$v) return false;

$this->p++;

return true;

}

public function JB(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (int)$v2 < (int)$v1;

if($this->stack[$this->p]!=((int)$v2<(int)$v1)) return false;

$this->p++;

return true;

}

public function JG(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (int)$v2 > (int)$v1;

if($this->stack[$this->p]!=((int)$v2>(int)$v1)) return false;

$this->p++;

return true;

}

public function JBE(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (int)$v2 <= (int)$v1;

if($this->stack[$this->p]!=((int)$v2<=(int)$v1)) return false;

$this->p++;

return true;

}

public function JGE(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (int)$v2 >= (int)$v1;

if($this->stack[$this->p]!=((int)$v2>=(int)$v1)) return false;

$this->p++;

return true;

}

public function JE(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (bool)($v2 == $v1);

if($this->stack[$this->p]!=($v2==$v1)) return false;

$this->p++;

return true;

}

public function JNE(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (bool)($v2 != $v1);

if($this->stack[$this->p]!=($v2!=$v1)) return false;

$this->p++;

return true;

}

public function JEO(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (bool)($v2 === $v1);

if($this->stack[$this->p]!=($v2===$v1)) return false;

$this->p++;

return true;

}

public function JNEO(array $opcode):bool {

$v1 = $this->stack[$this->p-1];

$v2 = $this->stack[$this->p-2];

$this->p -= 2;

$this->stack[$this->p] = (bool)($v2 !== $v1);

if($this->stack[$this->p]!=($v2!==$v1)) return false;

$this->p++;

return true;

}

public function JIF(array $opcode):bool {

$v = $this->stack[$this->p-1];

if($v){

$this->p--;
$this->f += 1;

}

return true;

}

public function JMP(array $opcode):bool {

$this->f = (int)$opcode[1];

if($this->f!=$opcode[1]) return false;

$this->f--;

return true;

}

public function FUNCTION(array $opcode):bool {

if($this->fn[$opcode[1]]) return true;

while(($fn=$this->opcode[$this->f])&&$fn[0]!="END"){

$this->f++;

}

$end = $this->opcode[$this->f];

$result = ["start"=>$opcode[3],"end"=>$end[1],"params"=>$opcode[2]];

$this->fn[$opcode[1]] = $result;

if($this->fn[$opcode[1]]!=$result) return false;

return true;

}

public function CALL(array $opcode):bool {

return true;

}

public function stack_delete(){

for($i=$this->p;$i<@count($this->stack);$i++){

unset($this->stack[$i]);

}

}

static function string_decode(string $str):array {

foreach(@explode("\n",$str) as $value){

$result[] = @explode("\x20",$value);

}

return $result;

}

}

?>