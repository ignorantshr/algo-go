#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

// 也叫波兰式
// 按照“右优先”原则确定的这些运算符的生效顺序和前缀表达式中各个运算符从右到左出现的次序是相同的

/* a+b*(c-d)-e/f 手算
+a[b*(c-d)-e/f]
+a-[b*(c-d)e/f]
+a-*b[(c-d)e/f]
+a-*b-cd[e/f]
+a-*b-cd/ef
*/

/* a+b*(c-d)-e/f 算法原理，和 中缀转后缀 类似
/ef
-*b/ef
-*b-cd/ef
+a-*b-cd/ef
*/

void infix2prefix(char* s, char* res) {
    ListStack opstack;
    int len = strlen(s);
    char numStr[10];
    int j = 0;
    res[0] = '\0';

    for (int i = len; i >= 0; i--) {
        char cur = s[i];
        if (cur >= '0' && cur <= '9') {
            numStr[j++] = cur;
        } else {
            if (j > 0) {
                numStr[j] = '\0';
                j = 0;
                strcat(res, numStr);
            }
            ElemType op;
            switch (cur) {
                case '+':
                case '-':
                    while (top(opstack, &op) && op != ')') {  // '*' '/' '+' '-'
                        pop(&opstack, &op);
                        char oprator[2];
                        snprintf(oprator, sizeof(oprator), "%c", op);
                        strcat(res, oprator);
                    }
                    push(&opstack, cur);
                    break;
                case '*':
                case '/':
                    while (top(opstack, &op) && (op == '*' || op == '/')) {
                        pop(&opstack, &op);
                        char oprator[2];
                        snprintf(oprator, sizeof(oprator), "%c", op);
                        strcat(res, oprator);
                    }
                    push(&opstack, cur);
                    break;
                case ')':
                    push(&opstack, cur);
                    break;
                case '(':
                    while (pop(&opstack, &op) && op != ')') {
                        char oprator[2];
                        snprintf(oprator, sizeof(oprator), "%c", op);
                        strcat(res, oprator);
                    }
                    break;
                default:
                    break;
            }
        }
    }

    if (j > 0) {
        numStr[j] = '\0';
        j = 0;
        strcat(res, numStr);
    }

    ElemType op;
    while (top(opstack, &op)) {
        pop(&opstack, &op);
        char oprator[2];
        snprintf(oprator, sizeof(oprator), "%c", op);
        strcat(res, oprator);
    }

    len = strlen(res);
    for (int i = 0; i < len / 2; i++) {
        char tmp = res[i];
        res[i] = res[len - i - 1];
        res[len - i - 1] = tmp;
    }
}

int main(int argc, char const* argv[]) {
    char res[100];

    infix2prefix("1+1", res);
    assert(!strcmp("+11", res));

    infix2prefix("1-1", res);
    assert(!strcmp("-11", res));

    infix2prefix("1*1", res);
    assert(!strcmp("*11", res));

    infix2prefix("1/1", res);
    assert(!strcmp("/11", res));

    infix2prefix("1+1*(2-3)", res);
    assert(!strcmp("+1*1-23", res));

    infix2prefix("1*(2+3)/(4-5)", res);
    assert(!strcmp("*1/+23-45", res));

    infix2prefix("((5/(7-(1+1)))*3)-(2+(1+1))", res);
    assert(!strcmp("-*/5-7+113+2+11", res));
    return 0;
}
