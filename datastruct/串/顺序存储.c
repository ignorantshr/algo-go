#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../util.c"

#define MaxSize 255

// 静态数组实现（定长顺序存储）
typedef struct sstr {
    char data[MaxSize];
    int len;
} sstr;

// 动态数组实现（malloc, 堆分配存储）
typedef struct hstr {
    char* data;
    int len;
} hstr;

void StrAssign(sstr* t, char* s) {
    // memcpy(t->data, s, len);
    int len = strlen(s);
    for (int i = 0; i < len; i++) {
        t->data[i] = s[i];
    }
    t->len = len;
    t->data[len] = '\0';
}

void Strcopy(sstr* t, sstr s) {
    // memcpy(t, s.data, s.len);
    for (int i = 0; i < s.len; i++) {
        t->data[i] = s.data[i];
    }
    t->len = s.len;
    t->data[t->len] = '\0';
}

bool StrEmpty(sstr s) {
    return s.len == 0;
}

int StrLength(sstr s) {
    return s.len;
}

void ClearString(sstr* s) {
    s->data[0] = '\0';
    s->len = 0;
}

void DestroyString(sstr* s) {
    // 系统自动回收
}

void Concat(sstr* t, sstr s1, sstr s2) {
    Strcopy(t, s1);
    for (int i = 0; i < s2.len; i++) {
        t->data[i + t->len] = s2.data[i];
    }
    t->len += +s2.len;
    t->data[t->len] = '\0';
}

bool SubString(sstr* sub, sstr s, int pos, int len) {
    if (s.len < pos + len) {
        return false;
    }

    int end = pos + len;
    for (int i = pos; i < end; i++) {
        sub->data[i - pos] = s.data[i];
    }
    sub->len = end - pos;
    return true;
}

int StrCompare(sstr s, sstr t) {
    int min_len = min(s.len, t.len);
    for (int i = 0; i < min_len; i++) {
        if (s.data[i] != t.data[i]) {
            return s.data[i] - t.data[i];
        }
    }
    return s.len - t.len;
}

int Index(sstr s, sstr t) {
    for (int i = 0; i <= s.len - t.len; i++) {
        if (s.data[i] == t.data[0]) {
            sstr sub;
            SubString(&sub, s, i, t.len);
            if (StrCompare(sub, t) == 0) {
                return i;
            }
        }
    }
    return -1;
}

// 朴素模式匹配算法
// 设模式串长度为 m，主串长度为 n.
// 1. 则匹配成功的最好时间复杂度：O(m)
// 2. 匹配失败的最好时间复杂度：O(n).
// 全部匹配失败且每个子串的第1
// 个字符就与模式串不匹配。要匹配n − m + 1 次（主串内有n − m + 1
// 个长度为 m的子串，都要依次匹配），时间复杂度为
// O(n-m+1)=O(n-m),考虑到许多情况下主串长度远大于模式串长度，时间复杂度可进一步约等于O(n)
// 3. 最坏时间复杂度：O(mn)
// 每个子串的前 m−1 个字符都和模式串匹配，只有第 m
// 个字符不匹配。匹配成功/匹配失败 最多需要（n−m + 1）×m次比较。
int Index2(sstr s, sstr t) {
    int i = 0, j = 0, k = 0;
    while (i < s.len && j < t.len) {
        if (s.data[i] == t.data[j]) {
            i++;
            j++;
        } else {
            k++;
            i = k;
            j = 0;
        }
    }
    if (j == t.len) {
        return k;
    }
    return -1;
}

// kmp,O(n+m)
// next[j]=S的最长相等前后缀长度
// https://blog.csdn.net/v_JULY_v/article/details/7041827
void gen_next(sstr s, int next[]) {
    next[0] = -1;
    int j = 0;
    int k = -1;  // 代表 next[j] = k, 即
    // S 的前 k 个字符 S(0,1,...,k-1) 和
    // Sj 的后 k 个字符 S(j-k,j-k+1,...,j-1) 是相等的

    while (j < s.len - 1) {
        if (k == -1 || s.data[k] == s.data[j]) {
            k++;
            j++;

            // =====优化代码=====
            if (s.data[k] == s.data[j]) {
                // 不能出现 S[j] = S[ next[j] ]，因为当 S[j] 失配时，回溯的 j
                // 的位置还是同一个字符，必然不匹配。
                // 所以出现这种情况时需要继续递归，
                // k = next[k] = next[next[k]]
                next[j] = next[k];
                // =====优化代码=====
            } else {
                next[j] = k;  // next[j+1] = next[j]+1 = k+1
            }
        } else {
            k = next[k];  // 不断地向前寻找匹配的前缀字符串
        }
    }
}

int IndexKMP(sstr s, sstr t) {
    int i = 0, j = 0;
    int next[t.len];
    gen_next(t, next);
    while (i < s.len && j < t.len) {
        if (j == -1 || s.data[i] == t.data[j]) {
            i++;
            j++;
        } else {
            // 不匹配时 i 处于不匹配的位置，让 j 回溯
            j = next[j];
        }
    }
    if (j == t.len) {
        return i - j;
    }
    return -1;
}

void printssr(sstr s) {
    printf("%s\n", s.data);
}

void testing(char* s1, char* s2, int index);

int main(int argc, char const* argv[]) {
    hstr hs;
    hs.data = (char*)malloc(MaxSize * sizeof(char));
    hs.len = 0;

    testing("123456", "34", 2);
    testing("123456", "24", -1);
    testing("123423", "2", 1);
    testing("1234235", "235", 4);
    testing("1234235232", "232", 7);
    testing("123454321", "3456345703", -1);
    return 0;
}

void testing(char* s1, char* s2, int index) {
    printf("===== %s, %s, %d =====\n", s1, s2, index);
    sstr ss;
    StrAssign(&ss, s1);

    sstr ss2;
    StrAssign(&ss2, s2);
    assert(Index(ss, ss2) == index);
    assert(Index2(ss, ss2) == index);
    assert(IndexKMP(ss, ss2) == index);
    printf("===== ok =====\n");
}
