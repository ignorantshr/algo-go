#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

// 也叫波兰式
// 按照“右优先”原则确定的这些运算符的生效顺序和前缀表达式中各个运算符从右到左出现的次序是相同的

// a+b*(c-d)-e/f
// +a-*b-cd/ef
int CalculatePolishFormula(char* expression) {
    ListStack stack = NULL;
    initListStack(stack);

    char numStr[10];  // 存储提取出的数字字符串
    int j = 0;
    int len = strlen(expression);

    for (int i = len - 1; i >= 0; i--) {
        char cur = expression[i];
        if (cur >= '0' && cur <= '9') {
            numStr[j++] = cur;
        } else {
            if (j > 0) {
                numStr[j] = '\0';
                int nl = j;
                // 因为从右向左遍历，所以需要反转数字得到正确的数字
                for (int m = 0; m < nl / 2; m++) {
                    char tmp = numStr[m];
                    numStr[m] = numStr[nl - m - 1];
                    numStr[nl - m - 1] = tmp;
                }

                j = 0;
                push(&stack, atoi(numStr));
            }
            if (cur == '|') {
                continue;
            }
            ElemType left;
            ElemType right;
            pop(&stack, &left);  // left pop first
            pop(&stack, &right);
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
    assert(CalculatePolishFormula("+|1|1") == 2);
    assert(CalculatePolishFormula("-|1|1") == 0);
    assert(CalculatePolishFormula("*|2|10") == 20);
    assert(CalculatePolishFormula("/|10|2") == 5);
    assert(CalculatePolishFormula("-|*|/|15|-|7|+|1|1|3|+|2|+|1|1") == 5);
    return 0;
}
