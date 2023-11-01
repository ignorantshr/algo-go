#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

// 也叫逆波兰式

int CalculateInversePolishFormula(char* expression) {
    ListStack stack = NULL;
    initListStack(stack);

    char numStr[100];  // 存储提取出的数字字符串
    int j = 0;
    int len = strlen(expression);
    for (int i = 0; i < len; i++) {
        char cur = expression[i];
        if (cur >= '0' && cur <= '9') {
            numStr[j++] = cur;
        } else {
            if (j > 0) {  // last is digital
                numStr[j] = '\0';
                j = 0;
                push(&stack, atoi(numStr));
            }
            if (cur == '|') {
                continue;
            }
            // execute expression
            ElemType left;
            ElemType right;
            pop(&stack, &right);  // right pop first
            pop(&stack, &left);
            int res;
            switch (cur) {
                case '+':
                    res = left + right;
                    push(&stack, res);
                    break;
                case '-':
                    res = left - right;
                    push(&stack, res);
                    break;
                case '*':
                    res = left * right;
                    push(&stack, res);
                    break;
                case '/':
                    res = left / right;
                    push(&stack, res);
                    break;
                default:
                    break;
            }
        }
    }

    int res;
    pop(&stack, &res);
    destroy(stack);
    return res;
}

int main(int argc, char const* argv[]) {
    assert(CalculateInversePolishFormula("1|1|+") == 2);
    assert(CalculateInversePolishFormula("1|1|-") == 0);
    assert(CalculateInversePolishFormula("2|10|*") == 20);
    assert(CalculateInversePolishFormula("10|2|/") == 5);
    assert(CalculateInversePolishFormula("15|7|1|1|+|-|/|3|*|2|1|1|+|+|-") ==
           5);
    return 0;
}
