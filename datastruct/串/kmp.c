#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void gen_next(char* s, int next[]) {
    int len = strlen(s);
    int j = 0;
    int k = -1;
    // next[j]=k, 代表以 s[j-1] 为结尾的最长前后缀匹配长度为 k
    // 即 s[0:k-1] 与 s[j-k:j-1] 相等

    next[0] = -1;
    // 已知 next[j]=k, 求 next[j+1]
    while (j < len - 1) {
        if (k == -1 || s[j] == s[k]) {
            // 如果 s[j] == s[k], 则说明 s[0:k] 与 s[j-k:j] 匹配，
            // next[j+1] = next[j] + 1
            next[++j] = ++k;
        } else {
            // 如果 s[j] != s[k], 则说明 s[0:k] 与 s[j-k:j] 不匹配
            // 继续回溯，寻找下一个可能的匹配使得 s[0:t] == s[j-t:j], t<k
            k = next[k];
        }
    }
}

int kmp(char* s, char* t) {
    int slen = strlen(s);
    int tlen = strlen(t);

    int next[tlen];
    gen_next(t, next);

    int i = 0;
    int j = 0;

    while (i < slen && j < tlen) {
        // j == -1 说明找不到字符 t[j] 与 s[i] 匹配
        if (j == -1 || s[i] == t[j]) {
            i++;
            j++;
        } else {
            j = next[j];
        }
    }
    if (j == tlen) {
        return i - j;
    }
    return -1;
}

void testing(char* s, char* t, int index) {
    printf("===== %s, %s, %d =====\n", s, t, index);
    assert(kmp(s, t) == index);
    printf("===== ok =====\n");
}

int main(int argc, char const* argv[]) {
    testing("123456", "34", 2);
    testing("123456", "24", -1);
    testing("123423", "2", 1);
    testing("1234235", "235", 4);
    testing("1234235232", "232", 7);
    testing("123454321", "3456345703", -1);
    return 0;
}
