# zcp
zcp programming language

还没完成，半成品。有兴趣的可以加我邮箱:2993150260@qq.com一起弄

目录支持的表达式:

true&&false

#a==888

["array":2333,666:["a","b",c]]!=null

(2+3)*4

a[2]["moid"][4*5]+666

"moid" ?* "m"

"moid" *? "d"




<---分割线--->

简单的zcp代码:

a = 2 + 3 * 4;

中间指令集结果(作者懒，没配图):

用堆栈方式，堆做为变量内存。栈用于运算

[[PUSH 2] [PUSH 3] [PUSH 4] [MUL] [ADD] [STORE a]]

等于:

PUSH 2
PUSH 3
PUSH 4
MUL
ADD
STORE a


PUSH向栈压入值。

运行PUSH 2，此时栈为:

0:2

PUSH 3为:

1:3
0:2

PUSH 4为:

2:4
1:3
0:2

MUL用于栈指针(栈顶，top)，向下两个值相乘。然后弹出(pop)这两个值，并把栈指针移下两个值，最后把结果压入栈，也就是

2:4  <---向下第一个值
1:3  <---向下第二个值
0:2

运算后MUL后，栈结构

1:12
0:2

ADD也和MUL一样，不过MUL是乘。ADD是相加

运行ADD指令后，栈结构:

0:14

STORE a是把当前栈顶元素(值)，加载到内存里。a就是变量名，也可以直接用于内存地址

这是简单的zcp代码(a = 2 + 3 * 4;)，从高级语言到低级语言(中间指令集，类似汇编，汇编jvm)的转换结果。

最后让虚拟机运行指令集就行了

官方群:486468643
