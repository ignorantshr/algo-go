#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

int cal(char op, int a, int b);

// 中缀转后缀+后缀表达式求值 两个算法结合
int CalculateInfix(char* expression) {
    ListStack numstack = NULL;
    ListStack opstack = NULL;
    int len = strlen(expression);
    char numStr[10];
    int j = 0;

    for (int i = 0; i < len; i++) {
        char cur = expression[i];
        if (cur >= '0' && cur <= '9') {
            numStr[j++] = cur;
        } else {
            if (j > 0) {
                numStr[j] = '\0';
                j = 0;
                push(&numstack, atoi(numStr));
            }
            ElemType op;
            ElemType left;
            ElemType right;
            switch (cur) {
                case '+':
                case '-':
                    while (top(opstack, &op) && op == '+' || op == '-') {
                        pop(&opstack, &op);
                        pop(&numstack, &right);
                        pop(&numstack, &left);
                        push(&numstack, cal(op, left, right));
                    }
                    push(&opstack, cur);
                    break;
                case '*':
                case '/':
                    while (top(opstack, &op) && op == '*' || op == '/') {
                        pop(&opstack, &op);
                        pop(&numstack, &right);
                        pop(&numstack, &left);
                        push(&numstack, cal(op, left, right));
                    }
                    push(&opstack, cur);
                    break;
                case '(':
                    push(&opstack, cur);
                    break;
                case ')':
                    while (pop(&opstack, &op) && op != '(') {
                        pop(&numstack, &right);
                        pop(&numstack, &left);
                        push(&numstack, cal(op, left, right));
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
        push(&numstack, atoi(numStr));
    }

    ElemType op;
    ElemType left;
    ElemType right;
    while (pop(&opstack, &op)) {
        pop(&numstack, &right);
        pop(&numstack, &left);
        push(&numstack, cal(op, left, right));
    }

    ElemType res;
    top(numstack, &res);
    return res;
}

int plus(int a, int b) {
    return a + b;
}

int minus(int a, int b) {
    return a - b;
}

int multiply(int a, int b) {
    return a * b;
}

int divide(int a, int b) {
    return a / b;
}

typedef struct iops {
    char op[4];
    int (*funcs[4])(int a, int b);
} iops;

iops ops = {{'+', '-', '*', '/'}, {plus, minus, multiply, divide}};

int cal(char op, int a, int b) {
    for (int i = 0; i < 4; i++) {
        if (ops.op[i] == op) {
            return ops.funcs[i](a, b);
        }
    }
    return 0;
}

int main(int argc, char const* argv[]) {
    assert(CalculateInfix("1+1") == 2);
    assert(CalculateInfix("1-1") == 0);
    assert(CalculateInfix("1+0") == 1);
    assert(CalculateInfix("2*30") == 60);
    assert(CalculateInfix("10/2") == 5);
    assert(CalculateInfix("((15/(7-(1+1)))*3)-(2+(1+1))") == 5);
    return 0;
}
