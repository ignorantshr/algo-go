#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

char getMathChar(char r) {
    switch (r) {
        case '}':
            return '{';
        case ')':
            return '(';
        case ']':
            return '[';
        default:
            return ' ';
    }
}

bool mathBraces(char* s, int n) {
    ListStack stack = NULL;
    initListStack(stack);
    bool match = true;

    for (int i = 0; i < n; i++) {
        ElemType t;
        switch (s[i]) {
            case '(':
            case '[':
            case '{':
                push(&stack, s[i]);
                break;
            case ')':
            case ']':
            case '}':
                if (!pop(&stack, &t)) {
                    match = false;
                    goto exit;
                }
                char m = getMathChar(s[i]);
                if (m == ' ' || t != m) {
                    match = false;
                    goto exit;
                }
                break;
            default:
                match = false;
                goto exit;
        }
    }

    ElemType e;
    if (top(stack, &e)) {
        match = false;
    }
exit:
    destroy(stack);
    return match;
}

int main(int argc, char const* argv[]) {
    assert(mathBraces("[]", 2));
    assert(mathBraces("{}", 2));
    assert(mathBraces("()", 2));
    assert(mathBraces("({})", 4));
    assert(mathBraces("({[]})", 6));
    assert(mathBraces("(){}[]", 6));
    assert(mathBraces("(){}[{}]", 8));
    assert(!mathBraces("(){}[{]", 7));
    assert(!mathBraces("(((){}[]", 8));
    assert(!mathBraces("(){}[]]", 7));
    return 0;
}
