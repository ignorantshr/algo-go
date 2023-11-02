#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

// 也叫波兰式
// 按照“右优先”原则确定的这些运算符的生效顺序和前缀表达式中各个运算符从右到左出现的次序是相同的

/* a+b*(c-d)-e/f
/ef
-*b/ef
-*b-cd/ef
+a-*b-cd/ef
 */
/*
+a-*b-cd/ef
 */